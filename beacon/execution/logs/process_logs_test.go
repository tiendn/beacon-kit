// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package logs_test

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	beacontypesv1 "github.com/itsdevbear/bolaris/beacon/core/types/v1"
	loghandler "github.com/itsdevbear/bolaris/beacon/execution/logs"
	"github.com/itsdevbear/bolaris/beacon/staking/logs"
	logmocks "github.com/itsdevbear/bolaris/beacon/staking/logs/mocks"
	"github.com/itsdevbear/bolaris/contracts/abi"
	enginetypes "github.com/itsdevbear/bolaris/engine/types"
	"github.com/stretchr/testify/require"
)

func TestProcessLogs(t *testing.T) {
	contractAddress := ethcommon.HexToAddress("0x1234")
	stakingAbi, err := abi.StakingMetaData.GetAbi()
	require.NoError(t, err)

	stakingLogRequest, err := logs.NewStakingRequest(
		contractAddress,
	)
	require.NoError(t, err)
	logFactory, err := loghandler.NewFactory(
		loghandler.WithRequest(stakingLogRequest),
	)
	require.NoError(t, err)

	blkNum := uint64(100)
	depositFactor := 3
	numDepositLogs := 10
	mockLogs, err := logmocks.CreateDepositLogs(
		numDepositLogs,
		depositFactor,
		contractAddress,
		blkNum,
	)
	require.NoError(t, err)

	blockHash := [32]byte{byte(blkNum)}
	vals, err := logFactory.ProcessLogs(mockLogs, blockHash)
	require.NoError(t, err)
	require.Len(t, vals, numDepositLogs)

	// Check if the values are returned in the correct order.
	for i, val := range vals {
		processedDeposit, ok := val.Interface().(*beacontypesv1.Deposit)
		require.True(t, ok)
		require.Equal(t, uint64(i*depositFactor), processedDeposit.GetAmount())
	}

	withdrawal := enginetypes.NewWithdrawal(
		[]byte("pubkey"),
		uint64(1000),
	)

	var log *ethtypes.Log
	log, err = logmocks.NewLogFromWithdrawal(
		stakingAbi.Events[logs.WithdrawalName],
		withdrawal,
	)
	require.NoError(t, err)

	log.Address = contractAddress
	log.BlockNumber = blkNum + 1
	blockHash = [32]byte{byte(blkNum + 1)}
	log.BlockHash = blockHash
	_, err = logFactory.ProcessLogs(
		append(mockLogs, *log),
		blockHash,
	)
	// This is an expected error as
	// the log is from a different block.
	require.Error(t, err)

	// This log is skipped because it is not
	// from the contract address of interest.
	log.Address = ethcommon.HexToAddress("0x5678")
	log.BlockNumber = blkNum
	blockHash = [32]byte{byte(blkNum)}
	log.BlockHash = blockHash
	mockLogs = append(mockLogs, *log)
	vals, err = logFactory.ProcessLogs(mockLogs, blockHash)
	require.NoError(t, err)
	require.Len(t, vals, numDepositLogs)

	log.Address = contractAddress
	log.BlockNumber = blkNum
	blockHash = [32]byte{byte(blkNum)}
	log.BlockHash = blockHash
	mockLogs = append(mockLogs, *log)
	_, err = logFactory.ProcessLogs(mockLogs, blockHash)
	// This is an expected error as currently we cannot
	// unmarsal a withdrawal log into a Withdrawal object.
	require.Error(t, err)
}

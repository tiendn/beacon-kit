// Code generated by fastssz. DO NOT EDIT.
// Hash: 94906e4acf612b85db780bf78626239b47764c89e8bc6c291231775d4e978f78
// Version: 0.1.3
package state

import (
	"github.com/berachain/beacon-kit/beacon/core/types"
	"github.com/berachain/beacon-kit/primitives"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconStateDeneb object
func (b *BeaconStateDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconStateDeneb object to a target array
func (b *BeaconStateDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(224)

	// Field (0) 'GenesisValidatorsRoot'
	dst = append(dst, b.GenesisValidatorsRoot[:]...)

	// Field (1) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (2) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(types.BeaconBlockHeader)
	}
	if dst, err = b.LatestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (3) 'BlockRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlockRoots) * 32

	// Offset (4) 'StateRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.StateRoots) * 32

	// Field (5) 'Eth1GenesisHash'
	dst = append(dst, b.Eth1GenesisHash[:]...)

	// Field (6) 'Eth1DepositIndex'
	dst = ssz.MarshalUint64(dst, b.Eth1DepositIndex)

	// Offset (7) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 121

	// Offset (8) 'RandaoMixes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.RandaoMixes) * 32

	// Field (9) 'NextWithdrawalIndex'
	dst = ssz.MarshalUint64(dst, b.NextWithdrawalIndex)

	// Field (10) 'NextWithdrawalValidatorIndex'
	dst = ssz.MarshalUint64(dst, b.NextWithdrawalValidatorIndex)

	// Field (3) 'BlockRoots'
	if size := len(b.BlockRoots); size > 8192 {
		err = ssz.ErrListTooBigFn("BeaconStateDeneb.BlockRoots", size, 8192)
		return
	}
	for ii := 0; ii < len(b.BlockRoots); ii++ {
		dst = append(dst, b.BlockRoots[ii][:]...)
	}

	// Field (4) 'StateRoots'
	if size := len(b.StateRoots); size > 8192 {
		err = ssz.ErrListTooBigFn("BeaconStateDeneb.StateRoots", size, 8192)
		return
	}
	for ii := 0; ii < len(b.StateRoots); ii++ {
		dst = append(dst, b.StateRoots[ii][:]...)
	}

	// Field (7) 'Validators'
	if size := len(b.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconStateDeneb.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (8) 'RandaoMixes'
	if size := len(b.RandaoMixes); size > 65536 {
		err = ssz.ErrListTooBigFn("BeaconStateDeneb.RandaoMixes", size, 65536)
		return
	}
	for ii := 0; ii < len(b.RandaoMixes); ii++ {
		dst = append(dst, b.RandaoMixes[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconStateDeneb object
func (b *BeaconStateDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 224 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o4, o7, o8 uint64

	// Field (0) 'GenesisValidatorsRoot'
	copy(b.GenesisValidatorsRoot[:], buf[0:32])

	// Field (1) 'Slot'
	b.Slot = primitives.Slot(ssz.UnmarshallUint64(buf[32:40]))

	// Field (2) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(types.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.UnmarshalSSZ(buf[40:152]); err != nil {
		return err
	}

	// Offset (3) 'BlockRoots'
	if o3 = ssz.ReadOffset(buf[152:156]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 224 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (4) 'StateRoots'
	if o4 = ssz.ReadOffset(buf[156:160]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (5) 'Eth1GenesisHash'
	copy(b.Eth1GenesisHash[:], buf[160:192])

	// Field (6) 'Eth1DepositIndex'
	b.Eth1DepositIndex = ssz.UnmarshallUint64(buf[192:200])

	// Offset (7) 'Validators'
	if o7 = ssz.ReadOffset(buf[200:204]); o7 > size || o4 > o7 {
		return ssz.ErrOffset
	}

	// Offset (8) 'RandaoMixes'
	if o8 = ssz.ReadOffset(buf[204:208]); o8 > size || o7 > o8 {
		return ssz.ErrOffset
	}

	// Field (9) 'NextWithdrawalIndex'
	b.NextWithdrawalIndex = ssz.UnmarshallUint64(buf[208:216])

	// Field (10) 'NextWithdrawalValidatorIndex'
	b.NextWithdrawalValidatorIndex = ssz.UnmarshallUint64(buf[216:224])

	// Field (3) 'BlockRoots'
	{
		buf = tail[o3:o4]
		num, err := ssz.DivideInt2(len(buf), 32, 8192)
		if err != nil {
			return err
		}
		b.BlockRoots = make([][32]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlockRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (4) 'StateRoots'
	{
		buf = tail[o4:o7]
		num, err := ssz.DivideInt2(len(buf), 32, 8192)
		if err != nil {
			return err
		}
		b.StateRoots = make([][32]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(b.StateRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (7) 'Validators'
	{
		buf = tail[o7:o8]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*types.Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(types.Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (8) 'RandaoMixes'
	{
		buf = tail[o8:]
		num, err := ssz.DivideInt2(len(buf), 32, 65536)
		if err != nil {
			return err
		}
		b.RandaoMixes = make([][32]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(b.RandaoMixes[ii][:], buf[ii*32:(ii+1)*32])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconStateDeneb object
func (b *BeaconStateDeneb) SizeSSZ() (size int) {
	size = 224

	// Field (3) 'BlockRoots'
	size += len(b.BlockRoots) * 32

	// Field (4) 'StateRoots'
	size += len(b.StateRoots) * 32

	// Field (7) 'Validators'
	size += len(b.Validators) * 121

	// Field (8) 'RandaoMixes'
	size += len(b.RandaoMixes) * 32

	return
}

// HashTreeRoot ssz hashes the BeaconStateDeneb object
func (b *BeaconStateDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconStateDeneb object with a hasher
func (b *BeaconStateDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisValidatorsRoot'
	hh.PutBytes(b.GenesisValidatorsRoot[:])

	// Field (1) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (2) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(types.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'BlockRoots'
	{
		if size := len(b.BlockRoots); size > 8192 {
			err = ssz.ErrListTooBigFn("BeaconStateDeneb.BlockRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlockRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.BlockRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, 8192)
	}

	// Field (4) 'StateRoots'
	{
		if size := len(b.StateRoots); size > 8192 {
			err = ssz.ErrListTooBigFn("BeaconStateDeneb.StateRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.StateRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.StateRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, 8192)
	}

	// Field (5) 'Eth1GenesisHash'
	hh.PutBytes(b.Eth1GenesisHash[:])

	// Field (6) 'Eth1DepositIndex'
	hh.PutUint64(b.Eth1DepositIndex)

	// Field (7) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (8) 'RandaoMixes'
	{
		if size := len(b.RandaoMixes); size > 65536 {
			err = ssz.ErrListTooBigFn("BeaconStateDeneb.RandaoMixes", size, 65536)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.RandaoMixes {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.RandaoMixes))
		hh.MerkleizeWithMixin(subIndx, numItems, 65536)
	}

	// Field (9) 'NextWithdrawalIndex'
	hh.PutUint64(b.NextWithdrawalIndex)

	// Field (10) 'NextWithdrawalValidatorIndex'
	hh.PutUint64(b.NextWithdrawalValidatorIndex)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconStateDeneb object
func (b *BeaconStateDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

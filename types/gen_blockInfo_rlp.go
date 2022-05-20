// Code generated by rlpgen. DO NOT EDIT.

//go:build !norlpgen
// +build !norlpgen

package types

import "github.com/kardiachain/go-kardia/lib/rlp"
import "io"

func (obj *storageBlockInfo) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteUint64(obj.GasUsed)
	if obj.Rewards == nil {
		w.Write(rlp.EmptyString)
	} else {
		if obj.Rewards.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(obj.Rewards)
	}
	_tmp1 := w.List()
	for _, _tmp2 := range obj.Receipts {
		if err := _tmp2.EncodeRLP(w); err != nil {
			return err
		}
	}
	w.ListEnd(_tmp1)
	w.WriteBytes(obj.Bloom[:])
	w.ListEnd(_tmp0)
	return w.Flush()
}

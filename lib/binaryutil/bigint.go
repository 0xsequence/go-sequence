package binaryutil

import (
	"encoding/binary"
	"io"
	"math/big"
)

func WriteOptionalBigInt(w io.Writer, n *big.Int) error {
	err := binary.Write(w, binary.LittleEndian, n != nil)
	if err != nil {
		return err
	}

	if n != nil {
		err = WriteBigInt(w, n)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteBigInt(w io.Writer, n *big.Int) error {
	err := binary.Write(w, binary.LittleEndian, int8(n.Sign()))
	if err != nil {
		return err
	}

	_, err = w.Write(n.Bytes())
	if err != nil {
		return err
	}

	return nil
}

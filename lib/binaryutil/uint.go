package binaryutil

import (
	"encoding/binary"
	"io"
)

func WriteOptionalUint64(w io.Writer, n *uint64) error {
	err := binary.Write(w, binary.LittleEndian, n != nil)
	if err != nil {
		return err
	}

	if n != nil {
		err = binary.Write(w, binary.LittleEndian, *n)
		if err != nil {
			return err
		}
	}

	return nil
}

package base

import (
	"strconv"
)

type IntNull struct {
	Value int
	Null bool
}

func (c IntNull) MarshalJSON() ([]byte, error) {
	if c.Null {
		return []byte("null"), nil
	}

	return []byte(strconv.FormatInt(int64(c.Value), 10)), nil
}

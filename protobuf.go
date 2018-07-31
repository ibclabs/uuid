package uuid

import (
	"bytes"
	"errors"
)

func (u UUID) Marshal() ([]byte, error) {
	if len(u) == 0 {
		return nil, nil
	}
	return []byte(u), nil
}

func (u UUID) MarshalTo(data []byte) (n int, err error) {
	if len(u) == 0 {
		return 0, nil
	}
	copy(data, u)
	return 16, nil
}

func (u *UUID) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	id := UUID(make([]byte, 16))
	copy(id, data)
	*u = id
	return nil
}

func (u *UUID) Size() int {
	if u == nil {
		return 0
	}
	if len(*u) == 0 {
		return 0
	}
	return 16
}

func (u UUID) MarshalJSON() ([]byte, error) {
	var buf [36]byte
	encodeHex(buf[:], u)

	b := make([]byte, 0, 38)
	b = append(b, '"')
	b = append(b, buf[:]...)
	b = append(b, '"')
	return b, nil
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	if len(data) != 38 {
		return errors.New("wrong len")
	}
	*u = ParseBytes(data[1:37])
	return nil
}

func (u UUID) Equal(o UUID) bool {
	return bytes.Equal(u, o)
}

func (u UUID) Compare(o UUID) int {
	return bytes.Compare(u, o)
}

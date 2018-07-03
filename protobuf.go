package uuid

import (
	"bytes"
	"encoding/hex"
	"errors"
)

func (uuid UUID) Marshal() ([]byte, error) {
	if len(uuid) == 0 {
		return nil, nil
	}
	return []byte(uuid), nil
}

func (uuid UUID) MarshalTo(data []byte) (n int, err error) {
	if len(uuid) == 0 {
		return 0, nil
	}
	copy(data, uuid)
	return 16, nil
}

func (uuid *UUID) Unmarshal(data []byte) error {
	if len(data) == 0 {
		uuid = nil
		return nil
	}
	id := UUID(make([]byte, 16))
	copy(id, data)
	*uuid = id
	return nil
}

func (uuid *UUID) Size() int {
	if uuid == nil {
		return 0
	}
	if len(*uuid) == 0 {
		return 0
	}
	return 16
}

func (uuid UUID) MarshalJSON() ([]byte, error) {
	var buf [36]byte
	encodeHex(buf[:], uuid)

	b := make([]byte, 0, 38)
	b = append(b, '"')
	b = append(b, buf[:]...)
	b = append(b, '"')
	return b, nil
}

func (uuid *UUID) UnmarshalJSON(data []byte) error {
	if len(data) != 38 {
		errors.New("wrong len")
	}
	var buf [16]byte
	decodeHex(data[1:37], buf[:])

	*uuid =  UUID(buf[:])
	return nil
}

func (uuid UUID) Equal(other UUID) bool {
	return bytes.Equal(uuid, other)
}

func (uuid UUID) Compare(other UUID) int {
	return bytes.Compare(uuid, other)
}

func decodeHex(src, dst []byte) {
	hex.Decode(dst[:4], src[:8])
	hex.Decode(dst[4:6], src[9:13])
	hex.Decode(dst[6:8], src[14:18])
	hex.Decode(dst[8:10], src[19:23])
	hex.Decode(dst[10:], src[24:])
}
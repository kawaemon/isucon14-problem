package json

import (
	"io"

	"github.com/bytedance/sonic"
)

func NewDecoder(r io.Reader) sonic.Decoder {
	return sonic.ConfigDefault.NewDecoder(r)
}

func NewEncoder(w io.Writer) sonic.Encoder {
	return sonic.ConfigDefault.NewEncoder(w)
}

func Unmarshal(b []byte, v any) error {
	return sonic.ConfigDefault.Unmarshal(b, v)
}

func Marshal(v any) ([]byte, error) {
	return sonic.ConfigDefault.Marshal(v)
}

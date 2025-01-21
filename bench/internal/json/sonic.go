package json

import (
	"io"

	"github.com/bytedance/sonic"
)

func NewDecoder(r io.Reader) sonic.Decoder {
	return sonic.ConfigFastest.NewDecoder(r)
}
func NewEncoder(w io.Writer) sonic.Encoder {
	return sonic.ConfigFastest.NewEncoder(w)
}

func Marshal[T any](v T) ([]byte, error) {
	return sonic.ConfigFastest.Marshal(v)
}
func Unmarshal[T any](v []byte, d T) error {
	return sonic.ConfigFastest.Unmarshal(v, d)
}

// Package config provides helpers for toml based config files
package config

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"
	"go.sancus.dev/config/expand"
)

// LoadFile reads a TOML encoded files applying shell
// expansions for environment variables into a
// configuration struct
func LoadFile(filename string, v any) error {
	s, err := expand.ExpandFile(filename, nil)
	if err != nil {
		return err
	}
	_, err = toml.Decode(s, v)
	return err
}

// WriteTo writes to a given io.Writer the TOML encoded
// representation of a configuration struct
func WriteTo(f io.Writer, v any) (int64, error) {
	var buf bytes.Buffer

	// encode
	enc := toml.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return 0, err
	}

	// write
	return buf.WriteTo(f)
}

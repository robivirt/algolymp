package postyk

import (
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
)

func encodeWindows1251String(src string) (string, error) {
	return charmap.Windows1251.NewEncoder().String(src)
}

func encodeWindows1251Bytes(src []byte) ([]byte, error) {
	if len(src) == 0 || !utf8.Valid(src) {
		return src, nil
	}

	return charmap.Windows1251.NewEncoder().Bytes(src)
}

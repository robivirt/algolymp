package postyk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeWindows1251String(t *testing.T) {
	t.Parallel()

	got, err := encodeWindows1251String("привет.cpp")
	require.NoError(t, err)
	require.Equal(t, string([]byte{0xef, 0xf0, 0xe8, 0xe2, 0xe5, 0xf2, '.', 'c', 'p', 'p'}), got)
}

func TestEncodeWindows1251Bytes(t *testing.T) {
	t.Parallel()

	got, err := encodeWindows1251Bytes([]byte("Привет\n"))
	require.NoError(t, err)
	require.Equal(t, []byte{0xcf, 0xf0, 0xe8, 0xe2, 0xe5, 0xf2, '\n'}, got)
}

func TestEncodeWindows1251BytesKeepsInvalidUTF8(t *testing.T) {
	t.Parallel()

	src := []byte{0xff, 0xfe, 0xfd}

	got, err := encodeWindows1251Bytes(src)
	require.NoError(t, err)
	require.Equal(t, src, got)
}

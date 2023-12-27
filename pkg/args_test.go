package godcmtk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testArg struct {
	Text     string  `arg:"-s"`
	Number   int     `arg:"-n"`
	Floating float64 `arg:"-f"`
	Flag     bool    `arg:"-b"`
}

func TestMarshalArgs(t *testing.T) {
	m := testArg{
		Text:     "str",
		Number:   3,
		Floating: 0.3,
		Flag:     true,
	}

	results, err := MarshalArgs(m)
	assert.Equal(t, []string{"-s", "str", "-n", "3", "-f", "0.3", "-b"}, results)
	assert.NoError(t, err)

	results, err = MarshalArgs(&m)
	assert.Equal(t, []string{"-s", "str", "-n", "3", "-f", "0.3", "-b"}, results)
	assert.NoError(t, err)
}

func TestMarshalEmptyArgs(t *testing.T) {
	m := testArg{
		Text:     "",
		Number:   0,
		Floating: 0.0,
		Flag:     false,
	}

	results, err := MarshalArgs(m)
	assert.Equal(t, []string{}, results)
	assert.NoError(t, err)

	results, err = MarshalArgs(&m)
	assert.Equal(t, []string{}, results)
	assert.NoError(t, err)
}

func TestMarshalNilArgs(t *testing.T) {
	results, err := MarshalArgs(nil)
	assert.Equal(t, []string{}, results)
	assert.NoError(t, err)
}

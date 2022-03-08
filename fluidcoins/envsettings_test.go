package fluidcoins

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetting_SetPath(t *testing.T) {
	var se SetupEnv

	path := "./"
	es := SetPath(path)
	es(&se)

	require.Equal(t, path, se.Path)
}

func TestSetting_SetFileName(t *testing.T) {
	var se SetupEnv

	fn := "credentials"
	es := SetFileName(fn)
	es(&se)

	require.Equal(t, fn, se.Filename)
}

func TestSetting_SetExtension(t *testing.T) {
	var se SetupEnv

	ex := "env"
	es := SetExtension(ex)
	es(&se)

	require.Equal(t, ex, se.Extension)
}

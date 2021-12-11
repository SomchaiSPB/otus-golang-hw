package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("copy whole file", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 0, 0)

		require.NoError(t, err)

		require.FileExistsf(t, "testdata/output.txt", "")

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		expected, err := os.Open("testdata/input.txt")
		require.NoError(t, err)

		go func() {
			defer actual.Close()
			defer expected.Close()
		}()

		actualByte, _ := io.ReadAll(actual)
		expectedByte, _ := io.ReadAll(expected)

		require.Equal(t, expectedByte, actualByte)

		_ = os.Remove("testdata/output.txt")
	})

	t.Run("copy limit 100", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 0, 100)

		require.NoError(t, err)

		require.FileExistsf(t, "testdata/output.txt", "")

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		go func() {
			defer actual.Close()
		}()

		actualByte, _ := io.ReadAll(actual)

		require.Equal(t, 100, len(actualByte))

		_ = os.Remove("testdata/output.txt")
	})

	t.Run("copy negative offset", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", -10, 0)

		require.Error(t, err)
	})

	t.Run("copy limit bigger than file", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 0, 9999999)

		require.NoError(t, err)

		require.FileExistsf(t, "testdata/output.txt", "")

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		go func() {
			defer actual.Close()
		}()

		actualByte, _ := io.ReadAll(actual)

		require.Equal(t, 6617, len(actualByte))
	})

	t.Run("copy offset 100 limit 100", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 100, 100)

		require.NoError(t, err)

		require.FileExistsf(t, "testdata/output.txt", "")

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		go func() {
			defer actual.Close()
		}()

		actualByte, _ := io.ReadAll(actual)

		require.Equal(t, 100, len(actualByte))
	})
}

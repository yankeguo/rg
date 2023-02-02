package xx

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGuard(t *testing.T) {
	var err error
	func() {
		defer Guard(&err)
		panic(errors.New("hello"))
	}()
	require.Error(t, err)
	require.Equal(t, "hello", err.Error())
}

func TestGuardNotError(t *testing.T) {
	var err error
	func() {
		defer Guard(&err)
		panic("hello")
	}()
	require.Error(t, err)
	require.Equal(t, "panic: hello", err.Error())
}

func TestMust0(t *testing.T) {
	var err error
	func() {
		defer Guard(&err)
		Must0(errors.New("hello"))
	}()
	require.Error(t, err)
	require.Equal(t, "hello", err.Error())
}

func TestMust(t *testing.T) {
	{
		var err error
		var (
			v1 int
		)
		func() {
			defer Guard(&err)
			v1 = Must(3, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
	}
	{
		var err error
		var (
			v1 int
		)
		func() {
			defer Guard(&err)
			v1 = Must(3, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 3, v1)
	}
}

func TestMust2(t *testing.T) {
	{
		var err error
		var (
			v1 int
			v2 int
		)
		func() {
			defer Guard(&err)
			v1, v2 = Must2(1, 2, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
		require.Equal(t, 0, v2)
	}
	{
		var err error
		var (
			v1 int
			v2 int
		)
		func() {
			defer Guard(&err)
			v1, v2 = Must2(1, 2, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, 2, v2)
	}
}

func TestMust3(t *testing.T) {
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3 = Must3(1, 2, 3, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
		require.Equal(t, 0, v2)
		require.Equal(t, 0, v3)
	}
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3 = Must3(1, 2, 3, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, 2, v2)
		require.Equal(t, 3, v3)
	}
}

func TestMust4(t *testing.T) {
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4 = Must4(1, 2, 3, 4, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
		require.Equal(t, 0, v2)
		require.Equal(t, 0, v3)
		require.Equal(t, 0, v4)
	}
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4 = Must4(1, 2, 3, 4, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, 2, v2)
		require.Equal(t, 3, v3)
		require.Equal(t, 4, v4)
	}
}

func TestMust5(t *testing.T) {
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
			v5 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4, v5 = Must5(1, 2, 3, 4, 5, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
		require.Equal(t, 0, v2)
		require.Equal(t, 0, v3)
		require.Equal(t, 0, v4)
		require.Equal(t, 0, v5)
	}
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
			v5 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4, v5 = Must5(1, 2, 3, 4, 5, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, 2, v2)
		require.Equal(t, 3, v3)
		require.Equal(t, 4, v4)
		require.Equal(t, 5, v5)
	}
}

func TestMust6(t *testing.T) {
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
			v5 int
			v6 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4, v5, v6 = Must6(1, 2, 3, 4, 5, 6, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
		require.Equal(t, 0, v2)
		require.Equal(t, 0, v3)
		require.Equal(t, 0, v4)
		require.Equal(t, 0, v5)
		require.Equal(t, 0, v6)
	}
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
			v5 int
			v6 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4, v5, v6 = Must6(1, 2, 3, 4, 5, 6, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, 2, v2)
		require.Equal(t, 3, v3)
		require.Equal(t, 4, v4)
		require.Equal(t, 5, v5)
		require.Equal(t, 6, v6)
	}
}

func TestMust7(t *testing.T) {
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
			v5 int
			v6 int
			v7 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4, v5, v6, v7 = Must7(1, 2, 3, 4, 5, 6, 7, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
		require.Equal(t, 0, v2)
		require.Equal(t, 0, v3)
		require.Equal(t, 0, v4)
		require.Equal(t, 0, v5)
		require.Equal(t, 0, v6)
		require.Equal(t, 0, v7)
	}
	{
		var err error
		var (
			v1 int
			v2 int
			v3 int
			v4 int
			v5 int
			v6 int
			v7 int
		)
		func() {
			defer Guard(&err)
			v1, v2, v3, v4, v5, v6, v7 = Must7(1, 2, 3, 4, 5, 6, 7, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, 2, v2)
		require.Equal(t, 3, v3)
		require.Equal(t, 4, v4)
		require.Equal(t, 5, v5)
		require.Equal(t, 6, v6)
		require.Equal(t, 7, v7)
	}
}

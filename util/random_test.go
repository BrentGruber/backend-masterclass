package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	rand1 := RandomInt(0, 1000)

	require.NotEmpty(t, rand1)
	require.GreaterOrEqual(t, rand1, int64(0))
	require.LessOrEqual(t, rand1, int64(1000))

	rand2 := RandomInt(500, 100000)

	require.NotEmpty(t, rand2)
	require.GreaterOrEqual(t, rand2, int64(500))
	require.LessOrEqual(t, rand2, int64(100000))

	require.NotEqual(t, rand1, rand2)
}

func TestRandomString(t *testing.T) {
	rand1 := RandomString(10)

	require.NotEmpty(t, rand1)
	require.Len(t, rand1, 10)

	rand2 := RandomString(10)

	require.NotEmpty(t, rand2)
	require.Len(t, rand2, 10)

	require.NotEqual(t, rand1, rand2)
}

func TestRandomOwner(t *testing.T) {
	rand1 := RandomOwner()

	require.NotEmpty(t, rand1)
	require.Len(t, rand1, 6)

	rand2 := RandomOwner()

	require.NotEmpty(t, rand2)
	require.Len(t, rand2, 6)

	require.NotEqual(t, rand1, rand2)
}

func TestRandomMoney(t *testing.T) {
	rand1 := RandomMoney()

	require.NotEmpty(t, rand1)
	require.GreaterOrEqual(t, rand1, int64(0))
	require.LessOrEqual(t, rand1, int64(1000))

	rand2 := RandomMoney()

	require.GreaterOrEqual(t, rand2, int64(0))
	require.LessOrEqual(t, rand2, int64(1000))

	require.NotEqual(t, rand1, rand2)
}

func TestRandomCurrency(t *testing.T) {
	rand1 := RandomCurrency()

	require.NotEmpty(t, rand1)
	require.Contains(t, []string{"EUR", "USD", "CAD"}, rand1)
}

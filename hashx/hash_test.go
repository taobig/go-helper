package hashx

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSha256Sum(t *testing.T) {
	t.Parallel()

	hash := Sha256Sum([]byte("hello"))
	hexHash := hex.EncodeToString(hash)
	require.Equal(t, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", hexHash)
}

func TestSha256Hex(t *testing.T) {
	t.Parallel()

	hexHash := Sha256Hex([]byte("hello"))
	require.Equal(t, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", hexHash)
}

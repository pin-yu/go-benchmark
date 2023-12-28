package gobenchmark

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	require.Equal(t, uintptr(1), unsafe.Sizeof(T1{}))
	require.Equal(t, uintptr(2), unsafe.Sizeof(T2{}))
	require.Equal(t, uintptr(3), unsafe.Sizeof(T3{}))
	require.Equal(t, uintptr(16), unsafe.Sizeof(T4{}))
	require.Equal(t, uintptr(24), unsafe.Sizeof(T5{}))
}

package file_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileArray(t *testing.T) {
	arr:=NewFileArray("a.txt", 8)
	arr.WriteAt(42, 7)
	arr.WriteAt(43,6)
	val:=arr.ReadAt(7)
	assert.Equal(t, val, int64(42))
	val=arr.ReadAt(6)
	assert.Equal(t, val, int64(43))

	arr.Remove()
}

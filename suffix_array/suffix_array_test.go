package suffix_array

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	f, err:=os.Create("t.txt")
	assert.NoError(t, err)
	_, err=f.WriteString("aaba")
	assert.NoError(t, err)
	s, err:=NewSuffixArray(f)
	assert.NoError(t, err)
	actual:=make([]int64, 4)
	for i:=int64(1); i < 5; i++ {
		actual[i-1] = s.p.ReadAt(i)
	}
	expected:=[]int64{3,0,1,2}
	assert.Equal(t, expected, actual)


	s.Terminate()
}

func TestCreate1(t *testing.T) {
	f, err:=os.Create("t.txt")
	assert.NoError(t, err)
	_, err=f.WriteString("abaab")
	assert.NoError(t, err)
	s, err:=NewSuffixArray(f)
	assert.NoError(t, err)
	actual:=make([]int64, 5)
	for i:=int64(1); i < 6; i++ {
		actual[i-1] = s.p.ReadAt(i)
	}
	expected:=[]int64{2, 3, 0, 4, 1}
	assert.Equal(t, expected, actual)
}

func TestFind(t *testing.T) {
	f, err:=os.Create("t.txt")
	assert.NoError(t, err)
	_, err=f.WriteString("abracadabra")
	assert.NoError(t, err)
	s, err:=NewSuffixArray(f)
	assert.NoError(t, err)

	pos:= s.FindSubstring("abr")
	assert.Equal(t, pos, int64(0))

	pos= s.FindSubstring("racad")
	assert.Equal(t, pos, int64(2))

	pos= s.FindSubstring("dabra")
	assert.Equal(t, pos, int64(6))


	pos= s.FindSubstring("dabraa")
	assert.Equal(t, pos, int64(-1))

	pos= s.FindSubstring("x")
	assert.Equal(t, pos, int64(-1))

	pos= s.FindSubstring("a")
	assert.Equal(t, pos, int64(5))

	pos= s.FindSubstring("bra")
	assert.Equal(t, pos, int64(1))

	pos= s.FindSubstring("abracadabra")
	assert.Equal(t, pos, int64(0))

	s.Terminate()
}
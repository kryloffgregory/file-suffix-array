package file_array

import (
	"encoding/binary"
	"os"
)

type Int64Array struct {
	file *os.File
	Size int64
}

func NewFileArray(fileName string, size int64) *Int64Array {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	for i := int64(0); i < size; i++ {
		_, err := file.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0})
		if err != nil {
			panic(err)
		}
	}

	return &Int64Array{
		file: file,
		Size: size,
	}
}

func (a *Int64Array) ReadAt(offset int64) int64 {
	buf := make([]byte, 8)
	_, err := a.file.ReadAt(buf, offset*8)
	if err != nil {
		panic(err)
	}

	return int64(binary.BigEndian.Uint64(buf))
}

func (a *Int64Array) WriteAt(val int64, offset int64) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(val))
	a.file.WriteAt(buf, offset*8)
}

func (a *Int64Array) Inc(offset int64) {
	a.WriteAt(a.ReadAt(offset)+1, offset)
}

func (a *Int64Array) Dec(offset int64) {
	a.WriteAt(a.ReadAt(offset)-1, offset)
}

func (a *Int64Array) Remove() {
	err := os.Remove(a.file.Name())
	if err != nil {
		panic(err)
	}
}

func (a *Int64Array) Clean() {
	for i := int64(0); i < a.Size; i++ {
		_, err := a.file.WriteAt([]byte{0, 0, 0, 0, 0, 0, 0, 0}, i*8)
		if err != nil {
			panic(err)
		}
	}
}

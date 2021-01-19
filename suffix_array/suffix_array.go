package suffix_array

import (
	"os"

	"fileSuffixArray/file_array"
)

type SuffixArray struct {
	p    *file_array.Int64Array
	file *os.File
}

func NewSuffixArray(file *os.File) (*SuffixArray, error) {
	file.Seek(0, 0)
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stat.Size()
	n := size + 1

	p := file_array.NewFileArray("p", n)
	c := file_array.NewFileArray("с", n)
	cnt := file_array.NewFileArray("cnt", 256)

	bte := make([]byte, 1)
	for i := int64(0); i < n; i++ {
		if i == size {
			bte = []byte{0}
		} else {
			_, err := file.Read(bte)
			if err != nil {
				return nil, err
			}
		}
		cnt.Inc(int64(bte[0]))
	}

	for i := int64(1); i < 256; i++ {
		cnt.WriteAt(cnt.ReadAt(i)+cnt.ReadAt(i-1), i)
	}

	file.Seek(0, 0)
	for i := int64(0); i < n; i++ {
		if i == size {
			bte = []byte{0}
		} else {
			_, err := file.Read(bte)
			if err != nil {
				return nil, err
			}
		}
		cnt.Dec(int64(bte[0]))
		p.WriteAt(i, cnt.ReadAt(int64(bte[0])))
	}

	c.WriteAt(0, p.ReadAt(0))
	classes := int64(1)
	for i := int64(1); i < n; i++ {
		idx1 := p.ReadAt(i)
		idx2 := p.ReadAt(i - 1)

		if idx1 == size || idx2 == size {
			classes++
		} else {
			file.ReadAt(bte, idx1)
			bte2 := make([]byte, 1)
			file.ReadAt(bte2, idx2)
			if bte[0] != bte2[0] {
				classes++
			}
		}

		c.WriteAt(classes-1, p.ReadAt(i))
	}

	pn := file_array.NewFileArray("pn", n)
	cn := file_array.NewFileArray("cn", n)

	for h := int64(0); (1 << h) < n; h++ {
		for i := int64(0); i < n; i++ {
			tmp := p.ReadAt(i) - (1 << h)
			if tmp < 0 {
				tmp += n
			}
			pn.WriteAt(tmp, i)
		}
		cnt.Clean()

		for i := int64(0); i < n; i++ {
			cnt.Inc(c.ReadAt(pn.ReadAt(i)))
		}

		for i := int64(1); i < n; i++ {
			cnt.WriteAt(cnt.ReadAt(i)+cnt.ReadAt(i-1), i)
		}
		for i := n - 1; i >= 0; i-- {
			pos := c.ReadAt(pn.ReadAt(i))
			cnt.Dec(pos)
			p.WriteAt(pn.ReadAt(i), cnt.ReadAt(pos))
		}

		cn.WriteAt(int64(0), p.ReadAt(0))
		classes := int64(1)
		for i := int64(1); i < n; i++ {
			mid1 := (p.ReadAt(i) + (1 << h)) % n
			mid2 := (p.ReadAt(i-1) + (1 << h)) % n
			if c.ReadAt(p.ReadAt(i)) != c.ReadAt(p.ReadAt(i-1)) || c.ReadAt(mid1) != c.ReadAt(mid2) {
				classes++
			}
			cn.WriteAt(classes-1, p.ReadAt(i))
		}

		for i := int64(0); i < n; i++ {
			c.WriteAt(cn.ReadAt(i), i)
		}
	}

	pn.Remove()
	cnt.Remove()
	cn.Remove()
	c.Remove()

	return &SuffixArray{p, file}, nil
}

func (sa *SuffixArray) FindSubstring(s string) (pos int64) {
	l := int64(1)
	r := sa.p.Size
	for r > l+1 {
		mid := (l + r) / 2
		cmpResult := sa.cmp(sa.p.ReadAt(mid), s)
		if cmpResult < 1 {
			l = mid
		} else {
			r = mid
		}
	}

	if sa.cmp(sa.p.ReadAt(l), s) == 0 {
		return sa.p.ReadAt(l)
	}

	return -1
}

func (sa *SuffixArray) cmp(pos int64, s string) int {
	for i := 0; i < len(s); i++ {
		if pos+int64(i) >= sa.p.Size-1 {
			return -1
		}

		bte := make([]byte, 1)
		_, err := sa.file.ReadAt(bte, pos+int64(i))
		if err != nil {
			panic(err)
		}

		if bte[0] < s[i] {
			return -1
		}

		if bte[0] > s[i] {
			return 1
		}
	}

	return 0
}

func (sa *SuffixArray) Terminate() {
	sa.p.Remove()
}

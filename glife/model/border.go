package model

type Fetcher func(x uint64, y uint64) bool

type Setter func(x uint64, y uint64, value bool)

type BorderStrategy interface {
	get(x uint64, y uint64, matrix BitMatrix, fetcher Fetcher) bool
	set(x uint64, y uint64, value bool, matrix BitMatrix, setter Setter)
	translateX(x uint64, width uint64) uint64
	translateY(y uint64, height uint64) uint64
}

type ToroidalBorderStrategy struct {
	BorderStrategy
}

func (s *ToroidalBorderStrategy) translateX(x uint64, width uint64) uint64 {
	return shift(x, width)
}

func (s *ToroidalBorderStrategy) translateY(y uint64, height uint64) uint64 {
	return shift(y, height)
}
func (s *ToroidalBorderStrategy) get(x uint64, y uint64, space BitMatrix, fetcher Fetcher) bool {
	return fetcher(s.translateX(x, space.Width()), s.translateY(y, space.Height()))
}

func (s *ToroidalBorderStrategy) set(x uint64, y uint64, value bool, space BitMatrix, setter Setter) {
	setter(s.translateX(x, space.Width()), s.translateY(y, space.Height()), value)
}

func shift(coord uint64, size uint64) uint64 {
	xlatn := coord
	for xlatn < 0 || xlatn >= size {
		switch {
		case xlatn < 0:
			xlatn += size
		case xlatn >= size:
			xlatn -= size
		default:
			panic("This should be impossible")
		}
	}
	return xlatn
}

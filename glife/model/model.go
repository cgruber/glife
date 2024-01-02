package model

type Field struct {
}

type Neighborhood struct {
}

type BitMatrix interface {
	Width() uint64
	Height() uint64
	Strategy() BorderStrategy
	Get(x uint64, y uint64) bool
	GetNeighborhood(x uint64, y uint64) Neighborhood
	Copy() BitMatrix
}

type MutableBitMatrix interface {
	BitMatrix
	Set(x uint64, y uint64, value bool)
	Toggle(x uint64, y uint64)
}

type ArrayBitMatrix struct {
	MutableBitMatrix
	width  uint64
	height uint64
	data   [][]bool
}

func (a *ArrayBitMatrix) Width() uint64 {
	return a.width
}

func (a *ArrayBitMatrix) Height() uint64 {
	return a.height
}

func (a *ArrayBitMatrix) strategy() BorderStrategy {
	return &ToroidalBorderStrategy{}
}

// NewArrayBitMatrixFromExisting Takes an existing BitMatrix and creates an ArrayBitMatrix from its contents.
func NewArrayBitMatrixFromExisting(matrix BitMatrix) MutableBitMatrix {
	data := make([][]bool, matrix.Width())
	for x := uint64(0); x < matrix.Width(); x++ {
		data[x] = make([]bool, matrix.Height())
		for y := uint64(0); y < matrix.Height(); y++ {
			data[x][y] = matrix.Get(x, y)
		}
	}
	return &ArrayBitMatrix{width: matrix.Width(), height: matrix.Height(), data: data}
}

// NewArrayBitMatrix Takes an existing BitMatrix and creates an ArrayBitMatrix from its contents.
func NewArrayBitMatrix(width uint64, height uint64) MutableBitMatrix {
	data := make([][]bool, width)
	for x := uint64(0); x < width; x++ {
		data[x] = make([]bool, height)
		for y := uint64(0); y < height; y++ {
			data[x][y] = false
		}
	}
	return &ArrayBitMatrix{width: width, height: height, data: data}
}

//override val strategy: BorderStrategy = ToroidalBorderStrategy
//) : MutableBitMatrix {
//constructor(matrix: BitMatrix) : this(matrix.width, matrix.height, matrix.strategy) {
//for (x in 0 until width) {
//for (y in 0 until height) {
//data[x][y] = matrix.get(x, y)
//}
//}
//}
//internal val data = Array(width) { BooleanArray(height) { DEAD } }
//override fun get(x: Int, y: Int): Bit = strategy.get(x, y, this) { x, y -> data[x][y] }
//override fun set(x: Int, y: Int, value: Bit) {
//strategy.set(x, y, value, this) { x, y, value -> data[x][y] = value }
//}
//
//override fun toggle(x: Int, y: Int) = set(x, y, ! get(x, y))
//override fun getNeighborhood(x: Int, y: Int): Neighborhood {
//return Neighborhood(
//get(x - 1, y - 1), get(x, y - 1), get(x + 1, y - 1),
//get(x - 1, y), get(x, y), get(x + 1, y),
//get(x - 1, y + 1), get(x, y + 1), get(x + 1, y + 1),
//)
//}
//
//override fun equals(other: Any?): Boolean {
//return when (other) {
//is ArrayBitMatrix -> data.contentDeepEquals(other.data)
//is BitMatrix -> {
//for (x in 0 until width) {
//for (y in 0 until height) {
//if (get(x, y) != other.get(x, y)) return false
//}
//}
//return true
//}
//else -> false
//}
//}
//
//override fun hashCode() = 31 * data.contentDeepHashCode()
//
///** Perform a deep value copy, rather than the default shallow copy from a data class */
//override fun copy() = ArrayBitMatrix(width, height).also { dupe ->
//for (x in 0 until width) {
//for (y in 0 until height) {
//dupe.set(x, y, this.get(x, y))
//}
//}
//}
//}

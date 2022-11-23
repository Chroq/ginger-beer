package testdata

type TestModel struct {
	Bool       bool
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Uintptr    uintptr
	Uint       uint
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Pointer    *int
	Int        int
	Byte       byte
	Rune       rune
	String     string
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	Interface  interface{}
	Chan       chan int
	ArrayInt   [1]int
	SliceInt   []int
	MapInt     map[int]int
	Struct     struct{}
	Func       func()
}

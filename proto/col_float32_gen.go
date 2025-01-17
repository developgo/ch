// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColFloat32 represents Float32 column.
type ColFloat32 []float32

// Compile-time assertions for ColFloat32.
var (
	_ ColInput  = ColFloat32{}
	_ ColResult = (*ColFloat32)(nil)
	_ Column    = (*ColFloat32)(nil)
)

// Type returns ColumnType of Float32.
func (ColFloat32) Type() ColumnType {
	return ColumnTypeFloat32
}

// Rows returns count of rows in column.
func (c ColFloat32) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColFloat32) Row(i int) float32 {
	return c[i]
}

// Append float32 to column.
func (c *ColFloat32) Append(v float32) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColFloat32) Reset() {
	*c = (*c)[:0]
}

// NewArrFloat32 returns new Array(Float32).
func NewArrFloat32() *ColArr {
	return &ColArr{
		Data: new(ColFloat32),
	}
}

// AppendFloat32 appends slice of float32 to Array(Float32).
func (c *ColArr) AppendFloat32(data []float32) {
	d := c.Data.(*ColFloat32)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

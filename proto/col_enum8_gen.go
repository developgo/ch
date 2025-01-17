// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColEnum8 represents Enum8 column.
type ColEnum8 []Enum8

// Compile-time assertions for ColEnum8.
var (
	_ ColInput  = ColEnum8{}
	_ ColResult = (*ColEnum8)(nil)
	_ Column    = (*ColEnum8)(nil)
)

// Type returns ColumnType of Enum8.
func (ColEnum8) Type() ColumnType {
	return ColumnTypeEnum8
}

// Rows returns count of rows in column.
func (c ColEnum8) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColEnum8) Row(i int) Enum8 {
	return c[i]
}

// Append Enum8 to column.
func (c *ColEnum8) Append(v Enum8) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColEnum8) Reset() {
	*c = (*c)[:0]
}

// NewArrEnum8 returns new Array(Enum8).
func NewArrEnum8() *ColArr {
	return &ColArr{
		Data: new(ColEnum8),
	}
}

// AppendEnum8 appends slice of Enum8 to Array(Enum8).
func (c *ColArr) AppendEnum8(data []Enum8) {
	d := c.Data.(*ColEnum8)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

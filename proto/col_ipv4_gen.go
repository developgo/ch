// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColIPv4 represents IPv4 column.
type ColIPv4 []IPv4

// Compile-time assertions for ColIPv4.
var (
	_ ColInput  = ColIPv4{}
	_ ColResult = (*ColIPv4)(nil)
	_ Column    = (*ColIPv4)(nil)
)

// Type returns ColumnType of IPv4.
func (ColIPv4) Type() ColumnType {
	return ColumnTypeIPv4
}

// Rows returns count of rows in column.
func (c ColIPv4) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColIPv4) Row(i int) IPv4 {
	return c[i]
}

// Append IPv4 to column.
func (c *ColIPv4) Append(v IPv4) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColIPv4) Reset() {
	*c = (*c)[:0]
}

// NewArrIPv4 returns new Array(IPv4).
func NewArrIPv4() *ColArr {
	return &ColArr{
		Data: new(ColIPv4),
	}
}

// AppendIPv4 appends slice of IPv4 to Array(IPv4).
func (c *ColArr) AppendIPv4(data []IPv4) {
	d := c.Data.(*ColIPv4)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

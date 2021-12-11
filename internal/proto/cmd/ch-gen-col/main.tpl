{{- /*gotype: github.com/go-faster/ch/internal/proto/cmd/ch-gen-col.Variant*/ -}}
// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
{{- if .Float }}
   "math"
{{- end }}
  "github.com/go-faster/errors"
)

// {{ .Type }} represents {{ .Name }} column.
type {{ .Type }} []{{ .ElemType }}

// Compile-time assertions for {{ .Type }}.
var (
  _ Input  = {{ .Type }}{}
  _ Result = (*{{ .Type }})(nil)
)

// Type returns ColumnType of {{ .Name }}.
func ({{ .Type }}) Type() ColumnType {
  return {{ .ColumnType }}
}

// Rows returns count of rows in column.
func (c {{ .Type }}) Rows() int {
  return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *{{ .Type }}) Reset() {
  *c = (*c)[:0]
}

// EncodeColumn encodes {{ .Name }} rows to *Buffer.
func (c {{ .Type }}) EncodeColumn(b *Buffer) {
  {{- if .Byte }}
  b.Buf = append(b.Buf, c...)
  {{- else if .SingleByte }}
  start := len(b.Buf)
  b.Buf = append(b.Buf, make([]byte, len(c))...)
  for i := range c {
    b.Buf[i + start] = {{ .UnsignedType }}(c[i])
  }
  {{- else }}
  const size = {{ .Bits }} / 8
  offset := len(b.Buf)
  b.Buf = append(b.Buf, make([]byte, size * len(c))...)
  for _, v := range c {
    {{ .BinPut }}(
      b.Buf[offset:offset+size],
    {{- if .Float }}
      math.{{ .Name }}bits(v),
    {{- else if .Signed }}
      {{ .UnsignedType }}(v),
    {{- else }}
      v,
    {{- end }}
    )
    offset += size
  }
  {{- end }}
}

// DecodeColumn decodes {{ .Name }} rows from *Reader.
func (c *{{ .Type }}) DecodeColumn(r *Reader, rows int) error {
  {{- if .SingleByte }}
  data, err := r.ReadRaw(rows)
  {{- else }}
  const size = {{ .Bits }} / 8
  data, err := r.ReadRaw(rows * size)
  {{- end }}
  if err != nil {
    return errors.Wrap(err, "read")
  }
  {{- if .Byte }}
  *c = append(*c, data...)
  {{- else if .SingleByte }}
  v := *c
  v = append(v, make([]{{ .ElemType }}, rows)...)
  for i := range data {
    v[i] = {{ .ElemType }}(data[i])
  }
  *c = v
  {{- else }}
  v := *c
  for i := 0; i < len(data); i += size {
    v = append(v,
    {{- if .Float }}
      math.{{ .Name }}frombits(bin.{{ .BinFunc }}(data[i:i+size])),
    {{- else if .Signed }}
     {{ .ElemType }}({{ .BinGet }}(data[i:i+size])),
    {{- else }}
      {{ .BinGet }}(data[i:i+size]),
    {{- end }}
    )
  }
  *c = v
  {{- end }}
  return nil
}

package builder

import (
	"bytes"
	//"fmt"
	"io"
)

type Writer interface {
	io.Writer
	Append(...interface{})
}

var _ Writer = NewWriter()

type BytesWriter struct {
	writer *bytes.Buffer
	buffer []byte
	args   []interface{}
}

func NewWriter() *BytesWriter {
	w := &BytesWriter{}
	w.writer = bytes.NewBuffer(w.buffer)
	return w
}

func (s *BytesWriter) Write(buf []byte) (int, error) {
	return s.writer.Write(buf)
}

func (s *BytesWriter) Append(args ...interface{}) {
	s.args = append(s.args, args...)
}

type Cond interface {
	WriteTo(Writer) error
	And(...Cond) Cond
	Or(...Cond) Cond
	IsValid() bool
}

type condEmpty struct{}

var _ Cond = condEmpty{}

func NewCond() Cond {
	return condEmpty{}
}

func (condEmpty) WriteTo(w Writer) error {
	return nil
}

func (condEmpty) And(conds ...Cond) Cond {
	return And(conds...)
}

func (condEmpty) Or(conds ...Cond) Cond {
	return Or(conds...)
}

func (condEmpty) IsValid() bool {
	return false
}

func condToSQL(cond Cond) (string, []interface{}, error) {
	if cond == nil || !cond.IsValid() {
		return "", nil, nil
	}

	w := NewWriter()
	if err := cond.WriteTo(w); err != nil {
		return "", nil, err
	}

	return w.writer.String(), w.args, nil
}

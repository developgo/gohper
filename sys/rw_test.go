package sys

import (
	"bytes"
	"testing"

	"github.com/cosiner/golib/test"
)

func TestBufReader(t *testing.T) {

}

func TestBufWriter(t *testing.T) {

}

func TestWrapBufWriter(t *testing.T) {

}

func TestWrapWriter(t *testing.T) {

}

func TestWriteVString(t *testing.T) {

}

func TestWriteLString(t *testing.T) {

}

func TestWriteV(t *testing.T) {

}

func TestWriteL(t *testing.T) {
	tt := test.WrapTest(t)
	buffer := bytes.NewBuffer(make([]byte, 0))
	bw := NewBufVWriter(buffer)
	bw.WriteLString("abc", "de")
	bw.Flush()
	tt.AssertTrue("WriteL", bytes.Equal([]byte("abcde"), buffer.Bytes()))
}

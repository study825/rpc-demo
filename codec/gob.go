package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (g GobCodec) Close() error {
	//TODO implement me
	panic("implement me")
}

func (g GobCodec) ReadHeader(header *Header) error {
	//TODO implement me
	panic("implement me")
}

func (g GobCodec) ReadBody(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (g GobCodec) Write(header *Header, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

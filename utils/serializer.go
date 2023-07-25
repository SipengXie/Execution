package utils

import (
	"io"
)

type Decoder interface {
	Decode(val interface{}) error // read
}

type Encoder interface {
	Encode(val interface{}) error // write
}

type Serializer interface {
	GetEncoder(writer io.Writer) Encoder                    // write to writer
	GetDecoder(reader io.Reader, inputLimit uint64) Decoder // read from reader
}

package data

// Stream is the generic representation of the functional signature
// that is defined by the io.Reader and io.Writer interface.
// It abstracts communication of any kind as a stream of bytes,
// which is either populated with information or copied from.
// In accordance with the io package, the implementation should return the number of buffer bytes
// that have been interacted with.
// On completion the error io.EOF must be returned.
type Stream func(buf []byte) (int, error)

// Write implements the io.Writer interface on the given stream.
// A call directly evokes the underlying function.
// By definition a generic stream does not differentiate itself as source (rx / io.Reader)
// or destination (tx / io.Writer).
func (s Stream) Write(buf []byte) (int, error) {
	return s(buf)
}

// Read implements the io.Reader interface on the given stream.
// A call directly evokes the underlying function.
// By definition a generic stream does not differentiate itself as source (rx / io.Reader) or
// destination (tx / io.Writer).
func (s Stream) Read(buf []byte) (int, error) {
	return s(buf)
}

// Codec defines a stream consumer.
// It may encode data onto the stream (tx / io.Writer) or
// decode from it (rx / io.Writer).
type Codec func(s Stream) error

// Handler specifies the bidirectional process of encoding and decoding data.
type Handler func(rx, tx Codec) error

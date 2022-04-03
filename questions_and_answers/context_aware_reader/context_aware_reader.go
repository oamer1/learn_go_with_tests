package contextawarereader

import "io"

func NewCancellableReader(rdr io.Reader) io.Reader {
	return rdr
}

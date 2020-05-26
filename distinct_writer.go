package drw

import (
	"bytes"
	"io"
)

// Writer that supposed to wrap original writer and write distinct values.
type Writer struct {
	*Opts
	OriginalWriter io.Writer
}

// NewWriter represents a distinct writer initializer.
func NewWriter(w io.Writer, delim byte, cache Cache) io.Writer {
	return &Writer{
		Opts: &Opts{
			Cache: cache,
		},
		OriginalWriter: w,
	}
}

// Writer implements io.Writer.
func (w *Writer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
	var dbuf bytes.Buffer
	for {
		line, err := buf.ReadBytes(w.Opts.Delimiter)
		// log.Println(string(line))
		if err == io.EOF {
			exists, err := w.Opts.Cache.Set(line)
			if err != nil {
				return 0, err
			}

			if !exists {
				dbuf.Write(line)
			}
			break
		}

		exists, err := w.Opts.Cache.Set(line)
		if err != nil {
			return 0, err
		}

		if exists {
			continue
		}
		dbuf.Write(line)
	}

	return w.OriginalWriter.Write(dbuf.Bytes())
}

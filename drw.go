package drw

// Opts of the distinct reader or writer.
type Opts struct {
	Cache     Cache
	Delimiter byte
}

// Cache that makes values distinct (more or less).
type Cache interface {
	Set([]byte) (exists bool, err error)
}

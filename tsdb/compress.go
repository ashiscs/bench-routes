package tsdb

import (
	"sync"
)

// Compressor implements the compression logic for tsdb files
type Compressor struct {
	dir string
	mtx sync.RWMutex
}

// OpenCompressor opens the compressor in the given directory
func OpenCompressor(dir string) *Compressor {
	return &Compressor{
		dir: dir,
	}
}

// Dir returns the directory of comporessor
func (cmp *Compressor) Dir() string {
	return cmp.dir
}

// Compress compresses the tsdb samples present in the given directory
func (cmp *Compressor) Compress() (bool, error) {
	return true, nil
}

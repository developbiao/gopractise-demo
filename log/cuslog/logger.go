package cuslog

import "sync"

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

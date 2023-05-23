package parallel

import (
	"testing"
	"time"
)

func TestCorge1(t *testing.T) {
	t.Parallel()
	time.Sleep(300 * time.Millisecond)
}

func TestCorge2(t *testing.T) {
	t.Parallel()
	time.Sleep(300 * time.Millisecond)
}
func TestCorge3(t *testing.T) {
	t.Parallel()
	time.Sleep(300 * time.Millisecond)
}

//go:build appengine || (!linux && !darwin && !freebsd && !openbsd && !netbsd)
// +build appengine !linux,!darwin,!freebsd,!openbsd,!netbsd

package walker

import (
	"os"
	"syscall"
)

func (w *walker) readdir(dirname string) error {
	f, err := os.OpenFile(dirname, syscall.O_DIRECT, 0666)
	if err != nil {
		return err
	}

	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return err
	}

	for _, fi := range list {
		if err = w.walk(dirname, fi); err != nil {
			return err
		}
	}
	return nil
}

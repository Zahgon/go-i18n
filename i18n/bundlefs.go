package i18n

import (
	"io/fs"
)

// LoadMessageFileFS is like LoadMessageFile but instead of reading from the
// hosts operating system's file system it reads from the fs file system.
func (b *Bundle) LoadMessageFileFS(fsys fs.FS, path string) (*MessageFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

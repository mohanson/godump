// Package path offers classes representing filesystem paths with semantics appropriate for different operating systems.
package path

import (
	"path/filepath"

	"github.com/libraries/go/doa"
)

// Path provide path-handling operations.
type Path struct {
	Full   string // The full path.
	Name   string // A string representing the final path component, excluding the drive and root, if any.
	Parent *Path  // The logical parent of the path.
	Stem   string // The final path component, without its suffix.
	Suffix string // The file extension of the final component, if any.
}

// String returns path's full name.
func (p *Path) String() string {
	return p.Full
}

// Join joins a path element.
func (p *Path) Join(elem string) *Path {
	return Wrap(filepath.Join(p.Full, elem))
}

// Wrap returns a new Path.
func Wrap(name string) *Path {
	path := &Path{}
	path.Full = filepath.ToSlash(doa.Try(filepath.Abs(name)))
	path.Name = filepath.Base(path.Full)
	if path.Full == "/" {
		path.Parent = nil
	} else {
		path.Parent = Wrap(filepath.Dir(path.Full))
	}
	path.Suffix = filepath.Ext(path.Name)
	path.Stem = path.Name[:len(path.Name)-len(path.Suffix)]
	return path
}

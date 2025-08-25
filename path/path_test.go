package path

import (
	"testing"

	"github.com/libraries/go/doa"
)

func TestPath(t *testing.T) {
	p0 := Wrap("/parent/stem.suffix")
	doa.Doa(p0.Full == "/parent/stem.suffix")
	doa.Doa(p0.Name == "stem.suffix")
	doa.Doa(p0.Parent.Full == "/parent")
	doa.Doa(p0.Parent.Parent.Full == "/")
	doa.Doa(p0.Parent.Parent.Parent == nil)
	doa.Doa(p0.Stem == "stem")
	doa.Doa(p0.Suffix == ".suffix")

	p1 := Wrap("/parent/name/")
	doa.Doa(p1.Full == "/parent/name")
	doa.Doa(p1.Name == "name")
}

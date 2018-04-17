package state

import (
	"github.com/snarlysodboxer/k8s-spec/spec"
)

type Interface interface {
	Create(*spec.SpecGroup) (int, error)
	Update(int, *spec.SpecGroup) (int, error)
	Read(int) (*spec.SpecGroup, error)
	List() (*map[int]string, error)
	Delete(int) error
}

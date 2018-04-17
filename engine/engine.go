package engine

import (
	"github.com/snarlysodboxer/k8s-spec/spec"
)

type Interface interface {
	Apply(*spec.SpecGroup) error
	GetUsingLabel(string, string, string) (map[string]string, error)
	Delete(*spec.SpecGroup) error
}

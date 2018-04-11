package kctl

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/spec"
)

type KubectlHandler interface {
	Apply(spec.SpecGroup) error
}

type Kubectl struct {
}

func (kubectl *Kubectl) Apply(specGroup *spec.SpecGroup) error {
	fmt.Println("Apply placeholder")
	return nil
}

package kctl

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/spec"
	"os/exec"
)

const (
	KubectlApply = `cat <<EOF | kubectl apply -f -
%s
EOF`
	KubectlDelete = `cat <<EOF | kubectl delete -f -
%s
EOF`
)

type KubectlHandler interface {
	Apply(spec.SpecGroup) error
}

type Kubectl struct {
}

func (kubectl *Kubectl) Apply(specGroup *spec.SpecGroup) error {
	for _, spec := range specGroup.Specs {
		output, err := exec.Command(
			"sh", "-c", fmt.Sprintf(KubectlApply, spec.Rendered),
		).CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
			return err
		}
		fmt.Println(string(output))
	}
	return nil
}

func (kubectl *Kubectl) Delete(specGroup *spec.SpecGroup) error {
	for _, spec := range specGroup.Specs {
		output, err := exec.Command(
			"sh", "-c", fmt.Sprintf(KubectlDelete, spec.Rendered),
		).CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
			return err
		}
		fmt.Println(string(output))
	}
	return nil
}

package kctl

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/spec"
	"os/exec"
	"strings"
)

const (
	KubectlApply = `cat <<EOF | kubectl apply -f -
%s
EOF`
	KubectlDelete = `cat <<EOF | kubectl delete -f -
%s
EOF`
	KubectlGetUsingLabel = `kubectl get %s -l %s=%s --ignore-not-found -o jsonpath="{range .items[*]}{.kind} {.metadata.name}{\"\n\"}{end}"`
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
		fmt.Println(string(output)) // TODO maybe pass this back, don't print it
	}
	return nil
}

func (kubectl *Kubectl) GetUsingLabel(kind, labelKey, labelValue string) (map[string]string, error) {
	output, err := exec.Command(
		"sh", "-c", fmt.Sprintf(KubectlGetUsingLabel, kind, labelKey, labelValue),
	).CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return nil, err
	}
	objects := make(map[string]string)
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line != "" {
			keyValue := strings.Split(line, " ")
			objects[keyValue[0]] = keyValue[1]
		}
	}
	return objects, nil
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
		fmt.Println(string(output)) // TODO maybe pass this back, don't print it
	}
	return nil
}

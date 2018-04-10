package spec

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

const (
	MetadataNameReplacer = `(.metadata.name | capture(\"(?<a>%s)(?<b>.*)\")) as \$names | (.metadata | select(.name == \"\(\$names.a + \$names.b)\") | .name) |= \"%s\(\$names.b)\"`
)

type SpecGroup struct {
	specs []SpecHandler
}

func (specGroup *SpecGroup) AddSpec(spec SpecHandler) {
	specGroup.specs = append(specGroup.specs, spec)
}

func (specGroup *SpecGroup) Activate() error {
	fmt.Println("Activate placeholder")
	return nil
}

func NewMetadataNameReplacer(changeString, replacementValue string) string {
	return fmt.Sprintf(MetadataNameReplacer, changeString, replacementValue)
}

type SpecHandler interface {
	Render() (string, error)
}

type Spec struct {
	Rendered  []byte
	filePath  string
	template  []byte
	replacers []string
}

func (spec *Spec) ReadTemplateFile(filePath string) {
	spec.filePath = filePath
	template, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	spec.template = template
}

func (spec *Spec) SetTemplateString(template []byte) {
	spec.template = template
}

func (spec *Spec) AddReplacer(replacer string) {
	spec.replacers = append(spec.replacers, replacer)
}

func (spec *Spec) Render() (string, error) {
	rendered := spec.template
	for _, replacer := range spec.replacers {
		output, err := exec.Command(
			// TODO support JSON as well
			"sh", "-c", fmt.Sprintf(`echo "%s" | yq -y "%s"`, rendered, replacer),
		).CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
			return "", err
		}
		rendered = output
	}
	spec.Rendered = rendered
	return string(spec.Rendered), nil
}

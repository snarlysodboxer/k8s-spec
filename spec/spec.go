package spec

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	MetadataNameReplacer                    = `(.metadata.name | capture(\"(?<a>%s)(?<b>.*)\")) as \$names | (.metadata | select(.name == \"\(\$names.a + \$names.b)\") | .name) |= \"%s\(\$names.b)\"`
	MetadataLabelsReplacer                  = `(.metadata.labels | select(.%s == \"%s\") | .%s) |= \"%s\"`
	SpecTemplateMetadataLabelsReplacer      = `(.spec.template.metadata.labels | select(.%s == \"%s\") | .%s) |= \"%s\"`
	SpecTemplateSpecContainersImageReplacer = `(.spec.template.spec.containers[] | select(.image == \"%s:%s\") | .image) |= \"%s:%s\"`
	SpecSelectorReplacer                    = `(.spec.selector | select(.%s == \"%s\") | .%s) |= \"%s\"`
	CommonLabelAdder                        = `.metadata.labels = (.metadata.labels + {\"%s\": \"%s\"})`
)

type SpecGroup struct {
	Specs            []*Spec
	commonLabelKey   string
	commonLabelValue string
}

func (specGroup *SpecGroup) AddSpec(specHandler SpecHandler) {
	specGroup.Specs = append(specGroup.Specs, specHandler.Get())
}

func (specGroup *SpecGroup) Render() (string, error) {
	allRendered := []string{}
	for _, spec := range specGroup.Specs {
		if specGroup.commonLabelKey != "" && specGroup.commonLabelValue != "" {
			spec.AddReplacer(fmt.Sprintf(CommonLabelAdder, specGroup.commonLabelKey, specGroup.commonLabelValue))
		}
		rendered, err := spec.Render()
		if err != nil {
			panic(err)
		}
		allRendered = append(allRendered, rendered)
	}
	ren := strings.Join(allRendered, "\n---\n\n")
	return string(ren), nil
}

func (specGroup *SpecGroup) SetCommonLabel(key, value string) {
	specGroup.commonLabelKey = key
	specGroup.commonLabelValue = value
}

func NewMetadataNameReplacer(changeString, replacementValue string) string {
	return fmt.Sprintf(MetadataNameReplacer, changeString, replacementValue)
}

func NewMetadataLabelsReplacer(changeString, labelKey, replacementValue string) string {
	return fmt.Sprintf(MetadataLabelsReplacer, labelKey, changeString, labelKey, replacementValue)
}

func NewSpecTemplateMetadataLabelsReplacer(changeString, labelKey, replacementValue string) string {
	return fmt.Sprintf(SpecTemplateMetadataLabelsReplacer, labelKey, changeString, labelKey, replacementValue)
}

func NewSpecTemplateSpecContainersImageReplacer(changeString, dockerRepo, replacementValue string) string {
	return fmt.Sprintf(SpecTemplateSpecContainersImageReplacer, dockerRepo, changeString, dockerRepo, replacementValue)
}

func NewSpecSelectorReplacer(changeString, labelKey, replacementValue string) string {
	return fmt.Sprintf(SpecSelectorReplacer, labelKey, changeString, labelKey, replacementValue)
}

type SpecHandler interface {
	Render() (string, error)
	Get() *Spec
}

type Spec struct {
	Rendered  []byte
	filePath  string
	template  []byte
	replacers []string
}

func (spec *Spec) Get() *Spec {
	return spec
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

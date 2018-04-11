package main

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/kctl"
	"github.com/snarlysodboxer/k8s-spec/spec"
)

type CustomKubectl struct {
	kctl.Kubectl
}

func (customKubectl *CustomKubectl) Apply(specGroup *spec.SpecGroup) error {
	fmt.Println("Custom Apply")
	return nil
}

func main() {
	// Create SpecGroup
	specGroup := &spec.SpecGroup{}

	// Setup Spec
	deploymentSpec := &spec.Spec{}
	deploymentSpec.ReadTemplateFile("./example-deployment.yml")

	// Add Built-in Replacers
	deploymentSpec.AddReplacer(spec.NewMetadataNameReplacer("CHANGEME", "my-app"))

	// Render Spec
	rendered, err := deploymentSpec.Render()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rendered the following spec:\n%s\n", rendered)

	// Add Spec to SpecGroup
	specGroup.AddSpec(deploymentSpec)

	// Apply SpecGroup to k8s
	customKubectl := &CustomKubectl{kctl.Kubectl{}}
	err = customKubectl.Apply(specGroup)
	if err != nil {
		panic(err)
	}
}

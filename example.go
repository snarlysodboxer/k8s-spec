package main

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/spec"
)

func main() {
	// Create SpecGroup
	specGroup := &spec.SpecGroup{}

	// Setup Spec
	deploymentSpec := &spec.Spec{}
	deploymentSpec.ReadTemplateFile("./example-deployment.yml")
	// // Or create a template string and add it
	// deploymentSpec.SetTemplateString(myTemplateString)

	// Built-in Replacers
	deploymentSpec.AddReplacer(spec.NewMetadataNameReplacer("CHANGEME", "my-app"))
	deploymentSpec.AddReplacer(spec.NewMetadataLabelsReplacer("CHANGEME", "my-app", "app"))
	deploymentSpec.AddReplacer(spec.NewSpecTemplateMetadataLabelsReplacer("CHANGEME", "my-app", "app"))
	deploymentSpec.AddReplacer(spec.NewSpecTemplateSpecContainersImageReplacer("CHANGEME", "0.0.1", "my-account/my-repo"))
	// // Custom Replacer
	// deploymentSpec.AddReplacer(`(.metadata.labels | select(.app == \"CHANGEME\") | .app) |= \"my-app\"`)

	// Render Spec
	rendered, err := deploymentSpec.Render()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rendered the following spec:\n%s\n", rendered)

	// Add Spec to SpecGroup
	specGroup.AddSpec(deploymentSpec)

	// Activate SpecGroup
	err = specGroup.Activate()
	if err != nil {
		panic(err)
	}
}

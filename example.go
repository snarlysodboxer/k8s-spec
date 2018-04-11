package main

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/kctl"
	"github.com/snarlysodboxer/k8s-spec/spec"
)

func main() {
	// Create SpecGroup
	specGroup := &spec.SpecGroup{}

	// Setup Spec for Deployment
	deploymentSpec := &spec.Spec{}
	deploymentSpec.ReadTemplateFile("./example-deployment.yml")
	// // Or create a template string and add it
	// deploymentSpec.SetTemplateString([]byte("my spec"))

	// Add Built-in Replacers
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

	// Setup Spec for Service
	serviceSpec := &spec.Spec{}
	serviceSpec.ReadTemplateFile("./example-service.yml")
	serviceSpec.AddReplacer(spec.NewMetadataNameReplacer("CHANGEME", "my-app"))
	serviceSpec.AddReplacer(spec.NewMetadataLabelsReplacer("CHANGEME", "my-app", "app"))
	serviceSpec.AddReplacer(spec.NewSpecSelectorReplacer("CHANGEME", "my-app", "app"))
	rendered, err = serviceSpec.Render()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rendered the following spec:\n%s\n", rendered)
	specGroup.AddSpec(serviceSpec)

	// Apply SpecGroup to k8s
	kubectl := &kctl.Kubectl{}
	err = kubectl.Apply(specGroup)
	if err != nil {
		panic(err)
	}

	err = kubectl.Delete(specGroup)
	if err != nil {
		panic(err)
	}
}

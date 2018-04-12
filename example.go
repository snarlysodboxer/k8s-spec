package main

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/kctl"
	"github.com/snarlysodboxer/k8s-spec/spec"
)

func main() {
	// Create SpecGroup
	specGroup := &spec.SpecGroup{}

	// Set a common label to be added to all Specs in the SpecGroup (i.e. if Spec discovery by label is desired)
	specGroup.SetCommonLabel("created-by", "k8s-spec")

	// Setup Spec for Deployment
	deploymentSpec := &spec.Spec{}
	specGroup.AddSpec(deploymentSpec)
	deploymentSpec.ReadTemplateFile("./example-deployment.yml")
	// // Or create a template string and add it
	// deploymentSpec.SetTemplateString([]byte("my spec"))
	// Built-in Replacers
	deploymentSpec.AddReplacer(spec.NewMetadataNameReplacer("CHANGEME", "my-app"))
	deploymentSpec.AddReplacer(spec.NewMetadataLabelsReplacer("CHANGEME", "my-app", "app"))
	deploymentSpec.AddReplacer(spec.NewSpecTemplateMetadataLabelsReplacer("CHANGEME", "my-app", "app"))
	deploymentSpec.AddReplacer(spec.NewSpecTemplateSpecContainersImageReplacer("CHANGEME", "0.0.1", "my-account/my-repo"))
	// Custom Replacer
	deploymentSpec.AddReplacer(`.spec.replicas = 2`)

	// Setup Spec for Service
	serviceSpec := &spec.Spec{}
	specGroup.AddSpec(serviceSpec)
	serviceSpec.ReadTemplateFile("./example-service.yml")
	serviceSpec.AddReplacer(spec.NewMetadataNameReplacer("CHANGEME", "my-app"))
	serviceSpec.AddReplacer(spec.NewMetadataLabelsReplacer("CHANGEME", "my-app", "app"))
	serviceSpec.AddReplacer(spec.NewSpecSelectorReplacer("CHANGEME", "my-app", "app"))

	// Render each Spec
	rendered, err := specGroup.Render()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rendered the following specs:\n\n%s\n", rendered)

	// Apply SpecGroup to Kuberentes
	kubectl := &kctl.Kubectl{}
	err = kubectl.Apply(specGroup)
	if err != nil {
		panic(err)
	}

	// Delete SpecGroup from Kuberentes
	err = kubectl.Delete(specGroup)
	if err != nil {
		panic(err)
	}
}

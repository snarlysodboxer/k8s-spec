package main

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/engine/kubectl"
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
	deploymentSpec.AddReplacer(spec.NewMetadataLabelsReplacer("CHANGEME", "app", "my-app"))
	deploymentSpec.AddReplacer(spec.NewSpecTemplateMetadataLabelsReplacer("CHANGEME", "app", "my-app"))
	deploymentSpec.AddReplacer(spec.NewSpecTemplateSpecContainersImageReplacer("CHANGEME", "my-account/my-repo", "0.0.1"))
	// Custom Replacer
	deploymentSpec.AddReplacer(`.spec.replicas = 2`)

	// Setup Spec for Service
	serviceSpec := &spec.Spec{}
	specGroup.AddSpec(serviceSpec)
	serviceSpec.ReadTemplateFile("./example-service.yml")
	serviceSpec.AddReplacer(spec.NewMetadataNameReplacer("CHANGEME", "my-app"))
	serviceSpec.AddReplacer(spec.NewMetadataLabelsReplacer("CHANGEME", "app", "my-app"))
	serviceSpec.AddReplacer(spec.NewSpecSelectorReplacer("CHANGEME", "app", "my-app"))

	// Render each Spec
	rendered, err := specGroup.Render()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rendered the following specs:\n\n%s\n", rendered)

	// Apply SpecGroup to Kuberentes
	engine := kubectl.NewEngine()
	err = engine.Apply(specGroup)
	if err != nil {
		panic(err)
	}

	// Get objects from Kubernetes
	objects, err := engine.GetUsingLabel("deployment,pod,svc", "app", "my-app")
	if err != nil {
		panic(err)
	}
	fmt.Println("Found the following objects:")
	for kind, name := range objects {
		fmt.Printf("\t%-26s%s\n", kind, name)
	}

	// Delete SpecGroup from Kubernetes
	err = engine.Delete(specGroup)
	if err != nil {
		panic(err)
	}
}

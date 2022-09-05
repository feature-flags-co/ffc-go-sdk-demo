package main

import (
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/ffc"
	"github.com/feature-flags-co/ffc-go-sdk/model"
)

func main() {

	// Step One:
	envSecret := "ZDMzLTY3NDEtNCUyMDIxMTAxNzIxNTYyNV9fMzZfXzQ2X185OF9fZGVmYXVsdF80ODEwNA=="

	// Step Two
	// if using dev env,you can create streaming builder use wss://api-dev.featureflag.co
	// if using product env, you can use default streaming builder when create config builder
	streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")

	// Step Three:
	// create config builder
	config := ffc.DefaultFFCConfigBuilder().
		SetOffline(true).
		UpdateProcessorFactory(streamingBuilder). // if using product env, don't need set this code
		Build()

	// Step Four:
	// create ffc client and invoke client's api
	client := ffc.NewClient(envSecret, config)
	fmt.Println(client)

	ffcUser := model.NewFFUserBuilder().
		UserName("userName").
		Country("country").
		Email("email").
		Custom("key", "value").Build()

	// invoke api
	flagStatue := client.VariationDetail("featureFlagKey", ffcUser, "defaultValue")
	fmt.Printf("flagstate %v", flagStatue)

	userTags := client.GetAllLatestFlagsVariations(ffcUser)
	fmt.Printf("userTags %v", userTags)

	//
}

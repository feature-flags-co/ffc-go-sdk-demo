# Go Server Side SDK Demo

## Introduction

This is the Go Server Side SDK Demo for the feature management platform [featureflag.co](https://featureflag.co/). It is
intended to introduce how to use Go Server Side SDK.

There are several main steps


##Step One:

Get the project Secret key from  [featureflag.co](https://featureflag.co/) platform.
like this 
```
    NWM4LTAzODgtNCUyMDIyMDcwNzE0MzUzN19fMTc3X18yMDZfXzQxNl9fZGVmYXVsdF8zNDY2Yw
```
##Step Two:

Create a StreamingBuilder, if using test env, you can like this:
```go
	streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")
```
if you are product env, you can use default streaming builder when create config builder.
##Step Three:

Create a ConfigBuilder,using StreamingBuilder, like this:

```go
	config := ffc.DefaultFFCConfigBuilder().
		SetOffline(false).
		UpdateProcessorFactory(streamingBuilder).
		Build()
	client = ffc.NewClient(envSecret, config)
```

if product env you can using below code :

```go
    config := ffc.DefaultFFCConfigBuilder().
		Build()
```
##Step Four:

Create ffc client and invoke clint api,get the data that you want.

```go

    client = ffc.NewClient(envSecret, config)

    ffcUser := model.NewFFUserBuilder().
    UserName("userName").
    Country("country").
    Email("email").
    Custom("key", "value").Build()

    // get the flag state
    flagtState := client.VariationDetail("featureFlagKey", ffcUser, "defaultValue")
    
    // get all user tags
    userTags := client.GetAllLatestFlagsVariations(ffcUser)

```


# Go Server Side SDK Demo

## Introduction

This is the Go Server Side SDK Demo for the feature management platform [featureflag.co](https://featureflag.co/). It is
intended to introduce how to use Go Server Side SDK. There are several steps to help you get started.

## Step One: Import
```go
import (
    "fmt"
    "github.com/feature-flags-co/ffc-go-sdk/ffc"
    "github.com/feature-flags-co/ffc-go-sdk/model"
)
```

## Step Two: Config

create a default config
```go
config := ffc.DefaultFFCConfigBuilder().Build()
```

this default config use our streaming server, If you deploy our streaming server locally, you can change like this
```go

streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://<your-streaming-server-address>")

config := ffc.DefaultFFCConfigBuilder().
	UpdateProcessorFactory(streamingBuilder).
	Build()
```

## Step Three: Get Client
create a client which you can interact with like this, we need an envSecret and your config. 
```go
client = ffc.NewClient(envSecret, config)
```
TIP: You can go to our [portal](https://portal.featureflag.co/) or your locally deployed portal to get an envSecret.

## Step Four: Evaluate Feature Flag

Create ffc client and invoke clint api,get the data that you want.

```go

// setup user
ffcUser := model.NewFFUserBuilder().
	UserName("userName").
	Country("country").
	Email("email").
	Custom("key", "value").
	Build()
	
// TODO: get flag value for this user

```


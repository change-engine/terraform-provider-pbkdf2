package main

import (
	"context"
	"flag"
	"log"

	"github.com/change-engine/terraform-provider-pbkdf2/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate tofu fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/change-engine/pbkdf2",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}

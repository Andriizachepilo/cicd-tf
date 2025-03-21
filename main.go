package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/Andriizachepilo/cicd-tf/internal/provider"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Set to true for debug logs.")
	flag.Parse()

    err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
        Address: "registry.terraform.io/Andriizachepilo/cicd",
        Debug: debug,
    })

    if err != nil {
        log.Fatal(err)
    }
}
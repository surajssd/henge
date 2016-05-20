package main

import (
	"flag"
	"fmt"

	"github.com/openshift/origin/pkg/generate/app"
	"github.com/rtnpro/henge/pkg/generate/dockercompose"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
)

func main() {

	flag.Parse()

	template, err := dockercompose.Generate(flag.Args()[0:]...)
	if err != nil {
		return
	}

	if errs := app.AsVersionedObjects(template.Objects, kapi.Scheme, kapi.Scheme, registered.EnabledVersions()...); len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("error: %v\n", err)
		}
	}

	// make it List instead of Template
	list := &kapi.List{Items: template.Objects}
	for _, obj := range list.Items {
		fmt.Printf("%#v\n", obj)
	}

}

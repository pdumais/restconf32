package main

import (
	"strings"
	"time"

	"github.com/freeconf/restconf"
	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/source"
)

// write your code to capture your domain however you want

type Animals struct {
	Cats string
	Dogs string
}

type Input struct {
	Cats string
	Dogs string
}

type Output struct {
	Cats string
	Dogs string
}

func (c *Animals) Start() {
	for {
		<-time.After(time.Second)
	}
}

// write mangement api to bridge from YANG to code
func manage(animal *Animals) node.Node {
	return &nodeutil.Extend{

		// use reflect when possible, here we're using to get/set speed AND
		// to read miles metrics.
		Base: nodeutil.ReflectChild(animal),

		// handle action request
		OnAction: func(parent node.Node, req node.ActionRequest) (node.Node, error) {
			switch req.Meta.Ident() {
			case "test":
				cats, _ := req.Input.Find("cats").Get()
				dogs, _ := req.Input.Find("dogs").Get()
				c := &Output{
					Cats: cats.String(),
					Dogs: dogs.String(),
				}
				return nodeutil.ReflectChild(c), nil
			}
			return nil, nil
		},
	}
}

// Connect everything together into a server to start up
func main() {

	// Your app
	a := &Animals{}

	// Device can hold multiple modules, here we are only adding one
	d := device.New(source.Path("."))
	if err := d.Add("animals", manage(a)); err != nil {
		panic(err)
	}

	// Select wire-protocol RESTCONF to serve the device.
	restconf.NewServer(d)

	// apply start-up config normally stored in a config file on disk
	config := `{
		"fc-restconf":{"web":{"port":":8080"}}
	}`

	if err := d.ApplyStartupConfig(strings.NewReader(config)); err != nil {
		panic(err)
	}

	// start your app
	a.Start()
}

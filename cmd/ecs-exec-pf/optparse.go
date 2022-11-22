package main

import (
	"log"

	"github.com/integrii/flaggy"
)

var version string

const (
	Description = "Remote Port forwarding using the ECS task container. (aws-cli wrapper)"
)

type Options struct {
	Cluster   string
	Task      string
	Container string
	Host	  string
	Port      uint16
	LocalPort uint16
}

func parseArgs() *Options {
	opts := &Options{}

	flaggy.SetDescription(Description)
	flaggy.SetVersion(version)
	flaggy.String(&opts.Cluster, "c", "cluster", "ECS cluster name.")
	flaggy.String(&opts.Task, "t", "task", "ECS task ID.")
	flaggy.String(&opts.Container, "n", "container", "Container name in ECS task.")
	flaggy.String(&opts.Host, "e", "endpoint", "Host/DB Endpoint.")
	flaggy.UInt16(&opts.Port, "p", "port", "Target remote port.")
	flaggy.UInt16(&opts.LocalPort, "l", "local-port", "Client local port.")
	flaggy.Parse()

	if opts.Cluster == "" {
		log.Fatal("'--cluster' is required")
	}

	if opts.Task == "" {
		log.Fatal("'--task' is required")
	}

	if opts.Host == "" {
		log.Fatal("'--endpoint' is required")
	}

	if opts.Port == 0 {
		log.Fatal("'--port' is required")
	}

	if opts.LocalPort == 0 {
		log.Fatal("'--local-port' is required")
	}

	return opts
}

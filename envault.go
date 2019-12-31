package main

import (
	"fmt"

	"github.com/pratishshr/envault/cli"
)

func main() {
	info := &cli.Info{
		Name:        "Envault",
		Version:     "1.1.6",
		Description: "Envault is a simple CLI tool which runs a process with secrets from AWS Secrets Manager.",
		AuthorName:  "Pratish Shrestha",
		AuthorEmail: "pratishshr@gmail.com",
	}

	err := cli.Initialize(info)

	if err != nil {
		fmt.Println(err)
	}
}

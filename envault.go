package main

import "github.com/pratishshr/envault/cli"

func main() {
	info := &cli.Info{
		Name:        "Envault",
		Version:     "0.0.1",
		Description: "Envault is a simple CLI tool to run a process with secrets from AWS Secrets Manager.",
		AuthorName:  "Pratish Shrestha",
		AuthorEmail: "pratishshr@gmail.com",
	}

	err := cli.Initialize(info)

	if err != nil {
		panic(err)
	}
}

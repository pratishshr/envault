# Envault 
![GitHub release](https://img.shields.io/github/release/pratishshr/envault.svg?style=flat-square)
![GitHub](https://img.shields.io/github/license/pratishshr/envault.svg?style=flat-square)

A simple CLI tool to run a process with secrets from AWS Secrets Manager.

## About

Envault focuses on integrating AWS Secrets Manager in your application with ease without having to write a single line of code in your source files. Simply run your commands with the Envault CLI and the secrets will be injected in that process only.

## Quick Start

#### Install Envault:

```
$ curl -sf https://raw.githubusercontent.com/pratishshr/envault/master/install.sh | sh
```

#### Setup environment:

```
$ export AWS_REGION={AWS region where you added your secret}
$ export AWS_ACCESS_KEY_ID=${Your AWS access key}
$ export AWS_SECRET_ACCESS_KEY=${Your AWS secret key}
```

Envault also supports AWS credentials from `~/.aws` with the `default` profile if you don't set the environment variables. <br>

#### List secrets:

```
$ envault list -secret=${path to your secret}
```

#### Run process with secrets:

```
$ envault run 'yarn build' -secret=${path to your secret}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

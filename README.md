# Envault
![GitHub release](https://img.shields.io/github/release/pratishshr/envault.svg?style=flat)
![Travis (.org)](https://img.shields.io/travis/pratishshr/envault.svg?style=flat)
![GitHub](https://img.shields.io/github/license/pratishshr/envault.svg?style=flat)

A simple CLI tool to run a process with secrets from AWS Secrets Manager.

## About

Envault focuses on integrating AWS Secrets Manager in your application with ease without having to write a single line of code in your source files. Simply run your commands with the Envault CLI and the secrets will be injected in that process.

## Table Of Contents
1. [Install Envault](#1-install-envault)
2. [Verify Installation](#2-verify-installation)
3. [AWS Credentials](#3-aws-credentials)
4. [Setup](#4-setup)
5. [List Secrets](#5-list-secrets)
6. [Run With Secrets](#6-run-with-secrets)
5. [Usage with CI/CD](#7-usage-with-cicd)
6. [Using custom .env files](#8-using-custom-env-files)

## Usage

### 1. Install Envault:

```
$ curl -sf https://raw.githubusercontent.com/pratishshr/envault/master/install.sh | sudo sh
```

### 2. Verify Installation:

```
$ envault
```

### 3. AWS Credentials

Before using envault, you have to provide your AWS credentials. This allows envault to fetch secrets from the AWS Secrets Manager. Also, make sure you have the correct access for your credentials.

Simply create `~/.aws/credentials` file for storing AWS credentials. <br/>
Example:

```
[example-profile]
aws_access_key_id = xxxxxx
aws_secret_access_key = xxxxxx
```
To know more about AWS configurations, view [Configuring the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)

### 4. Setup

Go to your project directory and run `setup` command to initiate the setup process.

```
$ envault setup
```

- Choose your AWS profile that was setup earlier. <br>
- Choose the AWS Region where your secrets are kept.
- You can also add a deployment environment associated with the secret name. You may add any number of environment you want.
- Set a default env

```
 Example:

 AWS profile: default
 Region: US West (Oregon)
 Add an environment (eg. dev): dev
 Secret Name: api/dev
 Add an environment (eg. dev): uat
 Secret Name: api/uat
```
`envault.json` file will be created in your project directory like below.
```json
{
  "profile": "default",
  "region": "us-west-2",
  "environments": {
    "dev": "api/dev",
    "uat": "api/uat"
  },
  "defaultEnv": "dev"
}
```

**If you do not want a project-specific config file, you can skip the above step.**

### 5. List secrets

```
$ envault list -e dev
$ envault list -e uat
```
Here `dev` and `uat` are the environments you specified in `envault.json`.


If you have not setup a `envault.json` file, you can still pass `--secret` or `-s` flag with the secrets path.
This will use the `default` profile from your `~/.aws/credentials` file.
```
$ envault list --secret=api/dev
$ envault list --secret=api/uat
```

### 6. Run with secrets

```
$ envault run 'yarn build' -e dev
```
This will inject the secrets from `dev` to the `yarn build` process.

Similarly, if you have not setup a `envault.json` file, you can still pass `--secret` or `-s` flag with the secrets path.
This will use the `default` profile from your `~/.aws/credentials` file.

```
$ envault run 'yarn build' --secret=api/dev
```

### 7. Usage with CI/CD:

Instead of setting up a `~/.aws/credentials` file. You can also use the following environment variables to set up your AWS credentials.

| Variable | Description |
|-----------|----------|
| AWS_ACCESS_KEY_ID | Your AWS access key|
| AWS_SECRET_ACCESS_KEY | Your AWS secret key|
| AWS_REGION | AWS region where you added your secret|
| ENVIRONMENT | Environment which you set in envault.json |


### 8. Using custom .env files
If you want to inject environment keys from a file instead of using AWS Secrets Manager. You can use the`-ef` flag.

```
$ envault run 'envault run 'go run main.go' -ef env/staging.env
```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

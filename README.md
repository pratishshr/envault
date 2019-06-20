# envault

[![PyPI](https://img.shields.io/pypi/v/envault.svg?style=for-the-badge)](https://pypi.org/project/envault/) [![PyPI - Downloads](https://img.shields.io/pypi/dd/envault.svg?style=for-the-badge)](https://pypi.org/project/envault/)

Envault is a simple CLI tool to run a process with secrets from HashiCorp Vault.

## Installation

`envault` requires `Python 3` and `pip` installed.

```sh
$ pip install envault
```

## Usage

### List secrets from a secret engine


> Note that `asm` is default engine

```
$ envault list -secret=<SECRET_NAME> \
               -region=<REGION_NAME>
               -accessid=<AWS_ACCESS_KEY_ID>
               -secretkey=<AWS_SECRET_ACCESS_KEY>
```

> Note that `KV version 2` follows the secrets path as: `${SECRET_ENGINE}/data/${SECRETS_PATH}`

```
$ envault list -server=https://vault.test-server.com \
               -secret=kv/data/api \
               -token=<VAULT_TOKEN>
               -engine=vault
```


### Run a process with injected environment variables from ASM

```
$ envault run 'node index.js' -secret=<SECRET_NAME> \
                              -region=<REGION_NAME>
                              -accessid=<AWS_ACCESS_KEY_ID>
                              -secretkey=<AWS_SECRET_ACCESS_KEY>
```


### Run a process with injected environment variables from vault

```
$ envault run 'node index.js' -server=https://vault.test-server.com \
                              -secret=kv/data/api \
                              -token=<VAULT_TOKEN>
```

### Environment Variables

Additionally you can also keep the following environment variables instead of passing it from the CLI.

> ASM configuration

|Variable|CLI Option|Description|
|--------|-----------|----------|
|SECRET_NAME| -secret | Path to the secrets |
|REGION_NAME| -region | AWS Secret manager region name |
|AWS_ACCESS_KEY_ID| -accessid | AWS Access Key ID |
|AWS_SECRET_ACCESS_KEY| -secretkey | AWS Secret Access Key |


> Vault congiguration

|Variable|CLI Option|Description|
|--------|-----------|----------|
|VAULT_TOKEN| -token | Vault token |
|VAULT_SERVER| -server | Server URI |
|VAULT_SECRETS_PATH| -secret | Path to the secrets |
|ENGINE| -engine | Secret engine name |

This is helpful when you need to run your scripts from a CI Server.
After setting these variables, you can simply run `envault` as:

```
$ envault run 'yarn build'
```

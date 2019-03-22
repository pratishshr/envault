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

> Note that `KV version 2` follows the secrets path as: `${SECRET_ENGINE}/data/${SECRETS_PATH}`

```
$ envault list -server=https://vault.test-server.com \
               -secret=kv/data/api \
               -token=<VAULT_TOKEN>
```

### Run a process with injected environment variables from vault

```
$ envault run 'node index.js' -server=https://vault.test-server.com \
                              -secret=kv/data/api \
                              -token=<VAULT_TOKEN>
```

### Environment Variables

Additionally you can also keep the following environment variables instead of passing it from the CLI.

|Variable|CLI Option|Description|
|--------|-----------|----------|
|VAULT_TOKEN| -token | Vault token |
|VAULT_SERVER| -server | Server URI |
|VAULT_SECRETS_PATH| -secret | Path to the secrets |

This is helpful when you need to run your scripts from a CI Server.
After setting these variables, you can simply run `envault` as:

```
$ envault run 'yarn build'
```

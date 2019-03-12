# envault

[![PyPI](https://img.shields.io/pypi/v/envault.svg?style=for-the-badge)](https://pypi.org/project/envault/) [![PyPI - Downloads](https://img.shields.io/pypi/dd/envault.svg?style=for-the-badge)](https://pypi.org/project/envault/)

Envault is a simple CLI tool to run processes with secrets from HashiCorp Vault.

## Usage

**envault** requires `Python 3` and `pip` installed.

### Install envault

```sh
$ pip install envault
```

### List secrets from a secret engine

\*\* Note that KV version 2 follows the secrets path as below:  
`${SECRET_ENGINE}/data/${SECRETS_PATH}`

```
$ envault list -server=https://vault.test-server.com
-secret=kv/data/api
-token=<VAULT_TOKEN>
```

### Run a process with injected environment variables from vault

```
$ envault run 'yarn build' -server=https://vault.test-server.com
-secret=kv/data/api
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

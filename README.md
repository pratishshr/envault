# envault

[WORK IN PROGRESS]

Envault is a simple CLI tool to run processes with secrets from HashiCorp Vault. It uses API calls to fetch data from vault.

## Usage

_envault_ requires Python 3 and Pip installed.

### 1. Install dependencies.

```sh
$ pip install envault
```

### 2. List

```
$ envault list -server=https://vault.lftechnology.com
-secret=secrets/data/Inhouse/lms-auth/dev/api
-token=s.gy8rv3Op1ClA1i0bqccHTzBw
```

### 3. Run

```
$ envault run ./script.sh -server=https://vault.lftechnology.com
-secret=secrets/data/Inhouse/lms-auth/dev/api
-token=s.gy8rv3Op1ClA1i0bqccHTzBw
```

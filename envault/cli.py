#!/usr/bin/env python3

import os

import click

from envault import vault, shell, __version__


def get_secrets(server, secret, token):
    """ Renew token and fetch secrets from Vault Server """
    token = token or os.environ.get("VAULT_TOKEN")
    server = server or os.environ.get("VAULT_SERVER")
    secret = secret or os.environ.get("VAULT_SECRETS_PATH")

    if not server:
        raise SystemExit(
            "Error: Vault Server URI is not present. Add '-server' flag or VAULT_SERVER variable in your environment"
        )

    if not token:
        raise SystemExit(
            "Error: Vault Token is not present. Add '-token' flag or VAULT_TOKEN variable in your environment"
        )

    vault.renew_token(server, token)

    return vault.get_secrets(server, secret, token)


@click.group()
@click.version_option(message=__version__)
def cli():
    pass


@cli.command("list")
@click.option("-server", help="Server URI")
@click.option("-secret", help="Path to the secrets")
@click.option("-token", help="Vault token")
def list(server, secret, token):
    """ List secrets from a given path """
    secrets = get_secrets(server, secret, token)

    for key, value in secrets.items():
        click.echo("{}={}".format(key, value))


@cli.command("run")
@click.option("-server", help="Server URI")
@click.option("-secret", help="Path to the secrets")
@click.option("-token", help="Vault token")
@click.argument("command")
def run(server, secret, token, command):
    """ Run a command with the injected env variables """
    secrets = get_secrets(server, secret, token)
    shell.run_with_env(command, secrets)


if __name__ == "__main__":
    cli()

#!/usr/bin/env python3

import os

import click


from envault import vault, shell, __version__

from envault.utils import config, yaml_file

from yaml import safe_load, dump

from pathlib import Path


def get_secrets(server, secret, token, profile):
    """ Renew token and fetch secrets from Vault Server """
    profile_configs = {
        "vault_token": None,
        "vault_server": None,
        "vault_secret_path": None,
    }
    if profile:
        profile_configs = config.get_profile_configs(profile)

    token = token or profile_configs.get("vault_token") or os.environ.get("VAULT_TOKEN")
    server = (
        server or profile_configs.get("vault_server") or os.environ.get("VAULT_SERVER")
    )
    secret = (
        secret
        or profile_configs.get("vault_secret_path")
        or os.environ.get("VAULT_SECRETS_PATH")
    )

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


@cli.command("init")
def init():
    """ Initialize envault config with vault server, token and secrets path """
    click.echo("Enter the profile name, server, token and path to vault secrets")
    profile_name = click.prompt("Profile Name", type=str)
    vault_server = click.prompt("Vault Server", type=str)
    vault_token = click.prompt("Vault Token", type=str)
    vault_secret_path = click.prompt("Path to vault secret", type=str)

    click.echo(
        """
        name: {name}
        vault_server: {server}
        vault_token: {token}
        vault_secret_path: {secret_path}
        """.format(
            name=profile_name,
            server=vault_server,
            token=vault_token,
            secret_path=vault_secret_path,
        )
    )

    config_file = config.create_config_file(
        vault_server, vault_token, vault_secret_path, profile_name
    )

    yaml_file.dump_data_to_yml(config_file)


@cli.command("list")
@click.option("-server", help="Server URI")
@click.option("-secret", help="Path to the secrets")
@click.option("-token", help="Vault token")
@click.option("-profile", help="Profile name stored in yml file")
def list(server, secret, token, profile):
    """ List secrets from a given path """
    secrets = get_secrets(server, secret, token, profile)

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

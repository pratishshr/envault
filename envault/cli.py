#!/usr/bin/env python3

import os

import click

from envault import vault, shell, __version__

from yaml import safe_load, dump


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

@cli.command("init")
@click.option("-profile", help="profile for the config")
def init(profile):
    if profile is None:
        profile = 'default'
    """ Initialize envault config with vault server, token and secrets path """
    click.echo("Enter the server, token and path to vault secrets")
    vault_server = click.prompt('Vault Server', type=str)
    vault_token = click.prompt('Vault Token', type=str)
    vault_secret_path = click.prompt('Path to vault secret', type=str)

    click.echo("server %s token %s path %s" %(vault_server, vault_token, vault_secret_path))

    new_profile = {
        profile: {
            'vault_server': vault_server,
            'vault_token': vault_token,
            'vault_secret_path': vault_secret_path
        }
    }

    if os.path.exists('envault.yml'):
        yaml_file = open('envault.yml', 'r')
        all_profile_configs = safe_load(yaml_file)
        if all_profile_configs is not None and profile in all_profile_configs:
            yaml_file = open('envault.yml', 'w')
            updated_profile = {
                **all_profile_configs,
                **new_profile
            }

            dump(updated_profile, yaml_file, default_flow_style = False)            
        else:
            yaml_file = open('envault.yml', 'a')
            dump(new_profile, yaml_file, default_flow_style = False)
    else:
        yaml_file = open('envault.yml', 'w+')
        dump(new_profile, yaml_file, default_flow_style = False)


@cli.command("list")
@click.option("-server", help="Server URI")
@click.option("-secret", help="Path to the secrets")
@click.option("-token", help="Vault token")
def list(server, secret, token):
    """ List secrets from a given path """
    config = vault.get_profile_configs()
    server = config.get('vault_server')
    secret = config.get('vault_secret_path')
    token = config.get('vault_token')
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
    config = vault.get_profile_configs()
    server = config.get('vault_server')
    secret = config.get('vault_secret_path')
    token = config.get('vault_token')

    secrets = get_secrets(server, secret, token)
    shell.run_with_env(command, secrets)


if __name__ == "__main__":
    cli()

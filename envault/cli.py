import click

from envault import vault
from envault import shell


@click.group()
def cli():
    pass


@cli.command("list")
@click.option("-secret", help="Path to the secrets")
@click.option("-token", help="Vault token")
def list(secret, token):
    """ List secrets from a given path """
    secrets = vault.get_secrets(secret, token)

    for key, value in secrets.items():
        click.echo("{}={}".format(key, value))


@cli.command("run")
@click.option("-secret", help="Path to the secrets")
@click.option("-token", help="Vault token")
@click.argument("command")
def run(secret, token, command):
    """ Run a command with the injected env variables """
    secrets = vault.get_secrets(secret, token)
    shell.run_with_env(command, secrets)


if __name__ == "__main__":
    cli()

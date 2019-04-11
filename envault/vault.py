import os
import click
import requests
from requests.exceptions import HTTPError

from yaml import safe_load


def renew_token(server_uri, token):
    """ Renew vault token """
    BASE_URI = "{}/v1".format(server_uri)
    headers = {"X-Vault-Token": token}
    response = requests.post(BASE_URI + "/auth/token/renew-self", headers=headers)

    try:
        response.raise_for_status()
    except HTTPError as e:
        raise SystemExit("Error: " + str(e))


def get_secrets(server_uri, secrets_path, token):
    """ Fetch secrets from vault server """
    BASE_URI = "{}/v1/".format(server_uri)

    headers = {"X-Vault-Token": token}
    response = requests.get(BASE_URI + secrets_path, headers=headers)

    try:
        response.raise_for_status()
        data = response.json()

        return data.get("data", {}).get("data")
    except HTTPError as e:
        raise SystemExit("Error: " + str(e))

def get_profile_configs(profile = 'default'):
    """ Extract vault configurations from yml file """

    if os.path.exists('envault.yml'):
        yaml_file = open('envault.yml', 'r')
    else:
        click.echo('envault.yml file not found')
    config = safe_load(yaml_file)
    click.echo(config)
    if profile in config:
        return config.get(profile)
    else:
        click.echo('No such profile found')
#!/usr/bin/env python3

import click
import requests
from requests.exceptions import HTTPError


def get_secrets(server_uri, secrets_path, token):
    """ Fetch secrets from vault server """
    BASE_URI = "{}/v1/".format(server_uri)

    headers = {"X-Vault-Token": token}
    request = requests.get(BASE_URI + secrets_path, headers=headers)
    data = request.json()

    return data["data"]["data"]

#!/usr/bin/env python3

import click
import requests
from requests.exceptions import HTTPError


def get_secrets(secrets_path, token):
    """ Fetch secrets from vault server """
    BASE_URI = "https://vault.lftechnology.com/v1/"

    headers = {"X-Vault-Token": token}
    request = requests.get(BASE_URI + secrets_path, headers=headers)
    data = request.json()

    return data["data"]["data"]

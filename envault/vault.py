import click
import requests
from requests.exceptions import HTTPError


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


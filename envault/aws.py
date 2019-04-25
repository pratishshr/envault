import json
import boto3
import click

from botocore.exceptions import ClientError


def get_secrets(aws_client_id, aws_secret_access_key, secret_name, region_name):
    """ Fetch aws secrets """
    session = boto3.session.Session()
    client = session.client(
        service_name="secretsmanager",
        region_name=region_name,
        aws_access_key_id=aws_client_id,
        aws_secret_access_key=aws_secret_access_key,
    )

    try:
        get_secret_value_response = client.get_secret_value(SecretId=secret_name)
    except ClientError as e:
        if e.response["Error"]["Code"] == "ResourceNotFoundException":
            raise SystemExit("The requested secret " + secret_name + " was not found")
        elif e.response["Error"]["Code"] == "InvalidRequestException":
            raise SystemExit("The request was invalid due to:", e)
        elif e.response["Error"]["Code"] == "InvalidParameterException":
            raise SystemExit("The request had invalid params:", e)
    else:
        if "SecretString" in get_secret_value_response:
            text_secret_data = get_secret_value_response["SecretString"]
            return json.loads(text_secret_data)
        else:
            return get_secret_value_response["SecretBinary"]

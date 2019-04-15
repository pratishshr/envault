import os

from envault.utils import yaml_file

from yaml import safe_load

import click


def create_config_file(vault_server, vault_token, vault_secret_path, name):
    new_profile = {
        "name": name,
        "vault_server": vault_server,
        "vault_token": vault_token,
        "vault_secret_path": vault_secret_path,
    }
    config_file = {"profiles": [new_profile]}

    if os.path.exists(yaml_file.get_yml_file()):
        existing_yml_data = yaml_file.load_data_from_yml()

        if existing_yml_data is not None and "profiles" in existing_yml_data:
            existing_profiles = existing_yml_data.get("profiles")

            for index, profile in enumerate(existing_profiles):
                if profile["name"] == name:
                    existing_profiles[index] = {
                        "name": name,
                        "vault_server": vault_server,
                        "vault_token": vault_token,
                        "vault_secret_path": vault_secret_path,
                    }

                    break

                else:
                    existing_profiles.append(new_profile)

            config_file = {"profiles": existing_profiles}

    return config_file


def get_profile_configs(name="default"):
    """ Extract vault configurations from yml file """

    if os.path.exists(yaml_file.get_yml_file()):
        existing_yml_data = yaml_file.load_data_from_yml()

        if existing_yml_data is not None and "profiles" in existing_yml_data:
            existing_profiles = existing_yml_data.get("profiles")

            for profile in existing_profiles:
                if profile["name"] == name:
                    return profile

            raise SystemExit(
                "Profile of name:{name} not present in yml file.".format(name=name)
            )

        else:
            raise SystemExit("Something is wrong in envault.yml file.")

    else:
        raise SystemExit("envault.yml file not found.")

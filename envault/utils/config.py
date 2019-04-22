import os
import click

from envault.utils import yaml


def create_config_file(vault_server, vault_token, vault_secret_path, name):
    """Create a new config file based on existing file and provided profile"""
    new_profile = {
        "name": name,
        "vault_server": vault_server,
        "vault_token": vault_token,
        "vault_secret_path": vault_secret_path,
    }
    config_file = {"profiles": [new_profile]}

    if not yaml.check_yml_file_exists():
        return config_file

    existing_yml_file = yaml.load_data_from_yml()

    if not is_config_file_valid(existing_yml_file):
        raise SystemExit("Invalid envault.yml file")

    profile_index = get_profile_index_and_values(existing_yml_file, name)["index"]
    existing_profiles = existing_yml_file.get("profiles")

    if profile_index != -1:
        existing_profiles[profile_index] = new_profile
    else:
        existing_profiles.append(new_profile)

    config_file = {"profiles": existing_profiles}

    return config_file


def get_profile_configs(name="default"):
    """ Extract vault configurations from yml file """
    if not yaml.check_yml_file_exists():
        raise SystemExit("envault.yml file not found")

    existing_yml_file = yaml.load_data_from_yml()

    if not is_config_file_valid(existing_yml_file) or existing_yml_file is None:
        raise SystemExit("Invalid envault.yml file")

    found_index_and_profile = get_profile_index_and_values(existing_yml_file, name)

    if found_index_and_profile["index"] == -1:
        raise SystemExit(
            "Profile of name:{name} not present in yml file.".format(name=name)
        )

    return found_index_and_profile["profile"]


def get_profile_index_and_values(config_file, profile_name):
    """Extract profile and its index from config file"""
    initial_profile = {"index": -1, "profile": None}

    profiles = config_file.get("profiles")

    for index, profile in enumerate(profiles):
        if profile["name"] == profile_name:
            initial_profile["index"] = index
            initial_profile["profile"] = profile

    return initial_profile


def is_config_file_valid(config_file):
    """Function to check if config file is valid. Valid if file has key 'profiles'"""
    if type(config_file) is dict and "profiles" in config_file.keys():
        return True

    return False

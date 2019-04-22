import os

from pathlib import Path
from yaml import safe_load, dump


def get_yml_file_path():
    """ Function to get the yml file """
    data_folder = Path(os.path.expanduser("~"))
    file_name = data_folder / ".envault.yml"

    return file_name


def dump_data_to_yml(dump_data, file_open_mode="w+"):
    """ Function to dump data in file of given path """
    yaml_file = open(get_yml_file_path(), file_open_mode)
    dump(dump_data, yaml_file, default_flow_style=False)


def load_data_from_yml():
    """ Function to load the data in yml file """
    yaml_file = open(get_yml_file_path(), "r")
    loaded_file = safe_load(yaml_file)

    return loaded_file


def check_yml_file_exists():
    """Function to check if yaml file exists"""
    if os.path.exists(get_yml_file_path()):
        return True

    return False

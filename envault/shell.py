import os
import subprocess


def run_with_env(command, env):
    """ Run command with injected environment variables """
    environment_variables = {**env}
    subprocess.Popen(command, shell=True, env=environment_variables)

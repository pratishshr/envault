import os
import subprocess


def run_with_env(command, env):
    """ Run command with injected environment variables """
    environment_variables = {**os.environ, **env}
    child = subprocess.Popen(command, shell=True, env=environment_variables)
    child.communicate()

    return child.returncode


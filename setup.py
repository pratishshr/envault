from setuptools import setup, find_packages

import envault

try:
    long_description = open("README.md", "r").read()
except IOError:
    long_description = (
        "A simple CLI tool to run processes with secrets from HashiCorp Vault."
    )

setup(
    name="envault",
    version=envault.__version__,
    author="Pratish Shrestha",
    author_email="pratishshr@gmail.com",
    packages=find_packages(),
    description="A simple CLI tool to run processes with secrets from HashiCorp Vault.",
    long_description=long_description,
    py_modules=["envault"],
    install_requires=["Click==7.0", "requests==2.21.0"],
    entry_points="""
        [console_scripts]
        envault=envault.cli:cli
    """,
    url="https://github.com/pratishshr/envault",
)

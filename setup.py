from setuptools import setup, find_packages

import envault

setup(
    name="envault",
    version=envault.__version__,
    author='Pratish Shrestha',
    author_email='pratishshr@gmail.com',
    packages=find_packages(),
    description='A simple CLI tool to run processes with secrets from HashiCorp Vault.',
    py_modules=["envault"],
    install_requires=["Click==7.0", "requests==2.21.0"],
    entry_points="""
        [console_scripts]
        envault=envault.cli:cli
    """,
    project_urls={
      'Source': 'https://github.com/pratishshr/envault',
      'Documentation': 'https://github.com/pratishshr/envault/blob/master/README.md',
    },
)

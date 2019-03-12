from setuptools import setup, find_packages

setup(
    name="envault",
    version=0.1,
    packages=find_packages(),
    py_modules=["envault"],
    install_requires=["Click==7.0", "requests==2.21.0"],
    entry_points="""
        [console_scripts]
        envault=envault.cli:cli
    """,
)

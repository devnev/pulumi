# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "0.0.0"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "plant Pulumi Package - Development Version"


setup(name='pulumi_plant',
      python_requires='>=3.8',
      version=VERSION,
      long_description=readme(),
      long_description_content_type='text/markdown',
      packages=find_packages(),
      package_data={
          'pulumi_plant': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi',
          'semver>=2.8.1'
      ],
      zip_safe=False)

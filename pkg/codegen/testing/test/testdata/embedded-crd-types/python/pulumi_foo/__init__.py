# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .component import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_foo.crd_k8s_amazonaws_com as __crd_k8s_amazonaws_com
    crd_k8s_amazonaws_com = __crd_k8s_amazonaws_com
else:
    crd_k8s_amazonaws_com = _utilities.lazy_import('pulumi_foo.crd_k8s_amazonaws_com')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "foo",
  "mod": "index",
  "fqn": "pulumi_foo",
  "classes": {
   "foo:index:Component": "Component"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "foo",
  "token": "pulumi:providers:foo",
  "fqn": "pulumi_foo",
  "class": "Provider"
 }
]
"""
)

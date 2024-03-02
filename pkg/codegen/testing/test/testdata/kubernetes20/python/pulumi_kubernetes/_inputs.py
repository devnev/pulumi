# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'KubeClientSettingsArgs',
]

@pulumi.input_type
class KubeClientSettingsArgs:
    def __init__(__self__, *,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Options for tuning the Kubernetes client used by a Provider.
        :param pulumi.Input[int] timeout: Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
        """
        if timeout is None:
            timeout = _utilities.get_env_int('PULUMI_K8S_CLIENT_TIMEOUT')
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)



# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AviatrixMicrosegPolicyListArgs', 'AviatrixMicrosegPolicyList']

@pulumi.input_type
class AviatrixMicrosegPolicyListArgs:
    def __init__(__self__, *,
                 policies: pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]]):
        """
        The set of arguments for constructing a AviatrixMicrosegPolicyList resource.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]] policies: List of micro-segmentation policies.
        """
        pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]]:
        """
        List of micro-segmentation policies.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _AviatrixMicrosegPolicyListState:
    def __init__(__self__, *,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]]] = None):
        """
        Input properties used for looking up and filtering AviatrixMicrosegPolicyList resources.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]] policies: List of micro-segmentation policies.
        """
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]]]:
        """
        List of micro-segmentation policies.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixMicrosegPolicyListPolicyArgs']]]]):
        pulumi.set(self, "policies", value)


class AviatrixMicrosegPolicyList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixMicrosegPolicyListPolicyArgs']]]]] = None,
                 __props__=None):
        """
        Create a AviatrixMicrosegPolicyList resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixMicrosegPolicyListPolicyArgs']]]] policies: List of micro-segmentation policies.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixMicrosegPolicyListArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixMicrosegPolicyList resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixMicrosegPolicyListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixMicrosegPolicyListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixMicrosegPolicyListPolicyArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixMicrosegPolicyListArgs.__new__(AviatrixMicrosegPolicyListArgs)

            if policies is None and not opts.urn:
                raise TypeError("Missing required property 'policies'")
            __props__.__dict__["policies"] = policies
        super(AviatrixMicrosegPolicyList, __self__).__init__(
            'aviatrix:index/aviatrixMicrosegPolicyList:AviatrixMicrosegPolicyList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixMicrosegPolicyListPolicyArgs']]]]] = None) -> 'AviatrixMicrosegPolicyList':
        """
        Get an existing AviatrixMicrosegPolicyList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixMicrosegPolicyListPolicyArgs']]]] policies: List of micro-segmentation policies.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixMicrosegPolicyListState.__new__(_AviatrixMicrosegPolicyListState)

        __props__.__dict__["policies"] = policies
        return AviatrixMicrosegPolicyList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence['outputs.AviatrixMicrosegPolicyListPolicy']]:
        """
        List of micro-segmentation policies.
        """
        return pulumi.get(self, "policies")

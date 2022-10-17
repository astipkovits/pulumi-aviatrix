# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixSegmentationNetworkDomainConnectionPolicyArgs', 'AviatrixSegmentationNetworkDomainConnectionPolicy']

@pulumi.input_type
class AviatrixSegmentationNetworkDomainConnectionPolicyArgs:
    def __init__(__self__, *,
                 domain_name1: pulumi.Input[str],
                 domain_name2: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixSegmentationNetworkDomainConnectionPolicy resource.
        :param pulumi.Input[str] domain_name1: Name of network domain that will be connected to domain 2.
        :param pulumi.Input[str] domain_name2: Name of network domain that will be connected to domain 1.
        """
        pulumi.set(__self__, "domain_name1", domain_name1)
        pulumi.set(__self__, "domain_name2", domain_name2)

    @property
    @pulumi.getter(name="domainName1")
    def domain_name1(self) -> pulumi.Input[str]:
        """
        Name of network domain that will be connected to domain 2.
        """
        return pulumi.get(self, "domain_name1")

    @domain_name1.setter
    def domain_name1(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name1", value)

    @property
    @pulumi.getter(name="domainName2")
    def domain_name2(self) -> pulumi.Input[str]:
        """
        Name of network domain that will be connected to domain 1.
        """
        return pulumi.get(self, "domain_name2")

    @domain_name2.setter
    def domain_name2(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name2", value)


@pulumi.input_type
class _AviatrixSegmentationNetworkDomainConnectionPolicyState:
    def __init__(__self__, *,
                 domain_name1: Optional[pulumi.Input[str]] = None,
                 domain_name2: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixSegmentationNetworkDomainConnectionPolicy resources.
        :param pulumi.Input[str] domain_name1: Name of network domain that will be connected to domain 2.
        :param pulumi.Input[str] domain_name2: Name of network domain that will be connected to domain 1.
        """
        if domain_name1 is not None:
            pulumi.set(__self__, "domain_name1", domain_name1)
        if domain_name2 is not None:
            pulumi.set(__self__, "domain_name2", domain_name2)

    @property
    @pulumi.getter(name="domainName1")
    def domain_name1(self) -> Optional[pulumi.Input[str]]:
        """
        Name of network domain that will be connected to domain 2.
        """
        return pulumi.get(self, "domain_name1")

    @domain_name1.setter
    def domain_name1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name1", value)

    @property
    @pulumi.getter(name="domainName2")
    def domain_name2(self) -> Optional[pulumi.Input[str]]:
        """
        Name of network domain that will be connected to domain 1.
        """
        return pulumi.get(self, "domain_name2")

    @domain_name2.setter
    def domain_name2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name2", value)


class AviatrixSegmentationNetworkDomainConnectionPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name1: Optional[pulumi.Input[str]] = None,
                 domain_name2: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixSegmentationNetworkDomainConnectionPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name1: Name of network domain that will be connected to domain 2.
        :param pulumi.Input[str] domain_name2: Name of network domain that will be connected to domain 1.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixSegmentationNetworkDomainConnectionPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixSegmentationNetworkDomainConnectionPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixSegmentationNetworkDomainConnectionPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixSegmentationNetworkDomainConnectionPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name1: Optional[pulumi.Input[str]] = None,
                 domain_name2: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixSegmentationNetworkDomainConnectionPolicyArgs.__new__(AviatrixSegmentationNetworkDomainConnectionPolicyArgs)

            if domain_name1 is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name1'")
            __props__.__dict__["domain_name1"] = domain_name1
            if domain_name2 is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name2'")
            __props__.__dict__["domain_name2"] = domain_name2
        super(AviatrixSegmentationNetworkDomainConnectionPolicy, __self__).__init__(
            'aviatrix:index/aviatrixSegmentationNetworkDomainConnectionPolicy:AviatrixSegmentationNetworkDomainConnectionPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name1: Optional[pulumi.Input[str]] = None,
            domain_name2: Optional[pulumi.Input[str]] = None) -> 'AviatrixSegmentationNetworkDomainConnectionPolicy':
        """
        Get an existing AviatrixSegmentationNetworkDomainConnectionPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name1: Name of network domain that will be connected to domain 2.
        :param pulumi.Input[str] domain_name2: Name of network domain that will be connected to domain 1.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixSegmentationNetworkDomainConnectionPolicyState.__new__(_AviatrixSegmentationNetworkDomainConnectionPolicyState)

        __props__.__dict__["domain_name1"] = domain_name1
        __props__.__dict__["domain_name2"] = domain_name2
        return AviatrixSegmentationNetworkDomainConnectionPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName1")
    def domain_name1(self) -> pulumi.Output[str]:
        """
        Name of network domain that will be connected to domain 2.
        """
        return pulumi.get(self, "domain_name1")

    @property
    @pulumi.getter(name="domainName2")
    def domain_name2(self) -> pulumi.Output[str]:
        """
        Name of network domain that will be connected to domain 1.
        """
        return pulumi.get(self, "domain_name2")

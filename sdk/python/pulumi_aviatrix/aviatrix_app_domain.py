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

__all__ = ['AviatrixAppDomainArgs', 'AviatrixAppDomain']

@pulumi.input_type
class AviatrixAppDomainArgs:
    def __init__(__self__, *,
                 selector: pulumi.Input['AviatrixAppDomainSelectorArgs'],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixAppDomain resource.
        :param pulumi.Input['AviatrixAppDomainSelectorArgs'] selector: Block containing match expressions to filter the App Domain.
        :param pulumi.Input[str] name: Name of the App Domain.
        """
        pulumi.set(__self__, "selector", selector)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def selector(self) -> pulumi.Input['AviatrixAppDomainSelectorArgs']:
        """
        Block containing match expressions to filter the App Domain.
        """
        return pulumi.get(self, "selector")

    @selector.setter
    def selector(self, value: pulumi.Input['AviatrixAppDomainSelectorArgs']):
        pulumi.set(self, "selector", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the App Domain.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AviatrixAppDomainState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 selector: Optional[pulumi.Input['AviatrixAppDomainSelectorArgs']] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixAppDomain resources.
        :param pulumi.Input[str] name: Name of the App Domain.
        :param pulumi.Input['AviatrixAppDomainSelectorArgs'] selector: Block containing match expressions to filter the App Domain.
        :param pulumi.Input[str] uuid: UUID of the App Domain.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if selector is not None:
            pulumi.set(__self__, "selector", selector)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the App Domain.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def selector(self) -> Optional[pulumi.Input['AviatrixAppDomainSelectorArgs']]:
        """
        Block containing match expressions to filter the App Domain.
        """
        return pulumi.get(self, "selector")

    @selector.setter
    def selector(self, value: Optional[pulumi.Input['AviatrixAppDomainSelectorArgs']]):
        pulumi.set(self, "selector", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the App Domain.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class AviatrixAppDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 selector: Optional[pulumi.Input[pulumi.InputType['AviatrixAppDomainSelectorArgs']]] = None,
                 __props__=None):
        """
        !> **WARNING** **aviatrix_app_domain** is part of the Micro-segmentation private preview feature for R2.22.0. If you wish to enable a private preview mode feature, please contact your sales representative or Aviatrix Support.
        The **aviatrix_app_domain** resource handles the creation and management of App Domains. Available as of Provider R2.22.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix App Domain
        test_app_domain_ip = aviatrix.AviatrixAppDomain("testAppDomainIp", selector=aviatrix.AviatrixAppDomainSelectorArgs(
            match_expressions=[
                aviatrix.AviatrixAppDomainSelectorMatchExpressionArgs(
                    account_name="devops",
                    region="us-west-2",
                    tags={
                        "k3": "v3",
                    },
                    type="vm",
                ),
                aviatrix.AviatrixAppDomainSelectorMatchExpressionArgs(
                    cidr="10.0.0.0/16",
                ),
            ],
        ))
        ```

        ## Import

        **aviatrix_app_domain** can be imported using the `uuid`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixAppDomain:AviatrixAppDomain test 41984f8b-5a37-4272-89b3-57c79e9ff77c
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the App Domain.
        :param pulumi.Input[pulumi.InputType['AviatrixAppDomainSelectorArgs']] selector: Block containing match expressions to filter the App Domain.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixAppDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        !> **WARNING** **aviatrix_app_domain** is part of the Micro-segmentation private preview feature for R2.22.0. If you wish to enable a private preview mode feature, please contact your sales representative or Aviatrix Support.
        The **aviatrix_app_domain** resource handles the creation and management of App Domains. Available as of Provider R2.22.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix App Domain
        test_app_domain_ip = aviatrix.AviatrixAppDomain("testAppDomainIp", selector=aviatrix.AviatrixAppDomainSelectorArgs(
            match_expressions=[
                aviatrix.AviatrixAppDomainSelectorMatchExpressionArgs(
                    account_name="devops",
                    region="us-west-2",
                    tags={
                        "k3": "v3",
                    },
                    type="vm",
                ),
                aviatrix.AviatrixAppDomainSelectorMatchExpressionArgs(
                    cidr="10.0.0.0/16",
                ),
            ],
        ))
        ```

        ## Import

        **aviatrix_app_domain** can be imported using the `uuid`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixAppDomain:AviatrixAppDomain test 41984f8b-5a37-4272-89b3-57c79e9ff77c
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixAppDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixAppDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 selector: Optional[pulumi.Input[pulumi.InputType['AviatrixAppDomainSelectorArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixAppDomainArgs.__new__(AviatrixAppDomainArgs)

            __props__.__dict__["name"] = name
            if selector is None and not opts.urn:
                raise TypeError("Missing required property 'selector'")
            __props__.__dict__["selector"] = selector
            __props__.__dict__["uuid"] = None
        super(AviatrixAppDomain, __self__).__init__(
            'aviatrix:index/aviatrixAppDomain:AviatrixAppDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            selector: Optional[pulumi.Input[pulumi.InputType['AviatrixAppDomainSelectorArgs']]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'AviatrixAppDomain':
        """
        Get an existing AviatrixAppDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the App Domain.
        :param pulumi.Input[pulumi.InputType['AviatrixAppDomainSelectorArgs']] selector: Block containing match expressions to filter the App Domain.
        :param pulumi.Input[str] uuid: UUID of the App Domain.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixAppDomainState.__new__(_AviatrixAppDomainState)

        __props__.__dict__["name"] = name
        __props__.__dict__["selector"] = selector
        __props__.__dict__["uuid"] = uuid
        return AviatrixAppDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the App Domain.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def selector(self) -> pulumi.Output['outputs.AviatrixAppDomainSelector']:
        """
        Block containing match expressions to filter the App Domain.
        """
        return pulumi.get(self, "selector")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        UUID of the App Domain.
        """
        return pulumi.get(self, "uuid")


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

__all__ = ['AviatrixFqdnArgs', 'AviatrixFqdn']

@pulumi.input_type
class AviatrixFqdnArgs:
    def __init__(__self__, *,
                 fqdn_tag: pulumi.Input[str],
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]]] = None,
                 fqdn_enabled: Optional[pulumi.Input[bool]] = None,
                 fqdn_mode: Optional[pulumi.Input[str]] = None,
                 gw_filter_tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]]] = None,
                 manage_domain_names: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AviatrixFqdn resource.
        :param pulumi.Input[str] fqdn_tag: FQDN Filter Tag Name.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]] domain_names: A list of one or more domain names/tag rules.
        :param pulumi.Input[bool] fqdn_enabled: FQDN Filter Tag Status. Valid values: true or false.
        :param pulumi.Input[str] fqdn_mode: Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]] gw_filter_tag_lists: A list of gateways to attach to the specific tag.
        :param pulumi.Input[bool] manage_domain_names: Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
               resources.
        """
        pulumi.set(__self__, "fqdn_tag", fqdn_tag)
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if fqdn_enabled is not None:
            pulumi.set(__self__, "fqdn_enabled", fqdn_enabled)
        if fqdn_mode is not None:
            pulumi.set(__self__, "fqdn_mode", fqdn_mode)
        if gw_filter_tag_lists is not None:
            pulumi.set(__self__, "gw_filter_tag_lists", gw_filter_tag_lists)
        if manage_domain_names is not None:
            pulumi.set(__self__, "manage_domain_names", manage_domain_names)

    @property
    @pulumi.getter(name="fqdnTag")
    def fqdn_tag(self) -> pulumi.Input[str]:
        """
        FQDN Filter Tag Name.
        """
        return pulumi.get(self, "fqdn_tag")

    @fqdn_tag.setter
    def fqdn_tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "fqdn_tag", value)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]]]:
        """
        A list of one or more domain names/tag rules.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="fqdnEnabled")
    def fqdn_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        FQDN Filter Tag Status. Valid values: true or false.
        """
        return pulumi.get(self, "fqdn_enabled")

    @fqdn_enabled.setter
    def fqdn_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fqdn_enabled", value)

    @property
    @pulumi.getter(name="fqdnMode")
    def fqdn_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        """
        return pulumi.get(self, "fqdn_mode")

    @fqdn_mode.setter
    def fqdn_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn_mode", value)

    @property
    @pulumi.getter(name="gwFilterTagLists")
    def gw_filter_tag_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]]]:
        """
        A list of gateways to attach to the specific tag.
        """
        return pulumi.get(self, "gw_filter_tag_lists")

    @gw_filter_tag_lists.setter
    def gw_filter_tag_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]]]):
        pulumi.set(self, "gw_filter_tag_lists", value)

    @property
    @pulumi.getter(name="manageDomainNames")
    def manage_domain_names(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
        resources.
        """
        return pulumi.get(self, "manage_domain_names")

    @manage_domain_names.setter
    def manage_domain_names(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage_domain_names", value)


@pulumi.input_type
class _AviatrixFqdnState:
    def __init__(__self__, *,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]]] = None,
                 fqdn_enabled: Optional[pulumi.Input[bool]] = None,
                 fqdn_mode: Optional[pulumi.Input[str]] = None,
                 fqdn_tag: Optional[pulumi.Input[str]] = None,
                 gw_filter_tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]]] = None,
                 manage_domain_names: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AviatrixFqdn resources.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]] domain_names: A list of one or more domain names/tag rules.
        :param pulumi.Input[bool] fqdn_enabled: FQDN Filter Tag Status. Valid values: true or false.
        :param pulumi.Input[str] fqdn_mode: Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        :param pulumi.Input[str] fqdn_tag: FQDN Filter Tag Name.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]] gw_filter_tag_lists: A list of gateways to attach to the specific tag.
        :param pulumi.Input[bool] manage_domain_names: Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
               resources.
        """
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if fqdn_enabled is not None:
            pulumi.set(__self__, "fqdn_enabled", fqdn_enabled)
        if fqdn_mode is not None:
            pulumi.set(__self__, "fqdn_mode", fqdn_mode)
        if fqdn_tag is not None:
            pulumi.set(__self__, "fqdn_tag", fqdn_tag)
        if gw_filter_tag_lists is not None:
            pulumi.set(__self__, "gw_filter_tag_lists", gw_filter_tag_lists)
        if manage_domain_names is not None:
            pulumi.set(__self__, "manage_domain_names", manage_domain_names)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]]]:
        """
        A list of one or more domain names/tag rules.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnDomainNameArgs']]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="fqdnEnabled")
    def fqdn_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        FQDN Filter Tag Status. Valid values: true or false.
        """
        return pulumi.get(self, "fqdn_enabled")

    @fqdn_enabled.setter
    def fqdn_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fqdn_enabled", value)

    @property
    @pulumi.getter(name="fqdnMode")
    def fqdn_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        """
        return pulumi.get(self, "fqdn_mode")

    @fqdn_mode.setter
    def fqdn_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn_mode", value)

    @property
    @pulumi.getter(name="fqdnTag")
    def fqdn_tag(self) -> Optional[pulumi.Input[str]]:
        """
        FQDN Filter Tag Name.
        """
        return pulumi.get(self, "fqdn_tag")

    @fqdn_tag.setter
    def fqdn_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn_tag", value)

    @property
    @pulumi.getter(name="gwFilterTagLists")
    def gw_filter_tag_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]]]:
        """
        A list of gateways to attach to the specific tag.
        """
        return pulumi.get(self, "gw_filter_tag_lists")

    @gw_filter_tag_lists.setter
    def gw_filter_tag_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFqdnGwFilterTagListArgs']]]]):
        pulumi.set(self, "gw_filter_tag_lists", value)

    @property
    @pulumi.getter(name="manageDomainNames")
    def manage_domain_names(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
        resources.
        """
        return pulumi.get(self, "manage_domain_names")

    @manage_domain_names.setter
    def manage_domain_names(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage_domain_names", value)


class AviatrixFqdn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnDomainNameArgs']]]]] = None,
                 fqdn_enabled: Optional[pulumi.Input[bool]] = None,
                 fqdn_mode: Optional[pulumi.Input[str]] = None,
                 fqdn_tag: Optional[pulumi.Input[str]] = None,
                 gw_filter_tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnGwFilterTagListArgs']]]]] = None,
                 manage_domain_names: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a AviatrixFqdn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnDomainNameArgs']]]] domain_names: A list of one or more domain names/tag rules.
        :param pulumi.Input[bool] fqdn_enabled: FQDN Filter Tag Status. Valid values: true or false.
        :param pulumi.Input[str] fqdn_mode: Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        :param pulumi.Input[str] fqdn_tag: FQDN Filter Tag Name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnGwFilterTagListArgs']]]] gw_filter_tag_lists: A list of gateways to attach to the specific tag.
        :param pulumi.Input[bool] manage_domain_names: Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
               resources.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixFqdnArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixFqdn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixFqdnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixFqdnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnDomainNameArgs']]]]] = None,
                 fqdn_enabled: Optional[pulumi.Input[bool]] = None,
                 fqdn_mode: Optional[pulumi.Input[str]] = None,
                 fqdn_tag: Optional[pulumi.Input[str]] = None,
                 gw_filter_tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnGwFilterTagListArgs']]]]] = None,
                 manage_domain_names: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixFqdnArgs.__new__(AviatrixFqdnArgs)

            __props__.__dict__["domain_names"] = domain_names
            __props__.__dict__["fqdn_enabled"] = fqdn_enabled
            __props__.__dict__["fqdn_mode"] = fqdn_mode
            if fqdn_tag is None and not opts.urn:
                raise TypeError("Missing required property 'fqdn_tag'")
            __props__.__dict__["fqdn_tag"] = fqdn_tag
            __props__.__dict__["gw_filter_tag_lists"] = gw_filter_tag_lists
            __props__.__dict__["manage_domain_names"] = manage_domain_names
        super(AviatrixFqdn, __self__).__init__(
            'aviatrix:index/aviatrixFqdn:AviatrixFqdn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnDomainNameArgs']]]]] = None,
            fqdn_enabled: Optional[pulumi.Input[bool]] = None,
            fqdn_mode: Optional[pulumi.Input[str]] = None,
            fqdn_tag: Optional[pulumi.Input[str]] = None,
            gw_filter_tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnGwFilterTagListArgs']]]]] = None,
            manage_domain_names: Optional[pulumi.Input[bool]] = None) -> 'AviatrixFqdn':
        """
        Get an existing AviatrixFqdn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnDomainNameArgs']]]] domain_names: A list of one or more domain names/tag rules.
        :param pulumi.Input[bool] fqdn_enabled: FQDN Filter Tag Status. Valid values: true or false.
        :param pulumi.Input[str] fqdn_mode: Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        :param pulumi.Input[str] fqdn_tag: FQDN Filter Tag Name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFqdnGwFilterTagListArgs']]]] gw_filter_tag_lists: A list of gateways to attach to the specific tag.
        :param pulumi.Input[bool] manage_domain_names: Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
               resources.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixFqdnState.__new__(_AviatrixFqdnState)

        __props__.__dict__["domain_names"] = domain_names
        __props__.__dict__["fqdn_enabled"] = fqdn_enabled
        __props__.__dict__["fqdn_mode"] = fqdn_mode
        __props__.__dict__["fqdn_tag"] = fqdn_tag
        __props__.__dict__["gw_filter_tag_lists"] = gw_filter_tag_lists
        __props__.__dict__["manage_domain_names"] = manage_domain_names
        return AviatrixFqdn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Output[Optional[Sequence['outputs.AviatrixFqdnDomainName']]]:
        """
        A list of one or more domain names/tag rules.
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="fqdnEnabled")
    def fqdn_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        FQDN Filter Tag Status. Valid values: true or false.
        """
        return pulumi.get(self, "fqdn_enabled")

    @property
    @pulumi.getter(name="fqdnMode")
    def fqdn_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        """
        return pulumi.get(self, "fqdn_mode")

    @property
    @pulumi.getter(name="fqdnTag")
    def fqdn_tag(self) -> pulumi.Output[str]:
        """
        FQDN Filter Tag Name.
        """
        return pulumi.get(self, "fqdn_tag")

    @property
    @pulumi.getter(name="gwFilterTagLists")
    def gw_filter_tag_lists(self) -> pulumi.Output[Optional[Sequence['outputs.AviatrixFqdnGwFilterTagList']]]:
        """
        A list of gateways to attach to the specific tag.
        """
        return pulumi.get(self, "gw_filter_tag_lists")

    @property
    @pulumi.getter(name="manageDomainNames")
    def manage_domain_names(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
        resources.
        """
        return pulumi.get(self, "manage_domain_names")

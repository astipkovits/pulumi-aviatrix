# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixFqdnTagRuleArgs', 'AviatrixFqdnTagRule']

@pulumi.input_type
class AviatrixFqdnTagRuleArgs:
    def __init__(__self__, *,
                 fqdn: pulumi.Input[str],
                 fqdn_tag_name: pulumi.Input[str],
                 port: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixFqdnTagRule resource.
        :param pulumi.Input[str] fqdn: FQDN.
        :param pulumi.Input[str] fqdn_tag_name: FQDN Filter Tag Name to attach this domain.
        :param pulumi.Input[str] port: Port.
        :param pulumi.Input[str] protocol: Protocol.
        :param pulumi.Input[str] action: What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
               Policy' if no value is provided.
        """
        pulumi.set(__self__, "fqdn", fqdn)
        pulumi.set(__self__, "fqdn_tag_name", fqdn_tag_name)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        if action is not None:
            pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Input[str]:
        """
        FQDN.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: pulumi.Input[str]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="fqdnTagName")
    def fqdn_tag_name(self) -> pulumi.Input[str]:
        """
        FQDN Filter Tag Name to attach this domain.
        """
        return pulumi.get(self, "fqdn_tag_name")

    @fqdn_tag_name.setter
    def fqdn_tag_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "fqdn_tag_name", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[str]:
        """
        Port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[str]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        Protocol.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
        Policy' if no value is provided.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)


@pulumi.input_type
class _AviatrixFqdnTagRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 fqdn_tag_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixFqdnTagRule resources.
        :param pulumi.Input[str] action: What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
               Policy' if no value is provided.
        :param pulumi.Input[str] fqdn: FQDN.
        :param pulumi.Input[str] fqdn_tag_name: FQDN Filter Tag Name to attach this domain.
        :param pulumi.Input[str] port: Port.
        :param pulumi.Input[str] protocol: Protocol.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if fqdn_tag_name is not None:
            pulumi.set(__self__, "fqdn_tag_name", fqdn_tag_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
        Policy' if no value is provided.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        FQDN.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="fqdnTagName")
    def fqdn_tag_name(self) -> Optional[pulumi.Input[str]]:
        """
        FQDN Filter Tag Name to attach this domain.
        """
        return pulumi.get(self, "fqdn_tag_name")

    @fqdn_tag_name.setter
    def fqdn_tag_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn_tag_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


class AviatrixFqdnTagRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 fqdn_tag_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixFqdnTagRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
               Policy' if no value is provided.
        :param pulumi.Input[str] fqdn: FQDN.
        :param pulumi.Input[str] fqdn_tag_name: FQDN Filter Tag Name to attach this domain.
        :param pulumi.Input[str] port: Port.
        :param pulumi.Input[str] protocol: Protocol.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixFqdnTagRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixFqdnTagRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixFqdnTagRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixFqdnTagRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 fqdn_tag_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixFqdnTagRuleArgs.__new__(AviatrixFqdnTagRuleArgs)

            __props__.__dict__["action"] = action
            if fqdn is None and not opts.urn:
                raise TypeError("Missing required property 'fqdn'")
            __props__.__dict__["fqdn"] = fqdn
            if fqdn_tag_name is None and not opts.urn:
                raise TypeError("Missing required property 'fqdn_tag_name'")
            __props__.__dict__["fqdn_tag_name"] = fqdn_tag_name
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
        super(AviatrixFqdnTagRule, __self__).__init__(
            'aviatrix:index/aviatrixFqdnTagRule:AviatrixFqdnTagRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            fqdn_tag_name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None) -> 'AviatrixFqdnTagRule':
        """
        Get an existing AviatrixFqdnTagRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
               Policy' if no value is provided.
        :param pulumi.Input[str] fqdn: FQDN.
        :param pulumi.Input[str] fqdn_tag_name: FQDN Filter Tag Name to attach this domain.
        :param pulumi.Input[str] port: Port.
        :param pulumi.Input[str] protocol: Protocol.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixFqdnTagRuleState.__new__(_AviatrixFqdnTagRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["fqdn_tag_name"] = fqdn_tag_name
        __props__.__dict__["port"] = port
        __props__.__dict__["protocol"] = protocol
        return AviatrixFqdnTagRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[str]]:
        """
        What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base
        Policy' if no value is provided.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        FQDN.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="fqdnTagName")
    def fqdn_tag_name(self) -> pulumi.Output[str]:
        """
        FQDN Filter Tag Name to attach this domain.
        """
        return pulumi.get(self, "fqdn_tag_name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        Port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Protocol.
        """
        return pulumi.get(self, "protocol")

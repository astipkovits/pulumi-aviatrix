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

__all__ = ['AviatrixFirewallArgs', 'AviatrixFirewall']

@pulumi.input_type
class AviatrixFirewallArgs:
    def __init__(__self__, *,
                 gw_name: pulumi.Input[str],
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]] = None):
        """
        The set of arguments for constructing a AviatrixFirewall resource.
        :param pulumi.Input[str] gw_name: Gateway name to attach firewall policy to.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]] policies: New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        pulumi.set(__self__, "gw_name", gw_name)
        if base_log_enabled is not None:
            pulumi.set(__self__, "base_log_enabled", base_log_enabled)
        if base_policy is not None:
            pulumi.set(__self__, "base_policy", base_policy)
        if manage_firewall_policies is not None:
            pulumi.set(__self__, "manage_firewall_policies", manage_firewall_policies)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> pulumi.Input[str]:
        """
        Gateway name to attach firewall policy to.
        """
        return pulumi.get(self, "gw_name")

    @gw_name.setter
    def gw_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "gw_name", value)

    @property
    @pulumi.getter(name="baseLogEnabled")
    def base_log_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        """
        return pulumi.get(self, "base_log_enabled")

    @base_log_enabled.setter
    def base_log_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "base_log_enabled", value)

    @property
    @pulumi.getter(name="basePolicy")
    def base_policy(self) -> Optional[pulumi.Input[str]]:
        """
        New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        """
        return pulumi.get(self, "base_policy")

    @base_policy.setter
    def base_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_policy", value)

    @property
    @pulumi.getter(name="manageFirewallPolicies")
    def manage_firewall_policies(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        """
        return pulumi.get(self, "manage_firewall_policies")

    @manage_firewall_policies.setter
    def manage_firewall_policies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage_firewall_policies", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]:
        """
        New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _AviatrixFirewallState:
    def __init__(__self__, *,
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]] = None):
        """
        Input properties used for looking up and filtering AviatrixFirewall resources.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        :param pulumi.Input[str] gw_name: Gateway name to attach firewall policy to.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]] policies: New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        if base_log_enabled is not None:
            pulumi.set(__self__, "base_log_enabled", base_log_enabled)
        if base_policy is not None:
            pulumi.set(__self__, "base_policy", base_policy)
        if gw_name is not None:
            pulumi.set(__self__, "gw_name", gw_name)
        if manage_firewall_policies is not None:
            pulumi.set(__self__, "manage_firewall_policies", manage_firewall_policies)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="baseLogEnabled")
    def base_log_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        """
        return pulumi.get(self, "base_log_enabled")

    @base_log_enabled.setter
    def base_log_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "base_log_enabled", value)

    @property
    @pulumi.getter(name="basePolicy")
    def base_policy(self) -> Optional[pulumi.Input[str]]:
        """
        New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        """
        return pulumi.get(self, "base_policy")

    @base_policy.setter
    def base_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_policy", value)

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway name to attach firewall policy to.
        """
        return pulumi.get(self, "gw_name")

    @gw_name.setter
    def gw_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gw_name", value)

    @property
    @pulumi.getter(name="manageFirewallPolicies")
    def manage_firewall_policies(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        """
        return pulumi.get(self, "manage_firewall_policies")

    @manage_firewall_policies.setter
    def manage_firewall_policies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage_firewall_policies", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]:
        """
        New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]):
        pulumi.set(self, "policies", value)


class AviatrixFirewall(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Firewall
        stateful_firewall1 = aviatrix.AviatrixFirewall("statefulFirewall1",
            base_log_enabled=True,
            base_policy="allow-all",
            gw_name="gateway-1",
            manage_firewall_policies=False)
        ```

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Firewall with in-line rules
        stateful_firewall1 = aviatrix.AviatrixFirewall("statefulFirewall1",
            gw_name="gateway-1",
            base_policy="allow-all",
            base_log_enabled=True,
            policies=[
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="all",
                    src_ip="10.17.0.224/32",
                    log_enabled=True,
                    dst_ip="10.12.0.172/32",
                    action="force-drop",
                    port="0:65535",
                    description="first_policy",
                ),
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="tcp",
                    src_ip="10.16.0.224/32",
                    log_enabled=False,
                    dst_ip="10.12.1.172/32",
                    action="force-drop",
                    port="325",
                    description="second_policy",
                ),
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="udp",
                    src_ip="10.14.0.225/32",
                    log_enabled=False,
                    dst_ip="10.13.1.173/32",
                    action="deny",
                    port="325",
                    description="third_policy",
                ),
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="tcp",
                    src_ip=aviatrix_firewall_tag["test"]["firewall_tag"],
                    log_enabled=False,
                    dst_ip="10.13.1.173/32",
                    action="deny",
                    port="325",
                    description="fourth_policy",
                ),
            ])
        ```

        ## Import

        **firewall** can be imported using the `gw_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixFirewall:AviatrixFirewall test gw_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        :param pulumi.Input[str] gw_name: Gateway name to attach firewall policy to.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]] policies: New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixFirewallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Firewall
        stateful_firewall1 = aviatrix.AviatrixFirewall("statefulFirewall1",
            base_log_enabled=True,
            base_policy="allow-all",
            gw_name="gateway-1",
            manage_firewall_policies=False)
        ```

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Firewall with in-line rules
        stateful_firewall1 = aviatrix.AviatrixFirewall("statefulFirewall1",
            gw_name="gateway-1",
            base_policy="allow-all",
            base_log_enabled=True,
            policies=[
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="all",
                    src_ip="10.17.0.224/32",
                    log_enabled=True,
                    dst_ip="10.12.0.172/32",
                    action="force-drop",
                    port="0:65535",
                    description="first_policy",
                ),
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="tcp",
                    src_ip="10.16.0.224/32",
                    log_enabled=False,
                    dst_ip="10.12.1.172/32",
                    action="force-drop",
                    port="325",
                    description="second_policy",
                ),
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="udp",
                    src_ip="10.14.0.225/32",
                    log_enabled=False,
                    dst_ip="10.13.1.173/32",
                    action="deny",
                    port="325",
                    description="third_policy",
                ),
                aviatrix.AviatrixFirewallPolicyArgs(
                    protocol="tcp",
                    src_ip=aviatrix_firewall_tag["test"]["firewall_tag"],
                    log_enabled=False,
                    dst_ip="10.13.1.173/32",
                    action="deny",
                    port="325",
                    description="fourth_policy",
                ),
            ])
        ```

        ## Import

        **firewall** can be imported using the `gw_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixFirewall:AviatrixFirewall test gw_name
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixFirewallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixFirewallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixFirewallArgs.__new__(AviatrixFirewallArgs)

            __props__.__dict__["base_log_enabled"] = base_log_enabled
            __props__.__dict__["base_policy"] = base_policy
            if gw_name is None and not opts.urn:
                raise TypeError("Missing required property 'gw_name'")
            __props__.__dict__["gw_name"] = gw_name
            __props__.__dict__["manage_firewall_policies"] = manage_firewall_policies
            __props__.__dict__["policies"] = policies
        super(AviatrixFirewall, __self__).__init__(
            'aviatrix:index/aviatrixFirewall:AviatrixFirewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            base_log_enabled: Optional[pulumi.Input[bool]] = None,
            base_policy: Optional[pulumi.Input[str]] = None,
            gw_name: Optional[pulumi.Input[str]] = None,
            manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]]] = None) -> 'AviatrixFirewall':
        """
        Get an existing AviatrixFirewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        :param pulumi.Input[str] gw_name: Gateway name to attach firewall policy to.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]] policies: New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixFirewallState.__new__(_AviatrixFirewallState)

        __props__.__dict__["base_log_enabled"] = base_log_enabled
        __props__.__dict__["base_policy"] = base_policy
        __props__.__dict__["gw_name"] = gw_name
        __props__.__dict__["manage_firewall_policies"] = manage_firewall_policies
        __props__.__dict__["policies"] = policies
        return AviatrixFirewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="baseLogEnabled")
    def base_log_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether enable logging or not. Valid Values: true, false. Default value: false.
        """
        return pulumi.get(self, "base_log_enabled")

    @property
    @pulumi.getter(name="basePolicy")
    def base_policy(self) -> pulumi.Output[Optional[str]]:
        """
        New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"
        """
        return pulumi.get(self, "base_policy")

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> pulumi.Output[str]:
        """
        Gateway name to attach firewall policy to.
        """
        return pulumi.get(self, "gw_name")

    @property
    @pulumi.getter(name="manageFirewallPolicies")
    def manage_firewall_policies(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable to manage firewall policies via in-line rules. If false, policies must be managed using `AviatrixFirewallPolicy` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.
        """
        return pulumi.get(self, "manage_firewall_policies")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence['outputs.AviatrixFirewallPolicy']]]:
        """
        New access policy for the gateway. Seven fields are required for each policy item: `src_ip`, `dst_ip`, `protocol`, `port`, `action`, `log_enabled` and `description`. No duplicate rules (with same `src_ip`, `dst_ip`, `protocol` and `port`) are allowed.
        """
        return pulumi.get(self, "policies")


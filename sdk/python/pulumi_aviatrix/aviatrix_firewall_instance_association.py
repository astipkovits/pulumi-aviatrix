# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixFirewallInstanceAssociationArgs', 'AviatrixFirewallInstanceAssociation']

@pulumi.input_type
class AviatrixFirewallInstanceAssociationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 attached: Optional[pulumi.Input[bool]] = None,
                 egress_interface: Optional[pulumi.Input[str]] = None,
                 firenet_gw_name: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 lan_interface: Optional[pulumi.Input[str]] = None,
                 management_interface: Optional[pulumi.Input[str]] = None,
                 vendor_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixFirewallInstanceAssociation resource.
        :param pulumi.Input[str] instance_id: ID of Firewall instance, or FQDN Gateway's gw_name.
        :param pulumi.Input[str] vpc_id: VPC ID.
        :param pulumi.Input[bool] attached: Switch to attach/detach firewall instance to/from fireNet.
        :param pulumi.Input[str] egress_interface: Egress interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] firenet_gw_name: Name of the gateway to launch the firewall instance.
        :param pulumi.Input[str] firewall_name: Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
               GCP
        :param pulumi.Input[str] lan_interface: Lan interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] management_interface: Management interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] vendor_type: Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
               'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if egress_interface is not None:
            pulumi.set(__self__, "egress_interface", egress_interface)
        if firenet_gw_name is not None:
            pulumi.set(__self__, "firenet_gw_name", firenet_gw_name)
        if firewall_name is not None:
            pulumi.set(__self__, "firewall_name", firewall_name)
        if lan_interface is not None:
            pulumi.set(__self__, "lan_interface", lan_interface)
        if management_interface is not None:
            pulumi.set(__self__, "management_interface", management_interface)
        if vendor_type is not None:
            pulumi.set(__self__, "vendor_type", vendor_type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of Firewall instance, or FQDN Gateway's gw_name.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def attached(self) -> Optional[pulumi.Input[bool]]:
        """
        Switch to attach/detach firewall instance to/from fireNet.
        """
        return pulumi.get(self, "attached")

    @attached.setter
    def attached(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "attached", value)

    @property
    @pulumi.getter(name="egressInterface")
    def egress_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Egress interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "egress_interface")

    @egress_interface.setter
    def egress_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "egress_interface", value)

    @property
    @pulumi.getter(name="firenetGwName")
    def firenet_gw_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the gateway to launch the firewall instance.
        """
        return pulumi.get(self, "firenet_gw_name")

    @firenet_gw_name.setter
    def firenet_gw_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firenet_gw_name", value)

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
        GCP
        """
        return pulumi.get(self, "firewall_name")

    @firewall_name.setter
    def firewall_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_name", value)

    @property
    @pulumi.getter(name="lanInterface")
    def lan_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Lan interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "lan_interface")

    @lan_interface.setter
    def lan_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lan_interface", value)

    @property
    @pulumi.getter(name="managementInterface")
    def management_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Management interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "management_interface")

    @management_interface.setter
    def management_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "management_interface", value)

    @property
    @pulumi.getter(name="vendorType")
    def vendor_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
        'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        """
        return pulumi.get(self, "vendor_type")

    @vendor_type.setter
    def vendor_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor_type", value)


@pulumi.input_type
class _AviatrixFirewallInstanceAssociationState:
    def __init__(__self__, *,
                 attached: Optional[pulumi.Input[bool]] = None,
                 egress_interface: Optional[pulumi.Input[str]] = None,
                 firenet_gw_name: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 lan_interface: Optional[pulumi.Input[str]] = None,
                 management_interface: Optional[pulumi.Input[str]] = None,
                 vendor_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixFirewallInstanceAssociation resources.
        :param pulumi.Input[bool] attached: Switch to attach/detach firewall instance to/from fireNet.
        :param pulumi.Input[str] egress_interface: Egress interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] firenet_gw_name: Name of the gateway to launch the firewall instance.
        :param pulumi.Input[str] firewall_name: Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
               GCP
        :param pulumi.Input[str] instance_id: ID of Firewall instance, or FQDN Gateway's gw_name.
        :param pulumi.Input[str] lan_interface: Lan interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] management_interface: Management interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] vendor_type: Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
               'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        :param pulumi.Input[str] vpc_id: VPC ID.
        """
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if egress_interface is not None:
            pulumi.set(__self__, "egress_interface", egress_interface)
        if firenet_gw_name is not None:
            pulumi.set(__self__, "firenet_gw_name", firenet_gw_name)
        if firewall_name is not None:
            pulumi.set(__self__, "firewall_name", firewall_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if lan_interface is not None:
            pulumi.set(__self__, "lan_interface", lan_interface)
        if management_interface is not None:
            pulumi.set(__self__, "management_interface", management_interface)
        if vendor_type is not None:
            pulumi.set(__self__, "vendor_type", vendor_type)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def attached(self) -> Optional[pulumi.Input[bool]]:
        """
        Switch to attach/detach firewall instance to/from fireNet.
        """
        return pulumi.get(self, "attached")

    @attached.setter
    def attached(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "attached", value)

    @property
    @pulumi.getter(name="egressInterface")
    def egress_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Egress interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "egress_interface")

    @egress_interface.setter
    def egress_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "egress_interface", value)

    @property
    @pulumi.getter(name="firenetGwName")
    def firenet_gw_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the gateway to launch the firewall instance.
        """
        return pulumi.get(self, "firenet_gw_name")

    @firenet_gw_name.setter
    def firenet_gw_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firenet_gw_name", value)

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
        GCP
        """
        return pulumi.get(self, "firewall_name")

    @firewall_name.setter
    def firewall_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of Firewall instance, or FQDN Gateway's gw_name.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="lanInterface")
    def lan_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Lan interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "lan_interface")

    @lan_interface.setter
    def lan_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lan_interface", value)

    @property
    @pulumi.getter(name="managementInterface")
    def management_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Management interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "management_interface")

    @management_interface.setter
    def management_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "management_interface", value)

    @property
    @pulumi.getter(name="vendorType")
    def vendor_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
        'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        """
        return pulumi.get(self, "vendor_type")

    @vendor_type.setter
    def vendor_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor_type", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class AviatrixFirewallInstanceAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attached: Optional[pulumi.Input[bool]] = None,
                 egress_interface: Optional[pulumi.Input[str]] = None,
                 firenet_gw_name: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 lan_interface: Optional[pulumi.Input[str]] = None,
                 management_interface: Optional[pulumi.Input[str]] = None,
                 vendor_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixFirewallInstanceAssociation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] attached: Switch to attach/detach firewall instance to/from fireNet.
        :param pulumi.Input[str] egress_interface: Egress interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] firenet_gw_name: Name of the gateway to launch the firewall instance.
        :param pulumi.Input[str] firewall_name: Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
               GCP
        :param pulumi.Input[str] instance_id: ID of Firewall instance, or FQDN Gateway's gw_name.
        :param pulumi.Input[str] lan_interface: Lan interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] management_interface: Management interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] vendor_type: Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
               'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        :param pulumi.Input[str] vpc_id: VPC ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixFirewallInstanceAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixFirewallInstanceAssociation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixFirewallInstanceAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixFirewallInstanceAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attached: Optional[pulumi.Input[bool]] = None,
                 egress_interface: Optional[pulumi.Input[str]] = None,
                 firenet_gw_name: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 lan_interface: Optional[pulumi.Input[str]] = None,
                 management_interface: Optional[pulumi.Input[str]] = None,
                 vendor_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixFirewallInstanceAssociationArgs.__new__(AviatrixFirewallInstanceAssociationArgs)

            __props__.__dict__["attached"] = attached
            __props__.__dict__["egress_interface"] = egress_interface
            __props__.__dict__["firenet_gw_name"] = firenet_gw_name
            __props__.__dict__["firewall_name"] = firewall_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["lan_interface"] = lan_interface
            __props__.__dict__["management_interface"] = management_interface
            __props__.__dict__["vendor_type"] = vendor_type
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(AviatrixFirewallInstanceAssociation, __self__).__init__(
            'aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attached: Optional[pulumi.Input[bool]] = None,
            egress_interface: Optional[pulumi.Input[str]] = None,
            firenet_gw_name: Optional[pulumi.Input[str]] = None,
            firewall_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            lan_interface: Optional[pulumi.Input[str]] = None,
            management_interface: Optional[pulumi.Input[str]] = None,
            vendor_type: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'AviatrixFirewallInstanceAssociation':
        """
        Get an existing AviatrixFirewallInstanceAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] attached: Switch to attach/detach firewall instance to/from fireNet.
        :param pulumi.Input[str] egress_interface: Egress interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] firenet_gw_name: Name of the gateway to launch the firewall instance.
        :param pulumi.Input[str] firewall_name: Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
               GCP
        :param pulumi.Input[str] instance_id: ID of Firewall instance, or FQDN Gateway's gw_name.
        :param pulumi.Input[str] lan_interface: Lan interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] management_interface: Management interface ID, required if it is a firewall instance.
        :param pulumi.Input[str] vendor_type: Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
               'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        :param pulumi.Input[str] vpc_id: VPC ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixFirewallInstanceAssociationState.__new__(_AviatrixFirewallInstanceAssociationState)

        __props__.__dict__["attached"] = attached
        __props__.__dict__["egress_interface"] = egress_interface
        __props__.__dict__["firenet_gw_name"] = firenet_gw_name
        __props__.__dict__["firewall_name"] = firewall_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["lan_interface"] = lan_interface
        __props__.__dict__["management_interface"] = management_interface
        __props__.__dict__["vendor_type"] = vendor_type
        __props__.__dict__["vpc_id"] = vpc_id
        return AviatrixFirewallInstanceAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attached(self) -> pulumi.Output[Optional[bool]]:
        """
        Switch to attach/detach firewall instance to/from fireNet.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="egressInterface")
    def egress_interface(self) -> pulumi.Output[Optional[str]]:
        """
        Egress interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "egress_interface")

    @property
    @pulumi.getter(name="firenetGwName")
    def firenet_gw_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the gateway to launch the firewall instance.
        """
        return pulumi.get(self, "firenet_gw_name")

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> pulumi.Output[Optional[str]]:
        """
        Firewall instance name, or FQDN Gateway's gw_name, required if it is a AWS or Azure firewall instance. Not allowed for
        GCP
        """
        return pulumi.get(self, "firewall_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of Firewall instance, or FQDN Gateway's gw_name.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="lanInterface")
    def lan_interface(self) -> pulumi.Output[Optional[str]]:
        """
        Lan interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "lan_interface")

    @property
    @pulumi.getter(name="managementInterface")
    def management_interface(self) -> pulumi.Output[Optional[str]]:
        """
        Management interface ID, required if it is a firewall instance.
        """
        return pulumi.get(self, "management_interface")

    @property
    @pulumi.getter(name="vendorType")
    def vendor_type(self) -> pulumi.Output[Optional[str]]:
        """
        Indication it is a firewall instance or FQDN gateway to be associated to fireNet. Valid values: 'Generic',
        'fqdn_gateway'. Value 'fqdn_gateway' is required for FQDN gateway.
        """
        return pulumi.get(self, "vendor_type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

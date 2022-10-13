# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixAwsTgwConnectArgs', 'AviatrixAwsTgwConnect']

@pulumi.input_type
class AviatrixAwsTgwConnectArgs:
    def __init__(__self__, *,
                 connection_name: pulumi.Input[str],
                 tgw_name: pulumi.Input[str],
                 transport_vpc_id: pulumi.Input[str],
                 network_domain_name: Optional[pulumi.Input[str]] = None,
                 security_domain_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixAwsTgwConnect resource.
        :param pulumi.Input[str] connection_name: Connection Name.
        :param pulumi.Input[str] tgw_name: AWS TGW Name.
        :param pulumi.Input[str] transport_vpc_id: Transport Attachment VPC ID.
        :param pulumi.Input[str] network_domain_name: Network Domain Name.
        :param pulumi.Input[str] security_domain_name: Security Domain Name.
        """
        pulumi.set(__self__, "connection_name", connection_name)
        pulumi.set(__self__, "tgw_name", tgw_name)
        pulumi.set(__self__, "transport_vpc_id", transport_vpc_id)
        if network_domain_name is not None:
            pulumi.set(__self__, "network_domain_name", network_domain_name)
        if security_domain_name is not None:
            warnings.warn("""Please use network_domain_name instead.""", DeprecationWarning)
            pulumi.log.warn("""security_domain_name is deprecated: Please use network_domain_name instead.""")
        if security_domain_name is not None:
            pulumi.set(__self__, "security_domain_name", security_domain_name)

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> pulumi.Input[str]:
        """
        Connection Name.
        """
        return pulumi.get(self, "connection_name")

    @connection_name.setter
    def connection_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_name", value)

    @property
    @pulumi.getter(name="tgwName")
    def tgw_name(self) -> pulumi.Input[str]:
        """
        AWS TGW Name.
        """
        return pulumi.get(self, "tgw_name")

    @tgw_name.setter
    def tgw_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "tgw_name", value)

    @property
    @pulumi.getter(name="transportVpcId")
    def transport_vpc_id(self) -> pulumi.Input[str]:
        """
        Transport Attachment VPC ID.
        """
        return pulumi.get(self, "transport_vpc_id")

    @transport_vpc_id.setter
    def transport_vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transport_vpc_id", value)

    @property
    @pulumi.getter(name="networkDomainName")
    def network_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Network Domain Name.
        """
        return pulumi.get(self, "network_domain_name")

    @network_domain_name.setter
    def network_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_domain_name", value)

    @property
    @pulumi.getter(name="securityDomainName")
    def security_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Security Domain Name.
        """
        return pulumi.get(self, "security_domain_name")

    @security_domain_name.setter
    def security_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_domain_name", value)


@pulumi.input_type
class _AviatrixAwsTgwConnectState:
    def __init__(__self__, *,
                 connect_attachment_id: Optional[pulumi.Input[str]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 network_domain_name: Optional[pulumi.Input[str]] = None,
                 security_domain_name: Optional[pulumi.Input[str]] = None,
                 tgw_name: Optional[pulumi.Input[str]] = None,
                 transport_attachment_id: Optional[pulumi.Input[str]] = None,
                 transport_vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixAwsTgwConnect resources.
        :param pulumi.Input[str] connect_attachment_id: Connect Attachment ID.
        :param pulumi.Input[str] connection_name: Connection Name.
        :param pulumi.Input[str] network_domain_name: Network Domain Name.
        :param pulumi.Input[str] security_domain_name: Security Domain Name.
        :param pulumi.Input[str] tgw_name: AWS TGW Name.
        :param pulumi.Input[str] transport_attachment_id: Transport Attachment ID.
        :param pulumi.Input[str] transport_vpc_id: Transport Attachment VPC ID.
        """
        if connect_attachment_id is not None:
            pulumi.set(__self__, "connect_attachment_id", connect_attachment_id)
        if connection_name is not None:
            pulumi.set(__self__, "connection_name", connection_name)
        if network_domain_name is not None:
            pulumi.set(__self__, "network_domain_name", network_domain_name)
        if security_domain_name is not None:
            warnings.warn("""Please use network_domain_name instead.""", DeprecationWarning)
            pulumi.log.warn("""security_domain_name is deprecated: Please use network_domain_name instead.""")
        if security_domain_name is not None:
            pulumi.set(__self__, "security_domain_name", security_domain_name)
        if tgw_name is not None:
            pulumi.set(__self__, "tgw_name", tgw_name)
        if transport_attachment_id is not None:
            pulumi.set(__self__, "transport_attachment_id", transport_attachment_id)
        if transport_vpc_id is not None:
            pulumi.set(__self__, "transport_vpc_id", transport_vpc_id)

    @property
    @pulumi.getter(name="connectAttachmentId")
    def connect_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        Connect Attachment ID.
        """
        return pulumi.get(self, "connect_attachment_id")

    @connect_attachment_id.setter
    def connect_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connect_attachment_id", value)

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> Optional[pulumi.Input[str]]:
        """
        Connection Name.
        """
        return pulumi.get(self, "connection_name")

    @connection_name.setter
    def connection_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_name", value)

    @property
    @pulumi.getter(name="networkDomainName")
    def network_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Network Domain Name.
        """
        return pulumi.get(self, "network_domain_name")

    @network_domain_name.setter
    def network_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_domain_name", value)

    @property
    @pulumi.getter(name="securityDomainName")
    def security_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Security Domain Name.
        """
        return pulumi.get(self, "security_domain_name")

    @security_domain_name.setter
    def security_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_domain_name", value)

    @property
    @pulumi.getter(name="tgwName")
    def tgw_name(self) -> Optional[pulumi.Input[str]]:
        """
        AWS TGW Name.
        """
        return pulumi.get(self, "tgw_name")

    @tgw_name.setter
    def tgw_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tgw_name", value)

    @property
    @pulumi.getter(name="transportAttachmentId")
    def transport_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        Transport Attachment ID.
        """
        return pulumi.get(self, "transport_attachment_id")

    @transport_attachment_id.setter
    def transport_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transport_attachment_id", value)

    @property
    @pulumi.getter(name="transportVpcId")
    def transport_vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Transport Attachment VPC ID.
        """
        return pulumi.get(self, "transport_vpc_id")

    @transport_vpc_id.setter
    def transport_vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transport_vpc_id", value)


class AviatrixAwsTgwConnect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 network_domain_name: Optional[pulumi.Input[str]] = None,
                 security_domain_name: Optional[pulumi.Input[str]] = None,
                 tgw_name: Optional[pulumi.Input[str]] = None,
                 transport_vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixAwsTgwConnect resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_name: Connection Name.
        :param pulumi.Input[str] network_domain_name: Network Domain Name.
        :param pulumi.Input[str] security_domain_name: Security Domain Name.
        :param pulumi.Input[str] tgw_name: AWS TGW Name.
        :param pulumi.Input[str] transport_vpc_id: Transport Attachment VPC ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixAwsTgwConnectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixAwsTgwConnect resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixAwsTgwConnectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixAwsTgwConnectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 network_domain_name: Optional[pulumi.Input[str]] = None,
                 security_domain_name: Optional[pulumi.Input[str]] = None,
                 tgw_name: Optional[pulumi.Input[str]] = None,
                 transport_vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixAwsTgwConnectArgs.__new__(AviatrixAwsTgwConnectArgs)

            if connection_name is None and not opts.urn:
                raise TypeError("Missing required property 'connection_name'")
            __props__.__dict__["connection_name"] = connection_name
            __props__.__dict__["network_domain_name"] = network_domain_name
            if security_domain_name is not None and not opts.urn:
                warnings.warn("""Please use network_domain_name instead.""", DeprecationWarning)
                pulumi.log.warn("""security_domain_name is deprecated: Please use network_domain_name instead.""")
            __props__.__dict__["security_domain_name"] = security_domain_name
            if tgw_name is None and not opts.urn:
                raise TypeError("Missing required property 'tgw_name'")
            __props__.__dict__["tgw_name"] = tgw_name
            if transport_vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'transport_vpc_id'")
            __props__.__dict__["transport_vpc_id"] = transport_vpc_id
            __props__.__dict__["connect_attachment_id"] = None
            __props__.__dict__["transport_attachment_id"] = None
        super(AviatrixAwsTgwConnect, __self__).__init__(
            'aviatrix:index/aviatrixAwsTgwConnect:AviatrixAwsTgwConnect',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connect_attachment_id: Optional[pulumi.Input[str]] = None,
            connection_name: Optional[pulumi.Input[str]] = None,
            network_domain_name: Optional[pulumi.Input[str]] = None,
            security_domain_name: Optional[pulumi.Input[str]] = None,
            tgw_name: Optional[pulumi.Input[str]] = None,
            transport_attachment_id: Optional[pulumi.Input[str]] = None,
            transport_vpc_id: Optional[pulumi.Input[str]] = None) -> 'AviatrixAwsTgwConnect':
        """
        Get an existing AviatrixAwsTgwConnect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connect_attachment_id: Connect Attachment ID.
        :param pulumi.Input[str] connection_name: Connection Name.
        :param pulumi.Input[str] network_domain_name: Network Domain Name.
        :param pulumi.Input[str] security_domain_name: Security Domain Name.
        :param pulumi.Input[str] tgw_name: AWS TGW Name.
        :param pulumi.Input[str] transport_attachment_id: Transport Attachment ID.
        :param pulumi.Input[str] transport_vpc_id: Transport Attachment VPC ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixAwsTgwConnectState.__new__(_AviatrixAwsTgwConnectState)

        __props__.__dict__["connect_attachment_id"] = connect_attachment_id
        __props__.__dict__["connection_name"] = connection_name
        __props__.__dict__["network_domain_name"] = network_domain_name
        __props__.__dict__["security_domain_name"] = security_domain_name
        __props__.__dict__["tgw_name"] = tgw_name
        __props__.__dict__["transport_attachment_id"] = transport_attachment_id
        __props__.__dict__["transport_vpc_id"] = transport_vpc_id
        return AviatrixAwsTgwConnect(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectAttachmentId")
    def connect_attachment_id(self) -> pulumi.Output[str]:
        """
        Connect Attachment ID.
        """
        return pulumi.get(self, "connect_attachment_id")

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> pulumi.Output[str]:
        """
        Connection Name.
        """
        return pulumi.get(self, "connection_name")

    @property
    @pulumi.getter(name="networkDomainName")
    def network_domain_name(self) -> pulumi.Output[Optional[str]]:
        """
        Network Domain Name.
        """
        return pulumi.get(self, "network_domain_name")

    @property
    @pulumi.getter(name="securityDomainName")
    def security_domain_name(self) -> pulumi.Output[Optional[str]]:
        """
        Security Domain Name.
        """
        return pulumi.get(self, "security_domain_name")

    @property
    @pulumi.getter(name="tgwName")
    def tgw_name(self) -> pulumi.Output[str]:
        """
        AWS TGW Name.
        """
        return pulumi.get(self, "tgw_name")

    @property
    @pulumi.getter(name="transportAttachmentId")
    def transport_attachment_id(self) -> pulumi.Output[str]:
        """
        Transport Attachment ID.
        """
        return pulumi.get(self, "transport_attachment_id")

    @property
    @pulumi.getter(name="transportVpcId")
    def transport_vpc_id(self) -> pulumi.Output[str]:
        """
        Transport Attachment VPC ID.
        """
        return pulumi.get(self, "transport_vpc_id")

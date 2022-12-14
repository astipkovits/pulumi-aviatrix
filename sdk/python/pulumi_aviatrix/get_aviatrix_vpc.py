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

__all__ = [
    'GetAviatrixVpcResult',
    'AwaitableGetAviatrixVpcResult',
    'get_aviatrix_vpc',
    'get_aviatrix_vpc_output',
]

@pulumi.output_type
class GetAviatrixVpcResult:
    """
    A collection of values returned by getAviatrixVpc.
    """
    def __init__(__self__, account_name=None, availability_domains=None, aviatrix_firenet_vpc=None, aviatrix_transit_vpc=None, azure_vnet_resource_id=None, cidr=None, cloud_type=None, fault_domains=None, id=None, name=None, num_of_subnet_pairs=None, private_subnets=None, public_subnets=None, region=None, resource_group=None, route_tables=None, route_tables_filter=None, subnet_size=None, subnets=None, vpc_id=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if availability_domains and not isinstance(availability_domains, list):
            raise TypeError("Expected argument 'availability_domains' to be a list")
        pulumi.set(__self__, "availability_domains", availability_domains)
        if aviatrix_firenet_vpc and not isinstance(aviatrix_firenet_vpc, bool):
            raise TypeError("Expected argument 'aviatrix_firenet_vpc' to be a bool")
        pulumi.set(__self__, "aviatrix_firenet_vpc", aviatrix_firenet_vpc)
        if aviatrix_transit_vpc and not isinstance(aviatrix_transit_vpc, bool):
            raise TypeError("Expected argument 'aviatrix_transit_vpc' to be a bool")
        pulumi.set(__self__, "aviatrix_transit_vpc", aviatrix_transit_vpc)
        if azure_vnet_resource_id and not isinstance(azure_vnet_resource_id, str):
            raise TypeError("Expected argument 'azure_vnet_resource_id' to be a str")
        pulumi.set(__self__, "azure_vnet_resource_id", azure_vnet_resource_id)
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if cloud_type and not isinstance(cloud_type, int):
            raise TypeError("Expected argument 'cloud_type' to be a int")
        pulumi.set(__self__, "cloud_type", cloud_type)
        if fault_domains and not isinstance(fault_domains, list):
            raise TypeError("Expected argument 'fault_domains' to be a list")
        pulumi.set(__self__, "fault_domains", fault_domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if num_of_subnet_pairs and not isinstance(num_of_subnet_pairs, int):
            raise TypeError("Expected argument 'num_of_subnet_pairs' to be a int")
        pulumi.set(__self__, "num_of_subnet_pairs", num_of_subnet_pairs)
        if private_subnets and not isinstance(private_subnets, list):
            raise TypeError("Expected argument 'private_subnets' to be a list")
        pulumi.set(__self__, "private_subnets", private_subnets)
        if public_subnets and not isinstance(public_subnets, list):
            raise TypeError("Expected argument 'public_subnets' to be a list")
        pulumi.set(__self__, "public_subnets", public_subnets)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        pulumi.set(__self__, "resource_group", resource_group)
        if route_tables and not isinstance(route_tables, list):
            raise TypeError("Expected argument 'route_tables' to be a list")
        pulumi.set(__self__, "route_tables", route_tables)
        if route_tables_filter and not isinstance(route_tables_filter, str):
            raise TypeError("Expected argument 'route_tables_filter' to be a str")
        pulumi.set(__self__, "route_tables_filter", route_tables_filter)
        if subnet_size and not isinstance(subnet_size, int):
            raise TypeError("Expected argument 'subnet_size' to be a int")
        pulumi.set(__self__, "subnet_size", subnet_size)
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        Account name of the VPC created.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="availabilityDomains")
    def availability_domains(self) -> Sequence[str]:
        """
        List of OCI availability domains.
        """
        return pulumi.get(self, "availability_domains")

    @property
    @pulumi.getter(name="aviatrixFirenetVpc")
    def aviatrix_firenet_vpc(self) -> bool:
        """
        Switch if the VPC created is an Aviatrix FireNet VPC or not.
        """
        return pulumi.get(self, "aviatrix_firenet_vpc")

    @property
    @pulumi.getter(name="aviatrixTransitVpc")
    def aviatrix_transit_vpc(self) -> bool:
        """
        Switch if the VPC created is an Aviatrix Transit VPC or not.
        """
        return pulumi.get(self, "aviatrix_transit_vpc")

    @property
    @pulumi.getter(name="azureVnetResourceId")
    def azure_vnet_resource_id(self) -> str:
        """
        Azure vnet resource ID.
        """
        return pulumi.get(self, "azure_vnet_resource_id")

    @property
    @pulumi.getter
    def cidr(self) -> str:
        """
        Private subnet CIDR.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> int:
        """
        Type of cloud service provider.
        """
        return pulumi.get(self, "cloud_type")

    @property
    @pulumi.getter(name="faultDomains")
    def fault_domains(self) -> Sequence[str]:
        """
        List of OCI fault domains.
        """
        return pulumi.get(self, "fault_domains")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Private subnet name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numOfSubnetPairs")
    def num_of_subnet_pairs(self) -> int:
        """
        Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider.
        """
        return pulumi.get(self, "num_of_subnet_pairs")

    @property
    @pulumi.getter(name="privateSubnets")
    def private_subnets(self) -> Sequence['outputs.GetAviatrixVpcPrivateSubnetResult']:
        """
        List of private subnet of the VPC(AWS, Azure) created.
        """
        return pulumi.get(self, "private_subnets")

    @property
    @pulumi.getter(name="publicSubnets")
    def public_subnets(self) -> Sequence['outputs.GetAviatrixVpcPublicSubnetResult']:
        """
        List of public subnet of the VPC(AWS, Azure) created.
        """
        return pulumi.get(self, "public_subnets")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region of the VPC created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Resource group of the Azure VPC created.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="routeTables")
    def route_tables(self) -> Sequence[str]:
        """
        List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure vpc.
        """
        return pulumi.get(self, "route_tables")

    @property
    @pulumi.getter(name="routeTablesFilter")
    def route_tables_filter(self) -> Optional[str]:
        return pulumi.get(self, "route_tables_filter")

    @property
    @pulumi.getter(name="subnetSize")
    def subnet_size(self) -> int:
        """
        Subnet size. Only supported for AWS, Azure provider.
        """
        return pulumi.get(self, "subnet_size")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence['outputs.GetAviatrixVpcSubnetResult']:
        """
        List of subnet of the VPC created.
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        ID of the VPC created.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetAviatrixVpcResult(GetAviatrixVpcResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAviatrixVpcResult(
            account_name=self.account_name,
            availability_domains=self.availability_domains,
            aviatrix_firenet_vpc=self.aviatrix_firenet_vpc,
            aviatrix_transit_vpc=self.aviatrix_transit_vpc,
            azure_vnet_resource_id=self.azure_vnet_resource_id,
            cidr=self.cidr,
            cloud_type=self.cloud_type,
            fault_domains=self.fault_domains,
            id=self.id,
            name=self.name,
            num_of_subnet_pairs=self.num_of_subnet_pairs,
            private_subnets=self.private_subnets,
            public_subnets=self.public_subnets,
            region=self.region,
            resource_group=self.resource_group,
            route_tables=self.route_tables,
            route_tables_filter=self.route_tables_filter,
            subnet_size=self.subnet_size,
            subnets=self.subnets,
            vpc_id=self.vpc_id)


def get_aviatrix_vpc(name: Optional[str] = None,
                     route_tables_filter: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAviatrixVpcResult:
    """
    The **aviatrix_vpc** data source provides details about a specific VPC created by the Aviatrix Controller.

    This data source can prove useful when a module accepts any form of VPC detail as an input variable. For example, requiring a subnet CIDR specification when creating a gateway.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aviatrix as aviatrix

    test = aviatrix.get_aviatrix_vpc(name="vpc-test")
    ```


    :param str name: Name of the Aviatrix VPC.
    :param str route_tables_filter: Filters the `route_tables` list to contain only public or private route tables. Valid values are 'private' or 'public'. If not set `route_tables` is not filtered.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['routeTablesFilter'] = route_tables_filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aviatrix:index/getAviatrixVpc:getAviatrixVpc', __args__, opts=opts, typ=GetAviatrixVpcResult).value

    return AwaitableGetAviatrixVpcResult(
        account_name=__ret__.account_name,
        availability_domains=__ret__.availability_domains,
        aviatrix_firenet_vpc=__ret__.aviatrix_firenet_vpc,
        aviatrix_transit_vpc=__ret__.aviatrix_transit_vpc,
        azure_vnet_resource_id=__ret__.azure_vnet_resource_id,
        cidr=__ret__.cidr,
        cloud_type=__ret__.cloud_type,
        fault_domains=__ret__.fault_domains,
        id=__ret__.id,
        name=__ret__.name,
        num_of_subnet_pairs=__ret__.num_of_subnet_pairs,
        private_subnets=__ret__.private_subnets,
        public_subnets=__ret__.public_subnets,
        region=__ret__.region,
        resource_group=__ret__.resource_group,
        route_tables=__ret__.route_tables,
        route_tables_filter=__ret__.route_tables_filter,
        subnet_size=__ret__.subnet_size,
        subnets=__ret__.subnets,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_aviatrix_vpc)
def get_aviatrix_vpc_output(name: Optional[pulumi.Input[str]] = None,
                            route_tables_filter: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAviatrixVpcResult]:
    """
    The **aviatrix_vpc** data source provides details about a specific VPC created by the Aviatrix Controller.

    This data source can prove useful when a module accepts any form of VPC detail as an input variable. For example, requiring a subnet CIDR specification when creating a gateway.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aviatrix as aviatrix

    test = aviatrix.get_aviatrix_vpc(name="vpc-test")
    ```


    :param str name: Name of the Aviatrix VPC.
    :param str route_tables_filter: Filters the `route_tables` list to contain only public or private route tables. Valid values are 'private' or 'public'. If not set `route_tables` is not filtered.
    """
    ...

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixVpnUserAcceleratorArgs', 'AviatrixVpnUserAccelerator']

@pulumi.input_type
class AviatrixVpnUserAcceleratorArgs:
    def __init__(__self__, *,
                 elb_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixVpnUserAccelerator resource.
        :param pulumi.Input[str] elb_name: Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        pulumi.set(__self__, "elb_name", elb_name)

    @property
    @pulumi.getter(name="elbName")
    def elb_name(self) -> pulumi.Input[str]:
        """
        Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        return pulumi.get(self, "elb_name")

    @elb_name.setter
    def elb_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "elb_name", value)


@pulumi.input_type
class _AviatrixVpnUserAcceleratorState:
    def __init__(__self__, *,
                 elb_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixVpnUserAccelerator resources.
        :param pulumi.Input[str] elb_name: Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        if elb_name is not None:
            pulumi.set(__self__, "elb_name", elb_name)

    @property
    @pulumi.getter(name="elbName")
    def elb_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        return pulumi.get(self, "elb_name")

    @elb_name.setter
    def elb_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "elb_name", value)


class AviatrixVpnUserAccelerator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 elb_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_vpn_user_accelerator** resource manages the [Aviatrix VPN User Accelerator](https://docs.aviatrix.com/HowTos/user_accelerator.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Vpn User Accelerator
        test_vpc_accelerator = aviatrix.AviatrixVpnUserAccelerator("testVpcAccelerator", elb_name="Aviatrix-vpc-abcd2134")
        ```

        ## Import

        **vpn_user_accelerator** can be imported using the `elb_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator test Aviatrix-vpc-abcd1234
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] elb_name: Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixVpnUserAcceleratorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_vpn_user_accelerator** resource manages the [Aviatrix VPN User Accelerator](https://docs.aviatrix.com/HowTos/user_accelerator.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Vpn User Accelerator
        test_vpc_accelerator = aviatrix.AviatrixVpnUserAccelerator("testVpcAccelerator", elb_name="Aviatrix-vpc-abcd2134")
        ```

        ## Import

        **vpn_user_accelerator** can be imported using the `elb_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator test Aviatrix-vpc-abcd1234
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixVpnUserAcceleratorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixVpnUserAcceleratorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 elb_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixVpnUserAcceleratorArgs.__new__(AviatrixVpnUserAcceleratorArgs)

            if elb_name is None and not opts.urn:
                raise TypeError("Missing required property 'elb_name'")
            __props__.__dict__["elb_name"] = elb_name
        super(AviatrixVpnUserAccelerator, __self__).__init__(
            'aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            elb_name: Optional[pulumi.Input[str]] = None) -> 'AviatrixVpnUserAccelerator':
        """
        Get an existing AviatrixVpnUserAccelerator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] elb_name: Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixVpnUserAcceleratorState.__new__(_AviatrixVpnUserAcceleratorState)

        __props__.__dict__["elb_name"] = elb_name
        return AviatrixVpnUserAccelerator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="elbName")
    def elb_name(self) -> pulumi.Output[str]:
        """
        Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
        """
        return pulumi.get(self, "elb_name")


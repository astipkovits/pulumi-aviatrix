# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixSegmentationSecurityDomainArgs', 'AviatrixSegmentationSecurityDomain']

@pulumi.input_type
class AviatrixSegmentationSecurityDomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixSegmentationSecurityDomain resource.
        :param pulumi.Input[str] domain_name: Name of the Security Domain.
        """
        pulumi.set(__self__, "domain_name", domain_name)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        Name of the Security Domain.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)


@pulumi.input_type
class _AviatrixSegmentationSecurityDomainState:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixSegmentationSecurityDomain resources.
        :param pulumi.Input[str] domain_name: Name of the Security Domain.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Security Domain.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)


class AviatrixSegmentationSecurityDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_segmentation_security_domain** resource handles creation of [Transit Segmentation](https://docs.aviatrix.com/HowTos/transit_segmentation_faq.html) Security Domains.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Segmentation Security Domain
        test_segmentation_security_domain = aviatrix.AviatrixSegmentationSecurityDomain("testSegmentationSecurityDomain", domain_name="domain-a")
        ```

        ## Import

        **aviatrix_segmentation_security_domain** can be imported using the `domain_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixSegmentationSecurityDomain:AviatrixSegmentationSecurityDomain test domain_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Name of the Security Domain.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixSegmentationSecurityDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_segmentation_security_domain** resource handles creation of [Transit Segmentation](https://docs.aviatrix.com/HowTos/transit_segmentation_faq.html) Security Domains.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Segmentation Security Domain
        test_segmentation_security_domain = aviatrix.AviatrixSegmentationSecurityDomain("testSegmentationSecurityDomain", domain_name="domain-a")
        ```

        ## Import

        **aviatrix_segmentation_security_domain** can be imported using the `domain_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixSegmentationSecurityDomain:AviatrixSegmentationSecurityDomain test domain_name
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixSegmentationSecurityDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixSegmentationSecurityDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixSegmentationSecurityDomainArgs.__new__(AviatrixSegmentationSecurityDomainArgs)

            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
        super(AviatrixSegmentationSecurityDomain, __self__).__init__(
            'aviatrix:index/aviatrixSegmentationSecurityDomain:AviatrixSegmentationSecurityDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name: Optional[pulumi.Input[str]] = None) -> 'AviatrixSegmentationSecurityDomain':
        """
        Get an existing AviatrixSegmentationSecurityDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Name of the Security Domain.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixSegmentationSecurityDomainState.__new__(_AviatrixSegmentationSecurityDomainState)

        __props__.__dict__["domain_name"] = domain_name
        return AviatrixSegmentationSecurityDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Name of the Security Domain.
        """
        return pulumi.get(self, "domain_name")


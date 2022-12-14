# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixCloudwatchAgentArgs', 'AviatrixCloudwatchAgent']

@pulumi.input_type
class AviatrixCloudwatchAgentArgs:
    def __init__(__self__, *,
                 cloudwatch_role_arn: pulumi.Input[str],
                 region: pulumi.Input[str],
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AviatrixCloudwatchAgent resource.
        :param pulumi.Input[str] cloudwatch_role_arn: CloudWatch role ARN.
        :param pulumi.Input[str] region: Name of AWS region.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        :param pulumi.Input[str] log_group_name: Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        """
        pulumi.set(__self__, "cloudwatch_role_arn", cloudwatch_role_arn)
        pulumi.set(__self__, "region", region)
        if excluded_gateways is not None:
            pulumi.set(__self__, "excluded_gateways", excluded_gateways)
        if log_group_name is not None:
            pulumi.set(__self__, "log_group_name", log_group_name)

    @property
    @pulumi.getter(name="cloudwatchRoleArn")
    def cloudwatch_role_arn(self) -> pulumi.Input[str]:
        """
        CloudWatch role ARN.
        """
        return pulumi.get(self, "cloudwatch_role_arn")

    @cloudwatch_role_arn.setter
    def cloudwatch_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloudwatch_role_arn", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Name of AWS region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="excludedGateways")
    def excluded_gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        """
        return pulumi.get(self, "excluded_gateways")

    @excluded_gateways.setter
    def excluded_gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_gateways", value)

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        """
        return pulumi.get(self, "log_group_name")

    @log_group_name.setter
    def log_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group_name", value)


@pulumi.input_type
class _AviatrixCloudwatchAgentState:
    def __init__(__self__, *,
                 cloudwatch_role_arn: Optional[pulumi.Input[str]] = None,
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixCloudwatchAgent resources.
        :param pulumi.Input[str] cloudwatch_role_arn: CloudWatch role ARN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        :param pulumi.Input[str] log_group_name: Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        :param pulumi.Input[str] region: Name of AWS region.
        :param pulumi.Input[str] status: The status of cloudwatch agent.
        """
        if cloudwatch_role_arn is not None:
            pulumi.set(__self__, "cloudwatch_role_arn", cloudwatch_role_arn)
        if excluded_gateways is not None:
            pulumi.set(__self__, "excluded_gateways", excluded_gateways)
        if log_group_name is not None:
            pulumi.set(__self__, "log_group_name", log_group_name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="cloudwatchRoleArn")
    def cloudwatch_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        CloudWatch role ARN.
        """
        return pulumi.get(self, "cloudwatch_role_arn")

    @cloudwatch_role_arn.setter
    def cloudwatch_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloudwatch_role_arn", value)

    @property
    @pulumi.getter(name="excludedGateways")
    def excluded_gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        """
        return pulumi.get(self, "excluded_gateways")

    @excluded_gateways.setter
    def excluded_gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_gateways", value)

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        """
        return pulumi.get(self, "log_group_name")

    @log_group_name.setter
    def log_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group_name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Name of AWS region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of cloudwatch agent.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class AviatrixCloudwatchAgent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudwatch_role_arn: Optional[pulumi.Input[str]] = None,
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_cloudwatch_agent** resource allows the enabling and disabling of cloudwatch agent.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Enable cloudwatch agent
        test_cloudwatch_agent = aviatrix.AviatrixCloudwatchAgent("testCloudwatchAgent",
            cloudwatch_role_arn="arn:aws:iam::469550033836:role/aviatrix-role-cloudwatch",
            excluded_gateways=[
                "a",
                "b",
            ],
            region="us-east-1")
        ```

        ## Import

        **cloudwatch_agent** can be imported using "cloudwatch_agent", e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixCloudwatchAgent:AviatrixCloudwatchAgent test cloudwatch_agent
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloudwatch_role_arn: CloudWatch role ARN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        :param pulumi.Input[str] log_group_name: Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        :param pulumi.Input[str] region: Name of AWS region.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixCloudwatchAgentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_cloudwatch_agent** resource allows the enabling and disabling of cloudwatch agent.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Enable cloudwatch agent
        test_cloudwatch_agent = aviatrix.AviatrixCloudwatchAgent("testCloudwatchAgent",
            cloudwatch_role_arn="arn:aws:iam::469550033836:role/aviatrix-role-cloudwatch",
            excluded_gateways=[
                "a",
                "b",
            ],
            region="us-east-1")
        ```

        ## Import

        **cloudwatch_agent** can be imported using "cloudwatch_agent", e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixCloudwatchAgent:AviatrixCloudwatchAgent test cloudwatch_agent
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixCloudwatchAgentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixCloudwatchAgentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudwatch_role_arn: Optional[pulumi.Input[str]] = None,
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixCloudwatchAgentArgs.__new__(AviatrixCloudwatchAgentArgs)

            if cloudwatch_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'cloudwatch_role_arn'")
            __props__.__dict__["cloudwatch_role_arn"] = cloudwatch_role_arn
            __props__.__dict__["excluded_gateways"] = excluded_gateways
            __props__.__dict__["log_group_name"] = log_group_name
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["status"] = None
        super(AviatrixCloudwatchAgent, __self__).__init__(
            'aviatrix:index/aviatrixCloudwatchAgent:AviatrixCloudwatchAgent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloudwatch_role_arn: Optional[pulumi.Input[str]] = None,
            excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            log_group_name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'AviatrixCloudwatchAgent':
        """
        Get an existing AviatrixCloudwatchAgent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloudwatch_role_arn: CloudWatch role ARN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        :param pulumi.Input[str] log_group_name: Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        :param pulumi.Input[str] region: Name of AWS region.
        :param pulumi.Input[str] status: The status of cloudwatch agent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixCloudwatchAgentState.__new__(_AviatrixCloudwatchAgentState)

        __props__.__dict__["cloudwatch_role_arn"] = cloudwatch_role_arn
        __props__.__dict__["excluded_gateways"] = excluded_gateways
        __props__.__dict__["log_group_name"] = log_group_name
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        return AviatrixCloudwatchAgent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudwatchRoleArn")
    def cloudwatch_role_arn(self) -> pulumi.Output[str]:
        """
        CloudWatch role ARN.
        """
        return pulumi.get(self, "cloudwatch_role_arn")

    @property
    @pulumi.getter(name="excludedGateways")
    def excluded_gateways(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        """
        return pulumi.get(self, "excluded_gateways")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        Log group name. "AVIATRIX-CLOUDWATCH-LOG" by default.
        """
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Name of AWS region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of cloudwatch agent.
        """
        return pulumi.get(self, "status")


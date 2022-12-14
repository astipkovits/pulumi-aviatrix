# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixRbacGroupPermissionAttachmentArgs', 'AviatrixRbacGroupPermissionAttachment']

@pulumi.input_type
class AviatrixRbacGroupPermissionAttachmentArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[str],
                 permission_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixRbacGroupPermissionAttachment resource.
        :param pulumi.Input[str] group_name: This parameter represents the name of a RBAC group.
        :param pulumi.Input[str] permission_name: This parameter represents the permission to attach to the RBAC group.
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "permission_name", permission_name)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        This parameter represents the name of a RBAC group.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="permissionName")
    def permission_name(self) -> pulumi.Input[str]:
        """
        This parameter represents the permission to attach to the RBAC group.
        """
        return pulumi.get(self, "permission_name")

    @permission_name.setter
    def permission_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission_name", value)


@pulumi.input_type
class _AviatrixRbacGroupPermissionAttachmentState:
    def __init__(__self__, *,
                 group_name: Optional[pulumi.Input[str]] = None,
                 permission_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixRbacGroupPermissionAttachment resources.
        :param pulumi.Input[str] group_name: This parameter represents the name of a RBAC group.
        :param pulumi.Input[str] permission_name: This parameter represents the permission to attach to the RBAC group.
        """
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if permission_name is not None:
            pulumi.set(__self__, "permission_name", permission_name)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        This parameter represents the name of a RBAC group.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="permissionName")
    def permission_name(self) -> Optional[pulumi.Input[str]]:
        """
        This parameter represents the permission to attach to the RBAC group.
        """
        return pulumi.get(self, "permission_name")

    @permission_name.setter
    def permission_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission_name", value)


class AviatrixRbacGroupPermissionAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 permission_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The **aviatrix_rbac_group_permission_attachment** resource allows the creation and management of permission attachments to Aviatrix (Role-Based Access Control) RBAC groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Rbac Group Permission Attachment
        test_attachment = aviatrix.AviatrixRbacGroupPermissionAttachment("testAttachment",
            group_name="write_only",
            permission_name="all_write")
        ```

        ## Import

        **rbac_group_permission_attachment** can be imported using the `group_name` and `permission_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment test group_name~permission_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: This parameter represents the name of a RBAC group.
        :param pulumi.Input[str] permission_name: This parameter represents the permission to attach to the RBAC group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixRbacGroupPermissionAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The **aviatrix_rbac_group_permission_attachment** resource allows the creation and management of permission attachments to Aviatrix (Role-Based Access Control) RBAC groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aviatrix as aviatrix

        # Create an Aviatrix Rbac Group Permission Attachment
        test_attachment = aviatrix.AviatrixRbacGroupPermissionAttachment("testAttachment",
            group_name="write_only",
            permission_name="all_write")
        ```

        ## Import

        **rbac_group_permission_attachment** can be imported using the `group_name` and `permission_name`, e.g.

        ```sh
         $ pulumi import aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment test group_name~permission_name
        ```

        :param str resource_name: The name of the resource.
        :param AviatrixRbacGroupPermissionAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixRbacGroupPermissionAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 permission_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixRbacGroupPermissionAttachmentArgs.__new__(AviatrixRbacGroupPermissionAttachmentArgs)

            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            if permission_name is None and not opts.urn:
                raise TypeError("Missing required property 'permission_name'")
            __props__.__dict__["permission_name"] = permission_name
        super(AviatrixRbacGroupPermissionAttachment, __self__).__init__(
            'aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            permission_name: Optional[pulumi.Input[str]] = None) -> 'AviatrixRbacGroupPermissionAttachment':
        """
        Get an existing AviatrixRbacGroupPermissionAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: This parameter represents the name of a RBAC group.
        :param pulumi.Input[str] permission_name: This parameter represents the permission to attach to the RBAC group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixRbacGroupPermissionAttachmentState.__new__(_AviatrixRbacGroupPermissionAttachmentState)

        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["permission_name"] = permission_name
        return AviatrixRbacGroupPermissionAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        This parameter represents the name of a RBAC group.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="permissionName")
    def permission_name(self) -> pulumi.Output[str]:
        """
        This parameter represents the permission to attach to the RBAC group.
        """
        return pulumi.get(self, "permission_name")


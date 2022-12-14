// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    /// <summary>
    /// The **aviatrix_rbac_group_permission_attachment** resource allows the creation and management of permission attachments to Aviatrix (Role-Based Access Control) RBAC groups.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Aviatrix Rbac Group Permission Attachment
    ///     var testAttachment = new Aviatrix.AviatrixRbacGroupPermissionAttachment("testAttachment", new()
    ///     {
    ///         GroupName = "write_only",
    ///         PermissionName = "all_write",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **rbac_group_permission_attachment** can be imported using the `group_name` and `permission_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment test group_name~permission_name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment")]
    public partial class AviatrixRbacGroupPermissionAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of a RBAC group.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the permission to attach to the RBAC group.
        /// </summary>
        [Output("permissionName")]
        public Output<string> PermissionName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixRbacGroupPermissionAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixRbacGroupPermissionAttachment(string name, AviatrixRbacGroupPermissionAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment", name, args ?? new AviatrixRbacGroupPermissionAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixRbacGroupPermissionAttachment(string name, Input<string> id, AviatrixRbacGroupPermissionAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/astipkovits",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixRbacGroupPermissionAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixRbacGroupPermissionAttachment Get(string name, Input<string> id, AviatrixRbacGroupPermissionAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixRbacGroupPermissionAttachment(name, id, state, options);
        }
    }

    public sealed class AviatrixRbacGroupPermissionAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a RBAC group.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// This parameter represents the permission to attach to the RBAC group.
        /// </summary>
        [Input("permissionName", required: true)]
        public Input<string> PermissionName { get; set; } = null!;

        public AviatrixRbacGroupPermissionAttachmentArgs()
        {
        }
        public static new AviatrixRbacGroupPermissionAttachmentArgs Empty => new AviatrixRbacGroupPermissionAttachmentArgs();
    }

    public sealed class AviatrixRbacGroupPermissionAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a RBAC group.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// This parameter represents the permission to attach to the RBAC group.
        /// </summary>
        [Input("permissionName")]
        public Input<string>? PermissionName { get; set; }

        public AviatrixRbacGroupPermissionAttachmentState()
        {
        }
        public static new AviatrixRbacGroupPermissionAttachmentState Empty => new AviatrixRbacGroupPermissionAttachmentState();
    }
}

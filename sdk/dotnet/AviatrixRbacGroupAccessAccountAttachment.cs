// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixRbacGroupAccessAccountAttachment:AviatrixRbacGroupAccessAccountAttachment")]
    public partial class AviatrixRbacGroupAccessAccountAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access account name.
        /// </summary>
        [Output("accessAccountName")]
        public Output<string> AccessAccountName { get; private set; } = null!;

        /// <summary>
        /// RBAC permission group name.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixRbacGroupAccessAccountAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixRbacGroupAccessAccountAttachment(string name, AviatrixRbacGroupAccessAccountAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixRbacGroupAccessAccountAttachment:AviatrixRbacGroupAccessAccountAttachment", name, args ?? new AviatrixRbacGroupAccessAccountAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixRbacGroupAccessAccountAttachment(string name, Input<string> id, AviatrixRbacGroupAccessAccountAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixRbacGroupAccessAccountAttachment:AviatrixRbacGroupAccessAccountAttachment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/astipkovits/pulumi-aviatrix/releases/download/${VERSION}/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixRbacGroupAccessAccountAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixRbacGroupAccessAccountAttachment Get(string name, Input<string> id, AviatrixRbacGroupAccessAccountAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixRbacGroupAccessAccountAttachment(name, id, state, options);
        }
    }

    public sealed class AviatrixRbacGroupAccessAccountAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access account name.
        /// </summary>
        [Input("accessAccountName", required: true)]
        public Input<string> AccessAccountName { get; set; } = null!;

        /// <summary>
        /// RBAC permission group name.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        public AviatrixRbacGroupAccessAccountAttachmentArgs()
        {
        }
        public static new AviatrixRbacGroupAccessAccountAttachmentArgs Empty => new AviatrixRbacGroupAccessAccountAttachmentArgs();
    }

    public sealed class AviatrixRbacGroupAccessAccountAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access account name.
        /// </summary>
        [Input("accessAccountName")]
        public Input<string>? AccessAccountName { get; set; }

        /// <summary>
        /// RBAC permission group name.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        public AviatrixRbacGroupAccessAccountAttachmentState()
        {
        }
        public static new AviatrixRbacGroupAccessAccountAttachmentState Empty => new AviatrixRbacGroupAccessAccountAttachmentState();
    }
}

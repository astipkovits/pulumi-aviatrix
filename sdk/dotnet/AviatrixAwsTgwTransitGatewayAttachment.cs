// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwTransitGatewayAttachment:AviatrixAwsTgwTransitGatewayAttachment")]
    public partial class AviatrixAwsTgwTransitGatewayAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Name of the AWS TGW.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;

        /// <summary>
        /// Name of the transit gateway to be attached to tgw.
        /// </summary>
        [Output("transitGatewayName")]
        public Output<string> TransitGatewayName { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("vpcAccountName")]
        public Output<string> VpcAccountName { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwTransitGatewayAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwTransitGatewayAttachment(string name, AviatrixAwsTgwTransitGatewayAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwTransitGatewayAttachment:AviatrixAwsTgwTransitGatewayAttachment", name, args ?? new AviatrixAwsTgwTransitGatewayAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwTransitGatewayAttachment(string name, Input<string> id, AviatrixAwsTgwTransitGatewayAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwTransitGatewayAttachment:AviatrixAwsTgwTransitGatewayAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsTgwTransitGatewayAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwTransitGatewayAttachment Get(string name, Input<string> id, AviatrixAwsTgwTransitGatewayAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwTransitGatewayAttachment(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwTransitGatewayAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Name of the AWS TGW.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        /// <summary>
        /// Name of the transit gateway to be attached to tgw.
        /// </summary>
        [Input("transitGatewayName", required: true)]
        public Input<string> TransitGatewayName { get; set; } = null!;

        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("vpcAccountName", required: true)]
        public Input<string> VpcAccountName { get; set; } = null!;

        /// <summary>
        /// This parameter represents the ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixAwsTgwTransitGatewayAttachmentArgs()
        {
        }
        public static new AviatrixAwsTgwTransitGatewayAttachmentArgs Empty => new AviatrixAwsTgwTransitGatewayAttachmentArgs();
    }

    public sealed class AviatrixAwsTgwTransitGatewayAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Name of the AWS TGW.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        /// <summary>
        /// Name of the transit gateway to be attached to tgw.
        /// </summary>
        [Input("transitGatewayName")]
        public Input<string>? TransitGatewayName { get; set; }

        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("vpcAccountName")]
        public Input<string>? VpcAccountName { get; set; }

        /// <summary>
        /// This parameter represents the ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixAwsTgwTransitGatewayAttachmentState()
        {
        }
        public static new AviatrixAwsTgwTransitGatewayAttachmentState Empty => new AviatrixAwsTgwTransitGatewayAttachmentState();
    }
}

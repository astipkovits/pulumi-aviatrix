// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwConnect:AviatrixAwsTgwConnect")]
    public partial class AviatrixAwsTgwConnect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Connect Attachment ID.
        /// </summary>
        [Output("connectAttachmentId")]
        public Output<string> ConnectAttachmentId { get; private set; } = null!;

        /// <summary>
        /// Connection Name.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// Network Domain Name.
        /// </summary>
        [Output("networkDomainName")]
        public Output<string?> NetworkDomainName { get; private set; } = null!;

        /// <summary>
        /// Security Domain Name.
        /// </summary>
        [Output("securityDomainName")]
        public Output<string?> SecurityDomainName { get; private set; } = null!;

        /// <summary>
        /// AWS TGW Name.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;

        /// <summary>
        /// Transport Attachment ID.
        /// </summary>
        [Output("transportAttachmentId")]
        public Output<string> TransportAttachmentId { get; private set; } = null!;

        /// <summary>
        /// Transport Attachment VPC ID.
        /// </summary>
        [Output("transportVpcId")]
        public Output<string> TransportVpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwConnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwConnect(string name, AviatrixAwsTgwConnectArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwConnect:AviatrixAwsTgwConnect", name, args ?? new AviatrixAwsTgwConnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwConnect(string name, Input<string> id, AviatrixAwsTgwConnectState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwConnect:AviatrixAwsTgwConnect", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsTgwConnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwConnect Get(string name, Input<string> id, AviatrixAwsTgwConnectState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwConnect(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwConnectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection Name.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Network Domain Name.
        /// </summary>
        [Input("networkDomainName")]
        public Input<string>? NetworkDomainName { get; set; }

        /// <summary>
        /// Security Domain Name.
        /// </summary>
        [Input("securityDomainName")]
        public Input<string>? SecurityDomainName { get; set; }

        /// <summary>
        /// AWS TGW Name.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        /// <summary>
        /// Transport Attachment VPC ID.
        /// </summary>
        [Input("transportVpcId", required: true)]
        public Input<string> TransportVpcId { get; set; } = null!;

        public AviatrixAwsTgwConnectArgs()
        {
        }
        public static new AviatrixAwsTgwConnectArgs Empty => new AviatrixAwsTgwConnectArgs();
    }

    public sealed class AviatrixAwsTgwConnectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connect Attachment ID.
        /// </summary>
        [Input("connectAttachmentId")]
        public Input<string>? ConnectAttachmentId { get; set; }

        /// <summary>
        /// Connection Name.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Network Domain Name.
        /// </summary>
        [Input("networkDomainName")]
        public Input<string>? NetworkDomainName { get; set; }

        /// <summary>
        /// Security Domain Name.
        /// </summary>
        [Input("securityDomainName")]
        public Input<string>? SecurityDomainName { get; set; }

        /// <summary>
        /// AWS TGW Name.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        /// <summary>
        /// Transport Attachment ID.
        /// </summary>
        [Input("transportAttachmentId")]
        public Input<string>? TransportAttachmentId { get; set; }

        /// <summary>
        /// Transport Attachment VPC ID.
        /// </summary>
        [Input("transportVpcId")]
        public Input<string>? TransportVpcId { get; set; }

        public AviatrixAwsTgwConnectState()
        {
        }
        public static new AviatrixAwsTgwConnectState Empty => new AviatrixAwsTgwConnectState();
    }
}

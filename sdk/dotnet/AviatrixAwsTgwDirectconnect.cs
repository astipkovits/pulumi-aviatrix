// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwDirectconnect:AviatrixAwsTgwDirectconnect")]
    public partial class AviatrixAwsTgwDirectconnect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Public IP address. Example: '40.0.0.0'.
        /// </summary>
        [Output("allowedPrefix")]
        public Output<string> AllowedPrefix { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of an Account in Aviatrix controller.
        /// </summary>
        [Output("directconnectAccountName")]
        public Output<string> DirectconnectAccountName { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of a Direct Connect Gateway ID.
        /// </summary>
        [Output("dxGatewayId")]
        public Output<string> DxGatewayId { get; private set; } = null!;

        /// <summary>
        /// Switch to enable/disable encrypted transit approval for direct connection. Valid values: true, false.
        /// </summary>
        [Output("enableLearnedCidrsApproval")]
        public Output<bool?> EnableLearnedCidrsApproval { get; private set; } = null!;

        /// <summary>
        /// The name of an Aviatrix network domain, to which the direct connect gateway will be attached.
        /// </summary>
        [Output("networkDomainName")]
        public Output<string?> NetworkDomainName { get; private set; } = null!;

        /// <summary>
        /// The name of an Aviatrix security domain, to which the direct connect gateway will be attached.
        /// </summary>
        [Output("securityDomainName")]
        public Output<string?> SecurityDomainName { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of an AWS TGW.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwDirectconnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwDirectconnect(string name, AviatrixAwsTgwDirectconnectArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwDirectconnect:AviatrixAwsTgwDirectconnect", name, args ?? new AviatrixAwsTgwDirectconnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwDirectconnect(string name, Input<string> id, AviatrixAwsTgwDirectconnectState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwDirectconnect:AviatrixAwsTgwDirectconnect", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsTgwDirectconnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwDirectconnect Get(string name, Input<string> id, AviatrixAwsTgwDirectconnectState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwDirectconnect(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwDirectconnectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public IP address. Example: '40.0.0.0'.
        /// </summary>
        [Input("allowedPrefix", required: true)]
        public Input<string> AllowedPrefix { get; set; } = null!;

        /// <summary>
        /// This parameter represents the name of an Account in Aviatrix controller.
        /// </summary>
        [Input("directconnectAccountName", required: true)]
        public Input<string> DirectconnectAccountName { get; set; } = null!;

        /// <summary>
        /// This parameter represents the name of a Direct Connect Gateway ID.
        /// </summary>
        [Input("dxGatewayId", required: true)]
        public Input<string> DxGatewayId { get; set; } = null!;

        /// <summary>
        /// Switch to enable/disable encrypted transit approval for direct connection. Valid values: true, false.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// The name of an Aviatrix network domain, to which the direct connect gateway will be attached.
        /// </summary>
        [Input("networkDomainName")]
        public Input<string>? NetworkDomainName { get; set; }

        /// <summary>
        /// The name of an Aviatrix security domain, to which the direct connect gateway will be attached.
        /// </summary>
        [Input("securityDomainName")]
        public Input<string>? SecurityDomainName { get; set; }

        /// <summary>
        /// This parameter represents the name of an AWS TGW.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        public AviatrixAwsTgwDirectconnectArgs()
        {
        }
        public static new AviatrixAwsTgwDirectconnectArgs Empty => new AviatrixAwsTgwDirectconnectArgs();
    }

    public sealed class AviatrixAwsTgwDirectconnectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public IP address. Example: '40.0.0.0'.
        /// </summary>
        [Input("allowedPrefix")]
        public Input<string>? AllowedPrefix { get; set; }

        /// <summary>
        /// This parameter represents the name of an Account in Aviatrix controller.
        /// </summary>
        [Input("directconnectAccountName")]
        public Input<string>? DirectconnectAccountName { get; set; }

        /// <summary>
        /// This parameter represents the name of a Direct Connect Gateway ID.
        /// </summary>
        [Input("dxGatewayId")]
        public Input<string>? DxGatewayId { get; set; }

        /// <summary>
        /// Switch to enable/disable encrypted transit approval for direct connection. Valid values: true, false.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// The name of an Aviatrix network domain, to which the direct connect gateway will be attached.
        /// </summary>
        [Input("networkDomainName")]
        public Input<string>? NetworkDomainName { get; set; }

        /// <summary>
        /// The name of an Aviatrix security domain, to which the direct connect gateway will be attached.
        /// </summary>
        [Input("securityDomainName")]
        public Input<string>? SecurityDomainName { get; set; }

        /// <summary>
        /// This parameter represents the name of an AWS TGW.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        public AviatrixAwsTgwDirectconnectState()
        {
        }
        public static new AviatrixAwsTgwDirectconnectState Empty => new AviatrixAwsTgwDirectconnectState();
    }
}

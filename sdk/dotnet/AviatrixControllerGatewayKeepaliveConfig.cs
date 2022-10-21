// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixControllerGatewayKeepaliveConfig:AviatrixControllerGatewayKeepaliveConfig")]
    public partial class AviatrixControllerGatewayKeepaliveConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Gateway keepalive speed.
        /// </summary>
        [Output("keepaliveSpeed")]
        public Output<string> KeepaliveSpeed { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixControllerGatewayKeepaliveConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixControllerGatewayKeepaliveConfig(string name, AviatrixControllerGatewayKeepaliveConfigArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixControllerGatewayKeepaliveConfig:AviatrixControllerGatewayKeepaliveConfig", name, args ?? new AviatrixControllerGatewayKeepaliveConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixControllerGatewayKeepaliveConfig(string name, Input<string> id, AviatrixControllerGatewayKeepaliveConfigState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixControllerGatewayKeepaliveConfig:AviatrixControllerGatewayKeepaliveConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixControllerGatewayKeepaliveConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixControllerGatewayKeepaliveConfig Get(string name, Input<string> id, AviatrixControllerGatewayKeepaliveConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixControllerGatewayKeepaliveConfig(name, id, state, options);
        }
    }

    public sealed class AviatrixControllerGatewayKeepaliveConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway keepalive speed.
        /// </summary>
        [Input("keepaliveSpeed", required: true)]
        public Input<string> KeepaliveSpeed { get; set; } = null!;

        public AviatrixControllerGatewayKeepaliveConfigArgs()
        {
        }
        public static new AviatrixControllerGatewayKeepaliveConfigArgs Empty => new AviatrixControllerGatewayKeepaliveConfigArgs();
    }

    public sealed class AviatrixControllerGatewayKeepaliveConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway keepalive speed.
        /// </summary>
        [Input("keepaliveSpeed")]
        public Input<string>? KeepaliveSpeed { get; set; }

        public AviatrixControllerGatewayKeepaliveConfigState()
        {
        }
        public static new AviatrixControllerGatewayKeepaliveConfigState Empty => new AviatrixControllerGatewayKeepaliveConfigState();
    }
}

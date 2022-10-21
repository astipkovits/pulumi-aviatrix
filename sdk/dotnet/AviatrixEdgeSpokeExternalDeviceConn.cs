// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn")]
    public partial class AviatrixEdgeSpokeExternalDeviceConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// BGP local AS number.
        /// </summary>
        [Output("bgpLocalAsNum")]
        public Output<string> BgpLocalAsNum { get; private set; } = null!;

        /// <summary>
        /// BGP remote AS number.
        /// </summary>
        [Output("bgpRemoteAsNum")]
        public Output<string> BgpRemoteAsNum { get; private set; } = null!;

        /// <summary>
        /// The name of the spoke external device connection which is going to be created.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// Connection type. Valid values: 'bgp'. Default value: 'bgp'.
        /// </summary>
        [Output("connectionType")]
        public Output<string?> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// Name of the BGP Spoke Gateway.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Local LAN IP.
        /// </summary>
        [Output("localLanIp")]
        public Output<string> LocalLanIp { get; private set; } = null!;

        /// <summary>
        /// Remote LAN IP.
        /// </summary>
        [Output("remoteLanIp")]
        public Output<string> RemoteLanIp { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC where the BGP Spoke Gateway is located.
        /// </summary>
        [Output("siteId")]
        public Output<string> SiteId { get; private set; } = null!;

        /// <summary>
        /// Tunnel Protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
        /// </summary>
        [Output("tunnelProtocol")]
        public Output<string?> TunnelProtocol { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixEdgeSpokeExternalDeviceConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixEdgeSpokeExternalDeviceConn(string name, AviatrixEdgeSpokeExternalDeviceConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn", name, args ?? new AviatrixEdgeSpokeExternalDeviceConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixEdgeSpokeExternalDeviceConn(string name, Input<string> id, AviatrixEdgeSpokeExternalDeviceConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixEdgeSpokeExternalDeviceConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixEdgeSpokeExternalDeviceConn Get(string name, Input<string> id, AviatrixEdgeSpokeExternalDeviceConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixEdgeSpokeExternalDeviceConn(name, id, state, options);
        }
    }

    public sealed class AviatrixEdgeSpokeExternalDeviceConnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP local AS number.
        /// </summary>
        [Input("bgpLocalAsNum", required: true)]
        public Input<string> BgpLocalAsNum { get; set; } = null!;

        /// <summary>
        /// BGP remote AS number.
        /// </summary>
        [Input("bgpRemoteAsNum", required: true)]
        public Input<string> BgpRemoteAsNum { get; set; } = null!;

        /// <summary>
        /// The name of the spoke external device connection which is going to be created.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Connection type. Valid values: 'bgp'. Default value: 'bgp'.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Name of the BGP Spoke Gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Local LAN IP.
        /// </summary>
        [Input("localLanIp", required: true)]
        public Input<string> LocalLanIp { get; set; } = null!;

        /// <summary>
        /// Remote LAN IP.
        /// </summary>
        [Input("remoteLanIp", required: true)]
        public Input<string> RemoteLanIp { get; set; } = null!;

        /// <summary>
        /// ID of the VPC where the BGP Spoke Gateway is located.
        /// </summary>
        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        /// <summary>
        /// Tunnel Protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
        /// </summary>
        [Input("tunnelProtocol")]
        public Input<string>? TunnelProtocol { get; set; }

        public AviatrixEdgeSpokeExternalDeviceConnArgs()
        {
        }
        public static new AviatrixEdgeSpokeExternalDeviceConnArgs Empty => new AviatrixEdgeSpokeExternalDeviceConnArgs();
    }

    public sealed class AviatrixEdgeSpokeExternalDeviceConnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP local AS number.
        /// </summary>
        [Input("bgpLocalAsNum")]
        public Input<string>? BgpLocalAsNum { get; set; }

        /// <summary>
        /// BGP remote AS number.
        /// </summary>
        [Input("bgpRemoteAsNum")]
        public Input<string>? BgpRemoteAsNum { get; set; }

        /// <summary>
        /// The name of the spoke external device connection which is going to be created.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Connection type. Valid values: 'bgp'. Default value: 'bgp'.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Name of the BGP Spoke Gateway.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// Local LAN IP.
        /// </summary>
        [Input("localLanIp")]
        public Input<string>? LocalLanIp { get; set; }

        /// <summary>
        /// Remote LAN IP.
        /// </summary>
        [Input("remoteLanIp")]
        public Input<string>? RemoteLanIp { get; set; }

        /// <summary>
        /// ID of the VPC where the BGP Spoke Gateway is located.
        /// </summary>
        [Input("siteId")]
        public Input<string>? SiteId { get; set; }

        /// <summary>
        /// Tunnel Protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
        /// </summary>
        [Input("tunnelProtocol")]
        public Input<string>? TunnelProtocol { get; set; }

        public AviatrixEdgeSpokeExternalDeviceConnState()
        {
        }
        public static new AviatrixEdgeSpokeExternalDeviceConnState Empty => new AviatrixEdgeSpokeExternalDeviceConnState();
    }
}

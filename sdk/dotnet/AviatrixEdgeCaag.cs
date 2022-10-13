// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixEdgeCaag:AviatrixEdgeCaag")]
    public partial class AviatrixEdgeCaag : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DNS server IP.
        /// </summary>
        [Output("dnsServerIp")]
        public Output<string?> DnsServerIp { get; private set; } = null!;

        /// <summary>
        /// Enable management over private network.
        /// </summary>
        [Output("enableOverPrivateNetwork")]
        public Output<bool?> EnableOverPrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// LAN interface IP / prefix.
        /// </summary>
        [Output("lanInterfaceIpPrefix")]
        public Output<string> LanInterfaceIpPrefix { get; private set; } = null!;

        /// <summary>
        /// Local AS number.
        /// </summary>
        [Output("localAsNumber")]
        public Output<string> LocalAsNumber { get; private set; } = null!;

        /// <summary>
        /// Management default gateway IP.
        /// </summary>
        [Output("managementDefaultGatewayIp")]
        public Output<string?> ManagementDefaultGatewayIp { get; private set; } = null!;

        /// <summary>
        /// Management egress gateway IP / prefix.
        /// </summary>
        [Output("managementEgressIpPrefix")]
        public Output<string?> ManagementEgressIpPrefix { get; private set; } = null!;

        /// <summary>
        /// Management interface configuration. Valid values: 'DHCP' and 'Static'.
        /// </summary>
        [Output("managementInterfaceConfig")]
        public Output<string> ManagementInterfaceConfig { get; private set; } = null!;

        /// <summary>
        /// Management interface IP / prefix.
        /// </summary>
        [Output("managementInterfaceIpPrefix")]
        public Output<string?> ManagementInterfaceIpPrefix { get; private set; } = null!;

        /// <summary>
        /// Edge as a CaaG name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// AS path prepend.
        /// </summary>
        [Output("prependAsPaths")]
        public Output<ImmutableArray<string>> PrependAsPaths { get; private set; } = null!;

        /// <summary>
        /// Secondary DNS server IP.
        /// </summary>
        [Output("secondaryDnsServerIp")]
        public Output<string?> SecondaryDnsServerIp { get; private set; } = null!;

        /// <summary>
        /// State of Edge as a CaaG.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// WAN default gateway IP.
        /// </summary>
        [Output("wanDefaultGatewayIp")]
        public Output<string> WanDefaultGatewayIp { get; private set; } = null!;

        /// <summary>
        /// WAN interface IP / prefix.
        /// </summary>
        [Output("wanInterfaceIpPrefix")]
        public Output<string> WanInterfaceIpPrefix { get; private set; } = null!;

        /// <summary>
        /// The location where the Edge as a CaaG ZTP file will be stored.
        /// </summary>
        [Output("ztpFileDownloadPath")]
        public Output<string> ZtpFileDownloadPath { get; private set; } = null!;

        /// <summary>
        /// ZTP file type.
        /// </summary>
        [Output("ztpFileType")]
        public Output<string> ZtpFileType { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixEdgeCaag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixEdgeCaag(string name, AviatrixEdgeCaagArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixEdgeCaag:AviatrixEdgeCaag", name, args ?? new AviatrixEdgeCaagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixEdgeCaag(string name, Input<string> id, AviatrixEdgeCaagState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixEdgeCaag:AviatrixEdgeCaag", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/astipkovits/pulumi-aviatrix/raw/main/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixEdgeCaag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixEdgeCaag Get(string name, Input<string> id, AviatrixEdgeCaagState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixEdgeCaag(name, id, state, options);
        }
    }

    public sealed class AviatrixEdgeCaagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS server IP.
        /// </summary>
        [Input("dnsServerIp")]
        public Input<string>? DnsServerIp { get; set; }

        /// <summary>
        /// Enable management over private network.
        /// </summary>
        [Input("enableOverPrivateNetwork")]
        public Input<bool>? EnableOverPrivateNetwork { get; set; }

        /// <summary>
        /// LAN interface IP / prefix.
        /// </summary>
        [Input("lanInterfaceIpPrefix", required: true)]
        public Input<string> LanInterfaceIpPrefix { get; set; } = null!;

        /// <summary>
        /// Local AS number.
        /// </summary>
        [Input("localAsNumber")]
        public Input<string>? LocalAsNumber { get; set; }

        /// <summary>
        /// Management default gateway IP.
        /// </summary>
        [Input("managementDefaultGatewayIp")]
        public Input<string>? ManagementDefaultGatewayIp { get; set; }

        /// <summary>
        /// Management egress gateway IP / prefix.
        /// </summary>
        [Input("managementEgressIpPrefix")]
        public Input<string>? ManagementEgressIpPrefix { get; set; }

        /// <summary>
        /// Management interface configuration. Valid values: 'DHCP' and 'Static'.
        /// </summary>
        [Input("managementInterfaceConfig", required: true)]
        public Input<string> ManagementInterfaceConfig { get; set; } = null!;

        /// <summary>
        /// Management interface IP / prefix.
        /// </summary>
        [Input("managementInterfaceIpPrefix")]
        public Input<string>? ManagementInterfaceIpPrefix { get; set; }

        /// <summary>
        /// Edge as a CaaG name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// AS path prepend.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// Secondary DNS server IP.
        /// </summary>
        [Input("secondaryDnsServerIp")]
        public Input<string>? SecondaryDnsServerIp { get; set; }

        /// <summary>
        /// WAN default gateway IP.
        /// </summary>
        [Input("wanDefaultGatewayIp", required: true)]
        public Input<string> WanDefaultGatewayIp { get; set; } = null!;

        /// <summary>
        /// WAN interface IP / prefix.
        /// </summary>
        [Input("wanInterfaceIpPrefix", required: true)]
        public Input<string> WanInterfaceIpPrefix { get; set; } = null!;

        /// <summary>
        /// The location where the Edge as a CaaG ZTP file will be stored.
        /// </summary>
        [Input("ztpFileDownloadPath", required: true)]
        public Input<string> ZtpFileDownloadPath { get; set; } = null!;

        /// <summary>
        /// ZTP file type.
        /// </summary>
        [Input("ztpFileType", required: true)]
        public Input<string> ZtpFileType { get; set; } = null!;

        public AviatrixEdgeCaagArgs()
        {
        }
        public static new AviatrixEdgeCaagArgs Empty => new AviatrixEdgeCaagArgs();
    }

    public sealed class AviatrixEdgeCaagState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS server IP.
        /// </summary>
        [Input("dnsServerIp")]
        public Input<string>? DnsServerIp { get; set; }

        /// <summary>
        /// Enable management over private network.
        /// </summary>
        [Input("enableOverPrivateNetwork")]
        public Input<bool>? EnableOverPrivateNetwork { get; set; }

        /// <summary>
        /// LAN interface IP / prefix.
        /// </summary>
        [Input("lanInterfaceIpPrefix")]
        public Input<string>? LanInterfaceIpPrefix { get; set; }

        /// <summary>
        /// Local AS number.
        /// </summary>
        [Input("localAsNumber")]
        public Input<string>? LocalAsNumber { get; set; }

        /// <summary>
        /// Management default gateway IP.
        /// </summary>
        [Input("managementDefaultGatewayIp")]
        public Input<string>? ManagementDefaultGatewayIp { get; set; }

        /// <summary>
        /// Management egress gateway IP / prefix.
        /// </summary>
        [Input("managementEgressIpPrefix")]
        public Input<string>? ManagementEgressIpPrefix { get; set; }

        /// <summary>
        /// Management interface configuration. Valid values: 'DHCP' and 'Static'.
        /// </summary>
        [Input("managementInterfaceConfig")]
        public Input<string>? ManagementInterfaceConfig { get; set; }

        /// <summary>
        /// Management interface IP / prefix.
        /// </summary>
        [Input("managementInterfaceIpPrefix")]
        public Input<string>? ManagementInterfaceIpPrefix { get; set; }

        /// <summary>
        /// Edge as a CaaG name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// AS path prepend.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// Secondary DNS server IP.
        /// </summary>
        [Input("secondaryDnsServerIp")]
        public Input<string>? SecondaryDnsServerIp { get; set; }

        /// <summary>
        /// State of Edge as a CaaG.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// WAN default gateway IP.
        /// </summary>
        [Input("wanDefaultGatewayIp")]
        public Input<string>? WanDefaultGatewayIp { get; set; }

        /// <summary>
        /// WAN interface IP / prefix.
        /// </summary>
        [Input("wanInterfaceIpPrefix")]
        public Input<string>? WanInterfaceIpPrefix { get; set; }

        /// <summary>
        /// The location where the Edge as a CaaG ZTP file will be stored.
        /// </summary>
        [Input("ztpFileDownloadPath")]
        public Input<string>? ZtpFileDownloadPath { get; set; }

        /// <summary>
        /// ZTP file type.
        /// </summary>
        [Input("ztpFileType")]
        public Input<string>? ZtpFileType { get; set; }

        public AviatrixEdgeCaagState()
        {
        }
        public static new AviatrixEdgeCaagState Empty => new AviatrixEdgeCaagState();
    }
}
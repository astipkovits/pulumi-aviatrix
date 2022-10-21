// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixTransitCloudnConn:AviatrixTransitCloudnConn")]
    public partial class AviatrixTransitCloudnConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set of approved cidrs. Requires 'enable_learned_cidrs_approval' to be true. Type: Set(String).
        /// </summary>
        [Output("approvedCidrs")]
        public Output<ImmutableArray<string>> ApprovedCidrs { get; private set; } = null!;

        /// <summary>
        /// Backup Aviatrix CloudN BGP ASN.
        /// </summary>
        [Output("backupCloudnAsNum")]
        public Output<string?> BackupCloudnAsNum { get; private set; } = null!;

        /// <summary>
        /// Backup Aviatrix CloudN IP Address.
        /// </summary>
        [Output("backupCloudnIp")]
        public Output<string?> BackupCloudnIp { get; private set; } = null!;

        /// <summary>
        /// Backup Aviatrix CloudN LAN Interface Neighbor's BGP ASN.
        /// </summary>
        [Output("backupCloudnNeighborAsNum")]
        public Output<string?> BackupCloudnNeighborAsNum { get; private set; } = null!;

        /// <summary>
        /// Backup Aviatrix CloudN LAN Interface Neighbor's IP Address.
        /// </summary>
        [Output("backupCloudnNeighborIp")]
        public Output<string?> BackupCloudnNeighborIp { get; private set; } = null!;

        /// <summary>
        /// Enable direct connect to Backup Aviatrix CloudN over private network.
        /// </summary>
        [Output("backupDirectConnect")]
        public Output<bool?> BackupDirectConnect { get; private set; } = null!;

        /// <summary>
        /// Enable Insane Mode for connection to Backup Aviatrix CloudN.
        /// </summary>
        [Output("backupInsaneMode")]
        public Output<bool?> BackupInsaneMode { get; private set; } = null!;

        /// <summary>
        /// BGP local ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Output("bgpLocalAsNum")]
        public Output<string> BgpLocalAsNum { get; private set; } = null!;

        /// <summary>
        /// Aviatrix CloudN BGP ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Output("cloudnAsNum")]
        public Output<string> CloudnAsNum { get; private set; } = null!;

        /// <summary>
        /// CloudN LAN Interface Neighbor's BGP ASN.
        /// </summary>
        [Output("cloudnNeighborAsNum")]
        public Output<string> CloudnNeighborAsNum { get; private set; } = null!;

        /// <summary>
        /// Aviatrix CloudN LAN Interface Neighbor's IP Address.
        /// </summary>
        [Output("cloudnNeighborIp")]
        public Output<string> CloudnNeighborIp { get; private set; } = null!;

        /// <summary>
        /// Aviatrix CloudN IP Address.
        /// </summary>
        [Output("cloudnRemoteIp")]
        public Output<string> CloudnRemoteIp { get; private set; } = null!;

        /// <summary>
        /// The name of the transit Aviatrix CloudN connection.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// Enable Direct Connect for private network infrastructure.
        /// </summary>
        [Output("directConnect")]
        public Output<bool?> DirectConnect { get; private set; } = null!;

        /// <summary>
        /// Enable connection to HA CloudN.
        /// </summary>
        [Output("enableHa")]
        public Output<bool?> EnableHa { get; private set; } = null!;

        /// <summary>
        /// Enable learned CIDRs approval.
        /// </summary>
        [Output("enableLearnedCidrsApproval")]
        public Output<bool?> EnableLearnedCidrsApproval { get; private set; } = null!;

        /// <summary>
        /// Enable load balancing between Aviatrix CloudN and Backup CloudN.
        /// </summary>
        [Output("enableLoadBalancing")]
        public Output<bool?> EnableLoadBalancing { get; private set; } = null!;

        /// <summary>
        /// The name of the Transit Gateway.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Enable Insane Mode for this connection.
        /// </summary>
        [Output("insaneMode")]
        public Output<bool?> InsaneMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC where the Transit Gateway is located.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixTransitCloudnConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixTransitCloudnConn(string name, AviatrixTransitCloudnConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixTransitCloudnConn:AviatrixTransitCloudnConn", name, args ?? new AviatrixTransitCloudnConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixTransitCloudnConn(string name, Input<string> id, AviatrixTransitCloudnConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixTransitCloudnConn:AviatrixTransitCloudnConn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixTransitCloudnConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixTransitCloudnConn Get(string name, Input<string> id, AviatrixTransitCloudnConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixTransitCloudnConn(name, id, state, options);
        }
    }

    public sealed class AviatrixTransitCloudnConnArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvedCidrs")]
        private InputList<string>? _approvedCidrs;

        /// <summary>
        /// Set of approved cidrs. Requires 'enable_learned_cidrs_approval' to be true. Type: Set(String).
        /// </summary>
        public InputList<string> ApprovedCidrs
        {
            get => _approvedCidrs ?? (_approvedCidrs = new InputList<string>());
            set => _approvedCidrs = value;
        }

        /// <summary>
        /// Backup Aviatrix CloudN BGP ASN.
        /// </summary>
        [Input("backupCloudnAsNum")]
        public Input<string>? BackupCloudnAsNum { get; set; }

        /// <summary>
        /// Backup Aviatrix CloudN IP Address.
        /// </summary>
        [Input("backupCloudnIp")]
        public Input<string>? BackupCloudnIp { get; set; }

        /// <summary>
        /// Backup Aviatrix CloudN LAN Interface Neighbor's BGP ASN.
        /// </summary>
        [Input("backupCloudnNeighborAsNum")]
        public Input<string>? BackupCloudnNeighborAsNum { get; set; }

        /// <summary>
        /// Backup Aviatrix CloudN LAN Interface Neighbor's IP Address.
        /// </summary>
        [Input("backupCloudnNeighborIp")]
        public Input<string>? BackupCloudnNeighborIp { get; set; }

        /// <summary>
        /// Enable direct connect to Backup Aviatrix CloudN over private network.
        /// </summary>
        [Input("backupDirectConnect")]
        public Input<bool>? BackupDirectConnect { get; set; }

        /// <summary>
        /// Enable Insane Mode for connection to Backup Aviatrix CloudN.
        /// </summary>
        [Input("backupInsaneMode")]
        public Input<bool>? BackupInsaneMode { get; set; }

        /// <summary>
        /// BGP local ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("bgpLocalAsNum", required: true)]
        public Input<string> BgpLocalAsNum { get; set; } = null!;

        /// <summary>
        /// Aviatrix CloudN BGP ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("cloudnAsNum", required: true)]
        public Input<string> CloudnAsNum { get; set; } = null!;

        /// <summary>
        /// CloudN LAN Interface Neighbor's BGP ASN.
        /// </summary>
        [Input("cloudnNeighborAsNum", required: true)]
        public Input<string> CloudnNeighborAsNum { get; set; } = null!;

        /// <summary>
        /// Aviatrix CloudN LAN Interface Neighbor's IP Address.
        /// </summary>
        [Input("cloudnNeighborIp", required: true)]
        public Input<string> CloudnNeighborIp { get; set; } = null!;

        /// <summary>
        /// Aviatrix CloudN IP Address.
        /// </summary>
        [Input("cloudnRemoteIp", required: true)]
        public Input<string> CloudnRemoteIp { get; set; } = null!;

        /// <summary>
        /// The name of the transit Aviatrix CloudN connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Enable Direct Connect for private network infrastructure.
        /// </summary>
        [Input("directConnect")]
        public Input<bool>? DirectConnect { get; set; }

        /// <summary>
        /// Enable connection to HA CloudN.
        /// </summary>
        [Input("enableHa")]
        public Input<bool>? EnableHa { get; set; }

        /// <summary>
        /// Enable learned CIDRs approval.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Enable load balancing between Aviatrix CloudN and Backup CloudN.
        /// </summary>
        [Input("enableLoadBalancing")]
        public Input<bool>? EnableLoadBalancing { get; set; }

        /// <summary>
        /// The name of the Transit Gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Enable Insane Mode for this connection.
        /// </summary>
        [Input("insaneMode")]
        public Input<bool>? InsaneMode { get; set; }

        /// <summary>
        /// The ID of the VPC where the Transit Gateway is located.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixTransitCloudnConnArgs()
        {
        }
        public static new AviatrixTransitCloudnConnArgs Empty => new AviatrixTransitCloudnConnArgs();
    }

    public sealed class AviatrixTransitCloudnConnState : global::Pulumi.ResourceArgs
    {
        [Input("approvedCidrs")]
        private InputList<string>? _approvedCidrs;

        /// <summary>
        /// Set of approved cidrs. Requires 'enable_learned_cidrs_approval' to be true. Type: Set(String).
        /// </summary>
        public InputList<string> ApprovedCidrs
        {
            get => _approvedCidrs ?? (_approvedCidrs = new InputList<string>());
            set => _approvedCidrs = value;
        }

        /// <summary>
        /// Backup Aviatrix CloudN BGP ASN.
        /// </summary>
        [Input("backupCloudnAsNum")]
        public Input<string>? BackupCloudnAsNum { get; set; }

        /// <summary>
        /// Backup Aviatrix CloudN IP Address.
        /// </summary>
        [Input("backupCloudnIp")]
        public Input<string>? BackupCloudnIp { get; set; }

        /// <summary>
        /// Backup Aviatrix CloudN LAN Interface Neighbor's BGP ASN.
        /// </summary>
        [Input("backupCloudnNeighborAsNum")]
        public Input<string>? BackupCloudnNeighborAsNum { get; set; }

        /// <summary>
        /// Backup Aviatrix CloudN LAN Interface Neighbor's IP Address.
        /// </summary>
        [Input("backupCloudnNeighborIp")]
        public Input<string>? BackupCloudnNeighborIp { get; set; }

        /// <summary>
        /// Enable direct connect to Backup Aviatrix CloudN over private network.
        /// </summary>
        [Input("backupDirectConnect")]
        public Input<bool>? BackupDirectConnect { get; set; }

        /// <summary>
        /// Enable Insane Mode for connection to Backup Aviatrix CloudN.
        /// </summary>
        [Input("backupInsaneMode")]
        public Input<bool>? BackupInsaneMode { get; set; }

        /// <summary>
        /// BGP local ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("bgpLocalAsNum")]
        public Input<string>? BgpLocalAsNum { get; set; }

        /// <summary>
        /// Aviatrix CloudN BGP ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("cloudnAsNum")]
        public Input<string>? CloudnAsNum { get; set; }

        /// <summary>
        /// CloudN LAN Interface Neighbor's BGP ASN.
        /// </summary>
        [Input("cloudnNeighborAsNum")]
        public Input<string>? CloudnNeighborAsNum { get; set; }

        /// <summary>
        /// Aviatrix CloudN LAN Interface Neighbor's IP Address.
        /// </summary>
        [Input("cloudnNeighborIp")]
        public Input<string>? CloudnNeighborIp { get; set; }

        /// <summary>
        /// Aviatrix CloudN IP Address.
        /// </summary>
        [Input("cloudnRemoteIp")]
        public Input<string>? CloudnRemoteIp { get; set; }

        /// <summary>
        /// The name of the transit Aviatrix CloudN connection.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Enable Direct Connect for private network infrastructure.
        /// </summary>
        [Input("directConnect")]
        public Input<bool>? DirectConnect { get; set; }

        /// <summary>
        /// Enable connection to HA CloudN.
        /// </summary>
        [Input("enableHa")]
        public Input<bool>? EnableHa { get; set; }

        /// <summary>
        /// Enable learned CIDRs approval.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Enable load balancing between Aviatrix CloudN and Backup CloudN.
        /// </summary>
        [Input("enableLoadBalancing")]
        public Input<bool>? EnableLoadBalancing { get; set; }

        /// <summary>
        /// The name of the Transit Gateway.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// Enable Insane Mode for this connection.
        /// </summary>
        [Input("insaneMode")]
        public Input<bool>? InsaneMode { get; set; }

        /// <summary>
        /// The ID of the VPC where the Transit Gateway is located.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixTransitCloudnConnState()
        {
        }
        public static new AviatrixTransitCloudnConnState Empty => new AviatrixTransitCloudnConnState();
    }
}

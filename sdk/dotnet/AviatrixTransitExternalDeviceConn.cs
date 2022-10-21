// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixTransitExternalDeviceConn:AviatrixTransitExternalDeviceConn")]
    public partial class AviatrixTransitExternalDeviceConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set of approved cidrs. Requires 'enable_learned_cidrs_approval' to be true. Type: Set(String).
        /// </summary>
        [Output("approvedCidrs")]
        public Output<ImmutableArray<string>> ApprovedCidrs { get; private set; } = null!;

        /// <summary>
        /// Backup BGP MD5 authentication key.
        /// </summary>
        [Output("backupBgpMd5Key")]
        public Output<string?> BackupBgpMd5Key { get; private set; } = null!;

        /// <summary>
        /// Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Output("backupBgpRemoteAsNum")]
        public Output<string?> BackupBgpRemoteAsNum { get; private set; } = null!;

        /// <summary>
        /// Backup direct connect for backup external device.
        /// </summary>
        [Output("backupDirectConnect")]
        public Output<bool?> BackupDirectConnect { get; private set; } = null!;

        /// <summary>
        /// Backup Local LAN IP. Required for GCP BGP over LAN Connection with HA enabled.
        /// </summary>
        [Output("backupLocalLanIp")]
        public Output<string> BackupLocalLanIp { get; private set; } = null!;

        /// <summary>
        /// Source CIDR for the tunnel from the backup Aviatrix transit gateway.
        /// </summary>
        [Output("backupLocalTunnelCidr")]
        public Output<string> BackupLocalTunnelCidr { get; private set; } = null!;

        /// <summary>
        /// Backup pre shared key.
        /// </summary>
        [Output("backupPreSharedKey")]
        public Output<string?> BackupPreSharedKey { get; private set; } = null!;

        /// <summary>
        /// Backup remote gateway IP.
        /// </summary>
        [Output("backupRemoteGatewayIp")]
        public Output<string?> BackupRemoteGatewayIp { get; private set; } = null!;

        /// <summary>
        /// Backup Remote LAN IP.
        /// </summary>
        [Output("backupRemoteLanIp")]
        public Output<string?> BackupRemoteLanIp { get; private set; } = null!;

        /// <summary>
        /// Destination CIDR for the tunnel to the backup external device.
        /// </summary>
        [Output("backupRemoteTunnelCidr")]
        public Output<string> BackupRemoteTunnelCidr { get; private set; } = null!;

        /// <summary>
        /// BGP local ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Output("bgpLocalAsNum")]
        public Output<string?> BgpLocalAsNum { get; private set; } = null!;

        /// <summary>
        /// BGP MD5 authentication key.
        /// </summary>
        [Output("bgpMd5Key")]
        public Output<string?> BgpMd5Key { get; private set; } = null!;

        /// <summary>
        /// BGP remote ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Output("bgpRemoteAsNum")]
        public Output<string?> BgpRemoteAsNum { get; private set; } = null!;

        /// <summary>
        /// The name of the transit external device connection which is going to be created.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'.
        /// </summary>
        [Output("connectionType")]
        public Output<string?> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption.
        /// </summary>
        [Output("customAlgorithms")]
        public Output<bool?> CustomAlgorithms { get; private set; } = null!;

        /// <summary>
        /// Set true for private network infrastructure.
        /// </summary>
        [Output("directConnect")]
        public Output<bool?> DirectConnect { get; private set; } = null!;

        /// <summary>
        /// Switch to enable BGP LAN ActiveMesh. Only valid for GCP with Remote Gateway HA enabled. Default: false. Available as of
        /// provider version R2.21+.
        /// </summary>
        [Output("enableBgpLanActivemesh")]
        public Output<bool?> EnableBgpLanActivemesh { get; private set; } = null!;

        /// <summary>
        /// Switch to allow this connection to communicate with a Network Domain via Connection Policy.
        /// </summary>
        [Output("enableEdgeSegmentation")]
        public Output<bool?> EnableEdgeSegmentation { get; private set; } = null!;

        /// <summary>
        /// Enable Event Triggered HA.
        /// </summary>
        [Output("enableEventTriggeredHa")]
        public Output<bool?> EnableEventTriggeredHa { get; private set; } = null!;

        /// <summary>
        /// Set as true if use IKEv2.
        /// </summary>
        [Output("enableIkev2")]
        public Output<bool?> EnableIkev2 { get; private set; } = null!;

        /// <summary>
        /// Enable Jumbo Frame for the transit external device connection. Valid values: true, false.
        /// </summary>
        [Output("enableJumboFrame")]
        public Output<bool?> EnableJumboFrame { get; private set; } = null!;

        /// <summary>
        /// Enable learned CIDR approval for the connection. Only valid with 'connection_type' = 'bgp'. Requires the
        /// transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default
        /// value: false. Available as of provider version R2.18+.
        /// </summary>
        [Output("enableLearnedCidrsApproval")]
        public Output<bool?> EnableLearnedCidrsApproval { get; private set; } = null!;

        /// <summary>
        /// Name of the Transit Gateway.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Set as true if there are two external devices.
        /// </summary>
        [Output("haEnabled")]
        public Output<bool?> HaEnabled { get; private set; } = null!;

        /// <summary>
        /// Local LAN IP. Required for GCP BGP over LAN Connection.
        /// </summary>
        [Output("localLanIp")]
        public Output<string> LocalLanIp { get; private set; } = null!;

        /// <summary>
        /// Source CIDR for the tunnel from the Aviatrix transit gateway.
        /// </summary>
        [Output("localTunnelCidr")]
        public Output<string> LocalTunnelCidr { get; private set; } = null!;

        /// <summary>
        /// Configure manual BGP advertised CIDRs for this connection. Only valid with 'connection_type' = 'bgp'. Available as of
        /// provider version R2.18+.
        /// </summary>
        [Output("manualBgpAdvertisedCidrs")]
        public Output<ImmutableArray<string>> ManualBgpAdvertisedCidrs { get; private set; } = null!;

        /// <summary>
        /// Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'.
        /// </summary>
        [Output("phase1Authentication")]
        public Output<string?> Phase1Authentication { get; private set; } = null!;

        /// <summary>
        /// Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'.
        /// </summary>
        [Output("phase1DhGroups")]
        public Output<string?> Phase1DhGroups { get; private set; } = null!;

        /// <summary>
        /// Phase one Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC' and 'AES-256-CBC', 'AES-128-GCM-64',
        /// 'AES-128-GCM-96', 'AES-128-GCM-128', 'AES-256-GCM-64', 'AES-256-GCM-96', and 'AES-256-GCM-128'.
        /// </summary>
        [Output("phase1Encryption")]
        public Output<string?> Phase1Encryption { get; private set; } = null!;

        /// <summary>
        /// Phase 1 remote identifier of the IPsec tunnel.
        /// </summary>
        [Output("phase1RemoteIdentifiers")]
        public Output<ImmutableArray<string>> Phase1RemoteIdentifiers { get; private set; } = null!;

        /// <summary>
        /// Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'.
        /// </summary>
        [Output("phase2Authentication")]
        public Output<string?> Phase2Authentication { get; private set; } = null!;

        /// <summary>
        /// Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'.
        /// </summary>
        [Output("phase2DhGroups")]
        public Output<string?> Phase2DhGroups { get; private set; } = null!;

        /// <summary>
        /// Phase two Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC', 'AES-256-CBC', 'AES-128-GCM-64',
        /// 'AES-128-GCM-96', 'AES-128-GCM-128', 'AES-256-GCM-64', 'AES-256-GCM-96', 'AES-256-GCM-128', and 'NULL-ENCR'.
        /// </summary>
        [Output("phase2Encryption")]
        public Output<string?> Phase2Encryption { get; private set; } = null!;

        /// <summary>
        /// If left blank, the pre-shared key will be auto generated.
        /// </summary>
        [Output("preSharedKey")]
        public Output<string?> PreSharedKey { get; private set; } = null!;

        /// <summary>
        /// Connection AS Path Prepend customized by specifying AS PATH for a BGP connection.
        /// </summary>
        [Output("prependAsPaths")]
        public Output<ImmutableArray<string>> PrependAsPaths { get; private set; } = null!;

        /// <summary>
        /// Remote Gateway IP.
        /// </summary>
        [Output("remoteGatewayIp")]
        public Output<string?> RemoteGatewayIp { get; private set; } = null!;

        /// <summary>
        /// Remote LAN IP.
        /// </summary>
        [Output("remoteLanIp")]
        public Output<string?> RemoteLanIp { get; private set; } = null!;

        /// <summary>
        /// Remote CIDRs joined as a string with ','. Required for a 'static' type connection.
        /// </summary>
        [Output("remoteSubnet")]
        public Output<string?> RemoteSubnet { get; private set; } = null!;

        /// <summary>
        /// Destination CIDR for the tunnel to the external device.
        /// </summary>
        [Output("remoteTunnelCidr")]
        public Output<string> RemoteTunnelCidr { get; private set; } = null!;

        /// <summary>
        /// Name of the remote VPC for a LAN BGP connection. Only valid when 'connection_type' = 'bgp' and tunnel_protocol' = 'LAN'
        /// with an Azure transit gateway. Must be in the form "&lt;VNET-name&gt;:&lt;resource-group-name&gt;". Available as of provider version
        /// R2.18+.
        /// </summary>
        [Output("remoteVpcName")]
        public Output<string?> RemoteVpcName { get; private set; } = null!;

        /// <summary>
        /// Only valid for Transit Gateway's with Active-Standby Mode enabled. Valid values: true, false. Default: false.
        /// </summary>
        [Output("switchToHaStandbyGateway")]
        public Output<bool?> SwitchToHaStandbyGateway { get; private set; } = null!;

        /// <summary>
        /// Tunnel Protocol. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive.
        /// </summary>
        [Output("tunnelProtocol")]
        public Output<string?> TunnelProtocol { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC where the Transit Gateway is located. For GCP BGP over LAN connection, it is in the format of
        /// 'vpc_name~-~account_name'.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixTransitExternalDeviceConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixTransitExternalDeviceConn(string name, AviatrixTransitExternalDeviceConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixTransitExternalDeviceConn:AviatrixTransitExternalDeviceConn", name, args ?? new AviatrixTransitExternalDeviceConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixTransitExternalDeviceConn(string name, Input<string> id, AviatrixTransitExternalDeviceConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixTransitExternalDeviceConn:AviatrixTransitExternalDeviceConn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixTransitExternalDeviceConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixTransitExternalDeviceConn Get(string name, Input<string> id, AviatrixTransitExternalDeviceConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixTransitExternalDeviceConn(name, id, state, options);
        }
    }

    public sealed class AviatrixTransitExternalDeviceConnArgs : global::Pulumi.ResourceArgs
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
        /// Backup BGP MD5 authentication key.
        /// </summary>
        [Input("backupBgpMd5Key")]
        public Input<string>? BackupBgpMd5Key { get; set; }

        /// <summary>
        /// Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("backupBgpRemoteAsNum")]
        public Input<string>? BackupBgpRemoteAsNum { get; set; }

        /// <summary>
        /// Backup direct connect for backup external device.
        /// </summary>
        [Input("backupDirectConnect")]
        public Input<bool>? BackupDirectConnect { get; set; }

        /// <summary>
        /// Backup Local LAN IP. Required for GCP BGP over LAN Connection with HA enabled.
        /// </summary>
        [Input("backupLocalLanIp")]
        public Input<string>? BackupLocalLanIp { get; set; }

        /// <summary>
        /// Source CIDR for the tunnel from the backup Aviatrix transit gateway.
        /// </summary>
        [Input("backupLocalTunnelCidr")]
        public Input<string>? BackupLocalTunnelCidr { get; set; }

        /// <summary>
        /// Backup pre shared key.
        /// </summary>
        [Input("backupPreSharedKey")]
        public Input<string>? BackupPreSharedKey { get; set; }

        /// <summary>
        /// Backup remote gateway IP.
        /// </summary>
        [Input("backupRemoteGatewayIp")]
        public Input<string>? BackupRemoteGatewayIp { get; set; }

        /// <summary>
        /// Backup Remote LAN IP.
        /// </summary>
        [Input("backupRemoteLanIp")]
        public Input<string>? BackupRemoteLanIp { get; set; }

        /// <summary>
        /// Destination CIDR for the tunnel to the backup external device.
        /// </summary>
        [Input("backupRemoteTunnelCidr")]
        public Input<string>? BackupRemoteTunnelCidr { get; set; }

        /// <summary>
        /// BGP local ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("bgpLocalAsNum")]
        public Input<string>? BgpLocalAsNum { get; set; }

        /// <summary>
        /// BGP MD5 authentication key.
        /// </summary>
        [Input("bgpMd5Key")]
        public Input<string>? BgpMd5Key { get; set; }

        /// <summary>
        /// BGP remote ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("bgpRemoteAsNum")]
        public Input<string>? BgpRemoteAsNum { get; set; }

        /// <summary>
        /// The name of the transit external device connection which is going to be created.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption.
        /// </summary>
        [Input("customAlgorithms")]
        public Input<bool>? CustomAlgorithms { get; set; }

        /// <summary>
        /// Set true for private network infrastructure.
        /// </summary>
        [Input("directConnect")]
        public Input<bool>? DirectConnect { get; set; }

        /// <summary>
        /// Switch to enable BGP LAN ActiveMesh. Only valid for GCP with Remote Gateway HA enabled. Default: false. Available as of
        /// provider version R2.21+.
        /// </summary>
        [Input("enableBgpLanActivemesh")]
        public Input<bool>? EnableBgpLanActivemesh { get; set; }

        /// <summary>
        /// Switch to allow this connection to communicate with a Network Domain via Connection Policy.
        /// </summary>
        [Input("enableEdgeSegmentation")]
        public Input<bool>? EnableEdgeSegmentation { get; set; }

        /// <summary>
        /// Enable Event Triggered HA.
        /// </summary>
        [Input("enableEventTriggeredHa")]
        public Input<bool>? EnableEventTriggeredHa { get; set; }

        /// <summary>
        /// Set as true if use IKEv2.
        /// </summary>
        [Input("enableIkev2")]
        public Input<bool>? EnableIkev2 { get; set; }

        /// <summary>
        /// Enable Jumbo Frame for the transit external device connection. Valid values: true, false.
        /// </summary>
        [Input("enableJumboFrame")]
        public Input<bool>? EnableJumboFrame { get; set; }

        /// <summary>
        /// Enable learned CIDR approval for the connection. Only valid with 'connection_type' = 'bgp'. Requires the
        /// transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default
        /// value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Name of the Transit Gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Set as true if there are two external devices.
        /// </summary>
        [Input("haEnabled")]
        public Input<bool>? HaEnabled { get; set; }

        /// <summary>
        /// Local LAN IP. Required for GCP BGP over LAN Connection.
        /// </summary>
        [Input("localLanIp")]
        public Input<string>? LocalLanIp { get; set; }

        /// <summary>
        /// Source CIDR for the tunnel from the Aviatrix transit gateway.
        /// </summary>
        [Input("localTunnelCidr")]
        public Input<string>? LocalTunnelCidr { get; set; }

        [Input("manualBgpAdvertisedCidrs")]
        private InputList<string>? _manualBgpAdvertisedCidrs;

        /// <summary>
        /// Configure manual BGP advertised CIDRs for this connection. Only valid with 'connection_type' = 'bgp'. Available as of
        /// provider version R2.18+.
        /// </summary>
        public InputList<string> ManualBgpAdvertisedCidrs
        {
            get => _manualBgpAdvertisedCidrs ?? (_manualBgpAdvertisedCidrs = new InputList<string>());
            set => _manualBgpAdvertisedCidrs = value;
        }

        /// <summary>
        /// Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'.
        /// </summary>
        [Input("phase1Authentication")]
        public Input<string>? Phase1Authentication { get; set; }

        /// <summary>
        /// Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'.
        /// </summary>
        [Input("phase1DhGroups")]
        public Input<string>? Phase1DhGroups { get; set; }

        /// <summary>
        /// Phase one Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC' and 'AES-256-CBC', 'AES-128-GCM-64',
        /// 'AES-128-GCM-96', 'AES-128-GCM-128', 'AES-256-GCM-64', 'AES-256-GCM-96', and 'AES-256-GCM-128'.
        /// </summary>
        [Input("phase1Encryption")]
        public Input<string>? Phase1Encryption { get; set; }

        [Input("phase1RemoteIdentifiers")]
        private InputList<string>? _phase1RemoteIdentifiers;

        /// <summary>
        /// Phase 1 remote identifier of the IPsec tunnel.
        /// </summary>
        public InputList<string> Phase1RemoteIdentifiers
        {
            get => _phase1RemoteIdentifiers ?? (_phase1RemoteIdentifiers = new InputList<string>());
            set => _phase1RemoteIdentifiers = value;
        }

        /// <summary>
        /// Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'.
        /// </summary>
        [Input("phase2Authentication")]
        public Input<string>? Phase2Authentication { get; set; }

        /// <summary>
        /// Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'.
        /// </summary>
        [Input("phase2DhGroups")]
        public Input<string>? Phase2DhGroups { get; set; }

        /// <summary>
        /// Phase two Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC', 'AES-256-CBC', 'AES-128-GCM-64',
        /// 'AES-128-GCM-96', 'AES-128-GCM-128', 'AES-256-GCM-64', 'AES-256-GCM-96', 'AES-256-GCM-128', and 'NULL-ENCR'.
        /// </summary>
        [Input("phase2Encryption")]
        public Input<string>? Phase2Encryption { get; set; }

        /// <summary>
        /// If left blank, the pre-shared key will be auto generated.
        /// </summary>
        [Input("preSharedKey")]
        public Input<string>? PreSharedKey { get; set; }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// Connection AS Path Prepend customized by specifying AS PATH for a BGP connection.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// Remote Gateway IP.
        /// </summary>
        [Input("remoteGatewayIp")]
        public Input<string>? RemoteGatewayIp { get; set; }

        /// <summary>
        /// Remote LAN IP.
        /// </summary>
        [Input("remoteLanIp")]
        public Input<string>? RemoteLanIp { get; set; }

        /// <summary>
        /// Remote CIDRs joined as a string with ','. Required for a 'static' type connection.
        /// </summary>
        [Input("remoteSubnet")]
        public Input<string>? RemoteSubnet { get; set; }

        /// <summary>
        /// Destination CIDR for the tunnel to the external device.
        /// </summary>
        [Input("remoteTunnelCidr")]
        public Input<string>? RemoteTunnelCidr { get; set; }

        /// <summary>
        /// Name of the remote VPC for a LAN BGP connection. Only valid when 'connection_type' = 'bgp' and tunnel_protocol' = 'LAN'
        /// with an Azure transit gateway. Must be in the form "&lt;VNET-name&gt;:&lt;resource-group-name&gt;". Available as of provider version
        /// R2.18+.
        /// </summary>
        [Input("remoteVpcName")]
        public Input<string>? RemoteVpcName { get; set; }

        /// <summary>
        /// Only valid for Transit Gateway's with Active-Standby Mode enabled. Valid values: true, false. Default: false.
        /// </summary>
        [Input("switchToHaStandbyGateway")]
        public Input<bool>? SwitchToHaStandbyGateway { get; set; }

        /// <summary>
        /// Tunnel Protocol. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive.
        /// </summary>
        [Input("tunnelProtocol")]
        public Input<string>? TunnelProtocol { get; set; }

        /// <summary>
        /// ID of the VPC where the Transit Gateway is located. For GCP BGP over LAN connection, it is in the format of
        /// 'vpc_name~-~account_name'.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixTransitExternalDeviceConnArgs()
        {
        }
        public static new AviatrixTransitExternalDeviceConnArgs Empty => new AviatrixTransitExternalDeviceConnArgs();
    }

    public sealed class AviatrixTransitExternalDeviceConnState : global::Pulumi.ResourceArgs
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
        /// Backup BGP MD5 authentication key.
        /// </summary>
        [Input("backupBgpMd5Key")]
        public Input<string>? BackupBgpMd5Key { get; set; }

        /// <summary>
        /// Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("backupBgpRemoteAsNum")]
        public Input<string>? BackupBgpRemoteAsNum { get; set; }

        /// <summary>
        /// Backup direct connect for backup external device.
        /// </summary>
        [Input("backupDirectConnect")]
        public Input<bool>? BackupDirectConnect { get; set; }

        /// <summary>
        /// Backup Local LAN IP. Required for GCP BGP over LAN Connection with HA enabled.
        /// </summary>
        [Input("backupLocalLanIp")]
        public Input<string>? BackupLocalLanIp { get; set; }

        /// <summary>
        /// Source CIDR for the tunnel from the backup Aviatrix transit gateway.
        /// </summary>
        [Input("backupLocalTunnelCidr")]
        public Input<string>? BackupLocalTunnelCidr { get; set; }

        /// <summary>
        /// Backup pre shared key.
        /// </summary>
        [Input("backupPreSharedKey")]
        public Input<string>? BackupPreSharedKey { get; set; }

        /// <summary>
        /// Backup remote gateway IP.
        /// </summary>
        [Input("backupRemoteGatewayIp")]
        public Input<string>? BackupRemoteGatewayIp { get; set; }

        /// <summary>
        /// Backup Remote LAN IP.
        /// </summary>
        [Input("backupRemoteLanIp")]
        public Input<string>? BackupRemoteLanIp { get; set; }

        /// <summary>
        /// Destination CIDR for the tunnel to the backup external device.
        /// </summary>
        [Input("backupRemoteTunnelCidr")]
        public Input<string>? BackupRemoteTunnelCidr { get; set; }

        /// <summary>
        /// BGP local ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("bgpLocalAsNum")]
        public Input<string>? BgpLocalAsNum { get; set; }

        /// <summary>
        /// BGP MD5 authentication key.
        /// </summary>
        [Input("bgpMd5Key")]
        public Input<string>? BgpMd5Key { get; set; }

        /// <summary>
        /// BGP remote ASN (Autonomous System Number). Integer between 1-4294967294.
        /// </summary>
        [Input("bgpRemoteAsNum")]
        public Input<string>? BgpRemoteAsNum { get; set; }

        /// <summary>
        /// The name of the transit external device connection which is going to be created.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption.
        /// </summary>
        [Input("customAlgorithms")]
        public Input<bool>? CustomAlgorithms { get; set; }

        /// <summary>
        /// Set true for private network infrastructure.
        /// </summary>
        [Input("directConnect")]
        public Input<bool>? DirectConnect { get; set; }

        /// <summary>
        /// Switch to enable BGP LAN ActiveMesh. Only valid for GCP with Remote Gateway HA enabled. Default: false. Available as of
        /// provider version R2.21+.
        /// </summary>
        [Input("enableBgpLanActivemesh")]
        public Input<bool>? EnableBgpLanActivemesh { get; set; }

        /// <summary>
        /// Switch to allow this connection to communicate with a Network Domain via Connection Policy.
        /// </summary>
        [Input("enableEdgeSegmentation")]
        public Input<bool>? EnableEdgeSegmentation { get; set; }

        /// <summary>
        /// Enable Event Triggered HA.
        /// </summary>
        [Input("enableEventTriggeredHa")]
        public Input<bool>? EnableEventTriggeredHa { get; set; }

        /// <summary>
        /// Set as true if use IKEv2.
        /// </summary>
        [Input("enableIkev2")]
        public Input<bool>? EnableIkev2 { get; set; }

        /// <summary>
        /// Enable Jumbo Frame for the transit external device connection. Valid values: true, false.
        /// </summary>
        [Input("enableJumboFrame")]
        public Input<bool>? EnableJumboFrame { get; set; }

        /// <summary>
        /// Enable learned CIDR approval for the connection. Only valid with 'connection_type' = 'bgp'. Requires the
        /// transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default
        /// value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Name of the Transit Gateway.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// Set as true if there are two external devices.
        /// </summary>
        [Input("haEnabled")]
        public Input<bool>? HaEnabled { get; set; }

        /// <summary>
        /// Local LAN IP. Required for GCP BGP over LAN Connection.
        /// </summary>
        [Input("localLanIp")]
        public Input<string>? LocalLanIp { get; set; }

        /// <summary>
        /// Source CIDR for the tunnel from the Aviatrix transit gateway.
        /// </summary>
        [Input("localTunnelCidr")]
        public Input<string>? LocalTunnelCidr { get; set; }

        [Input("manualBgpAdvertisedCidrs")]
        private InputList<string>? _manualBgpAdvertisedCidrs;

        /// <summary>
        /// Configure manual BGP advertised CIDRs for this connection. Only valid with 'connection_type' = 'bgp'. Available as of
        /// provider version R2.18+.
        /// </summary>
        public InputList<string> ManualBgpAdvertisedCidrs
        {
            get => _manualBgpAdvertisedCidrs ?? (_manualBgpAdvertisedCidrs = new InputList<string>());
            set => _manualBgpAdvertisedCidrs = value;
        }

        /// <summary>
        /// Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'.
        /// </summary>
        [Input("phase1Authentication")]
        public Input<string>? Phase1Authentication { get; set; }

        /// <summary>
        /// Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'.
        /// </summary>
        [Input("phase1DhGroups")]
        public Input<string>? Phase1DhGroups { get; set; }

        /// <summary>
        /// Phase one Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC' and 'AES-256-CBC', 'AES-128-GCM-64',
        /// 'AES-128-GCM-96', 'AES-128-GCM-128', 'AES-256-GCM-64', 'AES-256-GCM-96', and 'AES-256-GCM-128'.
        /// </summary>
        [Input("phase1Encryption")]
        public Input<string>? Phase1Encryption { get; set; }

        [Input("phase1RemoteIdentifiers")]
        private InputList<string>? _phase1RemoteIdentifiers;

        /// <summary>
        /// Phase 1 remote identifier of the IPsec tunnel.
        /// </summary>
        public InputList<string> Phase1RemoteIdentifiers
        {
            get => _phase1RemoteIdentifiers ?? (_phase1RemoteIdentifiers = new InputList<string>());
            set => _phase1RemoteIdentifiers = value;
        }

        /// <summary>
        /// Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'.
        /// </summary>
        [Input("phase2Authentication")]
        public Input<string>? Phase2Authentication { get; set; }

        /// <summary>
        /// Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'.
        /// </summary>
        [Input("phase2DhGroups")]
        public Input<string>? Phase2DhGroups { get; set; }

        /// <summary>
        /// Phase two Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC', 'AES-256-CBC', 'AES-128-GCM-64',
        /// 'AES-128-GCM-96', 'AES-128-GCM-128', 'AES-256-GCM-64', 'AES-256-GCM-96', 'AES-256-GCM-128', and 'NULL-ENCR'.
        /// </summary>
        [Input("phase2Encryption")]
        public Input<string>? Phase2Encryption { get; set; }

        /// <summary>
        /// If left blank, the pre-shared key will be auto generated.
        /// </summary>
        [Input("preSharedKey")]
        public Input<string>? PreSharedKey { get; set; }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// Connection AS Path Prepend customized by specifying AS PATH for a BGP connection.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// Remote Gateway IP.
        /// </summary>
        [Input("remoteGatewayIp")]
        public Input<string>? RemoteGatewayIp { get; set; }

        /// <summary>
        /// Remote LAN IP.
        /// </summary>
        [Input("remoteLanIp")]
        public Input<string>? RemoteLanIp { get; set; }

        /// <summary>
        /// Remote CIDRs joined as a string with ','. Required for a 'static' type connection.
        /// </summary>
        [Input("remoteSubnet")]
        public Input<string>? RemoteSubnet { get; set; }

        /// <summary>
        /// Destination CIDR for the tunnel to the external device.
        /// </summary>
        [Input("remoteTunnelCidr")]
        public Input<string>? RemoteTunnelCidr { get; set; }

        /// <summary>
        /// Name of the remote VPC for a LAN BGP connection. Only valid when 'connection_type' = 'bgp' and tunnel_protocol' = 'LAN'
        /// with an Azure transit gateway. Must be in the form "&lt;VNET-name&gt;:&lt;resource-group-name&gt;". Available as of provider version
        /// R2.18+.
        /// </summary>
        [Input("remoteVpcName")]
        public Input<string>? RemoteVpcName { get; set; }

        /// <summary>
        /// Only valid for Transit Gateway's with Active-Standby Mode enabled. Valid values: true, false. Default: false.
        /// </summary>
        [Input("switchToHaStandbyGateway")]
        public Input<bool>? SwitchToHaStandbyGateway { get; set; }

        /// <summary>
        /// Tunnel Protocol. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive.
        /// </summary>
        [Input("tunnelProtocol")]
        public Input<string>? TunnelProtocol { get; set; }

        /// <summary>
        /// ID of the VPC where the Transit Gateway is located. For GCP BGP over LAN connection, it is in the format of
        /// 'vpc_name~-~account_name'.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixTransitExternalDeviceConnState()
        {
        }
        public static new AviatrixTransitExternalDeviceConnState Empty => new AviatrixTransitExternalDeviceConnState();
    }
}

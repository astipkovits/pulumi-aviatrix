// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixFirenet:AviatrixFirenet")]
    public partial class AviatrixFirenet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as
        /// of provider version R2.19.2+.
        /// </summary>
        [Output("eastWestInspectionExcludedCidrs")]
        public Output<ImmutableArray<string>> EastWestInspectionExcludedCidrs { get; private set; } = null!;

        /// <summary>
        /// Enable/Disable egress through firewall.
        /// </summary>
        [Output("egressEnabled")]
        public Output<bool?> EgressEnabled { get; private set; } = null!;

        /// <summary>
        /// List of egress static cidrs.
        /// </summary>
        [Output("egressStaticCidrs")]
        public Output<ImmutableArray<string>> EgressStaticCidrs { get; private set; } = null!;

        /// <summary>
        /// List of firewall instances to be associated with fireNet.
        /// </summary>
        [Output("firewallInstanceAssociations")]
        public Output<ImmutableArray<Outputs.AviatrixFirenetFirewallInstanceAssociation>> FirewallInstanceAssociations { get; private set; } = null!;

        /// <summary>
        /// Hashing algorithm to load balance traffic across the firewall.
        /// </summary>
        [Output("hashingAlgorithm")]
        public Output<string?> HashingAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Enable/Disable traffic inspection.
        /// </summary>
        [Output("inspectionEnabled")]
        public Output<bool?> InspectionEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable Keep Alive via Firewall LAN Interface.
        /// </summary>
        [Output("keepAliveViaLanInterfaceEnabled")]
        public Output<bool?> KeepAliveViaLanInterfaceEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable this to manage firewall_instance_associations in-line. If this is false, associations must be managed via
        /// standalone aviatrix_firewall_instance_association resources. Type: boolean, Default: true, Valid values: true/false.
        /// </summary>
        [Output("manageFirewallInstanceAssociation")]
        public Output<bool?> ManageFirewallInstanceAssociation { get; private set; } = null!;

        /// <summary>
        /// Enable TGW segmentation for egress.
        /// </summary>
        [Output("tgwSegmentationForEgressEnabled")]
        public Output<bool?> TgwSegmentationForEgressEnabled { get; private set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFirenet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFirenet(string name, AviatrixFirenetArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirenet:AviatrixFirenet", name, args ?? new AviatrixFirenetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFirenet(string name, Input<string> id, AviatrixFirenetState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirenet:AviatrixFirenet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixFirenet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFirenet Get(string name, Input<string> id, AviatrixFirenetState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFirenet(name, id, state, options);
        }
    }

    public sealed class AviatrixFirenetArgs : global::Pulumi.ResourceArgs
    {
        [Input("eastWestInspectionExcludedCidrs")]
        private InputList<string>? _eastWestInspectionExcludedCidrs;

        /// <summary>
        /// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as
        /// of provider version R2.19.2+.
        /// </summary>
        public InputList<string> EastWestInspectionExcludedCidrs
        {
            get => _eastWestInspectionExcludedCidrs ?? (_eastWestInspectionExcludedCidrs = new InputList<string>());
            set => _eastWestInspectionExcludedCidrs = value;
        }

        /// <summary>
        /// Enable/Disable egress through firewall.
        /// </summary>
        [Input("egressEnabled")]
        public Input<bool>? EgressEnabled { get; set; }

        [Input("egressStaticCidrs")]
        private InputList<string>? _egressStaticCidrs;

        /// <summary>
        /// List of egress static cidrs.
        /// </summary>
        public InputList<string> EgressStaticCidrs
        {
            get => _egressStaticCidrs ?? (_egressStaticCidrs = new InputList<string>());
            set => _egressStaticCidrs = value;
        }

        [Input("firewallInstanceAssociations")]
        private InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationArgs>? _firewallInstanceAssociations;

        /// <summary>
        /// List of firewall instances to be associated with fireNet.
        /// </summary>
        [Obsolete(@"Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.")]
        public InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationArgs> FirewallInstanceAssociations
        {
            get => _firewallInstanceAssociations ?? (_firewallInstanceAssociations = new InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationArgs>());
            set => _firewallInstanceAssociations = value;
        }

        /// <summary>
        /// Hashing algorithm to load balance traffic across the firewall.
        /// </summary>
        [Input("hashingAlgorithm")]
        public Input<string>? HashingAlgorithm { get; set; }

        /// <summary>
        /// Enable/Disable traffic inspection.
        /// </summary>
        [Input("inspectionEnabled")]
        public Input<bool>? InspectionEnabled { get; set; }

        /// <summary>
        /// Enable Keep Alive via Firewall LAN Interface.
        /// </summary>
        [Input("keepAliveViaLanInterfaceEnabled")]
        public Input<bool>? KeepAliveViaLanInterfaceEnabled { get; set; }

        /// <summary>
        /// Enable this to manage firewall_instance_associations in-line. If this is false, associations must be managed via
        /// standalone aviatrix_firewall_instance_association resources. Type: boolean, Default: true, Valid values: true/false.
        /// </summary>
        [Input("manageFirewallInstanceAssociation")]
        public Input<bool>? ManageFirewallInstanceAssociation { get; set; }

        /// <summary>
        /// Enable TGW segmentation for egress.
        /// </summary>
        [Input("tgwSegmentationForEgressEnabled")]
        public Input<bool>? TgwSegmentationForEgressEnabled { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixFirenetArgs()
        {
        }
        public static new AviatrixFirenetArgs Empty => new AviatrixFirenetArgs();
    }

    public sealed class AviatrixFirenetState : global::Pulumi.ResourceArgs
    {
        [Input("eastWestInspectionExcludedCidrs")]
        private InputList<string>? _eastWestInspectionExcludedCidrs;

        /// <summary>
        /// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as
        /// of provider version R2.19.2+.
        /// </summary>
        public InputList<string> EastWestInspectionExcludedCidrs
        {
            get => _eastWestInspectionExcludedCidrs ?? (_eastWestInspectionExcludedCidrs = new InputList<string>());
            set => _eastWestInspectionExcludedCidrs = value;
        }

        /// <summary>
        /// Enable/Disable egress through firewall.
        /// </summary>
        [Input("egressEnabled")]
        public Input<bool>? EgressEnabled { get; set; }

        [Input("egressStaticCidrs")]
        private InputList<string>? _egressStaticCidrs;

        /// <summary>
        /// List of egress static cidrs.
        /// </summary>
        public InputList<string> EgressStaticCidrs
        {
            get => _egressStaticCidrs ?? (_egressStaticCidrs = new InputList<string>());
            set => _egressStaticCidrs = value;
        }

        [Input("firewallInstanceAssociations")]
        private InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationGetArgs>? _firewallInstanceAssociations;

        /// <summary>
        /// List of firewall instances to be associated with fireNet.
        /// </summary>
        [Obsolete(@"Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.")]
        public InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationGetArgs> FirewallInstanceAssociations
        {
            get => _firewallInstanceAssociations ?? (_firewallInstanceAssociations = new InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationGetArgs>());
            set => _firewallInstanceAssociations = value;
        }

        /// <summary>
        /// Hashing algorithm to load balance traffic across the firewall.
        /// </summary>
        [Input("hashingAlgorithm")]
        public Input<string>? HashingAlgorithm { get; set; }

        /// <summary>
        /// Enable/Disable traffic inspection.
        /// </summary>
        [Input("inspectionEnabled")]
        public Input<bool>? InspectionEnabled { get; set; }

        /// <summary>
        /// Enable Keep Alive via Firewall LAN Interface.
        /// </summary>
        [Input("keepAliveViaLanInterfaceEnabled")]
        public Input<bool>? KeepAliveViaLanInterfaceEnabled { get; set; }

        /// <summary>
        /// Enable this to manage firewall_instance_associations in-line. If this is false, associations must be managed via
        /// standalone aviatrix_firewall_instance_association resources. Type: boolean, Default: true, Valid values: true/false.
        /// </summary>
        [Input("manageFirewallInstanceAssociation")]
        public Input<bool>? ManageFirewallInstanceAssociation { get; set; }

        /// <summary>
        /// Enable TGW segmentation for egress.
        /// </summary>
        [Input("tgwSegmentationForEgressEnabled")]
        public Input<bool>? TgwSegmentationForEgressEnabled { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixFirenetState()
        {
        }
        public static new AviatrixFirenetState Empty => new AviatrixFirenetState();
    }
}

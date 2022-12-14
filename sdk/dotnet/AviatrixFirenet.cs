// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    /// <summary>
    /// The **aviatrix_firenet** resource allows the creation and management of [Aviatrix Firewall Networks](https://docs.aviatrix.com/HowTos/firewall_network_faq.html).
    /// 
    /// &gt; **NOTE:** This resource is used in conjunction with multiple other resources that may include, and are not limited to: **firewall_instance**, **firewall_instance_association**, **aws_tgw**, and **transit_gateway** resources or even **aviatrix_fqdn**, under the Aviatrix FireNet solution. Explicit dependencies may be set using `depends_on`. For more information on proper FireNet configuration, please see the workflow [here](https://docs.aviatrix.com/HowTos/firewall_network_workflow.html).
    /// 
    /// ## Import
    /// 
    /// **firenet** can be imported using the `vpc_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixFirenet:AviatrixFirenet test vpc_id
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixFirenet:AviatrixFirenet")]
    public partial class AviatrixFirenet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
        /// </summary>
        [Output("eastWestInspectionExcludedCidrs")]
        public Output<ImmutableArray<string>> EastWestInspectionExcludedCidrs { get; private set; } = null!;

        /// <summary>
        /// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
        /// </summary>
        [Output("egressEnabled")]
        public Output<bool?> EgressEnabled { get; private set; } = null!;

        /// <summary>
        /// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
        /// </summary>
        [Output("egressStaticCidrs")]
        public Output<ImmutableArray<string>> EgressStaticCidrs { get; private set; } = null!;

        /// <summary>
        /// Dynamic block of firewall instance(s) to be associated with the FireNet.
        /// </summary>
        [Output("firewallInstanceAssociations")]
        public Output<ImmutableArray<Outputs.AviatrixFirenetFirewallInstanceAssociation>> FirewallInstanceAssociations { get; private set; } = null!;

        /// <summary>
        /// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
        /// </summary>
        [Output("hashingAlgorithm")]
        public Output<string?> HashingAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
        /// </summary>
        [Output("inspectionEnabled")]
        public Output<bool?> InspectionEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
        /// </summary>
        [Output("keepAliveViaLanInterfaceEnabled")]
        public Output<bool?> KeepAliveViaLanInterfaceEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewall_instance_association` blocks can be used. If set to false, all firewall associations must be managed via standalone `aviatrix.AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
        /// </summary>
        [Output("manageFirewallInstanceAssociation")]
        public Output<bool?> ManageFirewallInstanceAssociation { get; private set; } = null!;

        /// <summary>
        /// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
        /// </summary>
        [Output("tgwSegmentationForEgressEnabled")]
        public Output<bool?> TgwSegmentationForEgressEnabled { get; private set; } = null!;

        /// <summary>
        /// VPC ID of the Security VPC.
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
                PluginDownloadURL = "github://api.github.com/astipkovits",
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
        /// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
        /// </summary>
        public InputList<string> EastWestInspectionExcludedCidrs
        {
            get => _eastWestInspectionExcludedCidrs ?? (_eastWestInspectionExcludedCidrs = new InputList<string>());
            set => _eastWestInspectionExcludedCidrs = value;
        }

        /// <summary>
        /// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("egressEnabled")]
        public Input<bool>? EgressEnabled { get; set; }

        [Input("egressStaticCidrs")]
        private InputList<string>? _egressStaticCidrs;

        /// <summary>
        /// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
        /// </summary>
        public InputList<string> EgressStaticCidrs
        {
            get => _egressStaticCidrs ?? (_egressStaticCidrs = new InputList<string>());
            set => _egressStaticCidrs = value;
        }

        [Input("firewallInstanceAssociations")]
        private InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationArgs>? _firewallInstanceAssociations;

        /// <summary>
        /// Dynamic block of firewall instance(s) to be associated with the FireNet.
        /// </summary>
        [Obsolete(@"Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.")]
        public InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationArgs> FirewallInstanceAssociations
        {
            get => _firewallInstanceAssociations ?? (_firewallInstanceAssociations = new InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationArgs>());
            set => _firewallInstanceAssociations = value;
        }

        /// <summary>
        /// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
        /// </summary>
        [Input("hashingAlgorithm")]
        public Input<string>? HashingAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("inspectionEnabled")]
        public Input<bool>? InspectionEnabled { get; set; }

        /// <summary>
        /// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
        /// </summary>
        [Input("keepAliveViaLanInterfaceEnabled")]
        public Input<bool>? KeepAliveViaLanInterfaceEnabled { get; set; }

        /// <summary>
        /// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewall_instance_association` blocks can be used. If set to false, all firewall associations must be managed via standalone `aviatrix.AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
        /// </summary>
        [Input("manageFirewallInstanceAssociation")]
        public Input<bool>? ManageFirewallInstanceAssociation { get; set; }

        /// <summary>
        /// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
        /// </summary>
        [Input("tgwSegmentationForEgressEnabled")]
        public Input<bool>? TgwSegmentationForEgressEnabled { get; set; }

        /// <summary>
        /// VPC ID of the Security VPC.
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
        /// Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.
        /// </summary>
        public InputList<string> EastWestInspectionExcludedCidrs
        {
            get => _eastWestInspectionExcludedCidrs ?? (_eastWestInspectionExcludedCidrs = new InputList<string>());
            set => _eastWestInspectionExcludedCidrs = value;
        }

        /// <summary>
        /// Enable/disable egress through firewall. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("egressEnabled")]
        public Input<bool>? EgressEnabled { get; set; }

        [Input("egressStaticCidrs")]
        private InputList<string>? _egressStaticCidrs;

        /// <summary>
        /// List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.
        /// </summary>
        public InputList<string> EgressStaticCidrs
        {
            get => _egressStaticCidrs ?? (_egressStaticCidrs = new InputList<string>());
            set => _egressStaticCidrs = value;
        }

        [Input("firewallInstanceAssociations")]
        private InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationGetArgs>? _firewallInstanceAssociations;

        /// <summary>
        /// Dynamic block of firewall instance(s) to be associated with the FireNet.
        /// </summary>
        [Obsolete(@"Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.")]
        public InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationGetArgs> FirewallInstanceAssociations
        {
            get => _firewallInstanceAssociations ?? (_firewallInstanceAssociations = new InputList<Inputs.AviatrixFirenetFirewallInstanceAssociationGetArgs>());
            set => _firewallInstanceAssociations = value;
        }

        /// <summary>
        /// Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".
        /// </summary>
        [Input("hashingAlgorithm")]
        public Input<string>? HashingAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable traffic inspection. Valid values: true, false. Default value: true.
        /// </summary>
        [Input("inspectionEnabled")]
        public Input<bool>? InspectionEnabled { get; set; }

        /// <summary>
        /// Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+.
        /// </summary>
        [Input("keepAliveViaLanInterfaceEnabled")]
        public Input<bool>? KeepAliveViaLanInterfaceEnabled { get; set; }

        /// <summary>
        /// Enable this attribute to manage firewall associations in-line. If set to true, in-line `firewall_instance_association` blocks can be used. If set to false, all firewall associations must be managed via standalone `aviatrix.AviatrixFirewallInstanceAssociation` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.
        /// </summary>
        [Input("manageFirewallInstanceAssociation")]
        public Input<bool>? ManageFirewallInstanceAssociation { get; set; }

        /// <summary>
        /// Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.
        /// </summary>
        [Input("tgwSegmentationForEgressEnabled")]
        public Input<bool>? TgwSegmentationForEgressEnabled { get; set; }

        /// <summary>
        /// VPC ID of the Security VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixFirenetState()
        {
        }
        public static new AviatrixFirenetState Empty => new AviatrixFirenetState();
    }
}

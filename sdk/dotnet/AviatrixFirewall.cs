// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixFirewall:AviatrixFirewall")]
    public partial class AviatrixFirewall : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        /// </summary>
        [Output("baseLogEnabled")]
        public Output<bool?> BaseLogEnabled { get; private set; } = null!;

        /// <summary>
        /// New base policy.
        /// </summary>
        [Output("basePolicy")]
        public Output<string?> BasePolicy { get; private set; } = null!;

        /// <summary>
        /// The name of gateway.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Enable to manage firewall policies via in-line rules. If false, policies must be managed using
        /// `aviatrix_firewall_policy` resources.
        /// </summary>
        [Output("manageFirewallPolicies")]
        public Output<bool?> ManageFirewallPolicies { get; private set; } = null!;

        /// <summary>
        /// New access policy for the gateway.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.AviatrixFirewallPolicy>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFirewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFirewall(string name, AviatrixFirewallArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewall:AviatrixFirewall", name, args ?? new AviatrixFirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFirewall(string name, Input<string> id, AviatrixFirewallState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewall:AviatrixFirewall", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixFirewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFirewall Get(string name, Input<string> id, AviatrixFirewallState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFirewall(name, id, state, options);
        }
    }

    public sealed class AviatrixFirewallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("baseLogEnabled")]
        public Input<bool>? BaseLogEnabled { get; set; }

        /// <summary>
        /// New base policy.
        /// </summary>
        [Input("basePolicy")]
        public Input<string>? BasePolicy { get; set; }

        /// <summary>
        /// The name of gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Enable to manage firewall policies via in-line rules. If false, policies must be managed using
        /// `aviatrix_firewall_policy` resources.
        /// </summary>
        [Input("manageFirewallPolicies")]
        public Input<bool>? ManageFirewallPolicies { get; set; }

        [Input("policies")]
        private InputList<Inputs.AviatrixFirewallPolicyArgs>? _policies;

        /// <summary>
        /// New access policy for the gateway.
        /// </summary>
        public InputList<Inputs.AviatrixFirewallPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.AviatrixFirewallPolicyArgs>());
            set => _policies = value;
        }

        public AviatrixFirewallArgs()
        {
        }
        public static new AviatrixFirewallArgs Empty => new AviatrixFirewallArgs();
    }

    public sealed class AviatrixFirewallState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        /// </summary>
        [Input("baseLogEnabled")]
        public Input<bool>? BaseLogEnabled { get; set; }

        /// <summary>
        /// New base policy.
        /// </summary>
        [Input("basePolicy")]
        public Input<string>? BasePolicy { get; set; }

        /// <summary>
        /// The name of gateway.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// Enable to manage firewall policies via in-line rules. If false, policies must be managed using
        /// `aviatrix_firewall_policy` resources.
        /// </summary>
        [Input("manageFirewallPolicies")]
        public Input<bool>? ManageFirewallPolicies { get; set; }

        [Input("policies")]
        private InputList<Inputs.AviatrixFirewallPolicyGetArgs>? _policies;

        /// <summary>
        /// New access policy for the gateway.
        /// </summary>
        public InputList<Inputs.AviatrixFirewallPolicyGetArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.AviatrixFirewallPolicyGetArgs>());
            set => _policies = value;
        }

        public AviatrixFirewallState()
        {
        }
        public static new AviatrixFirewallState Empty => new AviatrixFirewallState();
    }
}

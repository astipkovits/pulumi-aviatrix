// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixGatewayDnat:AviatrixGatewayDnat")]
    public partial class AviatrixGatewayDnat : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Computed attribute to store the previous connection policy.
        /// </summary>
        [Output("connectionPolicies")]
        public Output<ImmutableArray<Outputs.AviatrixGatewayDnatConnectionPolicy>> ConnectionPolicies { get; private set; } = null!;

        /// <summary>
        /// Policy rule to be applied to gateway.
        /// </summary>
        [Output("dnatPolicies")]
        public Output<ImmutableArray<Outputs.AviatrixGatewayDnatDnatPolicy>> DnatPolicies { get; private set; } = null!;

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Computed attribute to store the previous interface policy.
        /// </summary>
        [Output("interfacePolicies")]
        public Output<ImmutableArray<Outputs.AviatrixGatewayDnatInterfacePolicy>> InterfacePolicies { get; private set; } = null!;

        /// <summary>
        /// Whether to sync the policies to the HA gateway.
        /// </summary>
        [Output("syncToHa")]
        public Output<bool?> SyncToHa { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixGatewayDnat resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixGatewayDnat(string name, AviatrixGatewayDnatArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixGatewayDnat:AviatrixGatewayDnat", name, args ?? new AviatrixGatewayDnatArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixGatewayDnat(string name, Input<string> id, AviatrixGatewayDnatState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixGatewayDnat:AviatrixGatewayDnat", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixGatewayDnat resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixGatewayDnat Get(string name, Input<string> id, AviatrixGatewayDnatState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixGatewayDnat(name, id, state, options);
        }
    }

    public sealed class AviatrixGatewayDnatArgs : global::Pulumi.ResourceArgs
    {
        [Input("dnatPolicies", required: true)]
        private InputList<Inputs.AviatrixGatewayDnatDnatPolicyArgs>? _dnatPolicies;

        /// <summary>
        /// Policy rule to be applied to gateway.
        /// </summary>
        public InputList<Inputs.AviatrixGatewayDnatDnatPolicyArgs> DnatPolicies
        {
            get => _dnatPolicies ?? (_dnatPolicies = new InputList<Inputs.AviatrixGatewayDnatDnatPolicyArgs>());
            set => _dnatPolicies = value;
        }

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Whether to sync the policies to the HA gateway.
        /// </summary>
        [Input("syncToHa")]
        public Input<bool>? SyncToHa { get; set; }

        public AviatrixGatewayDnatArgs()
        {
        }
        public static new AviatrixGatewayDnatArgs Empty => new AviatrixGatewayDnatArgs();
    }

    public sealed class AviatrixGatewayDnatState : global::Pulumi.ResourceArgs
    {
        [Input("connectionPolicies")]
        private InputList<Inputs.AviatrixGatewayDnatConnectionPolicyGetArgs>? _connectionPolicies;

        /// <summary>
        /// Computed attribute to store the previous connection policy.
        /// </summary>
        public InputList<Inputs.AviatrixGatewayDnatConnectionPolicyGetArgs> ConnectionPolicies
        {
            get => _connectionPolicies ?? (_connectionPolicies = new InputList<Inputs.AviatrixGatewayDnatConnectionPolicyGetArgs>());
            set => _connectionPolicies = value;
        }

        [Input("dnatPolicies")]
        private InputList<Inputs.AviatrixGatewayDnatDnatPolicyGetArgs>? _dnatPolicies;

        /// <summary>
        /// Policy rule to be applied to gateway.
        /// </summary>
        public InputList<Inputs.AviatrixGatewayDnatDnatPolicyGetArgs> DnatPolicies
        {
            get => _dnatPolicies ?? (_dnatPolicies = new InputList<Inputs.AviatrixGatewayDnatDnatPolicyGetArgs>());
            set => _dnatPolicies = value;
        }

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        [Input("interfacePolicies")]
        private InputList<Inputs.AviatrixGatewayDnatInterfacePolicyGetArgs>? _interfacePolicies;

        /// <summary>
        /// Computed attribute to store the previous interface policy.
        /// </summary>
        public InputList<Inputs.AviatrixGatewayDnatInterfacePolicyGetArgs> InterfacePolicies
        {
            get => _interfacePolicies ?? (_interfacePolicies = new InputList<Inputs.AviatrixGatewayDnatInterfacePolicyGetArgs>());
            set => _interfacePolicies = value;
        }

        /// <summary>
        /// Whether to sync the policies to the HA gateway.
        /// </summary>
        [Input("syncToHa")]
        public Input<bool>? SyncToHa { get; set; }

        public AviatrixGatewayDnatState()
        {
        }
        public static new AviatrixGatewayDnatState Empty => new AviatrixGatewayDnatState();
    }
}

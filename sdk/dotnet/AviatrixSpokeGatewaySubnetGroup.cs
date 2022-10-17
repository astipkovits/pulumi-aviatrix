// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixSpokeGatewaySubnetGroup:AviatrixSpokeGatewaySubnetGroup")]
    public partial class AviatrixSpokeGatewaySubnetGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Spoke gateway name.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Subnet group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A set of subnets in the subnet group.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixSpokeGatewaySubnetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixSpokeGatewaySubnetGroup(string name, AviatrixSpokeGatewaySubnetGroupArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSpokeGatewaySubnetGroup:AviatrixSpokeGatewaySubnetGroup", name, args ?? new AviatrixSpokeGatewaySubnetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixSpokeGatewaySubnetGroup(string name, Input<string> id, AviatrixSpokeGatewaySubnetGroupState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSpokeGatewaySubnetGroup:AviatrixSpokeGatewaySubnetGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixSpokeGatewaySubnetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixSpokeGatewaySubnetGroup Get(string name, Input<string> id, AviatrixSpokeGatewaySubnetGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixSpokeGatewaySubnetGroup(name, id, state, options);
        }
    }

    public sealed class AviatrixSpokeGatewaySubnetGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Spoke gateway name.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// Subnet group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// A set of subnets in the subnet group.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public AviatrixSpokeGatewaySubnetGroupArgs()
        {
        }
        public static new AviatrixSpokeGatewaySubnetGroupArgs Empty => new AviatrixSpokeGatewaySubnetGroupArgs();
    }

    public sealed class AviatrixSpokeGatewaySubnetGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Spoke gateway name.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// Subnet group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// A set of subnets in the subnet group.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public AviatrixSpokeGatewaySubnetGroupState()
        {
        }
        public static new AviatrixSpokeGatewaySubnetGroupState Empty => new AviatrixSpokeGatewaySubnetGroupState();
    }
}
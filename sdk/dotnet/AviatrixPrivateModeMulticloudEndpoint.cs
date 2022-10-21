// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixPrivateModeMulticloudEndpoint:AviatrixPrivateModeMulticloudEndpoint")]
    public partial class AviatrixPrivateModeMulticloudEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the access account.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC with the Controller load balancer.
        /// </summary>
        [Output("controllerLbVpcId")]
        public Output<string> ControllerLbVpcId { get; private set; } = null!;

        /// <summary>
        /// DNS entry of this endpoint.
        /// </summary>
        [Output("dnsEntry")]
        public Output<string> DnsEntry { get; private set; } = null!;

        /// <summary>
        /// Name of the VPC region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixPrivateModeMulticloudEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixPrivateModeMulticloudEndpoint(string name, AviatrixPrivateModeMulticloudEndpointArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixPrivateModeMulticloudEndpoint:AviatrixPrivateModeMulticloudEndpoint", name, args ?? new AviatrixPrivateModeMulticloudEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixPrivateModeMulticloudEndpoint(string name, Input<string> id, AviatrixPrivateModeMulticloudEndpointState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixPrivateModeMulticloudEndpoint:AviatrixPrivateModeMulticloudEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixPrivateModeMulticloudEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixPrivateModeMulticloudEndpoint Get(string name, Input<string> id, AviatrixPrivateModeMulticloudEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixPrivateModeMulticloudEndpoint(name, id, state, options);
        }
    }

    public sealed class AviatrixPrivateModeMulticloudEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the access account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// ID of the VPC with the Controller load balancer.
        /// </summary>
        [Input("controllerLbVpcId", required: true)]
        public Input<string> ControllerLbVpcId { get; set; } = null!;

        /// <summary>
        /// Name of the VPC region.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixPrivateModeMulticloudEndpointArgs()
        {
        }
        public static new AviatrixPrivateModeMulticloudEndpointArgs Empty => new AviatrixPrivateModeMulticloudEndpointArgs();
    }

    public sealed class AviatrixPrivateModeMulticloudEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the access account.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// ID of the VPC with the Controller load balancer.
        /// </summary>
        [Input("controllerLbVpcId")]
        public Input<string>? ControllerLbVpcId { get; set; }

        /// <summary>
        /// DNS entry of this endpoint.
        /// </summary>
        [Input("dnsEntry")]
        public Input<string>? DnsEntry { get; set; }

        /// <summary>
        /// Name of the VPC region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixPrivateModeMulticloudEndpointState()
        {
        }
        public static new AviatrixPrivateModeMulticloudEndpointState Empty => new AviatrixPrivateModeMulticloudEndpointState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAzureVngConn:AviatrixAzureVngConn")]
    public partial class AviatrixAzureVngConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// VNG attached or not
        /// </summary>
        [Output("attached")]
        public Output<bool> Attached { get; private set; } = null!;

        /// <summary>
        /// Connection name
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// Primary gateway name
        /// </summary>
        [Output("primaryGatewayName")]
        public Output<string> PrimaryGatewayName { get; private set; } = null!;

        /// <summary>
        /// VNG name
        /// </summary>
        [Output("vngName")]
        public Output<string> VngName { get; private set; } = null!;

        /// <summary>
        /// VPC ID
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAzureVngConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAzureVngConn(string name, AviatrixAzureVngConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAzureVngConn:AviatrixAzureVngConn", name, args ?? new AviatrixAzureVngConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAzureVngConn(string name, Input<string> id, AviatrixAzureVngConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAzureVngConn:AviatrixAzureVngConn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAzureVngConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAzureVngConn Get(string name, Input<string> id, AviatrixAzureVngConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAzureVngConn(name, id, state, options);
        }
    }

    public sealed class AviatrixAzureVngConnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection name
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Primary gateway name
        /// </summary>
        [Input("primaryGatewayName", required: true)]
        public Input<string> PrimaryGatewayName { get; set; } = null!;

        public AviatrixAzureVngConnArgs()
        {
        }
        public static new AviatrixAzureVngConnArgs Empty => new AviatrixAzureVngConnArgs();
    }

    public sealed class AviatrixAzureVngConnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VNG attached or not
        /// </summary>
        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        /// <summary>
        /// Connection name
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Primary gateway name
        /// </summary>
        [Input("primaryGatewayName")]
        public Input<string>? PrimaryGatewayName { get; set; }

        /// <summary>
        /// VNG name
        /// </summary>
        [Input("vngName")]
        public Input<string>? VngName { get; set; }

        /// <summary>
        /// VPC ID
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixAzureVngConnState()
        {
        }
        public static new AviatrixAzureVngConnState Empty => new AviatrixAzureVngConnState();
    }
}

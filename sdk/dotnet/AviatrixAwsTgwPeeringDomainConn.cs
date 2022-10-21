// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwPeeringDomainConn:AviatrixAwsTgwPeeringDomainConn")]
    public partial class AviatrixAwsTgwPeeringDomainConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the source domain to make a connection.
        /// </summary>
        [Output("domainName1")]
        public Output<string> DomainName1 { get; private set; } = null!;

        /// <summary>
        /// The name of the destination domain to make a connection.
        /// </summary>
        [Output("domainName2")]
        public Output<string> DomainName2 { get; private set; } = null!;

        /// <summary>
        /// The AWS tgw name of the source domain to make a connection.
        /// </summary>
        [Output("tgwName1")]
        public Output<string> TgwName1 { get; private set; } = null!;

        /// <summary>
        /// The AWS tgw name of the destination domain to make a connection.
        /// </summary>
        [Output("tgwName2")]
        public Output<string> TgwName2 { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwPeeringDomainConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwPeeringDomainConn(string name, AviatrixAwsTgwPeeringDomainConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwPeeringDomainConn:AviatrixAwsTgwPeeringDomainConn", name, args ?? new AviatrixAwsTgwPeeringDomainConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwPeeringDomainConn(string name, Input<string> id, AviatrixAwsTgwPeeringDomainConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwPeeringDomainConn:AviatrixAwsTgwPeeringDomainConn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsTgwPeeringDomainConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwPeeringDomainConn Get(string name, Input<string> id, AviatrixAwsTgwPeeringDomainConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwPeeringDomainConn(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwPeeringDomainConnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the source domain to make a connection.
        /// </summary>
        [Input("domainName1", required: true)]
        public Input<string> DomainName1 { get; set; } = null!;

        /// <summary>
        /// The name of the destination domain to make a connection.
        /// </summary>
        [Input("domainName2", required: true)]
        public Input<string> DomainName2 { get; set; } = null!;

        /// <summary>
        /// The AWS tgw name of the source domain to make a connection.
        /// </summary>
        [Input("tgwName1", required: true)]
        public Input<string> TgwName1 { get; set; } = null!;

        /// <summary>
        /// The AWS tgw name of the destination domain to make a connection.
        /// </summary>
        [Input("tgwName2", required: true)]
        public Input<string> TgwName2 { get; set; } = null!;

        public AviatrixAwsTgwPeeringDomainConnArgs()
        {
        }
        public static new AviatrixAwsTgwPeeringDomainConnArgs Empty => new AviatrixAwsTgwPeeringDomainConnArgs();
    }

    public sealed class AviatrixAwsTgwPeeringDomainConnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the source domain to make a connection.
        /// </summary>
        [Input("domainName1")]
        public Input<string>? DomainName1 { get; set; }

        /// <summary>
        /// The name of the destination domain to make a connection.
        /// </summary>
        [Input("domainName2")]
        public Input<string>? DomainName2 { get; set; }

        /// <summary>
        /// The AWS tgw name of the source domain to make a connection.
        /// </summary>
        [Input("tgwName1")]
        public Input<string>? TgwName1 { get; set; }

        /// <summary>
        /// The AWS tgw name of the destination domain to make a connection.
        /// </summary>
        [Input("tgwName2")]
        public Input<string>? TgwName2 { get; set; }

        public AviatrixAwsTgwPeeringDomainConnState()
        {
        }
        public static new AviatrixAwsTgwPeeringDomainConnState Empty => new AviatrixAwsTgwPeeringDomainConnState();
    }
}

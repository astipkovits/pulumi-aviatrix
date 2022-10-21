// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixSegmentationSecurityDomainConnectionPolicy:AviatrixSegmentationSecurityDomainConnectionPolicy")]
    public partial class AviatrixSegmentationSecurityDomainConnectionPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of security domain that will be connected to domain 2.
        /// </summary>
        [Output("domainName1")]
        public Output<string> DomainName1 { get; private set; } = null!;

        /// <summary>
        /// Name of security domain that will be connected to domain 1.
        /// </summary>
        [Output("domainName2")]
        public Output<string> DomainName2 { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixSegmentationSecurityDomainConnectionPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixSegmentationSecurityDomainConnectionPolicy(string name, AviatrixSegmentationSecurityDomainConnectionPolicyArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSegmentationSecurityDomainConnectionPolicy:AviatrixSegmentationSecurityDomainConnectionPolicy", name, args ?? new AviatrixSegmentationSecurityDomainConnectionPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixSegmentationSecurityDomainConnectionPolicy(string name, Input<string> id, AviatrixSegmentationSecurityDomainConnectionPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSegmentationSecurityDomainConnectionPolicy:AviatrixSegmentationSecurityDomainConnectionPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixSegmentationSecurityDomainConnectionPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixSegmentationSecurityDomainConnectionPolicy Get(string name, Input<string> id, AviatrixSegmentationSecurityDomainConnectionPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixSegmentationSecurityDomainConnectionPolicy(name, id, state, options);
        }
    }

    public sealed class AviatrixSegmentationSecurityDomainConnectionPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of security domain that will be connected to domain 2.
        /// </summary>
        [Input("domainName1", required: true)]
        public Input<string> DomainName1 { get; set; } = null!;

        /// <summary>
        /// Name of security domain that will be connected to domain 1.
        /// </summary>
        [Input("domainName2", required: true)]
        public Input<string> DomainName2 { get; set; } = null!;

        public AviatrixSegmentationSecurityDomainConnectionPolicyArgs()
        {
        }
        public static new AviatrixSegmentationSecurityDomainConnectionPolicyArgs Empty => new AviatrixSegmentationSecurityDomainConnectionPolicyArgs();
    }

    public sealed class AviatrixSegmentationSecurityDomainConnectionPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of security domain that will be connected to domain 2.
        /// </summary>
        [Input("domainName1")]
        public Input<string>? DomainName1 { get; set; }

        /// <summary>
        /// Name of security domain that will be connected to domain 1.
        /// </summary>
        [Input("domainName2")]
        public Input<string>? DomainName2 { get; set; }

        public AviatrixSegmentationSecurityDomainConnectionPolicyState()
        {
        }
        public static new AviatrixSegmentationSecurityDomainConnectionPolicyState Empty => new AviatrixSegmentationSecurityDomainConnectionPolicyState();
    }
}

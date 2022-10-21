// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixFirewallManagementAccess:AviatrixFirewallManagementAccess")]
    public partial class AviatrixFirewallManagementAccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the resource to be enabled firewall management access.
        /// </summary>
        [Output("managementAccessResourceName")]
        public Output<string> ManagementAccessResourceName { get; private set; } = null!;

        /// <summary>
        /// Name of the transit firenet gateway.
        /// </summary>
        [Output("transitFirenetGatewayName")]
        public Output<string> TransitFirenetGatewayName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFirewallManagementAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFirewallManagementAccess(string name, AviatrixFirewallManagementAccessArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewallManagementAccess:AviatrixFirewallManagementAccess", name, args ?? new AviatrixFirewallManagementAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFirewallManagementAccess(string name, Input<string> id, AviatrixFirewallManagementAccessState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewallManagementAccess:AviatrixFirewallManagementAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixFirewallManagementAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFirewallManagementAccess Get(string name, Input<string> id, AviatrixFirewallManagementAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFirewallManagementAccess(name, id, state, options);
        }
    }

    public sealed class AviatrixFirewallManagementAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the resource to be enabled firewall management access.
        /// </summary>
        [Input("managementAccessResourceName", required: true)]
        public Input<string> ManagementAccessResourceName { get; set; } = null!;

        /// <summary>
        /// Name of the transit firenet gateway.
        /// </summary>
        [Input("transitFirenetGatewayName", required: true)]
        public Input<string> TransitFirenetGatewayName { get; set; } = null!;

        public AviatrixFirewallManagementAccessArgs()
        {
        }
        public static new AviatrixFirewallManagementAccessArgs Empty => new AviatrixFirewallManagementAccessArgs();
    }

    public sealed class AviatrixFirewallManagementAccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the resource to be enabled firewall management access.
        /// </summary>
        [Input("managementAccessResourceName")]
        public Input<string>? ManagementAccessResourceName { get; set; }

        /// <summary>
        /// Name of the transit firenet gateway.
        /// </summary>
        [Input("transitFirenetGatewayName")]
        public Input<string>? TransitFirenetGatewayName { get; set; }

        public AviatrixFirewallManagementAccessState()
        {
        }
        public static new AviatrixFirewallManagementAccessState Empty => new AviatrixFirewallManagementAccessState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwIntraDomainInspection:AviatrixAwsTgwIntraDomainInspection")]
    public partial class AviatrixAwsTgwIntraDomainInspection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Firewall domain name.
        /// </summary>
        [Output("firewallDomainName")]
        public Output<string> FirewallDomainName { get; private set; } = null!;

        /// <summary>
        /// Route domain name.
        /// </summary>
        [Output("routeDomainName")]
        public Output<string> RouteDomainName { get; private set; } = null!;

        /// <summary>
        /// AWS TGW name.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwIntraDomainInspection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwIntraDomainInspection(string name, AviatrixAwsTgwIntraDomainInspectionArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwIntraDomainInspection:AviatrixAwsTgwIntraDomainInspection", name, args ?? new AviatrixAwsTgwIntraDomainInspectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwIntraDomainInspection(string name, Input<string> id, AviatrixAwsTgwIntraDomainInspectionState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwIntraDomainInspection:AviatrixAwsTgwIntraDomainInspection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsTgwIntraDomainInspection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwIntraDomainInspection Get(string name, Input<string> id, AviatrixAwsTgwIntraDomainInspectionState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwIntraDomainInspection(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwIntraDomainInspectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Firewall domain name.
        /// </summary>
        [Input("firewallDomainName", required: true)]
        public Input<string> FirewallDomainName { get; set; } = null!;

        /// <summary>
        /// Route domain name.
        /// </summary>
        [Input("routeDomainName", required: true)]
        public Input<string> RouteDomainName { get; set; } = null!;

        /// <summary>
        /// AWS TGW name.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        public AviatrixAwsTgwIntraDomainInspectionArgs()
        {
        }
        public static new AviatrixAwsTgwIntraDomainInspectionArgs Empty => new AviatrixAwsTgwIntraDomainInspectionArgs();
    }

    public sealed class AviatrixAwsTgwIntraDomainInspectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Firewall domain name.
        /// </summary>
        [Input("firewallDomainName")]
        public Input<string>? FirewallDomainName { get; set; }

        /// <summary>
        /// Route domain name.
        /// </summary>
        [Input("routeDomainName")]
        public Input<string>? RouteDomainName { get; set; }

        /// <summary>
        /// AWS TGW name.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        public AviatrixAwsTgwIntraDomainInspectionState()
        {
        }
        public static new AviatrixAwsTgwIntraDomainInspectionState Empty => new AviatrixAwsTgwIntraDomainInspectionState();
    }
}

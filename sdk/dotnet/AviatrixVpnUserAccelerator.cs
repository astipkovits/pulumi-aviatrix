// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator")]
    public partial class AviatrixVpnUserAccelerator : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ELB to include into the VPN User Accelerator.
        /// </summary>
        [Output("elbName")]
        public Output<string> ElbName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixVpnUserAccelerator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixVpnUserAccelerator(string name, AviatrixVpnUserAcceleratorArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator", name, args ?? new AviatrixVpnUserAcceleratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixVpnUserAccelerator(string name, Input<string> id, AviatrixVpnUserAcceleratorState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixVpnUserAccelerator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixVpnUserAccelerator Get(string name, Input<string> id, AviatrixVpnUserAcceleratorState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixVpnUserAccelerator(name, id, state, options);
        }
    }

    public sealed class AviatrixVpnUserAcceleratorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ELB to include into the VPN User Accelerator.
        /// </summary>
        [Input("elbName", required: true)]
        public Input<string> ElbName { get; set; } = null!;

        public AviatrixVpnUserAcceleratorArgs()
        {
        }
        public static new AviatrixVpnUserAcceleratorArgs Empty => new AviatrixVpnUserAcceleratorArgs();
    }

    public sealed class AviatrixVpnUserAcceleratorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ELB to include into the VPN User Accelerator.
        /// </summary>
        [Input("elbName")]
        public Input<string>? ElbName { get; set; }

        public AviatrixVpnUserAcceleratorState()
        {
        }
        public static new AviatrixVpnUserAcceleratorState Empty => new AviatrixVpnUserAcceleratorState();
    }
}

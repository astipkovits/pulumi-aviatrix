// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixFilebeatForwarder:AviatrixFilebeatForwarder")]
    public partial class AviatrixFilebeatForwarder : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration file.
        /// </summary>
        [Output("configFile")]
        public Output<string?> ConfigFile { get; private set; } = null!;

        /// <summary>
        /// List of excluded gateways.
        /// </summary>
        [Output("excludedGateways")]
        public Output<ImmutableArray<string>> ExcludedGateways { get; private set; } = null!;

        /// <summary>
        /// Port number.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Server IP.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Enabled or not.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Trusted CA file.
        /// </summary>
        [Output("trustedCaFile")]
        public Output<string?> TrustedCaFile { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFilebeatForwarder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFilebeatForwarder(string name, AviatrixFilebeatForwarderArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFilebeatForwarder:AviatrixFilebeatForwarder", name, args ?? new AviatrixFilebeatForwarderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFilebeatForwarder(string name, Input<string> id, AviatrixFilebeatForwarderState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFilebeatForwarder:AviatrixFilebeatForwarder", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixFilebeatForwarder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFilebeatForwarder Get(string name, Input<string> id, AviatrixFilebeatForwarderState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFilebeatForwarder(name, id, state, options);
        }
    }

    public sealed class AviatrixFilebeatForwarderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration file.
        /// </summary>
        [Input("configFile")]
        public Input<string>? ConfigFile { get; set; }

        [Input("excludedGateways")]
        private InputList<string>? _excludedGateways;

        /// <summary>
        /// List of excluded gateways.
        /// </summary>
        public InputList<string> ExcludedGateways
        {
            get => _excludedGateways ?? (_excludedGateways = new InputList<string>());
            set => _excludedGateways = value;
        }

        /// <summary>
        /// Port number.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Server IP.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// Trusted CA file.
        /// </summary>
        [Input("trustedCaFile")]
        public Input<string>? TrustedCaFile { get; set; }

        public AviatrixFilebeatForwarderArgs()
        {
        }
        public static new AviatrixFilebeatForwarderArgs Empty => new AviatrixFilebeatForwarderArgs();
    }

    public sealed class AviatrixFilebeatForwarderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration file.
        /// </summary>
        [Input("configFile")]
        public Input<string>? ConfigFile { get; set; }

        [Input("excludedGateways")]
        private InputList<string>? _excludedGateways;

        /// <summary>
        /// List of excluded gateways.
        /// </summary>
        public InputList<string> ExcludedGateways
        {
            get => _excludedGateways ?? (_excludedGateways = new InputList<string>());
            set => _excludedGateways = value;
        }

        /// <summary>
        /// Port number.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Server IP.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Enabled or not.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Trusted CA file.
        /// </summary>
        [Input("trustedCaFile")]
        public Input<string>? TrustedCaFile { get; set; }

        public AviatrixFilebeatForwarderState()
        {
        }
        public static new AviatrixFilebeatForwarderState Empty => new AviatrixFilebeatForwarderState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent")]
    public partial class AviatrixDatadogAgent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API key.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// List of excluded gateways.
        /// </summary>
        [Output("excludedGateways")]
        public Output<ImmutableArray<string>> ExcludedGateways { get; private set; } = null!;

        /// <summary>
        /// Only export metrics without exporting logs.
        /// </summary>
        [Output("metricsOnly")]
        public Output<bool?> MetricsOnly { get; private set; } = null!;

        /// <summary>
        /// Site preference.
        /// </summary>
        [Output("site")]
        public Output<string?> Site { get; private set; } = null!;

        /// <summary>
        /// Enabled or not.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixDatadogAgent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixDatadogAgent(string name, AviatrixDatadogAgentArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent", name, args ?? new AviatrixDatadogAgentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixDatadogAgent(string name, Input<string> id, AviatrixDatadogAgentState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixDatadogAgent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixDatadogAgent Get(string name, Input<string> id, AviatrixDatadogAgentState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixDatadogAgent(name, id, state, options);
        }
    }

    public sealed class AviatrixDatadogAgentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API key.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

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
        /// Only export metrics without exporting logs.
        /// </summary>
        [Input("metricsOnly")]
        public Input<bool>? MetricsOnly { get; set; }

        /// <summary>
        /// Site preference.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        public AviatrixDatadogAgentArgs()
        {
        }
        public static new AviatrixDatadogAgentArgs Empty => new AviatrixDatadogAgentArgs();
    }

    public sealed class AviatrixDatadogAgentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API key.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

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
        /// Only export metrics without exporting logs.
        /// </summary>
        [Input("metricsOnly")]
        public Input<bool>? MetricsOnly { get; set; }

        /// <summary>
        /// Site preference.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// Enabled or not.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AviatrixDatadogAgentState()
        {
        }
        public static new AviatrixDatadogAgentState Empty => new AviatrixDatadogAgentState();
    }
}

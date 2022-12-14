// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    /// <summary>
    /// The **aviatrix_datadog_agent** resource allows the enabling and disabling of datadog agent.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Enable datadog agent
    ///     var testDatadogAgent = new Aviatrix.AviatrixDatadogAgent("testDatadogAgent", new()
    ///     {
    ///         ApiKey = "your_api_key",
    ///         ExcludedGateways = new[]
    ///         {
    ///             "a",
    ///             "b",
    ///         },
    ///         Site = "datadoghq.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **datadog_agent** can be imported using "datadog_agent", e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent test datadog_agent
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent")]
    public partial class AviatrixDatadogAgent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API key.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        /// </summary>
        [Output("excludedGateways")]
        public Output<ImmutableArray<string>> ExcludedGateways { get; private set; } = null!;

        /// <summary>
        /// Only export metrics without exporting logs. False by default.
        /// </summary>
        [Output("metricsOnly")]
        public Output<bool?> MetricsOnly { get; private set; } = null!;

        /// <summary>
        /// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
        /// </summary>
        [Output("site")]
        public Output<string?> Site { get; private set; } = null!;

        /// <summary>
        /// The status of datadog agent.
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
                AdditionalSecretOutputs =
                {
                    "apiKey",
                },
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
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// API key.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("excludedGateways")]
        private InputList<string>? _excludedGateways;

        /// <summary>
        /// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        /// </summary>
        public InputList<string> ExcludedGateways
        {
            get => _excludedGateways ?? (_excludedGateways = new InputList<string>());
            set => _excludedGateways = value;
        }

        /// <summary>
        /// Only export metrics without exporting logs. False by default.
        /// </summary>
        [Input("metricsOnly")]
        public Input<bool>? MetricsOnly { get; set; }

        /// <summary>
        /// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
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
        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// API key.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("excludedGateways")]
        private InputList<string>? _excludedGateways;

        /// <summary>
        /// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        /// </summary>
        public InputList<string> ExcludedGateways
        {
            get => _excludedGateways ?? (_excludedGateways = new InputList<string>());
            set => _excludedGateways = value;
        }

        /// <summary>
        /// Only export metrics without exporting logs. False by default.
        /// </summary>
        [Input("metricsOnly")]
        public Input<bool>? MetricsOnly { get; set; }

        /// <summary>
        /// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// The status of datadog agent.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AviatrixDatadogAgentState()
        {
        }
        public static new AviatrixDatadogAgentState Empty => new AviatrixDatadogAgentState();
    }
}

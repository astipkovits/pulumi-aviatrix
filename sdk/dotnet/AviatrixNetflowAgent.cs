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
    /// The **aviatrix_netflow_agent** resource allows the enabling and disabling of netflow agent.
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
    ///     // Enable netflow agent
    ///     var testNetflowAgent = new Aviatrix.AviatrixNetflowAgent("testNetflowAgent", new()
    ///     {
    ///         ExcludedGateways = new[]
    ///         {
    ///             "a",
    ///             "b",
    ///         },
    ///         Port = 10,
    ///         ServerIp = "1.2.3.4",
    ///         Version = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **netflow_agent** can be imported using "netflow_agent", e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixNetflowAgent:AviatrixNetflowAgent test netflow_agent
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixNetflowAgent:AviatrixNetflowAgent")]
    public partial class AviatrixNetflowAgent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
        /// </summary>
        [Output("excludedGateways")]
        public Output<ImmutableArray<string>> ExcludedGateways { get; private set; } = null!;

        /// <summary>
        /// Netflow server port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Netflow server IP address.
        /// </summary>
        [Output("serverIp")]
        public Output<string> ServerIp { get; private set; } = null!;

        /// <summary>
        /// The status of netflow agent.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Netflow version (5 or 9). 5 by default.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixNetflowAgent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixNetflowAgent(string name, AviatrixNetflowAgentArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixNetflowAgent:AviatrixNetflowAgent", name, args ?? new AviatrixNetflowAgentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixNetflowAgent(string name, Input<string> id, AviatrixNetflowAgentState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixNetflowAgent:AviatrixNetflowAgent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixNetflowAgent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixNetflowAgent Get(string name, Input<string> id, AviatrixNetflowAgentState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixNetflowAgent(name, id, state, options);
        }
    }

    public sealed class AviatrixNetflowAgentArgs : global::Pulumi.ResourceArgs
    {
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
        /// Netflow server port.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Netflow server IP address.
        /// </summary>
        [Input("serverIp", required: true)]
        public Input<string> ServerIp { get; set; } = null!;

        /// <summary>
        /// Netflow version (5 or 9). 5 by default.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AviatrixNetflowAgentArgs()
        {
        }
        public static new AviatrixNetflowAgentArgs Empty => new AviatrixNetflowAgentArgs();
    }

    public sealed class AviatrixNetflowAgentState : global::Pulumi.ResourceArgs
    {
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
        /// Netflow server port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Netflow server IP address.
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// The status of netflow agent.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Netflow version (5 or 9). 5 by default.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public AviatrixNetflowAgentState()
        {
        }
        public static new AviatrixNetflowAgentState Empty => new AviatrixNetflowAgentState();
    }
}

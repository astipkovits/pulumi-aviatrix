// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixSumologicForwarder:AviatrixSumologicForwarder")]
    public partial class AviatrixSumologicForwarder : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access ID.
        /// </summary>
        [Output("accessId")]
        public Output<string> AccessId { get; private set; } = null!;

        /// <summary>
        /// Access key.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// Custom configuration.
        /// </summary>
        [Output("customConfiguration")]
        public Output<string?> CustomConfiguration { get; private set; } = null!;

        /// <summary>
        /// List of excluded gateways.
        /// </summary>
        [Output("excludedGateways")]
        public Output<ImmutableArray<string>> ExcludedGateways { get; private set; } = null!;

        /// <summary>
        /// Source category.
        /// </summary>
        [Output("sourceCategory")]
        public Output<string?> SourceCategory { get; private set; } = null!;

        /// <summary>
        /// Enabled or not.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixSumologicForwarder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixSumologicForwarder(string name, AviatrixSumologicForwarderArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSumologicForwarder:AviatrixSumologicForwarder", name, args ?? new AviatrixSumologicForwarderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixSumologicForwarder(string name, Input<string> id, AviatrixSumologicForwarderState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSumologicForwarder:AviatrixSumologicForwarder", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixSumologicForwarder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixSumologicForwarder Get(string name, Input<string> id, AviatrixSumologicForwarderState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixSumologicForwarder(name, id, state, options);
        }
    }

    public sealed class AviatrixSumologicForwarderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access ID.
        /// </summary>
        [Input("accessId", required: true)]
        public Input<string> AccessId { get; set; } = null!;

        /// <summary>
        /// Access key.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// Custom configuration.
        /// </summary>
        [Input("customConfiguration")]
        public Input<string>? CustomConfiguration { get; set; }

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
        /// Source category.
        /// </summary>
        [Input("sourceCategory")]
        public Input<string>? SourceCategory { get; set; }

        public AviatrixSumologicForwarderArgs()
        {
        }
        public static new AviatrixSumologicForwarderArgs Empty => new AviatrixSumologicForwarderArgs();
    }

    public sealed class AviatrixSumologicForwarderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access ID.
        /// </summary>
        [Input("accessId")]
        public Input<string>? AccessId { get; set; }

        /// <summary>
        /// Access key.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// Custom configuration.
        /// </summary>
        [Input("customConfiguration")]
        public Input<string>? CustomConfiguration { get; set; }

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
        /// Source category.
        /// </summary>
        [Input("sourceCategory")]
        public Input<string>? SourceCategory { get; set; }

        /// <summary>
        /// Enabled or not.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AviatrixSumologicForwarderState()
        {
        }
        public static new AviatrixSumologicForwarderState Empty => new AviatrixSumologicForwarderState();
    }
}

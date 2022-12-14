// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class AviatrixAppDomainSelectorMatchExpressionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// - Account ID this expression matches.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// - Account name this expression matches.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// - CIDR block or IP Address this expression matches. `cidr` cannot be used with any other filters in the same `match_expressions` block.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// - Region this expression matches.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// - Resource ID this expression matches.
        /// </summary>
        [Input("resId")]
        public Input<string>? ResId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// - Map of tags this expression matches.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// - Type of resource this expression matches. Must be one of "vm", "vpc" or "subnet". `type` is required when `cidr` is not used.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// - Zone this expression matches.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AviatrixAppDomainSelectorMatchExpressionGetArgs()
        {
        }
        public static new AviatrixAppDomainSelectorMatchExpressionGetArgs Empty => new AviatrixAppDomainSelectorMatchExpressionGetArgs();
    }
}

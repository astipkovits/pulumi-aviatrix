// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixFqdn:AviatrixFqdn")]
    public partial class AviatrixFqdn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of one or more domain names/tag rules.
        /// </summary>
        [Output("domainNames")]
        public Output<ImmutableArray<Outputs.AviatrixFqdnDomainName>> DomainNames { get; private set; } = null!;

        /// <summary>
        /// FQDN Filter Tag Status. Valid values: true or false.
        /// </summary>
        [Output("fqdnEnabled")]
        public Output<bool?> FqdnEnabled { get; private set; } = null!;

        /// <summary>
        /// Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        /// </summary>
        [Output("fqdnMode")]
        public Output<string?> FqdnMode { get; private set; } = null!;

        /// <summary>
        /// FQDN Filter Tag Name.
        /// </summary>
        [Output("fqdnTag")]
        public Output<string> FqdnTag { get; private set; } = null!;

        /// <summary>
        /// A list of gateways to attach to the specific tag.
        /// </summary>
        [Output("gwFilterTagLists")]
        public Output<ImmutableArray<Outputs.AviatrixFqdnGwFilterTagList>> GwFilterTagLists { get; private set; } = null!;

        /// <summary>
        /// Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
        /// resources.
        /// </summary>
        [Output("manageDomainNames")]
        public Output<bool?> ManageDomainNames { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFqdn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFqdn(string name, AviatrixFqdnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFqdn:AviatrixFqdn", name, args ?? new AviatrixFqdnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFqdn(string name, Input<string> id, AviatrixFqdnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFqdn:AviatrixFqdn", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/astipkovits/pulumi-aviatrix/raw/main/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixFqdn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFqdn Get(string name, Input<string> id, AviatrixFqdnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFqdn(name, id, state, options);
        }
    }

    public sealed class AviatrixFqdnArgs : global::Pulumi.ResourceArgs
    {
        [Input("domainNames")]
        private InputList<Inputs.AviatrixFqdnDomainNameArgs>? _domainNames;

        /// <summary>
        /// A list of one or more domain names/tag rules.
        /// </summary>
        public InputList<Inputs.AviatrixFqdnDomainNameArgs> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<Inputs.AviatrixFqdnDomainNameArgs>());
            set => _domainNames = value;
        }

        /// <summary>
        /// FQDN Filter Tag Status. Valid values: true or false.
        /// </summary>
        [Input("fqdnEnabled")]
        public Input<bool>? FqdnEnabled { get; set; }

        /// <summary>
        /// Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        /// </summary>
        [Input("fqdnMode")]
        public Input<string>? FqdnMode { get; set; }

        /// <summary>
        /// FQDN Filter Tag Name.
        /// </summary>
        [Input("fqdnTag", required: true)]
        public Input<string> FqdnTag { get; set; } = null!;

        [Input("gwFilterTagLists")]
        private InputList<Inputs.AviatrixFqdnGwFilterTagListArgs>? _gwFilterTagLists;

        /// <summary>
        /// A list of gateways to attach to the specific tag.
        /// </summary>
        public InputList<Inputs.AviatrixFqdnGwFilterTagListArgs> GwFilterTagLists
        {
            get => _gwFilterTagLists ?? (_gwFilterTagLists = new InputList<Inputs.AviatrixFqdnGwFilterTagListArgs>());
            set => _gwFilterTagLists = value;
        }

        /// <summary>
        /// Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
        /// resources.
        /// </summary>
        [Input("manageDomainNames")]
        public Input<bool>? ManageDomainNames { get; set; }

        public AviatrixFqdnArgs()
        {
        }
        public static new AviatrixFqdnArgs Empty => new AviatrixFqdnArgs();
    }

    public sealed class AviatrixFqdnState : global::Pulumi.ResourceArgs
    {
        [Input("domainNames")]
        private InputList<Inputs.AviatrixFqdnDomainNameGetArgs>? _domainNames;

        /// <summary>
        /// A list of one or more domain names/tag rules.
        /// </summary>
        public InputList<Inputs.AviatrixFqdnDomainNameGetArgs> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<Inputs.AviatrixFqdnDomainNameGetArgs>());
            set => _domainNames = value;
        }

        /// <summary>
        /// FQDN Filter Tag Status. Valid values: true or false.
        /// </summary>
        [Input("fqdnEnabled")]
        public Input<bool>? FqdnEnabled { get; set; }

        /// <summary>
        /// Specify the tag color to be a white-list tag or black-list tag. 'white' or 'black'
        /// </summary>
        [Input("fqdnMode")]
        public Input<string>? FqdnMode { get; set; }

        /// <summary>
        /// FQDN Filter Tag Name.
        /// </summary>
        [Input("fqdnTag")]
        public Input<string>? FqdnTag { get; set; }

        [Input("gwFilterTagLists")]
        private InputList<Inputs.AviatrixFqdnGwFilterTagListGetArgs>? _gwFilterTagLists;

        /// <summary>
        /// A list of gateways to attach to the specific tag.
        /// </summary>
        public InputList<Inputs.AviatrixFqdnGwFilterTagListGetArgs> GwFilterTagLists
        {
            get => _gwFilterTagLists ?? (_gwFilterTagLists = new InputList<Inputs.AviatrixFqdnGwFilterTagListGetArgs>());
            set => _gwFilterTagLists = value;
        }

        /// <summary>
        /// Enable to manage domain name rules in-line. If false, domain name rules must be managed using `aviatrix_fqdn_tag_rule`
        /// resources.
        /// </summary>
        [Input("manageDomainNames")]
        public Input<bool>? ManageDomainNames { get; set; }

        public AviatrixFqdnState()
        {
        }
        public static new AviatrixFqdnState Empty => new AviatrixFqdnState();
    }
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload")]
    public partial class AviatrixVpnCertDownload : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the VPN Certificate download is enabled. Supported Values: "true", "false"
        /// </summary>
        [Output("downloadEnabled")]
        public Output<bool?> DownloadEnabled { get; private set; } = null!;

        /// <summary>
        /// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is
        /// supported. Example: ["saml_endpoint_1"].
        /// </summary>
        [Output("samlEndpoints")]
        public Output<ImmutableArray<string>> SamlEndpoints { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixVpnCertDownload resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixVpnCertDownload(string name, AviatrixVpnCertDownloadArgs? args = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload", name, args ?? new AviatrixVpnCertDownloadArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixVpnCertDownload(string name, Input<string> id, AviatrixVpnCertDownloadState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixVpnCertDownload resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixVpnCertDownload Get(string name, Input<string> id, AviatrixVpnCertDownloadState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixVpnCertDownload(name, id, state, options);
        }
    }

    public sealed class AviatrixVpnCertDownloadArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the VPN Certificate download is enabled. Supported Values: "true", "false"
        /// </summary>
        [Input("downloadEnabled")]
        public Input<bool>? DownloadEnabled { get; set; }

        [Input("samlEndpoints")]
        private InputList<string>? _samlEndpoints;

        /// <summary>
        /// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is
        /// supported. Example: ["saml_endpoint_1"].
        /// </summary>
        public InputList<string> SamlEndpoints
        {
            get => _samlEndpoints ?? (_samlEndpoints = new InputList<string>());
            set => _samlEndpoints = value;
        }

        public AviatrixVpnCertDownloadArgs()
        {
        }
        public static new AviatrixVpnCertDownloadArgs Empty => new AviatrixVpnCertDownloadArgs();
    }

    public sealed class AviatrixVpnCertDownloadState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the VPN Certificate download is enabled. Supported Values: "true", "false"
        /// </summary>
        [Input("downloadEnabled")]
        public Input<bool>? DownloadEnabled { get; set; }

        [Input("samlEndpoints")]
        private InputList<string>? _samlEndpoints;

        /// <summary>
        /// List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is
        /// supported. Example: ["saml_endpoint_1"].
        /// </summary>
        public InputList<string> SamlEndpoints
        {
            get => _samlEndpoints ?? (_samlEndpoints = new InputList<string>());
            set => _samlEndpoints = value;
        }

        public AviatrixVpnCertDownloadState()
        {
        }
        public static new AviatrixVpnCertDownloadState Empty => new AviatrixVpnCertDownloadState();
    }
}

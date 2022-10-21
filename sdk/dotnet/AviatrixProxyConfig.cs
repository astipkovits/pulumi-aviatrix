// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig")]
    public partial class AviatrixProxyConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// http proxy URL.
        /// </summary>
        [Output("httpProxy")]
        public Output<string> HttpProxy { get; private set; } = null!;

        /// <summary>
        /// https proxy URL.
        /// </summary>
        [Output("httpsProxy")]
        public Output<string> HttpsProxy { get; private set; } = null!;

        /// <summary>
        /// Server CA Certificate file.
        /// </summary>
        [Output("proxyCaCertificate")]
        public Output<string?> ProxyCaCertificate { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixProxyConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixProxyConfig(string name, AviatrixProxyConfigArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig", name, args ?? new AviatrixProxyConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixProxyConfig(string name, Input<string> id, AviatrixProxyConfigState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixProxyConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixProxyConfig Get(string name, Input<string> id, AviatrixProxyConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixProxyConfig(name, id, state, options);
        }
    }

    public sealed class AviatrixProxyConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// http proxy URL.
        /// </summary>
        [Input("httpProxy", required: true)]
        public Input<string> HttpProxy { get; set; } = null!;

        /// <summary>
        /// https proxy URL.
        /// </summary>
        [Input("httpsProxy", required: true)]
        public Input<string> HttpsProxy { get; set; } = null!;

        /// <summary>
        /// Server CA Certificate file.
        /// </summary>
        [Input("proxyCaCertificate")]
        public Input<string>? ProxyCaCertificate { get; set; }

        public AviatrixProxyConfigArgs()
        {
        }
        public static new AviatrixProxyConfigArgs Empty => new AviatrixProxyConfigArgs();
    }

    public sealed class AviatrixProxyConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// http proxy URL.
        /// </summary>
        [Input("httpProxy")]
        public Input<string>? HttpProxy { get; set; }

        /// <summary>
        /// https proxy URL.
        /// </summary>
        [Input("httpsProxy")]
        public Input<string>? HttpsProxy { get; set; }

        /// <summary>
        /// Server CA Certificate file.
        /// </summary>
        [Input("proxyCaCertificate")]
        public Input<string>? ProxyCaCertificate { get; set; }

        public AviatrixProxyConfigState()
        {
        }
        public static new AviatrixProxyConfigState Empty => new AviatrixProxyConfigState();
    }
}

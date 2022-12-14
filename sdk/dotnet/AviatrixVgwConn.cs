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
    /// The **aviatrix_vgw_conn** resource manages the connection between the Aviatrix transit gateway and AWS VGW for purposes of Transit Network.
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
    ///     // Create an Aviatrix VGW Connection
    ///     var testVgwConn = new Aviatrix.AviatrixVgwConn("testVgwConn", new()
    ///     {
    ///         BgpLocalAsNum = "65001",
    ///         BgpVgwAccount = "dev-account-1",
    ///         BgpVgwId = "vgw-abcd1234",
    ///         BgpVgwRegion = "us-east-1",
    ///         ConnName = "my-connection-vgw-to-tgw",
    ///         GwName = "my-transit-gw",
    ///         PrependAsPaths = new[]
    ///         {
    ///             "65001",
    ///             "65001",
    ///         },
    ///         VpcId = "vpc-abcd1234",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **vgw_conn** can be imported using the `conn_name` and `vpc_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixVgwConn:AviatrixVgwConn test conn_name~vpc_id
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixVgwConn:AviatrixVgwConn")]
    public partial class AviatrixVgwConn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
        /// </summary>
        [Output("bgpLocalAsNum")]
        public Output<string> BgpLocalAsNum { get; private set; } = null!;

        /// <summary>
        /// Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".
        /// </summary>
        [Output("bgpVgwAccount")]
        public Output<string> BgpVgwAccount { get; private set; } = null!;

        /// <summary>
        /// ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".
        /// </summary>
        [Output("bgpVgwId")]
        public Output<string> BgpVgwId { get; private set; } = null!;

        /// <summary>
        /// Region of AWS VGW that will be used for this connection. Example: "us-east-1".
        /// </summary>
        [Output("bgpVgwRegion")]
        public Output<string> BgpVgwRegion { get; private set; } = null!;

        /// <summary>
        /// The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".
        /// </summary>
        [Output("connName")]
        public Output<string> ConnName { get; private set; } = null!;

        /// <summary>
        /// Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
        /// </summary>
        [Output("enableEventTriggeredHa")]
        public Output<bool?> EnableEventTriggeredHa { get; private set; } = null!;

        /// <summary>
        /// Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Output("enableLearnedCidrsApproval")]
        public Output<bool?> EnableLearnedCidrsApproval { get; private set; } = null!;

        /// <summary>
        /// Name of the Transit Gateway. Example: "my-transit-gw".
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.
        /// </summary>
        [Output("manualBgpAdvertisedCidrs")]
        public Output<ImmutableArray<string>> ManualBgpAdvertisedCidrs { get; private set; } = null!;

        /// <summary>
        /// Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
        /// </summary>
        [Output("prependAsPaths")]
        public Output<ImmutableArray<string>> PrependAsPaths { get; private set; } = null!;

        /// <summary>
        /// VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixVgwConn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixVgwConn(string name, AviatrixVgwConnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVgwConn:AviatrixVgwConn", name, args ?? new AviatrixVgwConnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixVgwConn(string name, Input<string> id, AviatrixVgwConnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVgwConn:AviatrixVgwConn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixVgwConn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixVgwConn Get(string name, Input<string> id, AviatrixVgwConnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixVgwConn(name, id, state, options);
        }
    }

    public sealed class AviatrixVgwConnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
        /// </summary>
        [Input("bgpLocalAsNum", required: true)]
        public Input<string> BgpLocalAsNum { get; set; } = null!;

        /// <summary>
        /// Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".
        /// </summary>
        [Input("bgpVgwAccount", required: true)]
        public Input<string> BgpVgwAccount { get; set; } = null!;

        /// <summary>
        /// ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".
        /// </summary>
        [Input("bgpVgwId", required: true)]
        public Input<string> BgpVgwId { get; set; } = null!;

        /// <summary>
        /// Region of AWS VGW that will be used for this connection. Example: "us-east-1".
        /// </summary>
        [Input("bgpVgwRegion", required: true)]
        public Input<string> BgpVgwRegion { get; set; } = null!;

        /// <summary>
        /// The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".
        /// </summary>
        [Input("connName", required: true)]
        public Input<string> ConnName { get; set; } = null!;

        /// <summary>
        /// Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
        /// </summary>
        [Input("enableEventTriggeredHa")]
        public Input<bool>? EnableEventTriggeredHa { get; set; }

        /// <summary>
        /// Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Name of the Transit Gateway. Example: "my-transit-gw".
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        [Input("manualBgpAdvertisedCidrs")]
        private InputList<string>? _manualBgpAdvertisedCidrs;

        /// <summary>
        /// Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.
        /// </summary>
        public InputList<string> ManualBgpAdvertisedCidrs
        {
            get => _manualBgpAdvertisedCidrs ?? (_manualBgpAdvertisedCidrs = new InputList<string>());
            set => _manualBgpAdvertisedCidrs = value;
        }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public AviatrixVgwConnArgs()
        {
        }
        public static new AviatrixVgwConnArgs Empty => new AviatrixVgwConnArgs();
    }

    public sealed class AviatrixVgwConnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
        /// </summary>
        [Input("bgpLocalAsNum")]
        public Input<string>? BgpLocalAsNum { get; set; }

        /// <summary>
        /// Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".
        /// </summary>
        [Input("bgpVgwAccount")]
        public Input<string>? BgpVgwAccount { get; set; }

        /// <summary>
        /// ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".
        /// </summary>
        [Input("bgpVgwId")]
        public Input<string>? BgpVgwId { get; set; }

        /// <summary>
        /// Region of AWS VGW that will be used for this connection. Example: "us-east-1".
        /// </summary>
        [Input("bgpVgwRegion")]
        public Input<string>? BgpVgwRegion { get; set; }

        /// <summary>
        /// The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".
        /// </summary>
        [Input("connName")]
        public Input<string>? ConnName { get; set; }

        /// <summary>
        /// Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.
        /// </summary>
        [Input("enableEventTriggeredHa")]
        public Input<bool>? EnableEventTriggeredHa { get; set; }

        /// <summary>
        /// Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enableLearnedCidrsApproval")]
        public Input<bool>? EnableLearnedCidrsApproval { get; set; }

        /// <summary>
        /// Name of the Transit Gateway. Example: "my-transit-gw".
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        [Input("manualBgpAdvertisedCidrs")]
        private InputList<string>? _manualBgpAdvertisedCidrs;

        /// <summary>
        /// Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.
        /// </summary>
        public InputList<string> ManualBgpAdvertisedCidrs
        {
            get => _manualBgpAdvertisedCidrs ?? (_manualBgpAdvertisedCidrs = new InputList<string>());
            set => _manualBgpAdvertisedCidrs = value;
        }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixVgwConnState()
        {
        }
        public static new AviatrixVgwConnState Empty => new AviatrixVgwConnState();
    }
}

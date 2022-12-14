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
    /// The **aviatrix_vpc** resource allows the creation and management of Aviatrix-created VPCs of various cloud types.
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
    ///     // Create an AWS VPC
    ///     var awsVpc = new Aviatrix.AviatrixVpc("awsVpc", new()
    ///     {
    ///         AccountName = "devops",
    ///         AviatrixFirenetVpc = false,
    ///         AviatrixTransitVpc = false,
    ///         Cidr = "10.0.0.0/16",
    ///         CloudType = 1,
    ///         Region = "us-west-1",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a GCP VPC
    ///     var gcpVpc = new Aviatrix.AviatrixVpc("gcpVpc", new()
    ///     {
    ///         AccountName = "devops",
    ///         CloudType = 4,
    ///         Subnets = new[]
    ///         {
    ///             new Aviatrix.Inputs.AviatrixVpcSubnetArgs
    ///             {
    ///                 Cidr = "10.10.0.0/24",
    ///                 Name = "subnet-1",
    ///                 Region = "us-west1",
    ///             },
    ///             new Aviatrix.Inputs.AviatrixVpcSubnetArgs
    ///             {
    ///                 Cidr = "10.11.0.0/24",
    ///                 Name = "subnet-2",
    ///                 Region = "us-west2",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Azure VNet
    ///     var azureVnet = new Aviatrix.AviatrixVpc("azureVnet", new()
    ///     {
    ///         AccountName = "devops",
    ///         AviatrixFirenetVpc = false,
    ///         Cidr = "12.0.0.0/16",
    ///         CloudType = 8,
    ///         Region = "Central US",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an OCI VPC
    ///     var ociVpc = new Aviatrix.AviatrixVpc("ociVpc", new()
    ///     {
    ///         AccountName = "devops",
    ///         Cidr = "10.0.0.0/24",
    ///         CloudType = 16,
    ///         Region = "us-ashburn-1",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an AzureGov VNet
    ///     var azureVnet = new Aviatrix.AviatrixVpc("azureVnet", new()
    ///     {
    ///         AccountName = "devops",
    ///         AviatrixFirenetVpc = false,
    ///         Cidr = "12.0.0.0/16",
    ///         CloudType = 32,
    ///         Region = "USGov Arizona",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an AWSGov VPC
    ///     var awsgovVnet = new Aviatrix.AviatrixVpc("awsgovVnet", new()
    ///     {
    ///         AccountName = "devops",
    ///         AviatrixFirenetVpc = false,
    ///         AviatrixTransitVpc = false,
    ///         Cidr = "12.0.0.0/20",
    ///         CloudType = 256,
    ///         Region = "us-gov-west-1",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an AWS China VPC
    ///     var awsChinaVnet = new Aviatrix.AviatrixVpc("awsChinaVnet", new()
    ///     {
    ///         AccountName = "devops",
    ///         AviatrixTransitVpc = false,
    ///         Cidr = "12.0.0.0/20",
    ///         CloudType = 1024,
    ///         Region = "cn-north-1",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Azure China VNet
    ///     var azureChinaVnet = new Aviatrix.AviatrixVpc("azureChinaVnet", new()
    ///     {
    ///         AccountName = "devops",
    ///         Cidr = "12.0.0.0/16",
    ///         CloudType = 2048,
    ///         Region = "China North",
    ///     });
    /// 
    /// });
    /// ```
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aviatrix = Pulumi.Aviatrix;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create an Alibaba Cloud VPC
    ///     var aliyunVpc = new Aviatrix.AviatrixVpc("aliyunVpc", new()
    ///     {
    ///         AccountName = "devops",
    ///         Cidr = "10.0.0.0/20",
    ///         CloudType = 8192,
    ///         Region = "acs-us-west-1 (Silicon Valley)",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// **vpc** can be imported using the VPC's `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aviatrix:index/aviatrixVpc:AviatrixVpc test name
    /// ```
    /// </summary>
    [AviatrixResourceType("aviatrix:index/aviatrixVpc:AviatrixVpc")]
    public partial class AviatrixVpc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// List of OCI availability domains.
        /// </summary>
        [Output("availabilityDomains")]
        public Output<ImmutableArray<string>> AvailabilityDomains { get; private set; } = null!;

        /// <summary>
        /// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
        /// </summary>
        [Output("aviatrixFirenetVpc")]
        public Output<bool?> AviatrixFirenetVpc { get; private set; } = null!;

        /// <summary>
        /// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
        /// </summary>
        [Output("aviatrixTransitVpc")]
        public Output<bool?> AviatrixTransitVpc { get; private set; } = null!;

        /// <summary>
        /// Azure VNet resource ID.
        /// </summary>
        [Output("azureVnetResourceId")]
        public Output<string> AzureVnetResourceId { get; private set; } = null!;

        /// <summary>
        /// CIDR block.
        /// </summary>
        [Output("cidr")]
        public Output<string?> Cidr { get; private set; } = null!;

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
        /// </summary>
        [Output("cloudType")]
        public Output<int> CloudType { get; private set; } = null!;

        /// <summary>
        /// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloud_type = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Output("enableNativeGwlb")]
        public Output<bool?> EnableNativeGwlb { get; private set; } = null!;

        /// <summary>
        /// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Output("enablePrivateOobSubnet")]
        public Output<bool?> EnablePrivateOobSubnet { get; private set; } = null!;

        /// <summary>
        /// List of OCI fault domains.
        /// </summary>
        [Output("faultDomains")]
        public Output<ImmutableArray<string>> FaultDomains { get; private set; } = null!;

        /// <summary>
        /// Name of this subnet.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
        /// </summary>
        [Output("numOfSubnetPairs")]
        public Output<int?> NumOfSubnetPairs { get; private set; } = null!;

        /// <summary>
        /// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
        /// </summary>
        [Output("privateModeSubnets")]
        public Output<bool?> PrivateModeSubnets { get; private set; } = null!;

        /// <summary>
        /// List of private subnet of the VPC(AWS, Azure) to be created.
        /// </summary>
        [Output("privateSubnets")]
        public Output<ImmutableArray<Outputs.AviatrixVpcPrivateSubnet>> PrivateSubnets { get; private set; } = null!;

        /// <summary>
        /// List of public subnet of the VPC(AWS, Azure) to be created.
        /// </summary>
        [Output("publicSubnets")]
        public Output<ImmutableArray<Outputs.AviatrixVpcPublicSubnet>> PublicSubnets { get; private set; } = null!;

        /// <summary>
        /// Region of this subnet.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
        /// </summary>
        [Output("resourceGroup")]
        public Output<string> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
        /// </summary>
        [Output("routeTables")]
        public Output<ImmutableArray<string>> RouteTables { get; private set; } = null!;

        /// <summary>
        /// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
        /// </summary>
        [Output("subnetSize")]
        public Output<int?> SubnetSize { get; private set; } = null!;

        /// <summary>
        /// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<Outputs.AviatrixVpcSubnet>> Subnets { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC to be created.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixVpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixVpc(string name, AviatrixVpcArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVpc:AviatrixVpc", name, args ?? new AviatrixVpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixVpc(string name, Input<string> id, AviatrixVpcState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixVpc:AviatrixVpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixVpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixVpc Get(string name, Input<string> id, AviatrixVpcState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixVpc(name, id, state, options);
        }
    }

    public sealed class AviatrixVpcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
        /// </summary>
        [Input("aviatrixFirenetVpc")]
        public Input<bool>? AviatrixFirenetVpc { get; set; }

        /// <summary>
        /// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
        /// </summary>
        [Input("aviatrixTransitVpc")]
        public Input<bool>? AviatrixTransitVpc { get; set; }

        /// <summary>
        /// CIDR block.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
        /// </summary>
        [Input("cloudType", required: true)]
        public Input<int> CloudType { get; set; } = null!;

        /// <summary>
        /// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloud_type = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enableNativeGwlb")]
        public Input<bool>? EnableNativeGwlb { get; set; }

        /// <summary>
        /// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enablePrivateOobSubnet")]
        public Input<bool>? EnablePrivateOobSubnet { get; set; }

        /// <summary>
        /// Name of this subnet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
        /// </summary>
        [Input("numOfSubnetPairs")]
        public Input<int>? NumOfSubnetPairs { get; set; }

        /// <summary>
        /// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
        /// </summary>
        [Input("privateModeSubnets")]
        public Input<bool>? PrivateModeSubnets { get; set; }

        /// <summary>
        /// Region of this subnet.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
        /// </summary>
        [Input("subnetSize")]
        public Input<int>? SubnetSize { get; set; }

        [Input("subnets")]
        private InputList<Inputs.AviatrixVpcSubnetArgs>? _subnets;

        /// <summary>
        /// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
        /// </summary>
        public InputList<Inputs.AviatrixVpcSubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.AviatrixVpcSubnetArgs>());
            set => _subnets = value;
        }

        public AviatrixVpcArgs()
        {
        }
        public static new AviatrixVpcArgs Empty => new AviatrixVpcArgs();
    }

    public sealed class AviatrixVpcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("availabilityDomains")]
        private InputList<string>? _availabilityDomains;

        /// <summary>
        /// List of OCI availability domains.
        /// </summary>
        public InputList<string> AvailabilityDomains
        {
            get => _availabilityDomains ?? (_availabilityDomains = new InputList<string>());
            set => _availabilityDomains = value;
        }

        /// <summary>
        /// Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions. **Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
        /// </summary>
        [Input("aviatrixFirenetVpc")]
        public Input<bool>? AviatrixFirenetVpc { get; set; }

        /// <summary>
        /// Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions. **Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers.** Valid values: true, false. Default: false.
        /// </summary>
        [Input("aviatrixTransitVpc")]
        public Input<bool>? AviatrixTransitVpc { get; set; }

        /// <summary>
        /// Azure VNet resource ID.
        /// </summary>
        [Input("azureVnetResourceId")]
        public Input<string>? AzureVnetResourceId { get; set; }

        /// <summary>
        /// CIDR block.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloud_type = 1 (AWS). **This option is only applicable to TGW-integrated FireNet**. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enableNativeGwlb")]
        public Input<bool>? EnableNativeGwlb { get; set; }

        /// <summary>
        /// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
        /// </summary>
        [Input("enablePrivateOobSubnet")]
        public Input<bool>? EnablePrivateOobSubnet { get; set; }

        [Input("faultDomains")]
        private InputList<string>? _faultDomains;

        /// <summary>
        /// List of OCI fault domains.
        /// </summary>
        public InputList<string> FaultDomains
        {
            get => _faultDomains ?? (_faultDomains = new InputList<string>());
            set => _faultDomains = value;
        }

        /// <summary>
        /// Name of this subnet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
        /// </summary>
        [Input("numOfSubnetPairs")]
        public Input<int>? NumOfSubnetPairs { get; set; }

        /// <summary>
        /// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
        /// </summary>
        [Input("privateModeSubnets")]
        public Input<bool>? PrivateModeSubnets { get; set; }

        [Input("privateSubnets")]
        private InputList<Inputs.AviatrixVpcPrivateSubnetGetArgs>? _privateSubnets;

        /// <summary>
        /// List of private subnet of the VPC(AWS, Azure) to be created.
        /// </summary>
        public InputList<Inputs.AviatrixVpcPrivateSubnetGetArgs> PrivateSubnets
        {
            get => _privateSubnets ?? (_privateSubnets = new InputList<Inputs.AviatrixVpcPrivateSubnetGetArgs>());
            set => _privateSubnets = value;
        }

        [Input("publicSubnets")]
        private InputList<Inputs.AviatrixVpcPublicSubnetGetArgs>? _publicSubnets;

        /// <summary>
        /// List of public subnet of the VPC(AWS, Azure) to be created.
        /// </summary>
        public InputList<Inputs.AviatrixVpcPublicSubnetGetArgs> PublicSubnets
        {
            get => _publicSubnets ?? (_publicSubnets = new InputList<Inputs.AviatrixVpcPublicSubnetGetArgs>());
            set => _publicSubnets = value;
        }

        /// <summary>
        /// Region of this subnet.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        [Input("routeTables")]
        private InputList<string>? _routeTables;

        /// <summary>
        /// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
        /// </summary>
        public InputList<string> RouteTables
        {
            get => _routeTables ?? (_routeTables = new InputList<string>());
            set => _routeTables = value;
        }

        /// <summary>
        /// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
        /// </summary>
        [Input("subnetSize")]
        public Input<int>? SubnetSize { get; set; }

        [Input("subnets")]
        private InputList<Inputs.AviatrixVpcSubnetGetArgs>? _subnets;

        /// <summary>
        /// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
        /// </summary>
        public InputList<Inputs.AviatrixVpcSubnetGetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.AviatrixVpcSubnetGetArgs>());
            set => _subnets = value;
        }

        /// <summary>
        /// ID of the VPC to be created.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AviatrixVpcState()
        {
        }
        public static new AviatrixVpcState Empty => new AviatrixVpcState();
    }
}

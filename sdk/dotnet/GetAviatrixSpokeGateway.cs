// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixSpokeGateway
    {
        /// <summary>
        /// The **aviatrix_spoke_gateway** data source provides details about a specific spoke gateway created by the Aviatrix Controller.
        /// 
        /// This data source can prove useful when a module accepts a spoke gateway's detail as an input variable.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aviatrix = Pulumi.Aviatrix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aviatrix.GetAviatrixSpokeGateway.Invoke(new()
        ///     {
        ///         GwName = "gatewayname",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAviatrixSpokeGatewayResult> InvokeAsync(GetAviatrixSpokeGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixSpokeGatewayResult>("aviatrix:index/getAviatrixSpokeGateway:getAviatrixSpokeGateway", args ?? new GetAviatrixSpokeGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// The **aviatrix_spoke_gateway** data source provides details about a specific spoke gateway created by the Aviatrix Controller.
        /// 
        /// This data source can prove useful when a module accepts a spoke gateway's detail as an input variable.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aviatrix = Pulumi.Aviatrix;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Aviatrix.GetAviatrixSpokeGateway.Invoke(new()
        ///     {
        ///         GwName = "gatewayname",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAviatrixSpokeGatewayResult> Invoke(GetAviatrixSpokeGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAviatrixSpokeGatewayResult>("aviatrix:index/getAviatrixSpokeGateway:getAviatrixSpokeGateway", args ?? new GetAviatrixSpokeGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAviatrixSpokeGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Spoke gateway name. It can be used for getting spoke gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public string GwName { get; set; } = null!;

        public GetAviatrixSpokeGatewayArgs()
        {
        }
        public static new GetAviatrixSpokeGatewayArgs Empty => new GetAviatrixSpokeGatewayArgs();
    }

    public sealed class GetAviatrixSpokeGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Spoke gateway name. It can be used for getting spoke gateway.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        public GetAviatrixSpokeGatewayInvokeArgs()
        {
        }
        public static new GetAviatrixSpokeGatewayInvokeArgs Empty => new GetAviatrixSpokeGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetAviatrixSpokeGatewayResult
    {
        /// <summary>
        /// Aviatrix account name.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.
        /// </summary>
        public readonly bool AllocateNewEip;
        public readonly ImmutableArray<string> ApprovedLearnedCidrs;
        /// <summary>
        /// Availability domain for OCI.
        /// </summary>
        public readonly string AvailabilityDomain;
        public readonly string AzureEipNameResourceGroup;
        public readonly bool BgpEcmp;
        public readonly int BgpHoldTime;
        public readonly int BgpPollingTime;
        /// <summary>
        /// Cloud instance ID.
        /// </summary>
        public readonly string CloudInstanceId;
        /// <summary>
        /// Type of cloud service provider.
        /// </summary>
        public readonly int CloudType;
        /// <summary>
        /// A list of comma separated CIDRs to be customized for the spoke VPC routes.
        /// </summary>
        public readonly string CustomizedSpokeVpcRoutes;
        public readonly bool DisableRoutePropagation;
        /// <summary>
        /// The EIP address of the Spoke Gateway.
        /// </summary>
        public readonly string Eip;
        public readonly bool EnableActiveStandby;
        public readonly bool EnableActiveStandbyPreemptive;
        public readonly bool EnableAutoAdvertiseS2cCidrs;
        public readonly bool EnableBgp;
        /// <summary>
        /// Status of Encrypt Volume of spoke gateway.
        /// </summary>
        public readonly bool EnableEncryptVolume;
        public readonly bool EnableJumboFrame;
        public readonly bool EnableLearnedCidrsApproval;
        public readonly bool EnableMonitorGatewaySubnets;
        /// <summary>
        /// Status of private OOB for the spoke gateway.
        /// </summary>
        public readonly bool EnablePrivateOob;
        public readonly bool EnablePrivateVpcDefaultRoute;
        public readonly bool EnableSkipPublicRouteTableUpdate;
        public readonly bool EnableSpotInstance;
        /// <summary>
        /// Status of VPC Dns Server of spoke gateway.
        /// </summary>
        public readonly bool EnableVpcDnsServer;
        /// <summary>
        /// Fault domain for OCI.
        /// </summary>
        public readonly string FaultDomain;
        /// <summary>
        /// A list of comma separated CIDRs to be filtered from the spoke VPC route table.
        /// </summary>
        public readonly string FilteredSpokeVpcRoutes;
        /// <summary>
        /// Aviatrix spoke gateway name.
        /// </summary>
        public readonly string GwName;
        /// <summary>
        /// Size of spoke gateway instance.
        /// </summary>
        public readonly string GwSize;
        /// <summary>
        /// HA gateway availability domain for OCI.
        /// </summary>
        public readonly string HaAvailabilityDomain;
        public readonly string HaAzureEipNameResourceGroup;
        /// <summary>
        /// Cloud instance ID of HA spoke gateway.
        /// </summary>
        public readonly string HaCloudInstanceId;
        /// <summary>
        /// The EIP address of the HA Spoke Gateway.
        /// </summary>
        public readonly string HaEip;
        /// <summary>
        /// HA gateway fault domain for OCI.
        /// </summary>
        public readonly string HaFaultDomain;
        /// <summary>
        /// Aviatrix spoke gateway unique name of HA spoke gateway.
        /// </summary>
        public readonly string HaGwName;
        /// <summary>
        /// HA Gateway Size.
        /// </summary>
        public readonly string HaGwSize;
        /// <summary>
        /// The image version of the HA gateway.
        /// </summary>
        public readonly string HaImageVersion;
        /// <summary>
        /// AZ of subnet being created for Insane Mode Spoke HA Gateway.
        /// </summary>
        public readonly string HaInsaneModeAz;
        /// <summary>
        /// HA OOB availability zone.
        /// </summary>
        public readonly string HaOobAvailabilityZone;
        /// <summary>
        /// HA OOB management subnet.
        /// </summary>
        public readonly string HaOobManagementSubnet;
        /// <summary>
        /// Private IP address of HA spoke gateway.
        /// </summary>
        public readonly string HaPrivateIp;
        /// <summary>
        /// Public IP address of the HA spoke gateway.
        /// </summary>
        public readonly string HaPublicIp;
        public readonly string HaSecurityGroupId;
        /// <summary>
        /// The software version of the HA gateway.
        /// </summary>
        public readonly string HaSoftwareVersion;
        /// <summary>
        /// HA Subnet.
        /// </summary>
        public readonly string HaSubnet;
        /// <summary>
        /// HA Zone.
        /// </summary>
        public readonly string HaZone;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The image version of the gateway.
        /// </summary>
        public readonly string ImageVersion;
        /// <summary>
        /// A list of comma separated CIDRs to be advertised to on-prem as "Included CIDR List".
        /// </summary>
        public readonly string IncludedAdvertisedSpokeRoutes;
        /// <summary>
        /// Status of Insane Mode for Spoke Gateway.
        /// </summary>
        public readonly bool InsaneMode;
        /// <summary>
        /// AZ of subnet being created for Insane Mode spoke gateway.
        /// </summary>
        public readonly string InsaneModeAz;
        public readonly string LearnedCidrsApprovalMode;
        public readonly string LocalAsNumber;
        public readonly ImmutableArray<string> MonitorExcludeLists;
        /// <summary>
        /// OOB availability zone.
        /// </summary>
        public readonly string OobAvailabilityZone;
        /// <summary>
        /// OOB management subnet.
        /// </summary>
        public readonly string OobManagementSubnet;
        public readonly ImmutableArray<string> PrependAsPaths;
        /// <summary>
        /// Private IP address of the spoke gateway.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// Public IP of spoke gateway.
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// Security group used of the spoke gateway.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// Status of Single AZ HA of spoke gateway.
        /// </summary>
        public readonly bool SingleAzHa;
        /// <summary>
        /// Status of Single IP Source NAT mode of the spoke gateway.
        /// </summary>
        public readonly bool SingleIpSnat;
        /// <summary>
        /// The software version of the gateway.
        /// </summary>
        public readonly string SoftwareVersion;
        public readonly ImmutableArray<string> SpokeBgpManualAdvertiseCidrs;
        public readonly string SpotPrice;
        /// <summary>
        /// A VPC Network address range selected from one of the available network ranges.
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// Instance tag of cloud provider.
        /// </summary>
        public readonly ImmutableArray<string> TagLists;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Transit gateways attached to this spoke gateway.
        /// </summary>
        public readonly string TransitGw;
        public readonly int TunnelDetectionTime;
        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        public readonly string VpcReg;
        public readonly string Zone;

        [OutputConstructor]
        private GetAviatrixSpokeGatewayResult(
            string accountName,

            bool allocateNewEip,

            ImmutableArray<string> approvedLearnedCidrs,

            string availabilityDomain,

            string azureEipNameResourceGroup,

            bool bgpEcmp,

            int bgpHoldTime,

            int bgpPollingTime,

            string cloudInstanceId,

            int cloudType,

            string customizedSpokeVpcRoutes,

            bool disableRoutePropagation,

            string eip,

            bool enableActiveStandby,

            bool enableActiveStandbyPreemptive,

            bool enableAutoAdvertiseS2cCidrs,

            bool enableBgp,

            bool enableEncryptVolume,

            bool enableJumboFrame,

            bool enableLearnedCidrsApproval,

            bool enableMonitorGatewaySubnets,

            bool enablePrivateOob,

            bool enablePrivateVpcDefaultRoute,

            bool enableSkipPublicRouteTableUpdate,

            bool enableSpotInstance,

            bool enableVpcDnsServer,

            string faultDomain,

            string filteredSpokeVpcRoutes,

            string gwName,

            string gwSize,

            string haAvailabilityDomain,

            string haAzureEipNameResourceGroup,

            string haCloudInstanceId,

            string haEip,

            string haFaultDomain,

            string haGwName,

            string haGwSize,

            string haImageVersion,

            string haInsaneModeAz,

            string haOobAvailabilityZone,

            string haOobManagementSubnet,

            string haPrivateIp,

            string haPublicIp,

            string haSecurityGroupId,

            string haSoftwareVersion,

            string haSubnet,

            string haZone,

            string id,

            string imageVersion,

            string includedAdvertisedSpokeRoutes,

            bool insaneMode,

            string insaneModeAz,

            string learnedCidrsApprovalMode,

            string localAsNumber,

            ImmutableArray<string> monitorExcludeLists,

            string oobAvailabilityZone,

            string oobManagementSubnet,

            ImmutableArray<string> prependAsPaths,

            string privateIp,

            string publicIp,

            string securityGroupId,

            bool singleAzHa,

            bool singleIpSnat,

            string softwareVersion,

            ImmutableArray<string> spokeBgpManualAdvertiseCidrs,

            string spotPrice,

            string subnet,

            ImmutableArray<string> tagLists,

            ImmutableDictionary<string, string> tags,

            string transitGw,

            int tunnelDetectionTime,

            string vpcId,

            string vpcReg,

            string zone)
        {
            AccountName = accountName;
            AllocateNewEip = allocateNewEip;
            ApprovedLearnedCidrs = approvedLearnedCidrs;
            AvailabilityDomain = availabilityDomain;
            AzureEipNameResourceGroup = azureEipNameResourceGroup;
            BgpEcmp = bgpEcmp;
            BgpHoldTime = bgpHoldTime;
            BgpPollingTime = bgpPollingTime;
            CloudInstanceId = cloudInstanceId;
            CloudType = cloudType;
            CustomizedSpokeVpcRoutes = customizedSpokeVpcRoutes;
            DisableRoutePropagation = disableRoutePropagation;
            Eip = eip;
            EnableActiveStandby = enableActiveStandby;
            EnableActiveStandbyPreemptive = enableActiveStandbyPreemptive;
            EnableAutoAdvertiseS2cCidrs = enableAutoAdvertiseS2cCidrs;
            EnableBgp = enableBgp;
            EnableEncryptVolume = enableEncryptVolume;
            EnableJumboFrame = enableJumboFrame;
            EnableLearnedCidrsApproval = enableLearnedCidrsApproval;
            EnableMonitorGatewaySubnets = enableMonitorGatewaySubnets;
            EnablePrivateOob = enablePrivateOob;
            EnablePrivateVpcDefaultRoute = enablePrivateVpcDefaultRoute;
            EnableSkipPublicRouteTableUpdate = enableSkipPublicRouteTableUpdate;
            EnableSpotInstance = enableSpotInstance;
            EnableVpcDnsServer = enableVpcDnsServer;
            FaultDomain = faultDomain;
            FilteredSpokeVpcRoutes = filteredSpokeVpcRoutes;
            GwName = gwName;
            GwSize = gwSize;
            HaAvailabilityDomain = haAvailabilityDomain;
            HaAzureEipNameResourceGroup = haAzureEipNameResourceGroup;
            HaCloudInstanceId = haCloudInstanceId;
            HaEip = haEip;
            HaFaultDomain = haFaultDomain;
            HaGwName = haGwName;
            HaGwSize = haGwSize;
            HaImageVersion = haImageVersion;
            HaInsaneModeAz = haInsaneModeAz;
            HaOobAvailabilityZone = haOobAvailabilityZone;
            HaOobManagementSubnet = haOobManagementSubnet;
            HaPrivateIp = haPrivateIp;
            HaPublicIp = haPublicIp;
            HaSecurityGroupId = haSecurityGroupId;
            HaSoftwareVersion = haSoftwareVersion;
            HaSubnet = haSubnet;
            HaZone = haZone;
            Id = id;
            ImageVersion = imageVersion;
            IncludedAdvertisedSpokeRoutes = includedAdvertisedSpokeRoutes;
            InsaneMode = insaneMode;
            InsaneModeAz = insaneModeAz;
            LearnedCidrsApprovalMode = learnedCidrsApprovalMode;
            LocalAsNumber = localAsNumber;
            MonitorExcludeLists = monitorExcludeLists;
            OobAvailabilityZone = oobAvailabilityZone;
            OobManagementSubnet = oobManagementSubnet;
            PrependAsPaths = prependAsPaths;
            PrivateIp = privateIp;
            PublicIp = publicIp;
            SecurityGroupId = securityGroupId;
            SingleAzHa = singleAzHa;
            SingleIpSnat = singleIpSnat;
            SoftwareVersion = softwareVersion;
            SpokeBgpManualAdvertiseCidrs = spokeBgpManualAdvertiseCidrs;
            SpotPrice = spotPrice;
            Subnet = subnet;
            TagLists = tagLists;
            Tags = tags;
            TransitGw = transitGw;
            TunnelDetectionTime = tunnelDetectionTime;
            VpcId = vpcId;
            VpcReg = vpcReg;
            Zone = zone;
        }
    }
}

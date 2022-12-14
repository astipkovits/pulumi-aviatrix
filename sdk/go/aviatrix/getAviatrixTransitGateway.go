// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_transit_gateway** data source provides details about a specific transit gateway created by the Aviatrix Controller.
//
// This data source can prove useful when a module accepts a transit gateway's detail as an input variable. For example, requiring the transit gateway's name for a spoke gateway's attachment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/astipkovits/pulumi-aviatrix/sdk/go/aviatrix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = aviatrix.LookupAviatrixTransitGateway(ctx, &GetAviatrixTransitGatewayArgs{
//				GwName: "gatewayname",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAviatrixTransitGateway(ctx *pulumi.Context, args *LookupAviatrixTransitGatewayArgs, opts ...pulumi.InvokeOption) (*LookupAviatrixTransitGatewayResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAviatrixTransitGatewayResult
	err := ctx.Invoke("aviatrix:index/getAviatrixTransitGateway:getAviatrixTransitGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAviatrixTransitGateway.
type LookupAviatrixTransitGatewayArgs struct {
	// Transit gateway name.
	GwName string `pulumi:"gwName"`
}

// A collection of values returned by getAviatrixTransitGateway.
type LookupAviatrixTransitGatewayResult struct {
	// Aviatrix account name.
	AccountName string `pulumi:"accountName"`
	// When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.
	AllocateNewEip bool `pulumi:"allocateNewEip"`
	// Approved learned CIDRs.
	ApprovedLearnedCidrs []string `pulumi:"approvedLearnedCidrs"`
	// Availability domain for OCI.
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// The name of the public IP address and its resource group in Azure to assign to this Transit Gateway.
	AzureEipNameResourceGroup string `pulumi:"azureEipNameResourceGroup"`
	// Status of Equal Cost Multi Path (ECMP) routing for the next hop.
	BgpEcmp bool `pulumi:"bgpEcmp"`
	// BGP Hold Time.
	BgpHoldTime int `pulumi:"bgpHoldTime"`
	// Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer.
	BgpLanInterfaces []GetAviatrixTransitGatewayBgpLanInterface `pulumi:"bgpLanInterfaces"`
	// List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP and Azure.
	BgpLanIpLists []string `pulumi:"bgpLanIpLists"`
	// Intended CIDR list to advertise to VGW.
	BgpManualSpokeAdvertiseCidrs string `pulumi:"bgpManualSpokeAdvertiseCidrs"`
	// BGP route polling time.
	BgpPollingTime string `pulumi:"bgpPollingTime"`
	// Instance ID of the transit gateway.
	CloudInstanceId string `pulumi:"cloudInstanceId"`
	// Type of cloud service provider.
	// * `connectedTransit"` -  Status of Connected Transit of transit gateway.
	CloudType        int  `pulumi:"cloudType"`
	ConnectedTransit bool `pulumi:"connectedTransit"`
	// A list of comma separated CIDRs to be customized for the spoke VPC routes.
	CustomizedSpokeVpcRoutes string `pulumi:"customizedSpokeVpcRoutes"`
	// A list of CIDRs to be customized for the transit VPC routes.
	CustomizedTransitVpcRoutes []string `pulumi:"customizedTransitVpcRoutes"`
	// The EIP address of the Transit Gateway.
	Eip string `pulumi:"eip"`
	// Status of Active-Standby Mode.
	EnableActiveStandby bool `pulumi:"enableActiveStandby"`
	// Status of Preemptive Mode for Active-Standby.
	EnableActiveStandbyPreemptive bool `pulumi:"enableActiveStandbyPreemptive"`
	// Status of Advertise Transit VPC network CIDR of the transit gateway.
	EnableAdvertiseTransitCidr bool `pulumi:"enableAdvertiseTransitCidr"`
	// Status of BGP over LAN functionality.
	EnableBgpOverLan bool `pulumi:"enableBgpOverLan"`
	// Status of Egress Transit FireNet being enabled on the gateway.
	EnableEgressTransitFirenet bool `pulumi:"enableEgressTransitFirenet"`
	// Status of Encrypt Gateway EBS Volume of the transit gateway.
	EnableEncryptVolume bool `pulumi:"enableEncryptVolume"`
	// Status of FireNet Interfaces of the transit gateway.
	EnableFirenet bool `pulumi:"enableFirenet"`
	// Status of AWS Gateway Load Balancer.
	EnableGatewayLoadBalancer bool `pulumi:"enableGatewayLoadBalancer"`
	// Sign of readiness for TGW connection.
	EnableHybridConnection bool `pulumi:"enableHybridConnection"`
	// Status of jumbo frame support.
	EnableJumboFrame bool `pulumi:"enableJumboFrame"`
	// Status of Encrypted Transit Approval for transit gateway.
	EnableLearnedCidrsApproval bool `pulumi:"enableLearnedCidrsApproval"`
	// Status of monitor gateway subnets.
	EnableMonitorGatewaySubnets bool `pulumi:"enableMonitorGatewaySubnets"`
	// Status of multi-tier transit mode on transit gateway.
	EnableMultiTierTransit bool `pulumi:"enableMultiTierTransit"`
	// Status of private OOB for the transit gateway.
	EnablePrivateOob bool `pulumi:"enablePrivateOob"`
	// Status of segmentation.
	EnableSegmentation bool `pulumi:"enableSegmentation"`
	// Status of spot instance.
	EnableSpotInstance bool `pulumi:"enableSpotInstance"`
	// Status of Transit FireNet Interfaces of the transit gateway.
	EnableTransitFirenet bool `pulumi:"enableTransitFirenet"`
	// Status of transit summarize CIDR to TGW.
	EnableTransitSummarizeCidrToTgw bool `pulumi:"enableTransitSummarizeCidrToTgw"`
	// Status of Vpc Dns Server of the transit Gateway.
	EnableVpcDnsServer bool `pulumi:"enableVpcDnsServer"`
	// A list of comma separated CIDRs to be advertised to on-prem as "Excluded CIDR List".
	ExcludedAdvertisedSpokeRoutes string `pulumi:"excludedAdvertisedSpokeRoutes"`
	// Fault domain for OCI.
	FaultDomain string `pulumi:"faultDomain"`
	// A list of comma separated CIDRs to be filtered from the spoke VPC route table.
	FilteredSpokeVpcRoutes string `pulumi:"filteredSpokeVpcRoutes"`
	// Aviatrix transit gateway name.
	GwName string `pulumi:"gwName"`
	// Size of transit gateway instance.
	GwSize string `pulumi:"gwSize"`
	// HA gateway availability domain for OCI.
	HaAvailabilityDomain string `pulumi:"haAvailabilityDomain"`
	// The name of the public IP address and its resource group in Azure to assign to the HA Transit Gateway.
	HaAzureEipNameResourceGroup string `pulumi:"haAzureEipNameResourceGroup"`
	// Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP HA Transit.
	HaBgpLanInterfaces []GetAviatrixTransitGatewayHaBgpLanInterface `pulumi:"haBgpLanInterfaces"`
	// List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP and Azure.
	HaBgpLanIpLists []string `pulumi:"haBgpLanIpLists"`
	// Cloud instance ID of HA transit gateway.
	HaCloudInstanceId string `pulumi:"haCloudInstanceId"`
	// The EIP address of the HA Transit Gateway.
	HaEip string `pulumi:"haEip"`
	// HA gateway fault domain for OCI.
	HaFaultDomain string `pulumi:"haFaultDomain"`
	// Aviatrix transit gateway unique name of HA transit gateway.
	// * `haGwSize"` - HA Gateway Size.
	HaGwName string `pulumi:"haGwName"`
	HaGwSize string `pulumi:"haGwSize"`
	// The image version of the HA gateway.
	HaImageVersion string `pulumi:"haImageVersion"`
	// AZ of subnet being created for Insane Mode Transit HA Gateway.
	HaInsaneModeAz string `pulumi:"haInsaneModeAz"`
	// Transit gateway lan interface cidr for the HA gateway.
	HaLanInterfaceCidr string `pulumi:"haLanInterfaceCidr"`
	// HA OOB availability zone.
	HaOobAvailabilityZone string `pulumi:"haOobAvailabilityZone"`
	// HA OOB management subnet.
	HaOobManagementSubnet string `pulumi:"haOobManagementSubnet"`
	// Private IP address that assigned to the HA Transit Gateway.
	HaPrivateIp string `pulumi:"haPrivateIp"`
	// Public IP address that assigned to the HA Transit Gateway.
	HaPublicIp string `pulumi:"haPublicIp"`
	// HA security group used for the transit gateway.
	HaSecurityGroupId string `pulumi:"haSecurityGroupId"`
	// The software version of the HA gateway.
	HaSoftwareVersion string `pulumi:"haSoftwareVersion"`
	// HA Subnet.
	HaSubnet string `pulumi:"haSubnet"`
	// HA Zone.
	HaZone string `pulumi:"haZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The image version of the gateway.
	ImageVersion string `pulumi:"imageVersion"`
	// Status of Insane Mode of the transit gateway.
	InsaneMode bool `pulumi:"insaneMode"`
	// AZ of subnet being created for Insane Mode transit gateway.
	InsaneModeAz string `pulumi:"insaneModeAz"`
	// Transit gateway lan interface cidr.
	LanInterfaceCidr string `pulumi:"lanInterfaceCidr"`
	// LAN Private Subnet for GCP Transit FireNet.
	LanPrivateSubnet string `pulumi:"lanPrivateSubnet"`
	// LAN VPC ID for GCP Transit FireNet.
	LanVpcId string `pulumi:"lanVpcId"`
	// Learned CIDRs approval mode.
	LearnedCidrsApprovalMode string `pulumi:"learnedCidrsApprovalMode"`
	// Local ASN number.
	LocalAsNumber string `pulumi:"localAsNumber"`
	// A set of monitored instance IDs.
	MonitorExcludeLists []string `pulumi:"monitorExcludeLists"`
	// OOB availability zone.
	OobAvailabilityZone string `pulumi:"oobAvailabilityZone"`
	// OOB management subnet.
	OobManagementSubnet string `pulumi:"oobManagementSubnet"`
	// List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.
	PrependAsPaths []string `pulumi:"prependAsPaths"`
	// Private IP address of the transit gateway created.
	PrivateIp string `pulumi:"privateIp"`
	// Public IP address of the Transit Gateway created.
	PublicIp string `pulumi:"publicIp"`
	// Security group used for the transit gateway.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Status of Single AZ HA of transit gateway.
	SingleAzHa bool `pulumi:"singleAzHa"`
	// Status of Single IP Source Nat mode of the transit gateway.
	SingleIpSnat bool `pulumi:"singleIpSnat"`
	// The software version of the gateway.
	SoftwareVersion string `pulumi:"softwareVersion"`
	// Price for spot instance.
	SpotPrice string `pulumi:"spotPrice"`
	// Subnet Info.
	Subnet string `pulumi:"subnet"`
	// Instance tag of cloud provider.
	TagLists            []string          `pulumi:"tagLists"`
	Tags                map[string]string `pulumi:"tags"`
	TunnelDetectionTime int               `pulumi:"tunnelDetectionTime"`
	// VPC-ID of GCP cloud provider.
	VpcId string `pulumi:"vpcId"`
	// Region of cloud provider.
	VpcReg string `pulumi:"vpcReg"`
	// Availability Zone for Azure.
	Zone string `pulumi:"zone"`
}

func LookupAviatrixTransitGatewayOutput(ctx *pulumi.Context, args LookupAviatrixTransitGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupAviatrixTransitGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAviatrixTransitGatewayResult, error) {
			args := v.(LookupAviatrixTransitGatewayArgs)
			r, err := LookupAviatrixTransitGateway(ctx, &args, opts...)
			var s LookupAviatrixTransitGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAviatrixTransitGatewayResultOutput)
}

// A collection of arguments for invoking getAviatrixTransitGateway.
type LookupAviatrixTransitGatewayOutputArgs struct {
	// Transit gateway name.
	GwName pulumi.StringInput `pulumi:"gwName"`
}

func (LookupAviatrixTransitGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixTransitGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getAviatrixTransitGateway.
type LookupAviatrixTransitGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupAviatrixTransitGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixTransitGatewayResult)(nil)).Elem()
}

func (o LookupAviatrixTransitGatewayResultOutput) ToLookupAviatrixTransitGatewayResultOutput() LookupAviatrixTransitGatewayResultOutput {
	return o
}

func (o LookupAviatrixTransitGatewayResultOutput) ToLookupAviatrixTransitGatewayResultOutputWithContext(ctx context.Context) LookupAviatrixTransitGatewayResultOutput {
	return o
}

// Aviatrix account name.
func (o LookupAviatrixTransitGatewayResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.
func (o LookupAviatrixTransitGatewayResultOutput) AllocateNewEip() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.AllocateNewEip }).(pulumi.BoolOutput)
}

// Approved learned CIDRs.
func (o LookupAviatrixTransitGatewayResultOutput) ApprovedLearnedCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.ApprovedLearnedCidrs }).(pulumi.StringArrayOutput)
}

// Availability domain for OCI.
func (o LookupAviatrixTransitGatewayResultOutput) AvailabilityDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.AvailabilityDomain }).(pulumi.StringOutput)
}

// The name of the public IP address and its resource group in Azure to assign to this Transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) AzureEipNameResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.AzureEipNameResourceGroup }).(pulumi.StringOutput)
}

// Status of Equal Cost Multi Path (ECMP) routing for the next hop.
func (o LookupAviatrixTransitGatewayResultOutput) BgpEcmp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.BgpEcmp }).(pulumi.BoolOutput)
}

// BGP Hold Time.
func (o LookupAviatrixTransitGatewayResultOutput) BgpHoldTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) int { return v.BgpHoldTime }).(pulumi.IntOutput)
}

// Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer.
func (o LookupAviatrixTransitGatewayResultOutput) BgpLanInterfaces() GetAviatrixTransitGatewayBgpLanInterfaceArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []GetAviatrixTransitGatewayBgpLanInterface {
		return v.BgpLanInterfaces
	}).(GetAviatrixTransitGatewayBgpLanInterfaceArrayOutput)
}

// List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP and Azure.
func (o LookupAviatrixTransitGatewayResultOutput) BgpLanIpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.BgpLanIpLists }).(pulumi.StringArrayOutput)
}

// Intended CIDR list to advertise to VGW.
func (o LookupAviatrixTransitGatewayResultOutput) BgpManualSpokeAdvertiseCidrs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.BgpManualSpokeAdvertiseCidrs }).(pulumi.StringOutput)
}

// BGP route polling time.
func (o LookupAviatrixTransitGatewayResultOutput) BgpPollingTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.BgpPollingTime }).(pulumi.StringOutput)
}

// Instance ID of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) CloudInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.CloudInstanceId }).(pulumi.StringOutput)
}

// Type of cloud service provider.
// * `connectedTransit"` -  Status of Connected Transit of transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) int { return v.CloudType }).(pulumi.IntOutput)
}

func (o LookupAviatrixTransitGatewayResultOutput) ConnectedTransit() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.ConnectedTransit }).(pulumi.BoolOutput)
}

// A list of comma separated CIDRs to be customized for the spoke VPC routes.
func (o LookupAviatrixTransitGatewayResultOutput) CustomizedSpokeVpcRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.CustomizedSpokeVpcRoutes }).(pulumi.StringOutput)
}

// A list of CIDRs to be customized for the transit VPC routes.
func (o LookupAviatrixTransitGatewayResultOutput) CustomizedTransitVpcRoutes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.CustomizedTransitVpcRoutes }).(pulumi.StringArrayOutput)
}

// The EIP address of the Transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) Eip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.Eip }).(pulumi.StringOutput)
}

// Status of Active-Standby Mode.
func (o LookupAviatrixTransitGatewayResultOutput) EnableActiveStandby() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableActiveStandby }).(pulumi.BoolOutput)
}

// Status of Preemptive Mode for Active-Standby.
func (o LookupAviatrixTransitGatewayResultOutput) EnableActiveStandbyPreemptive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableActiveStandbyPreemptive }).(pulumi.BoolOutput)
}

// Status of Advertise Transit VPC network CIDR of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableAdvertiseTransitCidr() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableAdvertiseTransitCidr }).(pulumi.BoolOutput)
}

// Status of BGP over LAN functionality.
func (o LookupAviatrixTransitGatewayResultOutput) EnableBgpOverLan() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableBgpOverLan }).(pulumi.BoolOutput)
}

// Status of Egress Transit FireNet being enabled on the gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableEgressTransitFirenet() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableEgressTransitFirenet }).(pulumi.BoolOutput)
}

// Status of Encrypt Gateway EBS Volume of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableEncryptVolume() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableEncryptVolume }).(pulumi.BoolOutput)
}

// Status of FireNet Interfaces of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableFirenet() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableFirenet }).(pulumi.BoolOutput)
}

// Status of AWS Gateway Load Balancer.
func (o LookupAviatrixTransitGatewayResultOutput) EnableGatewayLoadBalancer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableGatewayLoadBalancer }).(pulumi.BoolOutput)
}

// Sign of readiness for TGW connection.
func (o LookupAviatrixTransitGatewayResultOutput) EnableHybridConnection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableHybridConnection }).(pulumi.BoolOutput)
}

// Status of jumbo frame support.
func (o LookupAviatrixTransitGatewayResultOutput) EnableJumboFrame() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableJumboFrame }).(pulumi.BoolOutput)
}

// Status of Encrypted Transit Approval for transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableLearnedCidrsApproval() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableLearnedCidrsApproval }).(pulumi.BoolOutput)
}

// Status of monitor gateway subnets.
func (o LookupAviatrixTransitGatewayResultOutput) EnableMonitorGatewaySubnets() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableMonitorGatewaySubnets }).(pulumi.BoolOutput)
}

// Status of multi-tier transit mode on transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableMultiTierTransit() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableMultiTierTransit }).(pulumi.BoolOutput)
}

// Status of private OOB for the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnablePrivateOob() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnablePrivateOob }).(pulumi.BoolOutput)
}

// Status of segmentation.
func (o LookupAviatrixTransitGatewayResultOutput) EnableSegmentation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableSegmentation }).(pulumi.BoolOutput)
}

// Status of spot instance.
func (o LookupAviatrixTransitGatewayResultOutput) EnableSpotInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableSpotInstance }).(pulumi.BoolOutput)
}

// Status of Transit FireNet Interfaces of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableTransitFirenet() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableTransitFirenet }).(pulumi.BoolOutput)
}

// Status of transit summarize CIDR to TGW.
func (o LookupAviatrixTransitGatewayResultOutput) EnableTransitSummarizeCidrToTgw() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableTransitSummarizeCidrToTgw }).(pulumi.BoolOutput)
}

// Status of Vpc Dns Server of the transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) EnableVpcDnsServer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.EnableVpcDnsServer }).(pulumi.BoolOutput)
}

// A list of comma separated CIDRs to be advertised to on-prem as "Excluded CIDR List".
func (o LookupAviatrixTransitGatewayResultOutput) ExcludedAdvertisedSpokeRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.ExcludedAdvertisedSpokeRoutes }).(pulumi.StringOutput)
}

// Fault domain for OCI.
func (o LookupAviatrixTransitGatewayResultOutput) FaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.FaultDomain }).(pulumi.StringOutput)
}

// A list of comma separated CIDRs to be filtered from the spoke VPC route table.
func (o LookupAviatrixTransitGatewayResultOutput) FilteredSpokeVpcRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.FilteredSpokeVpcRoutes }).(pulumi.StringOutput)
}

// Aviatrix transit gateway name.
func (o LookupAviatrixTransitGatewayResultOutput) GwName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.GwName }).(pulumi.StringOutput)
}

// Size of transit gateway instance.
func (o LookupAviatrixTransitGatewayResultOutput) GwSize() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.GwSize }).(pulumi.StringOutput)
}

// HA gateway availability domain for OCI.
func (o LookupAviatrixTransitGatewayResultOutput) HaAvailabilityDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaAvailabilityDomain }).(pulumi.StringOutput)
}

// The name of the public IP address and its resource group in Azure to assign to the HA Transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaAzureEipNameResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaAzureEipNameResourceGroup }).(pulumi.StringOutput)
}

// Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP HA Transit.
func (o LookupAviatrixTransitGatewayResultOutput) HaBgpLanInterfaces() GetAviatrixTransitGatewayHaBgpLanInterfaceArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []GetAviatrixTransitGatewayHaBgpLanInterface {
		return v.HaBgpLanInterfaces
	}).(GetAviatrixTransitGatewayHaBgpLanInterfaceArrayOutput)
}

// List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP and Azure.
func (o LookupAviatrixTransitGatewayResultOutput) HaBgpLanIpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.HaBgpLanIpLists }).(pulumi.StringArrayOutput)
}

// Cloud instance ID of HA transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaCloudInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaCloudInstanceId }).(pulumi.StringOutput)
}

// The EIP address of the HA Transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaEip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaEip }).(pulumi.StringOutput)
}

// HA gateway fault domain for OCI.
func (o LookupAviatrixTransitGatewayResultOutput) HaFaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaFaultDomain }).(pulumi.StringOutput)
}

// Aviatrix transit gateway unique name of HA transit gateway.
// * `haGwSize"` - HA Gateway Size.
func (o LookupAviatrixTransitGatewayResultOutput) HaGwName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaGwName }).(pulumi.StringOutput)
}

func (o LookupAviatrixTransitGatewayResultOutput) HaGwSize() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaGwSize }).(pulumi.StringOutput)
}

// The image version of the HA gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaImageVersion }).(pulumi.StringOutput)
}

// AZ of subnet being created for Insane Mode Transit HA Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaInsaneModeAz() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaInsaneModeAz }).(pulumi.StringOutput)
}

// Transit gateway lan interface cidr for the HA gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaLanInterfaceCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaLanInterfaceCidr }).(pulumi.StringOutput)
}

// HA OOB availability zone.
func (o LookupAviatrixTransitGatewayResultOutput) HaOobAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaOobAvailabilityZone }).(pulumi.StringOutput)
}

// HA OOB management subnet.
func (o LookupAviatrixTransitGatewayResultOutput) HaOobManagementSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaOobManagementSubnet }).(pulumi.StringOutput)
}

// Private IP address that assigned to the HA Transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaPrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaPrivateIp }).(pulumi.StringOutput)
}

// Public IP address that assigned to the HA Transit Gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaPublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaPublicIp }).(pulumi.StringOutput)
}

// HA security group used for the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaSecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaSecurityGroupId }).(pulumi.StringOutput)
}

// The software version of the HA gateway.
func (o LookupAviatrixTransitGatewayResultOutput) HaSoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaSoftwareVersion }).(pulumi.StringOutput)
}

// HA Subnet.
func (o LookupAviatrixTransitGatewayResultOutput) HaSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaSubnet }).(pulumi.StringOutput)
}

// HA Zone.
func (o LookupAviatrixTransitGatewayResultOutput) HaZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.HaZone }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAviatrixTransitGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// The image version of the gateway.
func (o LookupAviatrixTransitGatewayResultOutput) ImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.ImageVersion }).(pulumi.StringOutput)
}

// Status of Insane Mode of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) InsaneMode() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.InsaneMode }).(pulumi.BoolOutput)
}

// AZ of subnet being created for Insane Mode transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) InsaneModeAz() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.InsaneModeAz }).(pulumi.StringOutput)
}

// Transit gateway lan interface cidr.
func (o LookupAviatrixTransitGatewayResultOutput) LanInterfaceCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.LanInterfaceCidr }).(pulumi.StringOutput)
}

// LAN Private Subnet for GCP Transit FireNet.
func (o LookupAviatrixTransitGatewayResultOutput) LanPrivateSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.LanPrivateSubnet }).(pulumi.StringOutput)
}

// LAN VPC ID for GCP Transit FireNet.
func (o LookupAviatrixTransitGatewayResultOutput) LanVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.LanVpcId }).(pulumi.StringOutput)
}

// Learned CIDRs approval mode.
func (o LookupAviatrixTransitGatewayResultOutput) LearnedCidrsApprovalMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.LearnedCidrsApprovalMode }).(pulumi.StringOutput)
}

// Local ASN number.
func (o LookupAviatrixTransitGatewayResultOutput) LocalAsNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.LocalAsNumber }).(pulumi.StringOutput)
}

// A set of monitored instance IDs.
func (o LookupAviatrixTransitGatewayResultOutput) MonitorExcludeLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.MonitorExcludeLists }).(pulumi.StringArrayOutput)
}

// OOB availability zone.
func (o LookupAviatrixTransitGatewayResultOutput) OobAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.OobAvailabilityZone }).(pulumi.StringOutput)
}

// OOB management subnet.
func (o LookupAviatrixTransitGatewayResultOutput) OobManagementSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.OobManagementSubnet }).(pulumi.StringOutput)
}

// List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.
func (o LookupAviatrixTransitGatewayResultOutput) PrependAsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.PrependAsPaths }).(pulumi.StringArrayOutput)
}

// Private IP address of the transit gateway created.
func (o LookupAviatrixTransitGatewayResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// Public IP address of the Transit Gateway created.
func (o LookupAviatrixTransitGatewayResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

// Security group used for the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Status of Single AZ HA of transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) SingleAzHa() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.SingleAzHa }).(pulumi.BoolOutput)
}

// Status of Single IP Source Nat mode of the transit gateway.
func (o LookupAviatrixTransitGatewayResultOutput) SingleIpSnat() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) bool { return v.SingleIpSnat }).(pulumi.BoolOutput)
}

// The software version of the gateway.
func (o LookupAviatrixTransitGatewayResultOutput) SoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.SoftwareVersion }).(pulumi.StringOutput)
}

// Price for spot instance.
func (o LookupAviatrixTransitGatewayResultOutput) SpotPrice() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.SpotPrice }).(pulumi.StringOutput)
}

// Subnet Info.
func (o LookupAviatrixTransitGatewayResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// Instance tag of cloud provider.
func (o LookupAviatrixTransitGatewayResultOutput) TagLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) []string { return v.TagLists }).(pulumi.StringArrayOutput)
}

func (o LookupAviatrixTransitGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAviatrixTransitGatewayResultOutput) TunnelDetectionTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) int { return v.TunnelDetectionTime }).(pulumi.IntOutput)
}

// VPC-ID of GCP cloud provider.
func (o LookupAviatrixTransitGatewayResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.VpcId }).(pulumi.StringOutput)
}

// Region of cloud provider.
func (o LookupAviatrixTransitGatewayResultOutput) VpcReg() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.VpcReg }).(pulumi.StringOutput)
}

// Availability Zone for Azure.
func (o LookupAviatrixTransitGatewayResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixTransitGatewayResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAviatrixTransitGatewayResultOutput{})
}

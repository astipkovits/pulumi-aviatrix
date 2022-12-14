// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_vpc** data source provides details about a specific VPC created by the Aviatrix Controller.
//
// This data source can prove useful when a module accepts any form of VPC detail as an input variable. For example, requiring a subnet CIDR specification when creating a gateway.
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
//			_, err = aviatrix.LookupAviatrixVpc(ctx, &GetAviatrixVpcArgs{
//				Name: "vpc-test",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAviatrixVpc(ctx *pulumi.Context, args *LookupAviatrixVpcArgs, opts ...pulumi.InvokeOption) (*LookupAviatrixVpcResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAviatrixVpcResult
	err := ctx.Invoke("aviatrix:index/getAviatrixVpc:getAviatrixVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAviatrixVpc.
type LookupAviatrixVpcArgs struct {
	// Name of the Aviatrix VPC.
	Name string `pulumi:"name"`
	// Filters the `routeTables` list to contain only public or private route tables. Valid values are 'private' or 'public'. If not set `routeTables` is not filtered.
	RouteTablesFilter *string `pulumi:"routeTablesFilter"`
}

// A collection of values returned by getAviatrixVpc.
type LookupAviatrixVpcResult struct {
	// Account name of the VPC created.
	AccountName string `pulumi:"accountName"`
	// List of OCI availability domains.
	AvailabilityDomains []string `pulumi:"availabilityDomains"`
	// Switch if the VPC created is an Aviatrix FireNet VPC or not.
	AviatrixFirenetVpc bool `pulumi:"aviatrixFirenetVpc"`
	// Switch if the VPC created is an Aviatrix Transit VPC or not.
	AviatrixTransitVpc bool `pulumi:"aviatrixTransitVpc"`
	// Azure vnet resource ID.
	AzureVnetResourceId string `pulumi:"azureVnetResourceId"`
	// Private subnet CIDR.
	Cidr string `pulumi:"cidr"`
	// Type of cloud service provider.
	CloudType int `pulumi:"cloudType"`
	// List of OCI fault domains.
	FaultDomains []string `pulumi:"faultDomains"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Private subnet name.
	Name string `pulumi:"name"`
	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider.
	NumOfSubnetPairs int `pulumi:"numOfSubnetPairs"`
	// List of private subnet of the VPC(AWS, Azure) created.
	PrivateSubnets []GetAviatrixVpcPrivateSubnet `pulumi:"privateSubnets"`
	// List of public subnet of the VPC(AWS, Azure) created.
	PublicSubnets []GetAviatrixVpcPublicSubnet `pulumi:"publicSubnets"`
	// Region of the VPC created.
	Region string `pulumi:"region"`
	// Resource group of the Azure VPC created.
	ResourceGroup string `pulumi:"resourceGroup"`
	// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure vpc.
	RouteTables       []string `pulumi:"routeTables"`
	RouteTablesFilter *string  `pulumi:"routeTablesFilter"`
	// Subnet size. Only supported for AWS, Azure provider.
	SubnetSize int `pulumi:"subnetSize"`
	// List of subnet of the VPC created.
	Subnets []GetAviatrixVpcSubnet `pulumi:"subnets"`
	// ID of the VPC created.
	VpcId string `pulumi:"vpcId"`
}

func LookupAviatrixVpcOutput(ctx *pulumi.Context, args LookupAviatrixVpcOutputArgs, opts ...pulumi.InvokeOption) LookupAviatrixVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAviatrixVpcResult, error) {
			args := v.(LookupAviatrixVpcArgs)
			r, err := LookupAviatrixVpc(ctx, &args, opts...)
			var s LookupAviatrixVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAviatrixVpcResultOutput)
}

// A collection of arguments for invoking getAviatrixVpc.
type LookupAviatrixVpcOutputArgs struct {
	// Name of the Aviatrix VPC.
	Name pulumi.StringInput `pulumi:"name"`
	// Filters the `routeTables` list to contain only public or private route tables. Valid values are 'private' or 'public'. If not set `routeTables` is not filtered.
	RouteTablesFilter pulumi.StringPtrInput `pulumi:"routeTablesFilter"`
}

func (LookupAviatrixVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixVpcArgs)(nil)).Elem()
}

// A collection of values returned by getAviatrixVpc.
type LookupAviatrixVpcResultOutput struct{ *pulumi.OutputState }

func (LookupAviatrixVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixVpcResult)(nil)).Elem()
}

func (o LookupAviatrixVpcResultOutput) ToLookupAviatrixVpcResultOutput() LookupAviatrixVpcResultOutput {
	return o
}

func (o LookupAviatrixVpcResultOutput) ToLookupAviatrixVpcResultOutputWithContext(ctx context.Context) LookupAviatrixVpcResultOutput {
	return o
}

// Account name of the VPC created.
func (o LookupAviatrixVpcResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// List of OCI availability domains.
func (o LookupAviatrixVpcResultOutput) AvailabilityDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []string { return v.AvailabilityDomains }).(pulumi.StringArrayOutput)
}

// Switch if the VPC created is an Aviatrix FireNet VPC or not.
func (o LookupAviatrixVpcResultOutput) AviatrixFirenetVpc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) bool { return v.AviatrixFirenetVpc }).(pulumi.BoolOutput)
}

// Switch if the VPC created is an Aviatrix Transit VPC or not.
func (o LookupAviatrixVpcResultOutput) AviatrixTransitVpc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) bool { return v.AviatrixTransitVpc }).(pulumi.BoolOutput)
}

// Azure vnet resource ID.
func (o LookupAviatrixVpcResultOutput) AzureVnetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.AzureVnetResourceId }).(pulumi.StringOutput)
}

// Private subnet CIDR.
func (o LookupAviatrixVpcResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// Type of cloud service provider.
func (o LookupAviatrixVpcResultOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) int { return v.CloudType }).(pulumi.IntOutput)
}

// List of OCI fault domains.
func (o LookupAviatrixVpcResultOutput) FaultDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []string { return v.FaultDomains }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAviatrixVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

// Private subnet name.
func (o LookupAviatrixVpcResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider.
func (o LookupAviatrixVpcResultOutput) NumOfSubnetPairs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) int { return v.NumOfSubnetPairs }).(pulumi.IntOutput)
}

// List of private subnet of the VPC(AWS, Azure) created.
func (o LookupAviatrixVpcResultOutput) PrivateSubnets() GetAviatrixVpcPrivateSubnetArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []GetAviatrixVpcPrivateSubnet { return v.PrivateSubnets }).(GetAviatrixVpcPrivateSubnetArrayOutput)
}

// List of public subnet of the VPC(AWS, Azure) created.
func (o LookupAviatrixVpcResultOutput) PublicSubnets() GetAviatrixVpcPublicSubnetArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []GetAviatrixVpcPublicSubnet { return v.PublicSubnets }).(GetAviatrixVpcPublicSubnetArrayOutput)
}

// Region of the VPC created.
func (o LookupAviatrixVpcResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.Region }).(pulumi.StringOutput)
}

// Resource group of the Azure VPC created.
func (o LookupAviatrixVpcResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure vpc.
func (o LookupAviatrixVpcResultOutput) RouteTables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []string { return v.RouteTables }).(pulumi.StringArrayOutput)
}

func (o LookupAviatrixVpcResultOutput) RouteTablesFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) *string { return v.RouteTablesFilter }).(pulumi.StringPtrOutput)
}

// Subnet size. Only supported for AWS, Azure provider.
func (o LookupAviatrixVpcResultOutput) SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) int { return v.SubnetSize }).(pulumi.IntOutput)
}

// List of subnet of the VPC created.
func (o LookupAviatrixVpcResultOutput) Subnets() GetAviatrixVpcSubnetArrayOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) []GetAviatrixVpcSubnet { return v.Subnets }).(GetAviatrixVpcSubnetArrayOutput)
}

// ID of the VPC created.
func (o LookupAviatrixVpcResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixVpcResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAviatrixVpcResultOutput{})
}

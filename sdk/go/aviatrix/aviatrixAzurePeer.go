// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_azure_peer** resource allows the creation and management of the Aviatrix-created peerings between Azure VNets.
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
//			_, err := aviatrix.NewAviatrixAzurePeer(ctx, "testAzurepeer", &aviatrix.AviatrixAzurePeerArgs{
//				AccountName1:           pulumi.String("test1-account"),
//				AccountName2:           pulumi.String("test2-account"),
//				VnetNameResourceGroup1: pulumi.String("Foo_VNet1:Bar_RG1:GUID1"),
//				VnetNameResourceGroup2: pulumi.String("Foo_VNet2:Bar_RG2:GUID2"),
//				VnetReg1:               pulumi.String("Central US"),
//				VnetReg2:               pulumi.String("East US"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// **azure_peer** can be imported using the `vnet_name_resource_group1` and `vnet_name_resource_group2`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixAzurePeer:AviatrixAzurePeer test vnet_name_resource_group1~vnet_name_resource_group2
//
// ```
type AviatrixAzurePeer struct {
	pulumi.CustomResourceState

	// Name of the Azure cloud account in the Aviatrix controller for VNet 1.
	AccountName1 pulumi.StringOutput `pulumi:"accountName1"`
	// Name of the Azure cloud account in the Aviatrix controller for VNet 2.
	AccountName2 pulumi.StringOutput `pulumi:"accountName2"`
	// List of VNet CIDR of vnet_name_resource_group1.
	VnetCidr1s pulumi.StringArrayOutput `pulumi:"vnetCidr1s"`
	// List of VNet CIDR of vnet_name_resource_group2.
	VnetCidr2s pulumi.StringArrayOutput `pulumi:"vnetCidr2s"`
	// Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".
	VnetNameResourceGroup1 pulumi.StringOutput `pulumi:"vnetNameResourceGroup1"`
	// Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".
	VnetNameResourceGroup2 pulumi.StringOutput `pulumi:"vnetNameResourceGroup2"`
	// Region of Azure VNet 1. Example: "East US 2".
	VnetReg1 pulumi.StringOutput `pulumi:"vnetReg1"`
	// Region of Azure VNet 2. Example: "East US 2".
	VnetReg2 pulumi.StringOutput `pulumi:"vnetReg2"`
}

// NewAviatrixAzurePeer registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAzurePeer(ctx *pulumi.Context,
	name string, args *AviatrixAzurePeerArgs, opts ...pulumi.ResourceOption) (*AviatrixAzurePeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName1 == nil {
		return nil, errors.New("invalid value for required argument 'AccountName1'")
	}
	if args.AccountName2 == nil {
		return nil, errors.New("invalid value for required argument 'AccountName2'")
	}
	if args.VnetNameResourceGroup1 == nil {
		return nil, errors.New("invalid value for required argument 'VnetNameResourceGroup1'")
	}
	if args.VnetNameResourceGroup2 == nil {
		return nil, errors.New("invalid value for required argument 'VnetNameResourceGroup2'")
	}
	if args.VnetReg1 == nil {
		return nil, errors.New("invalid value for required argument 'VnetReg1'")
	}
	if args.VnetReg2 == nil {
		return nil, errors.New("invalid value for required argument 'VnetReg2'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAzurePeer
	err := ctx.RegisterResource("aviatrix:index/aviatrixAzurePeer:AviatrixAzurePeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAzurePeer gets an existing AviatrixAzurePeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAzurePeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAzurePeerState, opts ...pulumi.ResourceOption) (*AviatrixAzurePeer, error) {
	var resource AviatrixAzurePeer
	err := ctx.ReadResource("aviatrix:index/aviatrixAzurePeer:AviatrixAzurePeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAzurePeer resources.
type aviatrixAzurePeerState struct {
	// Name of the Azure cloud account in the Aviatrix controller for VNet 1.
	AccountName1 *string `pulumi:"accountName1"`
	// Name of the Azure cloud account in the Aviatrix controller for VNet 2.
	AccountName2 *string `pulumi:"accountName2"`
	// List of VNet CIDR of vnet_name_resource_group1.
	VnetCidr1s []string `pulumi:"vnetCidr1s"`
	// List of VNet CIDR of vnet_name_resource_group2.
	VnetCidr2s []string `pulumi:"vnetCidr2s"`
	// Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".
	VnetNameResourceGroup1 *string `pulumi:"vnetNameResourceGroup1"`
	// Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".
	VnetNameResourceGroup2 *string `pulumi:"vnetNameResourceGroup2"`
	// Region of Azure VNet 1. Example: "East US 2".
	VnetReg1 *string `pulumi:"vnetReg1"`
	// Region of Azure VNet 2. Example: "East US 2".
	VnetReg2 *string `pulumi:"vnetReg2"`
}

type AviatrixAzurePeerState struct {
	// Name of the Azure cloud account in the Aviatrix controller for VNet 1.
	AccountName1 pulumi.StringPtrInput
	// Name of the Azure cloud account in the Aviatrix controller for VNet 2.
	AccountName2 pulumi.StringPtrInput
	// List of VNet CIDR of vnet_name_resource_group1.
	VnetCidr1s pulumi.StringArrayInput
	// List of VNet CIDR of vnet_name_resource_group2.
	VnetCidr2s pulumi.StringArrayInput
	// Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".
	VnetNameResourceGroup1 pulumi.StringPtrInput
	// Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".
	VnetNameResourceGroup2 pulumi.StringPtrInput
	// Region of Azure VNet 1. Example: "East US 2".
	VnetReg1 pulumi.StringPtrInput
	// Region of Azure VNet 2. Example: "East US 2".
	VnetReg2 pulumi.StringPtrInput
}

func (AviatrixAzurePeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAzurePeerState)(nil)).Elem()
}

type aviatrixAzurePeerArgs struct {
	// Name of the Azure cloud account in the Aviatrix controller for VNet 1.
	AccountName1 string `pulumi:"accountName1"`
	// Name of the Azure cloud account in the Aviatrix controller for VNet 2.
	AccountName2 string `pulumi:"accountName2"`
	// Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".
	VnetNameResourceGroup1 string `pulumi:"vnetNameResourceGroup1"`
	// Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".
	VnetNameResourceGroup2 string `pulumi:"vnetNameResourceGroup2"`
	// Region of Azure VNet 1. Example: "East US 2".
	VnetReg1 string `pulumi:"vnetReg1"`
	// Region of Azure VNet 2. Example: "East US 2".
	VnetReg2 string `pulumi:"vnetReg2"`
}

// The set of arguments for constructing a AviatrixAzurePeer resource.
type AviatrixAzurePeerArgs struct {
	// Name of the Azure cloud account in the Aviatrix controller for VNet 1.
	AccountName1 pulumi.StringInput
	// Name of the Azure cloud account in the Aviatrix controller for VNet 2.
	AccountName2 pulumi.StringInput
	// Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".
	VnetNameResourceGroup1 pulumi.StringInput
	// Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".
	VnetNameResourceGroup2 pulumi.StringInput
	// Region of Azure VNet 1. Example: "East US 2".
	VnetReg1 pulumi.StringInput
	// Region of Azure VNet 2. Example: "East US 2".
	VnetReg2 pulumi.StringInput
}

func (AviatrixAzurePeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAzurePeerArgs)(nil)).Elem()
}

type AviatrixAzurePeerInput interface {
	pulumi.Input

	ToAviatrixAzurePeerOutput() AviatrixAzurePeerOutput
	ToAviatrixAzurePeerOutputWithContext(ctx context.Context) AviatrixAzurePeerOutput
}

func (*AviatrixAzurePeer) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAzurePeer)(nil)).Elem()
}

func (i *AviatrixAzurePeer) ToAviatrixAzurePeerOutput() AviatrixAzurePeerOutput {
	return i.ToAviatrixAzurePeerOutputWithContext(context.Background())
}

func (i *AviatrixAzurePeer) ToAviatrixAzurePeerOutputWithContext(ctx context.Context) AviatrixAzurePeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAzurePeerOutput)
}

// AviatrixAzurePeerArrayInput is an input type that accepts AviatrixAzurePeerArray and AviatrixAzurePeerArrayOutput values.
// You can construct a concrete instance of `AviatrixAzurePeerArrayInput` via:
//
//	AviatrixAzurePeerArray{ AviatrixAzurePeerArgs{...} }
type AviatrixAzurePeerArrayInput interface {
	pulumi.Input

	ToAviatrixAzurePeerArrayOutput() AviatrixAzurePeerArrayOutput
	ToAviatrixAzurePeerArrayOutputWithContext(context.Context) AviatrixAzurePeerArrayOutput
}

type AviatrixAzurePeerArray []AviatrixAzurePeerInput

func (AviatrixAzurePeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAzurePeer)(nil)).Elem()
}

func (i AviatrixAzurePeerArray) ToAviatrixAzurePeerArrayOutput() AviatrixAzurePeerArrayOutput {
	return i.ToAviatrixAzurePeerArrayOutputWithContext(context.Background())
}

func (i AviatrixAzurePeerArray) ToAviatrixAzurePeerArrayOutputWithContext(ctx context.Context) AviatrixAzurePeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAzurePeerArrayOutput)
}

// AviatrixAzurePeerMapInput is an input type that accepts AviatrixAzurePeerMap and AviatrixAzurePeerMapOutput values.
// You can construct a concrete instance of `AviatrixAzurePeerMapInput` via:
//
//	AviatrixAzurePeerMap{ "key": AviatrixAzurePeerArgs{...} }
type AviatrixAzurePeerMapInput interface {
	pulumi.Input

	ToAviatrixAzurePeerMapOutput() AviatrixAzurePeerMapOutput
	ToAviatrixAzurePeerMapOutputWithContext(context.Context) AviatrixAzurePeerMapOutput
}

type AviatrixAzurePeerMap map[string]AviatrixAzurePeerInput

func (AviatrixAzurePeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAzurePeer)(nil)).Elem()
}

func (i AviatrixAzurePeerMap) ToAviatrixAzurePeerMapOutput() AviatrixAzurePeerMapOutput {
	return i.ToAviatrixAzurePeerMapOutputWithContext(context.Background())
}

func (i AviatrixAzurePeerMap) ToAviatrixAzurePeerMapOutputWithContext(ctx context.Context) AviatrixAzurePeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAzurePeerMapOutput)
}

type AviatrixAzurePeerOutput struct{ *pulumi.OutputState }

func (AviatrixAzurePeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAzurePeer)(nil)).Elem()
}

func (o AviatrixAzurePeerOutput) ToAviatrixAzurePeerOutput() AviatrixAzurePeerOutput {
	return o
}

func (o AviatrixAzurePeerOutput) ToAviatrixAzurePeerOutputWithContext(ctx context.Context) AviatrixAzurePeerOutput {
	return o
}

// Name of the Azure cloud account in the Aviatrix controller for VNet 1.
func (o AviatrixAzurePeerOutput) AccountName1() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringOutput { return v.AccountName1 }).(pulumi.StringOutput)
}

// Name of the Azure cloud account in the Aviatrix controller for VNet 2.
func (o AviatrixAzurePeerOutput) AccountName2() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringOutput { return v.AccountName2 }).(pulumi.StringOutput)
}

// List of VNet CIDR of vnet_name_resource_group1.
func (o AviatrixAzurePeerOutput) VnetCidr1s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringArrayOutput { return v.VnetCidr1s }).(pulumi.StringArrayOutput)
}

// List of VNet CIDR of vnet_name_resource_group2.
func (o AviatrixAzurePeerOutput) VnetCidr2s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringArrayOutput { return v.VnetCidr2s }).(pulumi.StringArrayOutput)
}

// Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".
func (o AviatrixAzurePeerOutput) VnetNameResourceGroup1() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringOutput { return v.VnetNameResourceGroup1 }).(pulumi.StringOutput)
}

// Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".
func (o AviatrixAzurePeerOutput) VnetNameResourceGroup2() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringOutput { return v.VnetNameResourceGroup2 }).(pulumi.StringOutput)
}

// Region of Azure VNet 1. Example: "East US 2".
func (o AviatrixAzurePeerOutput) VnetReg1() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringOutput { return v.VnetReg1 }).(pulumi.StringOutput)
}

// Region of Azure VNet 2. Example: "East US 2".
func (o AviatrixAzurePeerOutput) VnetReg2() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzurePeer) pulumi.StringOutput { return v.VnetReg2 }).(pulumi.StringOutput)
}

type AviatrixAzurePeerArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAzurePeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAzurePeer)(nil)).Elem()
}

func (o AviatrixAzurePeerArrayOutput) ToAviatrixAzurePeerArrayOutput() AviatrixAzurePeerArrayOutput {
	return o
}

func (o AviatrixAzurePeerArrayOutput) ToAviatrixAzurePeerArrayOutputWithContext(ctx context.Context) AviatrixAzurePeerArrayOutput {
	return o
}

func (o AviatrixAzurePeerArrayOutput) Index(i pulumi.IntInput) AviatrixAzurePeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAzurePeer {
		return vs[0].([]*AviatrixAzurePeer)[vs[1].(int)]
	}).(AviatrixAzurePeerOutput)
}

type AviatrixAzurePeerMapOutput struct{ *pulumi.OutputState }

func (AviatrixAzurePeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAzurePeer)(nil)).Elem()
}

func (o AviatrixAzurePeerMapOutput) ToAviatrixAzurePeerMapOutput() AviatrixAzurePeerMapOutput {
	return o
}

func (o AviatrixAzurePeerMapOutput) ToAviatrixAzurePeerMapOutputWithContext(ctx context.Context) AviatrixAzurePeerMapOutput {
	return o
}

func (o AviatrixAzurePeerMapOutput) MapIndex(k pulumi.StringInput) AviatrixAzurePeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAzurePeer {
		return vs[0].(map[string]*AviatrixAzurePeer)[vs[1].(string)]
	}).(AviatrixAzurePeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAzurePeerInput)(nil)).Elem(), &AviatrixAzurePeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAzurePeerArrayInput)(nil)).Elem(), AviatrixAzurePeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAzurePeerMapInput)(nil)).Elem(), AviatrixAzurePeerMap{})
	pulumi.RegisterOutputType(AviatrixAzurePeerOutput{})
	pulumi.RegisterOutputType(AviatrixAzurePeerArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAzurePeerMapOutput{})
}

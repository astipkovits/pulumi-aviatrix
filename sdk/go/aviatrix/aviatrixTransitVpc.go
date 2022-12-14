// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AviatrixTransitVpc resource creates and manages the Aviatrix Transit Network Gateways.
//
// !> **WARNING:** The `AviatrixTransitVpc` resource is deprecated as of **Release 2.0**. It is currently kept for backward-compatibility and will be removed in the future. Please use the transit gateway resource instead. If this is already in the state, please remove it from state file and import as `AviatrixTransitGateway`.
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
//			_, err := aviatrix.NewAviatrixTransitVpc(ctx, "testTransitGwAws", &aviatrix.AviatrixTransitVpcArgs{
//				AccountName:            pulumi.String("devops_aws"),
//				CloudType:              pulumi.Int(1),
//				ConnectedTransit:       pulumi.String("yes"),
//				EnableHybridConnection: pulumi.Bool(true),
//				GwName:                 pulumi.String("transit"),
//				HaGwSize:               pulumi.String("t2.micro"),
//				HaSubnet:               pulumi.String("10.1.0.0/24"),
//				Subnet:                 pulumi.String("10.1.0.0/24"),
//				TagLists: pulumi.StringArray{
//					pulumi.String("name:value"),
//					pulumi.String("name1:value1"),
//					pulumi.String("name2:value2"),
//				},
//				VpcId:   pulumi.String("vpc-abcd1234"),
//				VpcReg:  pulumi.String("us-east-1"),
//				VpcSize: pulumi.String("t2.micro"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aviatrix.NewAviatrixTransitVpc(ctx, "testTransitGwAzure", &aviatrix.AviatrixTransitVpcArgs{
//				AccountName:      pulumi.String("devops_azure"),
//				CloudType:        pulumi.Int(8),
//				ConnectedTransit: pulumi.String("yes"),
//				GwName:           pulumi.String("transit"),
//				HaGwSize:         pulumi.String("Standard_B1s"),
//				HaSubnet:         pulumi.String("10.30.0.0/24"),
//				Subnet:           pulumi.String("10.30.0.0/24"),
//				VpcId:            pulumi.String("vnet1:hello"),
//				VpcReg:           pulumi.String("West US"),
//				VpcSize:          pulumi.String("Standard_B1s"),
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
// Instance transit_vpc can be imported using the gw_name, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc test gw_name
//
// ```
type AviatrixTransitVpc struct {
	pulumi.CustomResourceState

	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Type of cloud service provider, requires an integer value. Use 1 for AWS.
	CloudType pulumi.IntOutput `pulumi:"cloudType"`
	// Specify Connected Transit status. Supported values: true, false.
	ConnectedTransit pulumi.StringPtrOutput `pulumi:"connectedTransit"`
	// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
	EnableFirenetInterfaces pulumi.BoolPtrOutput `pulumi:"enableFirenetInterfaces"`
	// Sign of readiness for TGW connection. Only supported for aws. Example: false.
	EnableHybridConnection pulumi.BoolPtrOutput `pulumi:"enableHybridConnection"`
	// Enable NAT for this container.
	EnableNat pulumi.StringPtrOutput `pulumi:"enableNat"`
	// Name of the gateway which is going to be created.
	GwName pulumi.StringOutput `pulumi:"gwName"`
	// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
	HaGwSize pulumi.StringPtrOutput `pulumi:"haGwSize"`
	// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insaneMode is enabled and haSubnet is set.
	HaInsaneModeAz pulumi.StringPtrOutput `pulumi:"haInsaneModeAz"`
	// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
	HaSubnet pulumi.StringPtrOutput `pulumi:"haSubnet"`
	// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
	InsaneMode pulumi.BoolPtrOutput `pulumi:"insaneMode"`
	// AZ of subnet being created for Insane Mode Transit Gateway. Required if insaneMode is enabled.
	InsaneModeAz pulumi.StringPtrOutput `pulumi:"insaneModeAz"`
	// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
	TagLists pulumi.StringArrayOutput `pulumi:"tagLists"`
	// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
	VpcReg pulumi.StringOutput `pulumi:"vpcReg"`
	// Size of the gateway instance. Example: AWS: "t2.large", etc...
	VpcSize pulumi.StringOutput `pulumi:"vpcSize"`
}

// NewAviatrixTransitVpc registers a new resource with the given unique name, arguments, and options.
func NewAviatrixTransitVpc(ctx *pulumi.Context,
	name string, args *AviatrixTransitVpcArgs, opts ...pulumi.ResourceOption) (*AviatrixTransitVpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.CloudType == nil {
		return nil, errors.New("invalid value for required argument 'CloudType'")
	}
	if args.GwName == nil {
		return nil, errors.New("invalid value for required argument 'GwName'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VpcReg == nil {
		return nil, errors.New("invalid value for required argument 'VpcReg'")
	}
	if args.VpcSize == nil {
		return nil, errors.New("invalid value for required argument 'VpcSize'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixTransitVpc
	err := ctx.RegisterResource("aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixTransitVpc gets an existing AviatrixTransitVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixTransitVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixTransitVpcState, opts ...pulumi.ResourceOption) (*AviatrixTransitVpc, error) {
	var resource AviatrixTransitVpc
	err := ctx.ReadResource("aviatrix:index/aviatrixTransitVpc:AviatrixTransitVpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixTransitVpc resources.
type aviatrixTransitVpcState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName *string `pulumi:"accountName"`
	// Type of cloud service provider, requires an integer value. Use 1 for AWS.
	CloudType *int `pulumi:"cloudType"`
	// Specify Connected Transit status. Supported values: true, false.
	ConnectedTransit *string `pulumi:"connectedTransit"`
	// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
	EnableFirenetInterfaces *bool `pulumi:"enableFirenetInterfaces"`
	// Sign of readiness for TGW connection. Only supported for aws. Example: false.
	EnableHybridConnection *bool `pulumi:"enableHybridConnection"`
	// Enable NAT for this container.
	EnableNat *string `pulumi:"enableNat"`
	// Name of the gateway which is going to be created.
	GwName *string `pulumi:"gwName"`
	// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
	HaGwSize *string `pulumi:"haGwSize"`
	// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insaneMode is enabled and haSubnet is set.
	HaInsaneModeAz *string `pulumi:"haInsaneModeAz"`
	// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
	HaSubnet *string `pulumi:"haSubnet"`
	// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
	InsaneMode *bool `pulumi:"insaneMode"`
	// AZ of subnet being created for Insane Mode Transit Gateway. Required if insaneMode is enabled.
	InsaneModeAz *string `pulumi:"insaneModeAz"`
	// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
	Subnet *string `pulumi:"subnet"`
	// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
	TagLists []string `pulumi:"tagLists"`
	// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
	VpcId *string `pulumi:"vpcId"`
	// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
	VpcReg *string `pulumi:"vpcReg"`
	// Size of the gateway instance. Example: AWS: "t2.large", etc...
	VpcSize *string `pulumi:"vpcSize"`
}

type AviatrixTransitVpcState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringPtrInput
	// Type of cloud service provider, requires an integer value. Use 1 for AWS.
	CloudType pulumi.IntPtrInput
	// Specify Connected Transit status. Supported values: true, false.
	ConnectedTransit pulumi.StringPtrInput
	// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
	EnableFirenetInterfaces pulumi.BoolPtrInput
	// Sign of readiness for TGW connection. Only supported for aws. Example: false.
	EnableHybridConnection pulumi.BoolPtrInput
	// Enable NAT for this container.
	EnableNat pulumi.StringPtrInput
	// Name of the gateway which is going to be created.
	GwName pulumi.StringPtrInput
	// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
	HaGwSize pulumi.StringPtrInput
	// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insaneMode is enabled and haSubnet is set.
	HaInsaneModeAz pulumi.StringPtrInput
	// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
	HaSubnet pulumi.StringPtrInput
	// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
	InsaneMode pulumi.BoolPtrInput
	// AZ of subnet being created for Insane Mode Transit Gateway. Required if insaneMode is enabled.
	InsaneModeAz pulumi.StringPtrInput
	// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
	Subnet pulumi.StringPtrInput
	// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
	TagLists pulumi.StringArrayInput
	// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
	VpcId pulumi.StringPtrInput
	// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
	VpcReg pulumi.StringPtrInput
	// Size of the gateway instance. Example: AWS: "t2.large", etc...
	VpcSize pulumi.StringPtrInput
}

func (AviatrixTransitVpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixTransitVpcState)(nil)).Elem()
}

type aviatrixTransitVpcArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName string `pulumi:"accountName"`
	// Type of cloud service provider, requires an integer value. Use 1 for AWS.
	CloudType int `pulumi:"cloudType"`
	// Specify Connected Transit status. Supported values: true, false.
	ConnectedTransit *string `pulumi:"connectedTransit"`
	// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
	EnableFirenetInterfaces *bool `pulumi:"enableFirenetInterfaces"`
	// Sign of readiness for TGW connection. Only supported for aws. Example: false.
	EnableHybridConnection *bool `pulumi:"enableHybridConnection"`
	// Enable NAT for this container.
	EnableNat *string `pulumi:"enableNat"`
	// Name of the gateway which is going to be created.
	GwName string `pulumi:"gwName"`
	// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
	HaGwSize *string `pulumi:"haGwSize"`
	// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insaneMode is enabled and haSubnet is set.
	HaInsaneModeAz *string `pulumi:"haInsaneModeAz"`
	// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
	HaSubnet *string `pulumi:"haSubnet"`
	// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
	InsaneMode *bool `pulumi:"insaneMode"`
	// AZ of subnet being created for Insane Mode Transit Gateway. Required if insaneMode is enabled.
	InsaneModeAz *string `pulumi:"insaneModeAz"`
	// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
	Subnet string `pulumi:"subnet"`
	// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
	TagLists []string `pulumi:"tagLists"`
	// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
	VpcId string `pulumi:"vpcId"`
	// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
	VpcReg string `pulumi:"vpcReg"`
	// Size of the gateway instance. Example: AWS: "t2.large", etc...
	VpcSize string `pulumi:"vpcSize"`
}

// The set of arguments for constructing a AviatrixTransitVpc resource.
type AviatrixTransitVpcArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringInput
	// Type of cloud service provider, requires an integer value. Use 1 for AWS.
	CloudType pulumi.IntInput
	// Specify Connected Transit status. Supported values: true, false.
	ConnectedTransit pulumi.StringPtrInput
	// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
	EnableFirenetInterfaces pulumi.BoolPtrInput
	// Sign of readiness for TGW connection. Only supported for aws. Example: false.
	EnableHybridConnection pulumi.BoolPtrInput
	// Enable NAT for this container.
	EnableNat pulumi.StringPtrInput
	// Name of the gateway which is going to be created.
	GwName pulumi.StringInput
	// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
	HaGwSize pulumi.StringPtrInput
	// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insaneMode is enabled and haSubnet is set.
	HaInsaneModeAz pulumi.StringPtrInput
	// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
	HaSubnet pulumi.StringPtrInput
	// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
	InsaneMode pulumi.BoolPtrInput
	// AZ of subnet being created for Insane Mode Transit Gateway. Required if insaneMode is enabled.
	InsaneModeAz pulumi.StringPtrInput
	// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
	Subnet pulumi.StringInput
	// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
	TagLists pulumi.StringArrayInput
	// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
	VpcId pulumi.StringInput
	// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
	VpcReg pulumi.StringInput
	// Size of the gateway instance. Example: AWS: "t2.large", etc...
	VpcSize pulumi.StringInput
}

func (AviatrixTransitVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixTransitVpcArgs)(nil)).Elem()
}

type AviatrixTransitVpcInput interface {
	pulumi.Input

	ToAviatrixTransitVpcOutput() AviatrixTransitVpcOutput
	ToAviatrixTransitVpcOutputWithContext(ctx context.Context) AviatrixTransitVpcOutput
}

func (*AviatrixTransitVpc) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixTransitVpc)(nil)).Elem()
}

func (i *AviatrixTransitVpc) ToAviatrixTransitVpcOutput() AviatrixTransitVpcOutput {
	return i.ToAviatrixTransitVpcOutputWithContext(context.Background())
}

func (i *AviatrixTransitVpc) ToAviatrixTransitVpcOutputWithContext(ctx context.Context) AviatrixTransitVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTransitVpcOutput)
}

// AviatrixTransitVpcArrayInput is an input type that accepts AviatrixTransitVpcArray and AviatrixTransitVpcArrayOutput values.
// You can construct a concrete instance of `AviatrixTransitVpcArrayInput` via:
//
//	AviatrixTransitVpcArray{ AviatrixTransitVpcArgs{...} }
type AviatrixTransitVpcArrayInput interface {
	pulumi.Input

	ToAviatrixTransitVpcArrayOutput() AviatrixTransitVpcArrayOutput
	ToAviatrixTransitVpcArrayOutputWithContext(context.Context) AviatrixTransitVpcArrayOutput
}

type AviatrixTransitVpcArray []AviatrixTransitVpcInput

func (AviatrixTransitVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixTransitVpc)(nil)).Elem()
}

func (i AviatrixTransitVpcArray) ToAviatrixTransitVpcArrayOutput() AviatrixTransitVpcArrayOutput {
	return i.ToAviatrixTransitVpcArrayOutputWithContext(context.Background())
}

func (i AviatrixTransitVpcArray) ToAviatrixTransitVpcArrayOutputWithContext(ctx context.Context) AviatrixTransitVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTransitVpcArrayOutput)
}

// AviatrixTransitVpcMapInput is an input type that accepts AviatrixTransitVpcMap and AviatrixTransitVpcMapOutput values.
// You can construct a concrete instance of `AviatrixTransitVpcMapInput` via:
//
//	AviatrixTransitVpcMap{ "key": AviatrixTransitVpcArgs{...} }
type AviatrixTransitVpcMapInput interface {
	pulumi.Input

	ToAviatrixTransitVpcMapOutput() AviatrixTransitVpcMapOutput
	ToAviatrixTransitVpcMapOutputWithContext(context.Context) AviatrixTransitVpcMapOutput
}

type AviatrixTransitVpcMap map[string]AviatrixTransitVpcInput

func (AviatrixTransitVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixTransitVpc)(nil)).Elem()
}

func (i AviatrixTransitVpcMap) ToAviatrixTransitVpcMapOutput() AviatrixTransitVpcMapOutput {
	return i.ToAviatrixTransitVpcMapOutputWithContext(context.Background())
}

func (i AviatrixTransitVpcMap) ToAviatrixTransitVpcMapOutputWithContext(ctx context.Context) AviatrixTransitVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTransitVpcMapOutput)
}

type AviatrixTransitVpcOutput struct{ *pulumi.OutputState }

func (AviatrixTransitVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixTransitVpc)(nil)).Elem()
}

func (o AviatrixTransitVpcOutput) ToAviatrixTransitVpcOutput() AviatrixTransitVpcOutput {
	return o
}

func (o AviatrixTransitVpcOutput) ToAviatrixTransitVpcOutputWithContext(ctx context.Context) AviatrixTransitVpcOutput {
	return o
}

// This parameter represents the name of a Cloud-Account in Aviatrix controller.
func (o AviatrixTransitVpcOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Type of cloud service provider, requires an integer value. Use 1 for AWS.
func (o AviatrixTransitVpcOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.IntOutput { return v.CloudType }).(pulumi.IntOutput)
}

// Specify Connected Transit status. Supported values: true, false.
func (o AviatrixTransitVpcOutput) ConnectedTransit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringPtrOutput { return v.ConnectedTransit }).(pulumi.StringPtrOutput)
}

// Sign of readiness for FireNet connection. Valid values: true and false. Default: false.
func (o AviatrixTransitVpcOutput) EnableFirenetInterfaces() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.BoolPtrOutput { return v.EnableFirenetInterfaces }).(pulumi.BoolPtrOutput)
}

// Sign of readiness for TGW connection. Only supported for aws. Example: false.
func (o AviatrixTransitVpcOutput) EnableHybridConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.BoolPtrOutput { return v.EnableHybridConnection }).(pulumi.BoolPtrOutput)
}

// Enable NAT for this container.
func (o AviatrixTransitVpcOutput) EnableNat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringPtrOutput { return v.EnableNat }).(pulumi.StringPtrOutput)
}

// Name of the gateway which is going to be created.
func (o AviatrixTransitVpcOutput) GwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringOutput { return v.GwName }).(pulumi.StringOutput)
}

// HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
func (o AviatrixTransitVpcOutput) HaGwSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringPtrOutput { return v.HaGwSize }).(pulumi.StringPtrOutput)
}

// AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insaneMode is enabled and haSubnet is set.
func (o AviatrixTransitVpcOutput) HaInsaneModeAz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringPtrOutput { return v.HaInsaneModeAz }).(pulumi.StringPtrOutput)
}

// HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.
func (o AviatrixTransitVpcOutput) HaSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringPtrOutput { return v.HaSubnet }).(pulumi.StringPtrOutput)
}

// Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.
func (o AviatrixTransitVpcOutput) InsaneMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.BoolPtrOutput { return v.InsaneMode }).(pulumi.BoolPtrOutput)
}

// AZ of subnet being created for Insane Mode Transit Gateway. Required if insaneMode is enabled.
func (o AviatrixTransitVpcOutput) InsaneModeAz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringPtrOutput { return v.InsaneModeAz }).(pulumi.StringPtrOutput)
}

// Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.
func (o AviatrixTransitVpcOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]
func (o AviatrixTransitVpcOutput) TagLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringArrayOutput { return v.TagLists }).(pulumi.StringArrayOutput)
}

// VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...
func (o AviatrixTransitVpcOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...
func (o AviatrixTransitVpcOutput) VpcReg() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringOutput { return v.VpcReg }).(pulumi.StringOutput)
}

// Size of the gateway instance. Example: AWS: "t2.large", etc...
func (o AviatrixTransitVpcOutput) VpcSize() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTransitVpc) pulumi.StringOutput { return v.VpcSize }).(pulumi.StringOutput)
}

type AviatrixTransitVpcArrayOutput struct{ *pulumi.OutputState }

func (AviatrixTransitVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixTransitVpc)(nil)).Elem()
}

func (o AviatrixTransitVpcArrayOutput) ToAviatrixTransitVpcArrayOutput() AviatrixTransitVpcArrayOutput {
	return o
}

func (o AviatrixTransitVpcArrayOutput) ToAviatrixTransitVpcArrayOutputWithContext(ctx context.Context) AviatrixTransitVpcArrayOutput {
	return o
}

func (o AviatrixTransitVpcArrayOutput) Index(i pulumi.IntInput) AviatrixTransitVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixTransitVpc {
		return vs[0].([]*AviatrixTransitVpc)[vs[1].(int)]
	}).(AviatrixTransitVpcOutput)
}

type AviatrixTransitVpcMapOutput struct{ *pulumi.OutputState }

func (AviatrixTransitVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixTransitVpc)(nil)).Elem()
}

func (o AviatrixTransitVpcMapOutput) ToAviatrixTransitVpcMapOutput() AviatrixTransitVpcMapOutput {
	return o
}

func (o AviatrixTransitVpcMapOutput) ToAviatrixTransitVpcMapOutputWithContext(ctx context.Context) AviatrixTransitVpcMapOutput {
	return o
}

func (o AviatrixTransitVpcMapOutput) MapIndex(k pulumi.StringInput) AviatrixTransitVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixTransitVpc {
		return vs[0].(map[string]*AviatrixTransitVpc)[vs[1].(string)]
	}).(AviatrixTransitVpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTransitVpcInput)(nil)).Elem(), &AviatrixTransitVpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTransitVpcArrayInput)(nil)).Elem(), AviatrixTransitVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTransitVpcMapInput)(nil)).Elem(), AviatrixTransitVpcMap{})
	pulumi.RegisterOutputType(AviatrixTransitVpcOutput{})
	pulumi.RegisterOutputType(AviatrixTransitVpcArrayOutput{})
	pulumi.RegisterOutputType(AviatrixTransitVpcMapOutput{})
}

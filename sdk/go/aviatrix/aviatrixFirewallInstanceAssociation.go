// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_firewall_instance_association** resource allows for the creation and management of a firewall instance association. To use this resource you must also have an `AviatrixFirenet` resource with it's `manageFirewallInstanceAssociation` attribute set to false.
//
// Available in provider version R2.17.1+.
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
//			_, err := aviatrix.NewAviatrixFirewallInstanceAssociation(ctx, "firewallInstanceAssociation1", &aviatrix.AviatrixFirewallInstanceAssociationArgs{
//				VpcId:               pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Vpc_id),
//				FirenetGwName:       pulumi.Any(aviatrix_transit_gateway.Transit_gateway_1.Gw_name),
//				InstanceId:          pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Instance_id),
//				FirewallName:        pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Firewall_name),
//				LanInterface:        pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Lan_interface),
//				ManagementInterface: pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Management_interface),
//				EgressInterface:     pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Egress_interface),
//				Attached:            pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixFirewallInstanceAssociation(ctx, "firewallInstanceAssociation1", &aviatrix.AviatrixFirewallInstanceAssociationArgs{
//				VpcId:               pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Vpc_id),
//				FirenetGwName:       pulumi.Any(aviatrix_transit_gateway.Transit_gateway_1.Gw_name),
//				InstanceId:          pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Instance_id),
//				LanInterface:        pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Lan_interface),
//				ManagementInterface: pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Management_interface),
//				EgressInterface:     pulumi.Any(aviatrix_firewall_instance.Firewall_instance_1.Egress_interface),
//				Attached:            pulumi.Bool(true),
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
// **firewall_instance_association** can be imported using the `vpc_id`, `firenet_gw_name` and `instance_id` in the form `vpc_id~~firenet_gw_name~~instance_id` e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation test vpc_id~~firenet_gw_name~~instance_id
//
// ```
//
//	When using a Native GWLB VPC where there is no `firenet_gw_name` but the ID is in the same form e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation test vpc_id~~~~instance_id
//
// ```
type AviatrixFirewallInstanceAssociation struct {
	pulumi.CustomResourceState

	// Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
	Attached pulumi.BoolPtrOutput `pulumi:"attached"`
	// Egress interface ID. **Required if it is a firewall instance.**
	EgressInterface pulumi.StringPtrOutput `pulumi:"egressInterface"`
	// Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
	FirenetGwName pulumi.StringPtrOutput `pulumi:"firenetGwName"`
	// Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
	FirewallName pulumi.StringPtrOutput `pulumi:"firewallName"`
	// ID of Firewall instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Lan interface ID. **Required if it is a firewall instance.**
	LanInterface pulumi.StringPtrOutput `pulumi:"lanInterface"`
	// Management interface ID. **Required if it is a firewall instance.**
	ManagementInterface pulumi.StringPtrOutput `pulumi:"managementInterface"`
	// Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
	VendorType pulumi.StringPtrOutput `pulumi:"vendorType"`
	// VPC ID of the Security VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewAviatrixFirewallInstanceAssociation registers a new resource with the given unique name, arguments, and options.
func NewAviatrixFirewallInstanceAssociation(ctx *pulumi.Context,
	name string, args *AviatrixFirewallInstanceAssociationArgs, opts ...pulumi.ResourceOption) (*AviatrixFirewallInstanceAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixFirewallInstanceAssociation
	err := ctx.RegisterResource("aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixFirewallInstanceAssociation gets an existing AviatrixFirewallInstanceAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixFirewallInstanceAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixFirewallInstanceAssociationState, opts ...pulumi.ResourceOption) (*AviatrixFirewallInstanceAssociation, error) {
	var resource AviatrixFirewallInstanceAssociation
	err := ctx.ReadResource("aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixFirewallInstanceAssociation resources.
type aviatrixFirewallInstanceAssociationState struct {
	// Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
	Attached *bool `pulumi:"attached"`
	// Egress interface ID. **Required if it is a firewall instance.**
	EgressInterface *string `pulumi:"egressInterface"`
	// Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
	FirenetGwName *string `pulumi:"firenetGwName"`
	// Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
	FirewallName *string `pulumi:"firewallName"`
	// ID of Firewall instance.
	InstanceId *string `pulumi:"instanceId"`
	// Lan interface ID. **Required if it is a firewall instance.**
	LanInterface *string `pulumi:"lanInterface"`
	// Management interface ID. **Required if it is a firewall instance.**
	ManagementInterface *string `pulumi:"managementInterface"`
	// Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
	VendorType *string `pulumi:"vendorType"`
	// VPC ID of the Security VPC.
	VpcId *string `pulumi:"vpcId"`
}

type AviatrixFirewallInstanceAssociationState struct {
	// Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
	Attached pulumi.BoolPtrInput
	// Egress interface ID. **Required if it is a firewall instance.**
	EgressInterface pulumi.StringPtrInput
	// Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
	FirenetGwName pulumi.StringPtrInput
	// Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
	FirewallName pulumi.StringPtrInput
	// ID of Firewall instance.
	InstanceId pulumi.StringPtrInput
	// Lan interface ID. **Required if it is a firewall instance.**
	LanInterface pulumi.StringPtrInput
	// Management interface ID. **Required if it is a firewall instance.**
	ManagementInterface pulumi.StringPtrInput
	// Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
	VendorType pulumi.StringPtrInput
	// VPC ID of the Security VPC.
	VpcId pulumi.StringPtrInput
}

func (AviatrixFirewallInstanceAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixFirewallInstanceAssociationState)(nil)).Elem()
}

type aviatrixFirewallInstanceAssociationArgs struct {
	// Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
	Attached *bool `pulumi:"attached"`
	// Egress interface ID. **Required if it is a firewall instance.**
	EgressInterface *string `pulumi:"egressInterface"`
	// Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
	FirenetGwName *string `pulumi:"firenetGwName"`
	// Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
	FirewallName *string `pulumi:"firewallName"`
	// ID of Firewall instance.
	InstanceId string `pulumi:"instanceId"`
	// Lan interface ID. **Required if it is a firewall instance.**
	LanInterface *string `pulumi:"lanInterface"`
	// Management interface ID. **Required if it is a firewall instance.**
	ManagementInterface *string `pulumi:"managementInterface"`
	// Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
	VendorType *string `pulumi:"vendorType"`
	// VPC ID of the Security VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a AviatrixFirewallInstanceAssociation resource.
type AviatrixFirewallInstanceAssociationArgs struct {
	// Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
	Attached pulumi.BoolPtrInput
	// Egress interface ID. **Required if it is a firewall instance.**
	EgressInterface pulumi.StringPtrInput
	// Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
	FirenetGwName pulumi.StringPtrInput
	// Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
	FirewallName pulumi.StringPtrInput
	// ID of Firewall instance.
	InstanceId pulumi.StringInput
	// Lan interface ID. **Required if it is a firewall instance.**
	LanInterface pulumi.StringPtrInput
	// Management interface ID. **Required if it is a firewall instance.**
	ManagementInterface pulumi.StringPtrInput
	// Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
	VendorType pulumi.StringPtrInput
	// VPC ID of the Security VPC.
	VpcId pulumi.StringInput
}

func (AviatrixFirewallInstanceAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixFirewallInstanceAssociationArgs)(nil)).Elem()
}

type AviatrixFirewallInstanceAssociationInput interface {
	pulumi.Input

	ToAviatrixFirewallInstanceAssociationOutput() AviatrixFirewallInstanceAssociationOutput
	ToAviatrixFirewallInstanceAssociationOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationOutput
}

func (*AviatrixFirewallInstanceAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixFirewallInstanceAssociation)(nil)).Elem()
}

func (i *AviatrixFirewallInstanceAssociation) ToAviatrixFirewallInstanceAssociationOutput() AviatrixFirewallInstanceAssociationOutput {
	return i.ToAviatrixFirewallInstanceAssociationOutputWithContext(context.Background())
}

func (i *AviatrixFirewallInstanceAssociation) ToAviatrixFirewallInstanceAssociationOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirewallInstanceAssociationOutput)
}

// AviatrixFirewallInstanceAssociationArrayInput is an input type that accepts AviatrixFirewallInstanceAssociationArray and AviatrixFirewallInstanceAssociationArrayOutput values.
// You can construct a concrete instance of `AviatrixFirewallInstanceAssociationArrayInput` via:
//
//	AviatrixFirewallInstanceAssociationArray{ AviatrixFirewallInstanceAssociationArgs{...} }
type AviatrixFirewallInstanceAssociationArrayInput interface {
	pulumi.Input

	ToAviatrixFirewallInstanceAssociationArrayOutput() AviatrixFirewallInstanceAssociationArrayOutput
	ToAviatrixFirewallInstanceAssociationArrayOutputWithContext(context.Context) AviatrixFirewallInstanceAssociationArrayOutput
}

type AviatrixFirewallInstanceAssociationArray []AviatrixFirewallInstanceAssociationInput

func (AviatrixFirewallInstanceAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixFirewallInstanceAssociation)(nil)).Elem()
}

func (i AviatrixFirewallInstanceAssociationArray) ToAviatrixFirewallInstanceAssociationArrayOutput() AviatrixFirewallInstanceAssociationArrayOutput {
	return i.ToAviatrixFirewallInstanceAssociationArrayOutputWithContext(context.Background())
}

func (i AviatrixFirewallInstanceAssociationArray) ToAviatrixFirewallInstanceAssociationArrayOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirewallInstanceAssociationArrayOutput)
}

// AviatrixFirewallInstanceAssociationMapInput is an input type that accepts AviatrixFirewallInstanceAssociationMap and AviatrixFirewallInstanceAssociationMapOutput values.
// You can construct a concrete instance of `AviatrixFirewallInstanceAssociationMapInput` via:
//
//	AviatrixFirewallInstanceAssociationMap{ "key": AviatrixFirewallInstanceAssociationArgs{...} }
type AviatrixFirewallInstanceAssociationMapInput interface {
	pulumi.Input

	ToAviatrixFirewallInstanceAssociationMapOutput() AviatrixFirewallInstanceAssociationMapOutput
	ToAviatrixFirewallInstanceAssociationMapOutputWithContext(context.Context) AviatrixFirewallInstanceAssociationMapOutput
}

type AviatrixFirewallInstanceAssociationMap map[string]AviatrixFirewallInstanceAssociationInput

func (AviatrixFirewallInstanceAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixFirewallInstanceAssociation)(nil)).Elem()
}

func (i AviatrixFirewallInstanceAssociationMap) ToAviatrixFirewallInstanceAssociationMapOutput() AviatrixFirewallInstanceAssociationMapOutput {
	return i.ToAviatrixFirewallInstanceAssociationMapOutputWithContext(context.Background())
}

func (i AviatrixFirewallInstanceAssociationMap) ToAviatrixFirewallInstanceAssociationMapOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirewallInstanceAssociationMapOutput)
}

type AviatrixFirewallInstanceAssociationOutput struct{ *pulumi.OutputState }

func (AviatrixFirewallInstanceAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixFirewallInstanceAssociation)(nil)).Elem()
}

func (o AviatrixFirewallInstanceAssociationOutput) ToAviatrixFirewallInstanceAssociationOutput() AviatrixFirewallInstanceAssociationOutput {
	return o
}

func (o AviatrixFirewallInstanceAssociationOutput) ToAviatrixFirewallInstanceAssociationOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationOutput {
	return o
}

// Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
func (o AviatrixFirewallInstanceAssociationOutput) Attached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.BoolPtrOutput { return v.Attached }).(pulumi.BoolPtrOutput)
}

// Egress interface ID. **Required if it is a firewall instance.**
func (o AviatrixFirewallInstanceAssociationOutput) EgressInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringPtrOutput { return v.EgressInterface }).(pulumi.StringPtrOutput)
}

// Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
func (o AviatrixFirewallInstanceAssociationOutput) FirenetGwName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringPtrOutput { return v.FirenetGwName }).(pulumi.StringPtrOutput)
}

// Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
func (o AviatrixFirewallInstanceAssociationOutput) FirewallName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringPtrOutput { return v.FirewallName }).(pulumi.StringPtrOutput)
}

// ID of Firewall instance.
func (o AviatrixFirewallInstanceAssociationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Lan interface ID. **Required if it is a firewall instance.**
func (o AviatrixFirewallInstanceAssociationOutput) LanInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringPtrOutput { return v.LanInterface }).(pulumi.StringPtrOutput)
}

// Management interface ID. **Required if it is a firewall instance.**
func (o AviatrixFirewallInstanceAssociationOutput) ManagementInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringPtrOutput { return v.ManagementInterface }).(pulumi.StringPtrOutput)
}

// Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
func (o AviatrixFirewallInstanceAssociationOutput) VendorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringPtrOutput { return v.VendorType }).(pulumi.StringPtrOutput)
}

// VPC ID of the Security VPC.
func (o AviatrixFirewallInstanceAssociationOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstanceAssociation) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type AviatrixFirewallInstanceAssociationArrayOutput struct{ *pulumi.OutputState }

func (AviatrixFirewallInstanceAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixFirewallInstanceAssociation)(nil)).Elem()
}

func (o AviatrixFirewallInstanceAssociationArrayOutput) ToAviatrixFirewallInstanceAssociationArrayOutput() AviatrixFirewallInstanceAssociationArrayOutput {
	return o
}

func (o AviatrixFirewallInstanceAssociationArrayOutput) ToAviatrixFirewallInstanceAssociationArrayOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationArrayOutput {
	return o
}

func (o AviatrixFirewallInstanceAssociationArrayOutput) Index(i pulumi.IntInput) AviatrixFirewallInstanceAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixFirewallInstanceAssociation {
		return vs[0].([]*AviatrixFirewallInstanceAssociation)[vs[1].(int)]
	}).(AviatrixFirewallInstanceAssociationOutput)
}

type AviatrixFirewallInstanceAssociationMapOutput struct{ *pulumi.OutputState }

func (AviatrixFirewallInstanceAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixFirewallInstanceAssociation)(nil)).Elem()
}

func (o AviatrixFirewallInstanceAssociationMapOutput) ToAviatrixFirewallInstanceAssociationMapOutput() AviatrixFirewallInstanceAssociationMapOutput {
	return o
}

func (o AviatrixFirewallInstanceAssociationMapOutput) ToAviatrixFirewallInstanceAssociationMapOutputWithContext(ctx context.Context) AviatrixFirewallInstanceAssociationMapOutput {
	return o
}

func (o AviatrixFirewallInstanceAssociationMapOutput) MapIndex(k pulumi.StringInput) AviatrixFirewallInstanceAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixFirewallInstanceAssociation {
		return vs[0].(map[string]*AviatrixFirewallInstanceAssociation)[vs[1].(string)]
	}).(AviatrixFirewallInstanceAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirewallInstanceAssociationInput)(nil)).Elem(), &AviatrixFirewallInstanceAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirewallInstanceAssociationArrayInput)(nil)).Elem(), AviatrixFirewallInstanceAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirewallInstanceAssociationMapInput)(nil)).Elem(), AviatrixFirewallInstanceAssociationMap{})
	pulumi.RegisterOutputType(AviatrixFirewallInstanceAssociationOutput{})
	pulumi.RegisterOutputType(AviatrixFirewallInstanceAssociationArrayOutput{})
	pulumi.RegisterOutputType(AviatrixFirewallInstanceAssociationMapOutput{})
}

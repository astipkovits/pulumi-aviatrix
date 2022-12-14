// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_vpn_user_accelerator** resource manages the [Aviatrix VPN User Accelerator](https://docs.aviatrix.com/HowTos/user_accelerator.html).
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
//			_, err := aviatrix.NewAviatrixVpnUserAccelerator(ctx, "testVpcAccelerator", &aviatrix.AviatrixVpnUserAcceleratorArgs{
//				ElbName: pulumi.String("Aviatrix-vpc-abcd2134"),
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
// **vpn_user_accelerator** can be imported using the `elb_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator test Aviatrix-vpc-abcd1234
//
// ```
type AviatrixVpnUserAccelerator struct {
	pulumi.CustomResourceState

	// Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
	ElbName pulumi.StringOutput `pulumi:"elbName"`
}

// NewAviatrixVpnUserAccelerator registers a new resource with the given unique name, arguments, and options.
func NewAviatrixVpnUserAccelerator(ctx *pulumi.Context,
	name string, args *AviatrixVpnUserAcceleratorArgs, opts ...pulumi.ResourceOption) (*AviatrixVpnUserAccelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElbName == nil {
		return nil, errors.New("invalid value for required argument 'ElbName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixVpnUserAccelerator
	err := ctx.RegisterResource("aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixVpnUserAccelerator gets an existing AviatrixVpnUserAccelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixVpnUserAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixVpnUserAcceleratorState, opts ...pulumi.ResourceOption) (*AviatrixVpnUserAccelerator, error) {
	var resource AviatrixVpnUserAccelerator
	err := ctx.ReadResource("aviatrix:index/aviatrixVpnUserAccelerator:AviatrixVpnUserAccelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixVpnUserAccelerator resources.
type aviatrixVpnUserAcceleratorState struct {
	// Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
	ElbName *string `pulumi:"elbName"`
}

type AviatrixVpnUserAcceleratorState struct {
	// Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
	ElbName pulumi.StringPtrInput
}

func (AviatrixVpnUserAcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixVpnUserAcceleratorState)(nil)).Elem()
}

type aviatrixVpnUserAcceleratorArgs struct {
	// Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
	ElbName string `pulumi:"elbName"`
}

// The set of arguments for constructing a AviatrixVpnUserAccelerator resource.
type AviatrixVpnUserAcceleratorArgs struct {
	// Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
	ElbName pulumi.StringInput
}

func (AviatrixVpnUserAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixVpnUserAcceleratorArgs)(nil)).Elem()
}

type AviatrixVpnUserAcceleratorInput interface {
	pulumi.Input

	ToAviatrixVpnUserAcceleratorOutput() AviatrixVpnUserAcceleratorOutput
	ToAviatrixVpnUserAcceleratorOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorOutput
}

func (*AviatrixVpnUserAccelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixVpnUserAccelerator)(nil)).Elem()
}

func (i *AviatrixVpnUserAccelerator) ToAviatrixVpnUserAcceleratorOutput() AviatrixVpnUserAcceleratorOutput {
	return i.ToAviatrixVpnUserAcceleratorOutputWithContext(context.Background())
}

func (i *AviatrixVpnUserAccelerator) ToAviatrixVpnUserAcceleratorOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpnUserAcceleratorOutput)
}

// AviatrixVpnUserAcceleratorArrayInput is an input type that accepts AviatrixVpnUserAcceleratorArray and AviatrixVpnUserAcceleratorArrayOutput values.
// You can construct a concrete instance of `AviatrixVpnUserAcceleratorArrayInput` via:
//
//	AviatrixVpnUserAcceleratorArray{ AviatrixVpnUserAcceleratorArgs{...} }
type AviatrixVpnUserAcceleratorArrayInput interface {
	pulumi.Input

	ToAviatrixVpnUserAcceleratorArrayOutput() AviatrixVpnUserAcceleratorArrayOutput
	ToAviatrixVpnUserAcceleratorArrayOutputWithContext(context.Context) AviatrixVpnUserAcceleratorArrayOutput
}

type AviatrixVpnUserAcceleratorArray []AviatrixVpnUserAcceleratorInput

func (AviatrixVpnUserAcceleratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixVpnUserAccelerator)(nil)).Elem()
}

func (i AviatrixVpnUserAcceleratorArray) ToAviatrixVpnUserAcceleratorArrayOutput() AviatrixVpnUserAcceleratorArrayOutput {
	return i.ToAviatrixVpnUserAcceleratorArrayOutputWithContext(context.Background())
}

func (i AviatrixVpnUserAcceleratorArray) ToAviatrixVpnUserAcceleratorArrayOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpnUserAcceleratorArrayOutput)
}

// AviatrixVpnUserAcceleratorMapInput is an input type that accepts AviatrixVpnUserAcceleratorMap and AviatrixVpnUserAcceleratorMapOutput values.
// You can construct a concrete instance of `AviatrixVpnUserAcceleratorMapInput` via:
//
//	AviatrixVpnUserAcceleratorMap{ "key": AviatrixVpnUserAcceleratorArgs{...} }
type AviatrixVpnUserAcceleratorMapInput interface {
	pulumi.Input

	ToAviatrixVpnUserAcceleratorMapOutput() AviatrixVpnUserAcceleratorMapOutput
	ToAviatrixVpnUserAcceleratorMapOutputWithContext(context.Context) AviatrixVpnUserAcceleratorMapOutput
}

type AviatrixVpnUserAcceleratorMap map[string]AviatrixVpnUserAcceleratorInput

func (AviatrixVpnUserAcceleratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixVpnUserAccelerator)(nil)).Elem()
}

func (i AviatrixVpnUserAcceleratorMap) ToAviatrixVpnUserAcceleratorMapOutput() AviatrixVpnUserAcceleratorMapOutput {
	return i.ToAviatrixVpnUserAcceleratorMapOutputWithContext(context.Background())
}

func (i AviatrixVpnUserAcceleratorMap) ToAviatrixVpnUserAcceleratorMapOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixVpnUserAcceleratorMapOutput)
}

type AviatrixVpnUserAcceleratorOutput struct{ *pulumi.OutputState }

func (AviatrixVpnUserAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixVpnUserAccelerator)(nil)).Elem()
}

func (o AviatrixVpnUserAcceleratorOutput) ToAviatrixVpnUserAcceleratorOutput() AviatrixVpnUserAcceleratorOutput {
	return o
}

func (o AviatrixVpnUserAcceleratorOutput) ToAviatrixVpnUserAcceleratorOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorOutput {
	return o
}

// Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134".
func (o AviatrixVpnUserAcceleratorOutput) ElbName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixVpnUserAccelerator) pulumi.StringOutput { return v.ElbName }).(pulumi.StringOutput)
}

type AviatrixVpnUserAcceleratorArrayOutput struct{ *pulumi.OutputState }

func (AviatrixVpnUserAcceleratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixVpnUserAccelerator)(nil)).Elem()
}

func (o AviatrixVpnUserAcceleratorArrayOutput) ToAviatrixVpnUserAcceleratorArrayOutput() AviatrixVpnUserAcceleratorArrayOutput {
	return o
}

func (o AviatrixVpnUserAcceleratorArrayOutput) ToAviatrixVpnUserAcceleratorArrayOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorArrayOutput {
	return o
}

func (o AviatrixVpnUserAcceleratorArrayOutput) Index(i pulumi.IntInput) AviatrixVpnUserAcceleratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixVpnUserAccelerator {
		return vs[0].([]*AviatrixVpnUserAccelerator)[vs[1].(int)]
	}).(AviatrixVpnUserAcceleratorOutput)
}

type AviatrixVpnUserAcceleratorMapOutput struct{ *pulumi.OutputState }

func (AviatrixVpnUserAcceleratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixVpnUserAccelerator)(nil)).Elem()
}

func (o AviatrixVpnUserAcceleratorMapOutput) ToAviatrixVpnUserAcceleratorMapOutput() AviatrixVpnUserAcceleratorMapOutput {
	return o
}

func (o AviatrixVpnUserAcceleratorMapOutput) ToAviatrixVpnUserAcceleratorMapOutputWithContext(ctx context.Context) AviatrixVpnUserAcceleratorMapOutput {
	return o
}

func (o AviatrixVpnUserAcceleratorMapOutput) MapIndex(k pulumi.StringInput) AviatrixVpnUserAcceleratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixVpnUserAccelerator {
		return vs[0].(map[string]*AviatrixVpnUserAccelerator)[vs[1].(string)]
	}).(AviatrixVpnUserAcceleratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpnUserAcceleratorInput)(nil)).Elem(), &AviatrixVpnUserAccelerator{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpnUserAcceleratorArrayInput)(nil)).Elem(), AviatrixVpnUserAcceleratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixVpnUserAcceleratorMapInput)(nil)).Elem(), AviatrixVpnUserAcceleratorMap{})
	pulumi.RegisterOutputType(AviatrixVpnUserAcceleratorOutput{})
	pulumi.RegisterOutputType(AviatrixVpnUserAcceleratorArrayOutput{})
	pulumi.RegisterOutputType(AviatrixVpnUserAcceleratorMapOutput{})
}

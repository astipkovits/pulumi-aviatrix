// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_tunnel** resource allows the creation and management of Aviatrix Encrypted Peering tunnels.
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
//			_, err := aviatrix.NewAviatrixTunnel(ctx, "testTunnel", &aviatrix.AviatrixTunnelArgs{
//				GwName1: pulumi.String("avtx-gw1"),
//				GwName2: pulumi.String("avtx-gw2"),
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
// **tunnel** can be imported using the `gw_name1` and `gw_name2`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixTunnel:AviatrixTunnel test gw_name1~gw_name2
//
// ```
type AviatrixTunnel struct {
	pulumi.CustomResourceState

	// Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false.
	EnableHa pulumi.BoolPtrOutput `pulumi:"enableHa"`
	// The first VPC Container name to make a peer pair.
	GwName1 pulumi.StringOutput `pulumi:"gwName1"`
	// The second VPC Container name to make a peer pair.
	GwName2 pulumi.StringOutput `pulumi:"gwName2"`
	// (Computed) Status of the HA tunnel.
	PeeringHastatus pulumi.StringOutput `pulumi:"peeringHastatus"`
	// (Computed) Name of the peering link.
	PeeringLink pulumi.StringOutput `pulumi:"peeringLink"`
	// (Computed) Status of the tunnel.
	PeeringState pulumi.StringOutput `pulumi:"peeringState"`
}

// NewAviatrixTunnel registers a new resource with the given unique name, arguments, and options.
func NewAviatrixTunnel(ctx *pulumi.Context,
	name string, args *AviatrixTunnelArgs, opts ...pulumi.ResourceOption) (*AviatrixTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GwName1 == nil {
		return nil, errors.New("invalid value for required argument 'GwName1'")
	}
	if args.GwName2 == nil {
		return nil, errors.New("invalid value for required argument 'GwName2'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixTunnel
	err := ctx.RegisterResource("aviatrix:index/aviatrixTunnel:AviatrixTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixTunnel gets an existing AviatrixTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixTunnelState, opts ...pulumi.ResourceOption) (*AviatrixTunnel, error) {
	var resource AviatrixTunnel
	err := ctx.ReadResource("aviatrix:index/aviatrixTunnel:AviatrixTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixTunnel resources.
type aviatrixTunnelState struct {
	// Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false.
	EnableHa *bool `pulumi:"enableHa"`
	// The first VPC Container name to make a peer pair.
	GwName1 *string `pulumi:"gwName1"`
	// The second VPC Container name to make a peer pair.
	GwName2 *string `pulumi:"gwName2"`
	// (Computed) Status of the HA tunnel.
	PeeringHastatus *string `pulumi:"peeringHastatus"`
	// (Computed) Name of the peering link.
	PeeringLink *string `pulumi:"peeringLink"`
	// (Computed) Status of the tunnel.
	PeeringState *string `pulumi:"peeringState"`
}

type AviatrixTunnelState struct {
	// Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false.
	EnableHa pulumi.BoolPtrInput
	// The first VPC Container name to make a peer pair.
	GwName1 pulumi.StringPtrInput
	// The second VPC Container name to make a peer pair.
	GwName2 pulumi.StringPtrInput
	// (Computed) Status of the HA tunnel.
	PeeringHastatus pulumi.StringPtrInput
	// (Computed) Name of the peering link.
	PeeringLink pulumi.StringPtrInput
	// (Computed) Status of the tunnel.
	PeeringState pulumi.StringPtrInput
}

func (AviatrixTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixTunnelState)(nil)).Elem()
}

type aviatrixTunnelArgs struct {
	// Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false.
	EnableHa *bool `pulumi:"enableHa"`
	// The first VPC Container name to make a peer pair.
	GwName1 string `pulumi:"gwName1"`
	// The second VPC Container name to make a peer pair.
	GwName2 string `pulumi:"gwName2"`
}

// The set of arguments for constructing a AviatrixTunnel resource.
type AviatrixTunnelArgs struct {
	// Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false.
	EnableHa pulumi.BoolPtrInput
	// The first VPC Container name to make a peer pair.
	GwName1 pulumi.StringInput
	// The second VPC Container name to make a peer pair.
	GwName2 pulumi.StringInput
}

func (AviatrixTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixTunnelArgs)(nil)).Elem()
}

type AviatrixTunnelInput interface {
	pulumi.Input

	ToAviatrixTunnelOutput() AviatrixTunnelOutput
	ToAviatrixTunnelOutputWithContext(ctx context.Context) AviatrixTunnelOutput
}

func (*AviatrixTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixTunnel)(nil)).Elem()
}

func (i *AviatrixTunnel) ToAviatrixTunnelOutput() AviatrixTunnelOutput {
	return i.ToAviatrixTunnelOutputWithContext(context.Background())
}

func (i *AviatrixTunnel) ToAviatrixTunnelOutputWithContext(ctx context.Context) AviatrixTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTunnelOutput)
}

// AviatrixTunnelArrayInput is an input type that accepts AviatrixTunnelArray and AviatrixTunnelArrayOutput values.
// You can construct a concrete instance of `AviatrixTunnelArrayInput` via:
//
//	AviatrixTunnelArray{ AviatrixTunnelArgs{...} }
type AviatrixTunnelArrayInput interface {
	pulumi.Input

	ToAviatrixTunnelArrayOutput() AviatrixTunnelArrayOutput
	ToAviatrixTunnelArrayOutputWithContext(context.Context) AviatrixTunnelArrayOutput
}

type AviatrixTunnelArray []AviatrixTunnelInput

func (AviatrixTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixTunnel)(nil)).Elem()
}

func (i AviatrixTunnelArray) ToAviatrixTunnelArrayOutput() AviatrixTunnelArrayOutput {
	return i.ToAviatrixTunnelArrayOutputWithContext(context.Background())
}

func (i AviatrixTunnelArray) ToAviatrixTunnelArrayOutputWithContext(ctx context.Context) AviatrixTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTunnelArrayOutput)
}

// AviatrixTunnelMapInput is an input type that accepts AviatrixTunnelMap and AviatrixTunnelMapOutput values.
// You can construct a concrete instance of `AviatrixTunnelMapInput` via:
//
//	AviatrixTunnelMap{ "key": AviatrixTunnelArgs{...} }
type AviatrixTunnelMapInput interface {
	pulumi.Input

	ToAviatrixTunnelMapOutput() AviatrixTunnelMapOutput
	ToAviatrixTunnelMapOutputWithContext(context.Context) AviatrixTunnelMapOutput
}

type AviatrixTunnelMap map[string]AviatrixTunnelInput

func (AviatrixTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixTunnel)(nil)).Elem()
}

func (i AviatrixTunnelMap) ToAviatrixTunnelMapOutput() AviatrixTunnelMapOutput {
	return i.ToAviatrixTunnelMapOutputWithContext(context.Background())
}

func (i AviatrixTunnelMap) ToAviatrixTunnelMapOutputWithContext(ctx context.Context) AviatrixTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixTunnelMapOutput)
}

type AviatrixTunnelOutput struct{ *pulumi.OutputState }

func (AviatrixTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixTunnel)(nil)).Elem()
}

func (o AviatrixTunnelOutput) ToAviatrixTunnelOutput() AviatrixTunnelOutput {
	return o
}

func (o AviatrixTunnelOutput) ToAviatrixTunnelOutputWithContext(ctx context.Context) AviatrixTunnelOutput {
	return o
}

// Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false.
func (o AviatrixTunnelOutput) EnableHa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixTunnel) pulumi.BoolPtrOutput { return v.EnableHa }).(pulumi.BoolPtrOutput)
}

// The first VPC Container name to make a peer pair.
func (o AviatrixTunnelOutput) GwName1() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTunnel) pulumi.StringOutput { return v.GwName1 }).(pulumi.StringOutput)
}

// The second VPC Container name to make a peer pair.
func (o AviatrixTunnelOutput) GwName2() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTunnel) pulumi.StringOutput { return v.GwName2 }).(pulumi.StringOutput)
}

// (Computed) Status of the HA tunnel.
func (o AviatrixTunnelOutput) PeeringHastatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTunnel) pulumi.StringOutput { return v.PeeringHastatus }).(pulumi.StringOutput)
}

// (Computed) Name of the peering link.
func (o AviatrixTunnelOutput) PeeringLink() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTunnel) pulumi.StringOutput { return v.PeeringLink }).(pulumi.StringOutput)
}

// (Computed) Status of the tunnel.
func (o AviatrixTunnelOutput) PeeringState() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixTunnel) pulumi.StringOutput { return v.PeeringState }).(pulumi.StringOutput)
}

type AviatrixTunnelArrayOutput struct{ *pulumi.OutputState }

func (AviatrixTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixTunnel)(nil)).Elem()
}

func (o AviatrixTunnelArrayOutput) ToAviatrixTunnelArrayOutput() AviatrixTunnelArrayOutput {
	return o
}

func (o AviatrixTunnelArrayOutput) ToAviatrixTunnelArrayOutputWithContext(ctx context.Context) AviatrixTunnelArrayOutput {
	return o
}

func (o AviatrixTunnelArrayOutput) Index(i pulumi.IntInput) AviatrixTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixTunnel {
		return vs[0].([]*AviatrixTunnel)[vs[1].(int)]
	}).(AviatrixTunnelOutput)
}

type AviatrixTunnelMapOutput struct{ *pulumi.OutputState }

func (AviatrixTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixTunnel)(nil)).Elem()
}

func (o AviatrixTunnelMapOutput) ToAviatrixTunnelMapOutput() AviatrixTunnelMapOutput {
	return o
}

func (o AviatrixTunnelMapOutput) ToAviatrixTunnelMapOutputWithContext(ctx context.Context) AviatrixTunnelMapOutput {
	return o
}

func (o AviatrixTunnelMapOutput) MapIndex(k pulumi.StringInput) AviatrixTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixTunnel {
		return vs[0].(map[string]*AviatrixTunnel)[vs[1].(string)]
	}).(AviatrixTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTunnelInput)(nil)).Elem(), &AviatrixTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTunnelArrayInput)(nil)).Elem(), AviatrixTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixTunnelMapInput)(nil)).Elem(), AviatrixTunnelMap{})
	pulumi.RegisterOutputType(AviatrixTunnelOutput{})
	pulumi.RegisterOutputType(AviatrixTunnelArrayOutput{})
	pulumi.RegisterOutputType(AviatrixTunnelMapOutput{})
}

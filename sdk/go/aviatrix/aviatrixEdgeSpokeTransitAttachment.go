// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_edge_spoke_transit_attachment** resource allows the creation and management of Aviatrix Edge as a Spoke to Transit gateway attachments. This resource is available as of provider version R2.23+.
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
//			_, err := aviatrix.NewAviatrixEdgeSpokeTransitAttachment(ctx, "testAttachment", &aviatrix.AviatrixEdgeSpokeTransitAttachmentArgs{
//				SpokeGwName:   pulumi.String("edge-as-a-spoke"),
//				TransitGwName: pulumi.String("transit-gw"),
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
// **spoke_transit_attachment** can be imported using the `spoke_gw_name` and `transit_gw_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixEdgeSpokeTransitAttachment:AviatrixEdgeSpokeTransitAttachment test spoke_gw_name~transit_gw_name
//
// ```
type AviatrixEdgeSpokeTransitAttachment struct {
	pulumi.CustomResourceState

	// Switch to enable insane mode. Valid values: true, false. Default: false.
	EnableInsaneMode pulumi.BoolPtrOutput `pulumi:"enableInsaneMode"`
	// Switch to enable jumbo frame. Valid values: true, false. Default: false.
	EnableJumboFrame pulumi.BoolPtrOutput `pulumi:"enableJumboFrame"`
	// Switch to enable over the private network. Valid values: true, false. Default: true.
	EnableOverPrivateNetwork pulumi.BoolPtrOutput `pulumi:"enableOverPrivateNetwork"`
	// Insane mode tunnel number. Default: 0.
	InsaneModeTunnelNumber pulumi.IntPtrOutput `pulumi:"insaneModeTunnelNumber"`
	// Name of the Edge as a Spoke to attach to transit network.
	SpokeGwName pulumi.StringOutput `pulumi:"spokeGwName"`
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.
	SpokePrependAsPaths pulumi.StringArrayOutput `pulumi:"spokePrependAsPaths"`
	// Name of the transit gateway to attach the Edge as a Spoke to.
	TransitGwName pulumi.StringOutput `pulumi:"transitGwName"`
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.
	TransitPrependAsPaths pulumi.StringArrayOutput `pulumi:"transitPrependAsPaths"`
}

// NewAviatrixEdgeSpokeTransitAttachment registers a new resource with the given unique name, arguments, and options.
func NewAviatrixEdgeSpokeTransitAttachment(ctx *pulumi.Context,
	name string, args *AviatrixEdgeSpokeTransitAttachmentArgs, opts ...pulumi.ResourceOption) (*AviatrixEdgeSpokeTransitAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SpokeGwName == nil {
		return nil, errors.New("invalid value for required argument 'SpokeGwName'")
	}
	if args.TransitGwName == nil {
		return nil, errors.New("invalid value for required argument 'TransitGwName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixEdgeSpokeTransitAttachment
	err := ctx.RegisterResource("aviatrix:index/aviatrixEdgeSpokeTransitAttachment:AviatrixEdgeSpokeTransitAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixEdgeSpokeTransitAttachment gets an existing AviatrixEdgeSpokeTransitAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixEdgeSpokeTransitAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixEdgeSpokeTransitAttachmentState, opts ...pulumi.ResourceOption) (*AviatrixEdgeSpokeTransitAttachment, error) {
	var resource AviatrixEdgeSpokeTransitAttachment
	err := ctx.ReadResource("aviatrix:index/aviatrixEdgeSpokeTransitAttachment:AviatrixEdgeSpokeTransitAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixEdgeSpokeTransitAttachment resources.
type aviatrixEdgeSpokeTransitAttachmentState struct {
	// Switch to enable insane mode. Valid values: true, false. Default: false.
	EnableInsaneMode *bool `pulumi:"enableInsaneMode"`
	// Switch to enable jumbo frame. Valid values: true, false. Default: false.
	EnableJumboFrame *bool `pulumi:"enableJumboFrame"`
	// Switch to enable over the private network. Valid values: true, false. Default: true.
	EnableOverPrivateNetwork *bool `pulumi:"enableOverPrivateNetwork"`
	// Insane mode tunnel number. Default: 0.
	InsaneModeTunnelNumber *int `pulumi:"insaneModeTunnelNumber"`
	// Name of the Edge as a Spoke to attach to transit network.
	SpokeGwName *string `pulumi:"spokeGwName"`
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.
	SpokePrependAsPaths []string `pulumi:"spokePrependAsPaths"`
	// Name of the transit gateway to attach the Edge as a Spoke to.
	TransitGwName *string `pulumi:"transitGwName"`
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.
	TransitPrependAsPaths []string `pulumi:"transitPrependAsPaths"`
}

type AviatrixEdgeSpokeTransitAttachmentState struct {
	// Switch to enable insane mode. Valid values: true, false. Default: false.
	EnableInsaneMode pulumi.BoolPtrInput
	// Switch to enable jumbo frame. Valid values: true, false. Default: false.
	EnableJumboFrame pulumi.BoolPtrInput
	// Switch to enable over the private network. Valid values: true, false. Default: true.
	EnableOverPrivateNetwork pulumi.BoolPtrInput
	// Insane mode tunnel number. Default: 0.
	InsaneModeTunnelNumber pulumi.IntPtrInput
	// Name of the Edge as a Spoke to attach to transit network.
	SpokeGwName pulumi.StringPtrInput
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.
	SpokePrependAsPaths pulumi.StringArrayInput
	// Name of the transit gateway to attach the Edge as a Spoke to.
	TransitGwName pulumi.StringPtrInput
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.
	TransitPrependAsPaths pulumi.StringArrayInput
}

func (AviatrixEdgeSpokeTransitAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixEdgeSpokeTransitAttachmentState)(nil)).Elem()
}

type aviatrixEdgeSpokeTransitAttachmentArgs struct {
	// Switch to enable insane mode. Valid values: true, false. Default: false.
	EnableInsaneMode *bool `pulumi:"enableInsaneMode"`
	// Switch to enable jumbo frame. Valid values: true, false. Default: false.
	EnableJumboFrame *bool `pulumi:"enableJumboFrame"`
	// Switch to enable over the private network. Valid values: true, false. Default: true.
	EnableOverPrivateNetwork *bool `pulumi:"enableOverPrivateNetwork"`
	// Insane mode tunnel number. Default: 0.
	InsaneModeTunnelNumber *int `pulumi:"insaneModeTunnelNumber"`
	// Name of the Edge as a Spoke to attach to transit network.
	SpokeGwName string `pulumi:"spokeGwName"`
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.
	SpokePrependAsPaths []string `pulumi:"spokePrependAsPaths"`
	// Name of the transit gateway to attach the Edge as a Spoke to.
	TransitGwName string `pulumi:"transitGwName"`
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.
	TransitPrependAsPaths []string `pulumi:"transitPrependAsPaths"`
}

// The set of arguments for constructing a AviatrixEdgeSpokeTransitAttachment resource.
type AviatrixEdgeSpokeTransitAttachmentArgs struct {
	// Switch to enable insane mode. Valid values: true, false. Default: false.
	EnableInsaneMode pulumi.BoolPtrInput
	// Switch to enable jumbo frame. Valid values: true, false. Default: false.
	EnableJumboFrame pulumi.BoolPtrInput
	// Switch to enable over the private network. Valid values: true, false. Default: true.
	EnableOverPrivateNetwork pulumi.BoolPtrInput
	// Insane mode tunnel number. Default: 0.
	InsaneModeTunnelNumber pulumi.IntPtrInput
	// Name of the Edge as a Spoke to attach to transit network.
	SpokeGwName pulumi.StringInput
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.
	SpokePrependAsPaths pulumi.StringArrayInput
	// Name of the transit gateway to attach the Edge as a Spoke to.
	TransitGwName pulumi.StringInput
	// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.
	TransitPrependAsPaths pulumi.StringArrayInput
}

func (AviatrixEdgeSpokeTransitAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixEdgeSpokeTransitAttachmentArgs)(nil)).Elem()
}

type AviatrixEdgeSpokeTransitAttachmentInput interface {
	pulumi.Input

	ToAviatrixEdgeSpokeTransitAttachmentOutput() AviatrixEdgeSpokeTransitAttachmentOutput
	ToAviatrixEdgeSpokeTransitAttachmentOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentOutput
}

func (*AviatrixEdgeSpokeTransitAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixEdgeSpokeTransitAttachment)(nil)).Elem()
}

func (i *AviatrixEdgeSpokeTransitAttachment) ToAviatrixEdgeSpokeTransitAttachmentOutput() AviatrixEdgeSpokeTransitAttachmentOutput {
	return i.ToAviatrixEdgeSpokeTransitAttachmentOutputWithContext(context.Background())
}

func (i *AviatrixEdgeSpokeTransitAttachment) ToAviatrixEdgeSpokeTransitAttachmentOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixEdgeSpokeTransitAttachmentOutput)
}

// AviatrixEdgeSpokeTransitAttachmentArrayInput is an input type that accepts AviatrixEdgeSpokeTransitAttachmentArray and AviatrixEdgeSpokeTransitAttachmentArrayOutput values.
// You can construct a concrete instance of `AviatrixEdgeSpokeTransitAttachmentArrayInput` via:
//
//	AviatrixEdgeSpokeTransitAttachmentArray{ AviatrixEdgeSpokeTransitAttachmentArgs{...} }
type AviatrixEdgeSpokeTransitAttachmentArrayInput interface {
	pulumi.Input

	ToAviatrixEdgeSpokeTransitAttachmentArrayOutput() AviatrixEdgeSpokeTransitAttachmentArrayOutput
	ToAviatrixEdgeSpokeTransitAttachmentArrayOutputWithContext(context.Context) AviatrixEdgeSpokeTransitAttachmentArrayOutput
}

type AviatrixEdgeSpokeTransitAttachmentArray []AviatrixEdgeSpokeTransitAttachmentInput

func (AviatrixEdgeSpokeTransitAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixEdgeSpokeTransitAttachment)(nil)).Elem()
}

func (i AviatrixEdgeSpokeTransitAttachmentArray) ToAviatrixEdgeSpokeTransitAttachmentArrayOutput() AviatrixEdgeSpokeTransitAttachmentArrayOutput {
	return i.ToAviatrixEdgeSpokeTransitAttachmentArrayOutputWithContext(context.Background())
}

func (i AviatrixEdgeSpokeTransitAttachmentArray) ToAviatrixEdgeSpokeTransitAttachmentArrayOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixEdgeSpokeTransitAttachmentArrayOutput)
}

// AviatrixEdgeSpokeTransitAttachmentMapInput is an input type that accepts AviatrixEdgeSpokeTransitAttachmentMap and AviatrixEdgeSpokeTransitAttachmentMapOutput values.
// You can construct a concrete instance of `AviatrixEdgeSpokeTransitAttachmentMapInput` via:
//
//	AviatrixEdgeSpokeTransitAttachmentMap{ "key": AviatrixEdgeSpokeTransitAttachmentArgs{...} }
type AviatrixEdgeSpokeTransitAttachmentMapInput interface {
	pulumi.Input

	ToAviatrixEdgeSpokeTransitAttachmentMapOutput() AviatrixEdgeSpokeTransitAttachmentMapOutput
	ToAviatrixEdgeSpokeTransitAttachmentMapOutputWithContext(context.Context) AviatrixEdgeSpokeTransitAttachmentMapOutput
}

type AviatrixEdgeSpokeTransitAttachmentMap map[string]AviatrixEdgeSpokeTransitAttachmentInput

func (AviatrixEdgeSpokeTransitAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixEdgeSpokeTransitAttachment)(nil)).Elem()
}

func (i AviatrixEdgeSpokeTransitAttachmentMap) ToAviatrixEdgeSpokeTransitAttachmentMapOutput() AviatrixEdgeSpokeTransitAttachmentMapOutput {
	return i.ToAviatrixEdgeSpokeTransitAttachmentMapOutputWithContext(context.Background())
}

func (i AviatrixEdgeSpokeTransitAttachmentMap) ToAviatrixEdgeSpokeTransitAttachmentMapOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixEdgeSpokeTransitAttachmentMapOutput)
}

type AviatrixEdgeSpokeTransitAttachmentOutput struct{ *pulumi.OutputState }

func (AviatrixEdgeSpokeTransitAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixEdgeSpokeTransitAttachment)(nil)).Elem()
}

func (o AviatrixEdgeSpokeTransitAttachmentOutput) ToAviatrixEdgeSpokeTransitAttachmentOutput() AviatrixEdgeSpokeTransitAttachmentOutput {
	return o
}

func (o AviatrixEdgeSpokeTransitAttachmentOutput) ToAviatrixEdgeSpokeTransitAttachmentOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentOutput {
	return o
}

// Switch to enable insane mode. Valid values: true, false. Default: false.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) EnableInsaneMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.BoolPtrOutput { return v.EnableInsaneMode }).(pulumi.BoolPtrOutput)
}

// Switch to enable jumbo frame. Valid values: true, false. Default: false.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) EnableJumboFrame() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.BoolPtrOutput { return v.EnableJumboFrame }).(pulumi.BoolPtrOutput)
}

// Switch to enable over the private network. Valid values: true, false. Default: true.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) EnableOverPrivateNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.BoolPtrOutput { return v.EnableOverPrivateNetwork }).(pulumi.BoolPtrOutput)
}

// Insane mode tunnel number. Default: 0.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) InsaneModeTunnelNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.IntPtrOutput { return v.InsaneModeTunnelNumber }).(pulumi.IntPtrOutput)
}

// Name of the Edge as a Spoke to attach to transit network.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) SpokeGwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.StringOutput { return v.SpokeGwName }).(pulumi.StringOutput)
}

// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) SpokePrependAsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.StringArrayOutput { return v.SpokePrependAsPaths }).(pulumi.StringArrayOutput)
}

// Name of the transit gateway to attach the Edge as a Spoke to.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) TransitGwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.StringOutput { return v.TransitGwName }).(pulumi.StringOutput)
}

// Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.
func (o AviatrixEdgeSpokeTransitAttachmentOutput) TransitPrependAsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeTransitAttachment) pulumi.StringArrayOutput { return v.TransitPrependAsPaths }).(pulumi.StringArrayOutput)
}

type AviatrixEdgeSpokeTransitAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AviatrixEdgeSpokeTransitAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixEdgeSpokeTransitAttachment)(nil)).Elem()
}

func (o AviatrixEdgeSpokeTransitAttachmentArrayOutput) ToAviatrixEdgeSpokeTransitAttachmentArrayOutput() AviatrixEdgeSpokeTransitAttachmentArrayOutput {
	return o
}

func (o AviatrixEdgeSpokeTransitAttachmentArrayOutput) ToAviatrixEdgeSpokeTransitAttachmentArrayOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentArrayOutput {
	return o
}

func (o AviatrixEdgeSpokeTransitAttachmentArrayOutput) Index(i pulumi.IntInput) AviatrixEdgeSpokeTransitAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixEdgeSpokeTransitAttachment {
		return vs[0].([]*AviatrixEdgeSpokeTransitAttachment)[vs[1].(int)]
	}).(AviatrixEdgeSpokeTransitAttachmentOutput)
}

type AviatrixEdgeSpokeTransitAttachmentMapOutput struct{ *pulumi.OutputState }

func (AviatrixEdgeSpokeTransitAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixEdgeSpokeTransitAttachment)(nil)).Elem()
}

func (o AviatrixEdgeSpokeTransitAttachmentMapOutput) ToAviatrixEdgeSpokeTransitAttachmentMapOutput() AviatrixEdgeSpokeTransitAttachmentMapOutput {
	return o
}

func (o AviatrixEdgeSpokeTransitAttachmentMapOutput) ToAviatrixEdgeSpokeTransitAttachmentMapOutputWithContext(ctx context.Context) AviatrixEdgeSpokeTransitAttachmentMapOutput {
	return o
}

func (o AviatrixEdgeSpokeTransitAttachmentMapOutput) MapIndex(k pulumi.StringInput) AviatrixEdgeSpokeTransitAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixEdgeSpokeTransitAttachment {
		return vs[0].(map[string]*AviatrixEdgeSpokeTransitAttachment)[vs[1].(string)]
	}).(AviatrixEdgeSpokeTransitAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixEdgeSpokeTransitAttachmentInput)(nil)).Elem(), &AviatrixEdgeSpokeTransitAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixEdgeSpokeTransitAttachmentArrayInput)(nil)).Elem(), AviatrixEdgeSpokeTransitAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixEdgeSpokeTransitAttachmentMapInput)(nil)).Elem(), AviatrixEdgeSpokeTransitAttachmentMap{})
	pulumi.RegisterOutputType(AviatrixEdgeSpokeTransitAttachmentOutput{})
	pulumi.RegisterOutputType(AviatrixEdgeSpokeTransitAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AviatrixEdgeSpokeTransitAttachmentMapOutput{})
}

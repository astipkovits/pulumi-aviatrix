// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_controller_bgp_max_as_limit_config** resource allows management of an Aviatrix Controller's BGP max AS limit for transit gateways. This resource is available as of provider version R2.18.1+.
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
//			_, err := aviatrix.NewAviatrixControllerBgpMaxAsLimitConfig(ctx, "testMaxAsLimit", &aviatrix.AviatrixControllerBgpMaxAsLimitConfigArgs{
//				MaxAsLimit: pulumi.Int(1),
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
// **aviatrix_controller_bgp_max_as_limit_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixControllerBgpMaxAsLimitConfig:AviatrixControllerBgpMaxAsLimitConfig test_max_as_limit 10-11-12-13
//
// ```
type AviatrixControllerBgpMaxAsLimitConfig struct {
	pulumi.CustomResourceState

	// The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254].
	MaxAsLimit pulumi.IntOutput `pulumi:"maxAsLimit"`
}

// NewAviatrixControllerBgpMaxAsLimitConfig registers a new resource with the given unique name, arguments, and options.
func NewAviatrixControllerBgpMaxAsLimitConfig(ctx *pulumi.Context,
	name string, args *AviatrixControllerBgpMaxAsLimitConfigArgs, opts ...pulumi.ResourceOption) (*AviatrixControllerBgpMaxAsLimitConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxAsLimit == nil {
		return nil, errors.New("invalid value for required argument 'MaxAsLimit'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixControllerBgpMaxAsLimitConfig
	err := ctx.RegisterResource("aviatrix:index/aviatrixControllerBgpMaxAsLimitConfig:AviatrixControllerBgpMaxAsLimitConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixControllerBgpMaxAsLimitConfig gets an existing AviatrixControllerBgpMaxAsLimitConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixControllerBgpMaxAsLimitConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixControllerBgpMaxAsLimitConfigState, opts ...pulumi.ResourceOption) (*AviatrixControllerBgpMaxAsLimitConfig, error) {
	var resource AviatrixControllerBgpMaxAsLimitConfig
	err := ctx.ReadResource("aviatrix:index/aviatrixControllerBgpMaxAsLimitConfig:AviatrixControllerBgpMaxAsLimitConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixControllerBgpMaxAsLimitConfig resources.
type aviatrixControllerBgpMaxAsLimitConfigState struct {
	// The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254].
	MaxAsLimit *int `pulumi:"maxAsLimit"`
}

type AviatrixControllerBgpMaxAsLimitConfigState struct {
	// The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254].
	MaxAsLimit pulumi.IntPtrInput
}

func (AviatrixControllerBgpMaxAsLimitConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerBgpMaxAsLimitConfigState)(nil)).Elem()
}

type aviatrixControllerBgpMaxAsLimitConfigArgs struct {
	// The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254].
	MaxAsLimit int `pulumi:"maxAsLimit"`
}

// The set of arguments for constructing a AviatrixControllerBgpMaxAsLimitConfig resource.
type AviatrixControllerBgpMaxAsLimitConfigArgs struct {
	// The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254].
	MaxAsLimit pulumi.IntInput
}

func (AviatrixControllerBgpMaxAsLimitConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerBgpMaxAsLimitConfigArgs)(nil)).Elem()
}

type AviatrixControllerBgpMaxAsLimitConfigInput interface {
	pulumi.Input

	ToAviatrixControllerBgpMaxAsLimitConfigOutput() AviatrixControllerBgpMaxAsLimitConfigOutput
	ToAviatrixControllerBgpMaxAsLimitConfigOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigOutput
}

func (*AviatrixControllerBgpMaxAsLimitConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerBgpMaxAsLimitConfig)(nil)).Elem()
}

func (i *AviatrixControllerBgpMaxAsLimitConfig) ToAviatrixControllerBgpMaxAsLimitConfigOutput() AviatrixControllerBgpMaxAsLimitConfigOutput {
	return i.ToAviatrixControllerBgpMaxAsLimitConfigOutputWithContext(context.Background())
}

func (i *AviatrixControllerBgpMaxAsLimitConfig) ToAviatrixControllerBgpMaxAsLimitConfigOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerBgpMaxAsLimitConfigOutput)
}

// AviatrixControllerBgpMaxAsLimitConfigArrayInput is an input type that accepts AviatrixControllerBgpMaxAsLimitConfigArray and AviatrixControllerBgpMaxAsLimitConfigArrayOutput values.
// You can construct a concrete instance of `AviatrixControllerBgpMaxAsLimitConfigArrayInput` via:
//
//	AviatrixControllerBgpMaxAsLimitConfigArray{ AviatrixControllerBgpMaxAsLimitConfigArgs{...} }
type AviatrixControllerBgpMaxAsLimitConfigArrayInput interface {
	pulumi.Input

	ToAviatrixControllerBgpMaxAsLimitConfigArrayOutput() AviatrixControllerBgpMaxAsLimitConfigArrayOutput
	ToAviatrixControllerBgpMaxAsLimitConfigArrayOutputWithContext(context.Context) AviatrixControllerBgpMaxAsLimitConfigArrayOutput
}

type AviatrixControllerBgpMaxAsLimitConfigArray []AviatrixControllerBgpMaxAsLimitConfigInput

func (AviatrixControllerBgpMaxAsLimitConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerBgpMaxAsLimitConfig)(nil)).Elem()
}

func (i AviatrixControllerBgpMaxAsLimitConfigArray) ToAviatrixControllerBgpMaxAsLimitConfigArrayOutput() AviatrixControllerBgpMaxAsLimitConfigArrayOutput {
	return i.ToAviatrixControllerBgpMaxAsLimitConfigArrayOutputWithContext(context.Background())
}

func (i AviatrixControllerBgpMaxAsLimitConfigArray) ToAviatrixControllerBgpMaxAsLimitConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerBgpMaxAsLimitConfigArrayOutput)
}

// AviatrixControllerBgpMaxAsLimitConfigMapInput is an input type that accepts AviatrixControllerBgpMaxAsLimitConfigMap and AviatrixControllerBgpMaxAsLimitConfigMapOutput values.
// You can construct a concrete instance of `AviatrixControllerBgpMaxAsLimitConfigMapInput` via:
//
//	AviatrixControllerBgpMaxAsLimitConfigMap{ "key": AviatrixControllerBgpMaxAsLimitConfigArgs{...} }
type AviatrixControllerBgpMaxAsLimitConfigMapInput interface {
	pulumi.Input

	ToAviatrixControllerBgpMaxAsLimitConfigMapOutput() AviatrixControllerBgpMaxAsLimitConfigMapOutput
	ToAviatrixControllerBgpMaxAsLimitConfigMapOutputWithContext(context.Context) AviatrixControllerBgpMaxAsLimitConfigMapOutput
}

type AviatrixControllerBgpMaxAsLimitConfigMap map[string]AviatrixControllerBgpMaxAsLimitConfigInput

func (AviatrixControllerBgpMaxAsLimitConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerBgpMaxAsLimitConfig)(nil)).Elem()
}

func (i AviatrixControllerBgpMaxAsLimitConfigMap) ToAviatrixControllerBgpMaxAsLimitConfigMapOutput() AviatrixControllerBgpMaxAsLimitConfigMapOutput {
	return i.ToAviatrixControllerBgpMaxAsLimitConfigMapOutputWithContext(context.Background())
}

func (i AviatrixControllerBgpMaxAsLimitConfigMap) ToAviatrixControllerBgpMaxAsLimitConfigMapOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerBgpMaxAsLimitConfigMapOutput)
}

type AviatrixControllerBgpMaxAsLimitConfigOutput struct{ *pulumi.OutputState }

func (AviatrixControllerBgpMaxAsLimitConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerBgpMaxAsLimitConfig)(nil)).Elem()
}

func (o AviatrixControllerBgpMaxAsLimitConfigOutput) ToAviatrixControllerBgpMaxAsLimitConfigOutput() AviatrixControllerBgpMaxAsLimitConfigOutput {
	return o
}

func (o AviatrixControllerBgpMaxAsLimitConfigOutput) ToAviatrixControllerBgpMaxAsLimitConfigOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigOutput {
	return o
}

// The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254].
func (o AviatrixControllerBgpMaxAsLimitConfigOutput) MaxAsLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *AviatrixControllerBgpMaxAsLimitConfig) pulumi.IntOutput { return v.MaxAsLimit }).(pulumi.IntOutput)
}

type AviatrixControllerBgpMaxAsLimitConfigArrayOutput struct{ *pulumi.OutputState }

func (AviatrixControllerBgpMaxAsLimitConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerBgpMaxAsLimitConfig)(nil)).Elem()
}

func (o AviatrixControllerBgpMaxAsLimitConfigArrayOutput) ToAviatrixControllerBgpMaxAsLimitConfigArrayOutput() AviatrixControllerBgpMaxAsLimitConfigArrayOutput {
	return o
}

func (o AviatrixControllerBgpMaxAsLimitConfigArrayOutput) ToAviatrixControllerBgpMaxAsLimitConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigArrayOutput {
	return o
}

func (o AviatrixControllerBgpMaxAsLimitConfigArrayOutput) Index(i pulumi.IntInput) AviatrixControllerBgpMaxAsLimitConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixControllerBgpMaxAsLimitConfig {
		return vs[0].([]*AviatrixControllerBgpMaxAsLimitConfig)[vs[1].(int)]
	}).(AviatrixControllerBgpMaxAsLimitConfigOutput)
}

type AviatrixControllerBgpMaxAsLimitConfigMapOutput struct{ *pulumi.OutputState }

func (AviatrixControllerBgpMaxAsLimitConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerBgpMaxAsLimitConfig)(nil)).Elem()
}

func (o AviatrixControllerBgpMaxAsLimitConfigMapOutput) ToAviatrixControllerBgpMaxAsLimitConfigMapOutput() AviatrixControllerBgpMaxAsLimitConfigMapOutput {
	return o
}

func (o AviatrixControllerBgpMaxAsLimitConfigMapOutput) ToAviatrixControllerBgpMaxAsLimitConfigMapOutputWithContext(ctx context.Context) AviatrixControllerBgpMaxAsLimitConfigMapOutput {
	return o
}

func (o AviatrixControllerBgpMaxAsLimitConfigMapOutput) MapIndex(k pulumi.StringInput) AviatrixControllerBgpMaxAsLimitConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixControllerBgpMaxAsLimitConfig {
		return vs[0].(map[string]*AviatrixControllerBgpMaxAsLimitConfig)[vs[1].(string)]
	}).(AviatrixControllerBgpMaxAsLimitConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerBgpMaxAsLimitConfigInput)(nil)).Elem(), &AviatrixControllerBgpMaxAsLimitConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerBgpMaxAsLimitConfigArrayInput)(nil)).Elem(), AviatrixControllerBgpMaxAsLimitConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerBgpMaxAsLimitConfigMapInput)(nil)).Elem(), AviatrixControllerBgpMaxAsLimitConfigMap{})
	pulumi.RegisterOutputType(AviatrixControllerBgpMaxAsLimitConfigOutput{})
	pulumi.RegisterOutputType(AviatrixControllerBgpMaxAsLimitConfigArrayOutput{})
	pulumi.RegisterOutputType(AviatrixControllerBgpMaxAsLimitConfigMapOutput{})
}

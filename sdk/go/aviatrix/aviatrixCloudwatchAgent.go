// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixCloudwatchAgent struct {
	pulumi.CustomResourceState

	// CloudWatch role ARN.
	CloudwatchRoleArn pulumi.StringOutput `pulumi:"cloudwatchRoleArn"`
	// List of excluded gateways.
	ExcludedGateways pulumi.StringArrayOutput `pulumi:"excludedGateways"`
	// Log group name.
	LogGroupName pulumi.StringPtrOutput `pulumi:"logGroupName"`
	// Name of AWS region.
	Region pulumi.StringOutput `pulumi:"region"`
	// Enabled or not.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAviatrixCloudwatchAgent registers a new resource with the given unique name, arguments, and options.
func NewAviatrixCloudwatchAgent(ctx *pulumi.Context,
	name string, args *AviatrixCloudwatchAgentArgs, opts ...pulumi.ResourceOption) (*AviatrixCloudwatchAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudwatchRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'CloudwatchRoleArn'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixCloudwatchAgent
	err := ctx.RegisterResource("aviatrix:index/aviatrixCloudwatchAgent:AviatrixCloudwatchAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixCloudwatchAgent gets an existing AviatrixCloudwatchAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixCloudwatchAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixCloudwatchAgentState, opts ...pulumi.ResourceOption) (*AviatrixCloudwatchAgent, error) {
	var resource AviatrixCloudwatchAgent
	err := ctx.ReadResource("aviatrix:index/aviatrixCloudwatchAgent:AviatrixCloudwatchAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixCloudwatchAgent resources.
type aviatrixCloudwatchAgentState struct {
	// CloudWatch role ARN.
	CloudwatchRoleArn *string `pulumi:"cloudwatchRoleArn"`
	// List of excluded gateways.
	ExcludedGateways []string `pulumi:"excludedGateways"`
	// Log group name.
	LogGroupName *string `pulumi:"logGroupName"`
	// Name of AWS region.
	Region *string `pulumi:"region"`
	// Enabled or not.
	Status *string `pulumi:"status"`
}

type AviatrixCloudwatchAgentState struct {
	// CloudWatch role ARN.
	CloudwatchRoleArn pulumi.StringPtrInput
	// List of excluded gateways.
	ExcludedGateways pulumi.StringArrayInput
	// Log group name.
	LogGroupName pulumi.StringPtrInput
	// Name of AWS region.
	Region pulumi.StringPtrInput
	// Enabled or not.
	Status pulumi.StringPtrInput
}

func (AviatrixCloudwatchAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixCloudwatchAgentState)(nil)).Elem()
}

type aviatrixCloudwatchAgentArgs struct {
	// CloudWatch role ARN.
	CloudwatchRoleArn string `pulumi:"cloudwatchRoleArn"`
	// List of excluded gateways.
	ExcludedGateways []string `pulumi:"excludedGateways"`
	// Log group name.
	LogGroupName *string `pulumi:"logGroupName"`
	// Name of AWS region.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a AviatrixCloudwatchAgent resource.
type AviatrixCloudwatchAgentArgs struct {
	// CloudWatch role ARN.
	CloudwatchRoleArn pulumi.StringInput
	// List of excluded gateways.
	ExcludedGateways pulumi.StringArrayInput
	// Log group name.
	LogGroupName pulumi.StringPtrInput
	// Name of AWS region.
	Region pulumi.StringInput
}

func (AviatrixCloudwatchAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixCloudwatchAgentArgs)(nil)).Elem()
}

type AviatrixCloudwatchAgentInput interface {
	pulumi.Input

	ToAviatrixCloudwatchAgentOutput() AviatrixCloudwatchAgentOutput
	ToAviatrixCloudwatchAgentOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentOutput
}

func (*AviatrixCloudwatchAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixCloudwatchAgent)(nil)).Elem()
}

func (i *AviatrixCloudwatchAgent) ToAviatrixCloudwatchAgentOutput() AviatrixCloudwatchAgentOutput {
	return i.ToAviatrixCloudwatchAgentOutputWithContext(context.Background())
}

func (i *AviatrixCloudwatchAgent) ToAviatrixCloudwatchAgentOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixCloudwatchAgentOutput)
}

// AviatrixCloudwatchAgentArrayInput is an input type that accepts AviatrixCloudwatchAgentArray and AviatrixCloudwatchAgentArrayOutput values.
// You can construct a concrete instance of `AviatrixCloudwatchAgentArrayInput` via:
//
//	AviatrixCloudwatchAgentArray{ AviatrixCloudwatchAgentArgs{...} }
type AviatrixCloudwatchAgentArrayInput interface {
	pulumi.Input

	ToAviatrixCloudwatchAgentArrayOutput() AviatrixCloudwatchAgentArrayOutput
	ToAviatrixCloudwatchAgentArrayOutputWithContext(context.Context) AviatrixCloudwatchAgentArrayOutput
}

type AviatrixCloudwatchAgentArray []AviatrixCloudwatchAgentInput

func (AviatrixCloudwatchAgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixCloudwatchAgent)(nil)).Elem()
}

func (i AviatrixCloudwatchAgentArray) ToAviatrixCloudwatchAgentArrayOutput() AviatrixCloudwatchAgentArrayOutput {
	return i.ToAviatrixCloudwatchAgentArrayOutputWithContext(context.Background())
}

func (i AviatrixCloudwatchAgentArray) ToAviatrixCloudwatchAgentArrayOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixCloudwatchAgentArrayOutput)
}

// AviatrixCloudwatchAgentMapInput is an input type that accepts AviatrixCloudwatchAgentMap and AviatrixCloudwatchAgentMapOutput values.
// You can construct a concrete instance of `AviatrixCloudwatchAgentMapInput` via:
//
//	AviatrixCloudwatchAgentMap{ "key": AviatrixCloudwatchAgentArgs{...} }
type AviatrixCloudwatchAgentMapInput interface {
	pulumi.Input

	ToAviatrixCloudwatchAgentMapOutput() AviatrixCloudwatchAgentMapOutput
	ToAviatrixCloudwatchAgentMapOutputWithContext(context.Context) AviatrixCloudwatchAgentMapOutput
}

type AviatrixCloudwatchAgentMap map[string]AviatrixCloudwatchAgentInput

func (AviatrixCloudwatchAgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixCloudwatchAgent)(nil)).Elem()
}

func (i AviatrixCloudwatchAgentMap) ToAviatrixCloudwatchAgentMapOutput() AviatrixCloudwatchAgentMapOutput {
	return i.ToAviatrixCloudwatchAgentMapOutputWithContext(context.Background())
}

func (i AviatrixCloudwatchAgentMap) ToAviatrixCloudwatchAgentMapOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixCloudwatchAgentMapOutput)
}

type AviatrixCloudwatchAgentOutput struct{ *pulumi.OutputState }

func (AviatrixCloudwatchAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixCloudwatchAgent)(nil)).Elem()
}

func (o AviatrixCloudwatchAgentOutput) ToAviatrixCloudwatchAgentOutput() AviatrixCloudwatchAgentOutput {
	return o
}

func (o AviatrixCloudwatchAgentOutput) ToAviatrixCloudwatchAgentOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentOutput {
	return o
}

// CloudWatch role ARN.
func (o AviatrixCloudwatchAgentOutput) CloudwatchRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudwatchAgent) pulumi.StringOutput { return v.CloudwatchRoleArn }).(pulumi.StringOutput)
}

// List of excluded gateways.
func (o AviatrixCloudwatchAgentOutput) ExcludedGateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixCloudwatchAgent) pulumi.StringArrayOutput { return v.ExcludedGateways }).(pulumi.StringArrayOutput)
}

// Log group name.
func (o AviatrixCloudwatchAgentOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixCloudwatchAgent) pulumi.StringPtrOutput { return v.LogGroupName }).(pulumi.StringPtrOutput)
}

// Name of AWS region.
func (o AviatrixCloudwatchAgentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudwatchAgent) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Enabled or not.
func (o AviatrixCloudwatchAgentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudwatchAgent) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AviatrixCloudwatchAgentArrayOutput struct{ *pulumi.OutputState }

func (AviatrixCloudwatchAgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixCloudwatchAgent)(nil)).Elem()
}

func (o AviatrixCloudwatchAgentArrayOutput) ToAviatrixCloudwatchAgentArrayOutput() AviatrixCloudwatchAgentArrayOutput {
	return o
}

func (o AviatrixCloudwatchAgentArrayOutput) ToAviatrixCloudwatchAgentArrayOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentArrayOutput {
	return o
}

func (o AviatrixCloudwatchAgentArrayOutput) Index(i pulumi.IntInput) AviatrixCloudwatchAgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixCloudwatchAgent {
		return vs[0].([]*AviatrixCloudwatchAgent)[vs[1].(int)]
	}).(AviatrixCloudwatchAgentOutput)
}

type AviatrixCloudwatchAgentMapOutput struct{ *pulumi.OutputState }

func (AviatrixCloudwatchAgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixCloudwatchAgent)(nil)).Elem()
}

func (o AviatrixCloudwatchAgentMapOutput) ToAviatrixCloudwatchAgentMapOutput() AviatrixCloudwatchAgentMapOutput {
	return o
}

func (o AviatrixCloudwatchAgentMapOutput) ToAviatrixCloudwatchAgentMapOutputWithContext(ctx context.Context) AviatrixCloudwatchAgentMapOutput {
	return o
}

func (o AviatrixCloudwatchAgentMapOutput) MapIndex(k pulumi.StringInput) AviatrixCloudwatchAgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixCloudwatchAgent {
		return vs[0].(map[string]*AviatrixCloudwatchAgent)[vs[1].(string)]
	}).(AviatrixCloudwatchAgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixCloudwatchAgentInput)(nil)).Elem(), &AviatrixCloudwatchAgent{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixCloudwatchAgentArrayInput)(nil)).Elem(), AviatrixCloudwatchAgentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixCloudwatchAgentMapInput)(nil)).Elem(), AviatrixCloudwatchAgentMap{})
	pulumi.RegisterOutputType(AviatrixCloudwatchAgentOutput{})
	pulumi.RegisterOutputType(AviatrixCloudwatchAgentArrayOutput{})
	pulumi.RegisterOutputType(AviatrixCloudwatchAgentMapOutput{})
}
// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_datadog_agent** resource allows the enabling and disabling of datadog agent.
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
//			_, err := aviatrix.NewAviatrixDatadogAgent(ctx, "testDatadogAgent", &aviatrix.AviatrixDatadogAgentArgs{
//				ApiKey: pulumi.String("your_api_key"),
//				ExcludedGateways: pulumi.StringArray{
//					pulumi.String("a"),
//					pulumi.String("b"),
//				},
//				Site: pulumi.String("datadoghq.com"),
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
// **datadog_agent** can be imported using "datadog_agent", e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent test datadog_agent
//
// ```
type AviatrixDatadogAgent struct {
	pulumi.CustomResourceState

	// API key.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
	ExcludedGateways pulumi.StringArrayOutput `pulumi:"excludedGateways"`
	// Only export metrics without exporting logs. False by default.
	MetricsOnly pulumi.BoolPtrOutput `pulumi:"metricsOnly"`
	// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
	Site pulumi.StringPtrOutput `pulumi:"site"`
	// The status of datadog agent.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAviatrixDatadogAgent registers a new resource with the given unique name, arguments, and options.
func NewAviatrixDatadogAgent(ctx *pulumi.Context,
	name string, args *AviatrixDatadogAgentArgs, opts ...pulumi.ResourceOption) (*AviatrixDatadogAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixDatadogAgent
	err := ctx.RegisterResource("aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixDatadogAgent gets an existing AviatrixDatadogAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixDatadogAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixDatadogAgentState, opts ...pulumi.ResourceOption) (*AviatrixDatadogAgent, error) {
	var resource AviatrixDatadogAgent
	err := ctx.ReadResource("aviatrix:index/aviatrixDatadogAgent:AviatrixDatadogAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixDatadogAgent resources.
type aviatrixDatadogAgentState struct {
	// API key.
	ApiKey *string `pulumi:"apiKey"`
	// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
	ExcludedGateways []string `pulumi:"excludedGateways"`
	// Only export metrics without exporting logs. False by default.
	MetricsOnly *bool `pulumi:"metricsOnly"`
	// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
	Site *string `pulumi:"site"`
	// The status of datadog agent.
	Status *string `pulumi:"status"`
}

type AviatrixDatadogAgentState struct {
	// API key.
	ApiKey pulumi.StringPtrInput
	// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
	ExcludedGateways pulumi.StringArrayInput
	// Only export metrics without exporting logs. False by default.
	MetricsOnly pulumi.BoolPtrInput
	// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
	Site pulumi.StringPtrInput
	// The status of datadog agent.
	Status pulumi.StringPtrInput
}

func (AviatrixDatadogAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixDatadogAgentState)(nil)).Elem()
}

type aviatrixDatadogAgentArgs struct {
	// API key.
	ApiKey string `pulumi:"apiKey"`
	// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
	ExcludedGateways []string `pulumi:"excludedGateways"`
	// Only export metrics without exporting logs. False by default.
	MetricsOnly *bool `pulumi:"metricsOnly"`
	// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
	Site *string `pulumi:"site"`
}

// The set of arguments for constructing a AviatrixDatadogAgent resource.
type AviatrixDatadogAgentArgs struct {
	// API key.
	ApiKey pulumi.StringInput
	// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
	ExcludedGateways pulumi.StringArrayInput
	// Only export metrics without exporting logs. False by default.
	MetricsOnly pulumi.BoolPtrInput
	// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
	Site pulumi.StringPtrInput
}

func (AviatrixDatadogAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixDatadogAgentArgs)(nil)).Elem()
}

type AviatrixDatadogAgentInput interface {
	pulumi.Input

	ToAviatrixDatadogAgentOutput() AviatrixDatadogAgentOutput
	ToAviatrixDatadogAgentOutputWithContext(ctx context.Context) AviatrixDatadogAgentOutput
}

func (*AviatrixDatadogAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixDatadogAgent)(nil)).Elem()
}

func (i *AviatrixDatadogAgent) ToAviatrixDatadogAgentOutput() AviatrixDatadogAgentOutput {
	return i.ToAviatrixDatadogAgentOutputWithContext(context.Background())
}

func (i *AviatrixDatadogAgent) ToAviatrixDatadogAgentOutputWithContext(ctx context.Context) AviatrixDatadogAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixDatadogAgentOutput)
}

// AviatrixDatadogAgentArrayInput is an input type that accepts AviatrixDatadogAgentArray and AviatrixDatadogAgentArrayOutput values.
// You can construct a concrete instance of `AviatrixDatadogAgentArrayInput` via:
//
//	AviatrixDatadogAgentArray{ AviatrixDatadogAgentArgs{...} }
type AviatrixDatadogAgentArrayInput interface {
	pulumi.Input

	ToAviatrixDatadogAgentArrayOutput() AviatrixDatadogAgentArrayOutput
	ToAviatrixDatadogAgentArrayOutputWithContext(context.Context) AviatrixDatadogAgentArrayOutput
}

type AviatrixDatadogAgentArray []AviatrixDatadogAgentInput

func (AviatrixDatadogAgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixDatadogAgent)(nil)).Elem()
}

func (i AviatrixDatadogAgentArray) ToAviatrixDatadogAgentArrayOutput() AviatrixDatadogAgentArrayOutput {
	return i.ToAviatrixDatadogAgentArrayOutputWithContext(context.Background())
}

func (i AviatrixDatadogAgentArray) ToAviatrixDatadogAgentArrayOutputWithContext(ctx context.Context) AviatrixDatadogAgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixDatadogAgentArrayOutput)
}

// AviatrixDatadogAgentMapInput is an input type that accepts AviatrixDatadogAgentMap and AviatrixDatadogAgentMapOutput values.
// You can construct a concrete instance of `AviatrixDatadogAgentMapInput` via:
//
//	AviatrixDatadogAgentMap{ "key": AviatrixDatadogAgentArgs{...} }
type AviatrixDatadogAgentMapInput interface {
	pulumi.Input

	ToAviatrixDatadogAgentMapOutput() AviatrixDatadogAgentMapOutput
	ToAviatrixDatadogAgentMapOutputWithContext(context.Context) AviatrixDatadogAgentMapOutput
}

type AviatrixDatadogAgentMap map[string]AviatrixDatadogAgentInput

func (AviatrixDatadogAgentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixDatadogAgent)(nil)).Elem()
}

func (i AviatrixDatadogAgentMap) ToAviatrixDatadogAgentMapOutput() AviatrixDatadogAgentMapOutput {
	return i.ToAviatrixDatadogAgentMapOutputWithContext(context.Background())
}

func (i AviatrixDatadogAgentMap) ToAviatrixDatadogAgentMapOutputWithContext(ctx context.Context) AviatrixDatadogAgentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixDatadogAgentMapOutput)
}

type AviatrixDatadogAgentOutput struct{ *pulumi.OutputState }

func (AviatrixDatadogAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixDatadogAgent)(nil)).Elem()
}

func (o AviatrixDatadogAgentOutput) ToAviatrixDatadogAgentOutput() AviatrixDatadogAgentOutput {
	return o
}

func (o AviatrixDatadogAgentOutput) ToAviatrixDatadogAgentOutputWithContext(ctx context.Context) AviatrixDatadogAgentOutput {
	return o
}

// API key.
func (o AviatrixDatadogAgentOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixDatadogAgent) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

// List of gateways to be excluded from logging. e.g.: ["gateway01", "gateway02", "gateway01-hagw"].
func (o AviatrixDatadogAgentOutput) ExcludedGateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixDatadogAgent) pulumi.StringArrayOutput { return v.ExcludedGateways }).(pulumi.StringArrayOutput)
}

// Only export metrics without exporting logs. False by default.
func (o AviatrixDatadogAgentOutput) MetricsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixDatadogAgent) pulumi.BoolPtrOutput { return v.MetricsOnly }).(pulumi.BoolPtrOutput)
}

// Site preference ("datadoghq.com" or" datadoghq.eu"). "datadoghq.com" by default.
func (o AviatrixDatadogAgentOutput) Site() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixDatadogAgent) pulumi.StringPtrOutput { return v.Site }).(pulumi.StringPtrOutput)
}

// The status of datadog agent.
func (o AviatrixDatadogAgentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixDatadogAgent) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AviatrixDatadogAgentArrayOutput struct{ *pulumi.OutputState }

func (AviatrixDatadogAgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixDatadogAgent)(nil)).Elem()
}

func (o AviatrixDatadogAgentArrayOutput) ToAviatrixDatadogAgentArrayOutput() AviatrixDatadogAgentArrayOutput {
	return o
}

func (o AviatrixDatadogAgentArrayOutput) ToAviatrixDatadogAgentArrayOutputWithContext(ctx context.Context) AviatrixDatadogAgentArrayOutput {
	return o
}

func (o AviatrixDatadogAgentArrayOutput) Index(i pulumi.IntInput) AviatrixDatadogAgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixDatadogAgent {
		return vs[0].([]*AviatrixDatadogAgent)[vs[1].(int)]
	}).(AviatrixDatadogAgentOutput)
}

type AviatrixDatadogAgentMapOutput struct{ *pulumi.OutputState }

func (AviatrixDatadogAgentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixDatadogAgent)(nil)).Elem()
}

func (o AviatrixDatadogAgentMapOutput) ToAviatrixDatadogAgentMapOutput() AviatrixDatadogAgentMapOutput {
	return o
}

func (o AviatrixDatadogAgentMapOutput) ToAviatrixDatadogAgentMapOutputWithContext(ctx context.Context) AviatrixDatadogAgentMapOutput {
	return o
}

func (o AviatrixDatadogAgentMapOutput) MapIndex(k pulumi.StringInput) AviatrixDatadogAgentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixDatadogAgent {
		return vs[0].(map[string]*AviatrixDatadogAgent)[vs[1].(string)]
	}).(AviatrixDatadogAgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixDatadogAgentInput)(nil)).Elem(), &AviatrixDatadogAgent{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixDatadogAgentArrayInput)(nil)).Elem(), AviatrixDatadogAgentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixDatadogAgentMapInput)(nil)).Elem(), AviatrixDatadogAgentMap{})
	pulumi.RegisterOutputType(AviatrixDatadogAgentOutput{})
	pulumi.RegisterOutputType(AviatrixDatadogAgentArrayOutput{})
	pulumi.RegisterOutputType(AviatrixDatadogAgentMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_proxy_config** resource allows management of an Aviatrix Controller's proxy configurations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/astipkovits/pulumi-aviatrix/sdk/go/aviatrix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aviatrix.NewAviatrixProxyConfig(ctx, "testProxyConfig", &aviatrix.AviatrixProxyConfigArgs{
//				HttpProxy:          pulumi.String("172.31.52.145:3127"),
//				HttpsProxy:         pulumi.String("172.31.52.145:3129"),
//				ProxyCaCertificate: readFileOrPanic("/path/to/ca.pem"),
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
// **controller_proxy_config** can be imported using controller IP, e.g. controller IP is 10.11.12.13
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig test 10-11-12-13
//
// ```
type AviatrixProxyConfig struct {
	pulumi.CustomResourceState

	// Http proxy URL.
	HttpProxy pulumi.StringOutput `pulumi:"httpProxy"`
	// Https proxy URL.
	HttpsProxy pulumi.StringOutput `pulumi:"httpsProxy"`
	// Server CA Certificate file. Use the `file` function to read from a file.
	ProxyCaCertificate pulumi.StringPtrOutput `pulumi:"proxyCaCertificate"`
}

// NewAviatrixProxyConfig registers a new resource with the given unique name, arguments, and options.
func NewAviatrixProxyConfig(ctx *pulumi.Context,
	name string, args *AviatrixProxyConfigArgs, opts ...pulumi.ResourceOption) (*AviatrixProxyConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HttpProxy == nil {
		return nil, errors.New("invalid value for required argument 'HttpProxy'")
	}
	if args.HttpsProxy == nil {
		return nil, errors.New("invalid value for required argument 'HttpsProxy'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixProxyConfig
	err := ctx.RegisterResource("aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixProxyConfig gets an existing AviatrixProxyConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixProxyConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixProxyConfigState, opts ...pulumi.ResourceOption) (*AviatrixProxyConfig, error) {
	var resource AviatrixProxyConfig
	err := ctx.ReadResource("aviatrix:index/aviatrixProxyConfig:AviatrixProxyConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixProxyConfig resources.
type aviatrixProxyConfigState struct {
	// Http proxy URL.
	HttpProxy *string `pulumi:"httpProxy"`
	// Https proxy URL.
	HttpsProxy *string `pulumi:"httpsProxy"`
	// Server CA Certificate file. Use the `file` function to read from a file.
	ProxyCaCertificate *string `pulumi:"proxyCaCertificate"`
}

type AviatrixProxyConfigState struct {
	// Http proxy URL.
	HttpProxy pulumi.StringPtrInput
	// Https proxy URL.
	HttpsProxy pulumi.StringPtrInput
	// Server CA Certificate file. Use the `file` function to read from a file.
	ProxyCaCertificate pulumi.StringPtrInput
}

func (AviatrixProxyConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixProxyConfigState)(nil)).Elem()
}

type aviatrixProxyConfigArgs struct {
	// Http proxy URL.
	HttpProxy string `pulumi:"httpProxy"`
	// Https proxy URL.
	HttpsProxy string `pulumi:"httpsProxy"`
	// Server CA Certificate file. Use the `file` function to read from a file.
	ProxyCaCertificate *string `pulumi:"proxyCaCertificate"`
}

// The set of arguments for constructing a AviatrixProxyConfig resource.
type AviatrixProxyConfigArgs struct {
	// Http proxy URL.
	HttpProxy pulumi.StringInput
	// Https proxy URL.
	HttpsProxy pulumi.StringInput
	// Server CA Certificate file. Use the `file` function to read from a file.
	ProxyCaCertificate pulumi.StringPtrInput
}

func (AviatrixProxyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixProxyConfigArgs)(nil)).Elem()
}

type AviatrixProxyConfigInput interface {
	pulumi.Input

	ToAviatrixProxyConfigOutput() AviatrixProxyConfigOutput
	ToAviatrixProxyConfigOutputWithContext(ctx context.Context) AviatrixProxyConfigOutput
}

func (*AviatrixProxyConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixProxyConfig)(nil)).Elem()
}

func (i *AviatrixProxyConfig) ToAviatrixProxyConfigOutput() AviatrixProxyConfigOutput {
	return i.ToAviatrixProxyConfigOutputWithContext(context.Background())
}

func (i *AviatrixProxyConfig) ToAviatrixProxyConfigOutputWithContext(ctx context.Context) AviatrixProxyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixProxyConfigOutput)
}

// AviatrixProxyConfigArrayInput is an input type that accepts AviatrixProxyConfigArray and AviatrixProxyConfigArrayOutput values.
// You can construct a concrete instance of `AviatrixProxyConfigArrayInput` via:
//
//	AviatrixProxyConfigArray{ AviatrixProxyConfigArgs{...} }
type AviatrixProxyConfigArrayInput interface {
	pulumi.Input

	ToAviatrixProxyConfigArrayOutput() AviatrixProxyConfigArrayOutput
	ToAviatrixProxyConfigArrayOutputWithContext(context.Context) AviatrixProxyConfigArrayOutput
}

type AviatrixProxyConfigArray []AviatrixProxyConfigInput

func (AviatrixProxyConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixProxyConfig)(nil)).Elem()
}

func (i AviatrixProxyConfigArray) ToAviatrixProxyConfigArrayOutput() AviatrixProxyConfigArrayOutput {
	return i.ToAviatrixProxyConfigArrayOutputWithContext(context.Background())
}

func (i AviatrixProxyConfigArray) ToAviatrixProxyConfigArrayOutputWithContext(ctx context.Context) AviatrixProxyConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixProxyConfigArrayOutput)
}

// AviatrixProxyConfigMapInput is an input type that accepts AviatrixProxyConfigMap and AviatrixProxyConfigMapOutput values.
// You can construct a concrete instance of `AviatrixProxyConfigMapInput` via:
//
//	AviatrixProxyConfigMap{ "key": AviatrixProxyConfigArgs{...} }
type AviatrixProxyConfigMapInput interface {
	pulumi.Input

	ToAviatrixProxyConfigMapOutput() AviatrixProxyConfigMapOutput
	ToAviatrixProxyConfigMapOutputWithContext(context.Context) AviatrixProxyConfigMapOutput
}

type AviatrixProxyConfigMap map[string]AviatrixProxyConfigInput

func (AviatrixProxyConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixProxyConfig)(nil)).Elem()
}

func (i AviatrixProxyConfigMap) ToAviatrixProxyConfigMapOutput() AviatrixProxyConfigMapOutput {
	return i.ToAviatrixProxyConfigMapOutputWithContext(context.Background())
}

func (i AviatrixProxyConfigMap) ToAviatrixProxyConfigMapOutputWithContext(ctx context.Context) AviatrixProxyConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixProxyConfigMapOutput)
}

type AviatrixProxyConfigOutput struct{ *pulumi.OutputState }

func (AviatrixProxyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixProxyConfig)(nil)).Elem()
}

func (o AviatrixProxyConfigOutput) ToAviatrixProxyConfigOutput() AviatrixProxyConfigOutput {
	return o
}

func (o AviatrixProxyConfigOutput) ToAviatrixProxyConfigOutputWithContext(ctx context.Context) AviatrixProxyConfigOutput {
	return o
}

// Http proxy URL.
func (o AviatrixProxyConfigOutput) HttpProxy() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixProxyConfig) pulumi.StringOutput { return v.HttpProxy }).(pulumi.StringOutput)
}

// Https proxy URL.
func (o AviatrixProxyConfigOutput) HttpsProxy() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixProxyConfig) pulumi.StringOutput { return v.HttpsProxy }).(pulumi.StringOutput)
}

// Server CA Certificate file. Use the `file` function to read from a file.
func (o AviatrixProxyConfigOutput) ProxyCaCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixProxyConfig) pulumi.StringPtrOutput { return v.ProxyCaCertificate }).(pulumi.StringPtrOutput)
}

type AviatrixProxyConfigArrayOutput struct{ *pulumi.OutputState }

func (AviatrixProxyConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixProxyConfig)(nil)).Elem()
}

func (o AviatrixProxyConfigArrayOutput) ToAviatrixProxyConfigArrayOutput() AviatrixProxyConfigArrayOutput {
	return o
}

func (o AviatrixProxyConfigArrayOutput) ToAviatrixProxyConfigArrayOutputWithContext(ctx context.Context) AviatrixProxyConfigArrayOutput {
	return o
}

func (o AviatrixProxyConfigArrayOutput) Index(i pulumi.IntInput) AviatrixProxyConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixProxyConfig {
		return vs[0].([]*AviatrixProxyConfig)[vs[1].(int)]
	}).(AviatrixProxyConfigOutput)
}

type AviatrixProxyConfigMapOutput struct{ *pulumi.OutputState }

func (AviatrixProxyConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixProxyConfig)(nil)).Elem()
}

func (o AviatrixProxyConfigMapOutput) ToAviatrixProxyConfigMapOutput() AviatrixProxyConfigMapOutput {
	return o
}

func (o AviatrixProxyConfigMapOutput) ToAviatrixProxyConfigMapOutputWithContext(ctx context.Context) AviatrixProxyConfigMapOutput {
	return o
}

func (o AviatrixProxyConfigMapOutput) MapIndex(k pulumi.StringInput) AviatrixProxyConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixProxyConfig {
		return vs[0].(map[string]*AviatrixProxyConfig)[vs[1].(string)]
	}).(AviatrixProxyConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixProxyConfigInput)(nil)).Elem(), &AviatrixProxyConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixProxyConfigArrayInput)(nil)).Elem(), AviatrixProxyConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixProxyConfigMapInput)(nil)).Elem(), AviatrixProxyConfigMap{})
	pulumi.RegisterOutputType(AviatrixProxyConfigOutput{})
	pulumi.RegisterOutputType(AviatrixProxyConfigArrayOutput{})
	pulumi.RegisterOutputType(AviatrixProxyConfigMapOutput{})
}

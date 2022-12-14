// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_gateway_image** data source provides the current image version that pairs with the given software version
// and cloud type.
//
// This data source is useful for getting the correct imageVersion for a gateway when upgrading the softwareVersion of
// the gateway.
func GetAviatrixGatewayImage(ctx *pulumi.Context, args *GetAviatrixGatewayImageArgs, opts ...pulumi.InvokeOption) (*GetAviatrixGatewayImageResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAviatrixGatewayImageResult
	err := ctx.Invoke("aviatrix:index/getAviatrixGatewayImage:getAviatrixGatewayImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAviatrixGatewayImage.
type GetAviatrixGatewayImageArgs struct {
	// Cloud type. Type: Integer. Example: 1 (AWS)
	CloudType int `pulumi:"cloudType"`
	// Software version. Type: String. Example: "6.4.2487"
	SoftwareVersion string `pulumi:"softwareVersion"`
}

// A collection of values returned by getAviatrixGatewayImage.
type GetAviatrixGatewayImageResult struct {
	CloudType int `pulumi:"cloudType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Image version that is compatible with the given cloudType and software_version.
	ImageVersion    string `pulumi:"imageVersion"`
	SoftwareVersion string `pulumi:"softwareVersion"`
}

func GetAviatrixGatewayImageOutput(ctx *pulumi.Context, args GetAviatrixGatewayImageOutputArgs, opts ...pulumi.InvokeOption) GetAviatrixGatewayImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAviatrixGatewayImageResult, error) {
			args := v.(GetAviatrixGatewayImageArgs)
			r, err := GetAviatrixGatewayImage(ctx, &args, opts...)
			var s GetAviatrixGatewayImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAviatrixGatewayImageResultOutput)
}

// A collection of arguments for invoking getAviatrixGatewayImage.
type GetAviatrixGatewayImageOutputArgs struct {
	// Cloud type. Type: Integer. Example: 1 (AWS)
	CloudType pulumi.IntInput `pulumi:"cloudType"`
	// Software version. Type: String. Example: "6.4.2487"
	SoftwareVersion pulumi.StringInput `pulumi:"softwareVersion"`
}

func (GetAviatrixGatewayImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAviatrixGatewayImageArgs)(nil)).Elem()
}

// A collection of values returned by getAviatrixGatewayImage.
type GetAviatrixGatewayImageResultOutput struct{ *pulumi.OutputState }

func (GetAviatrixGatewayImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAviatrixGatewayImageResult)(nil)).Elem()
}

func (o GetAviatrixGatewayImageResultOutput) ToGetAviatrixGatewayImageResultOutput() GetAviatrixGatewayImageResultOutput {
	return o
}

func (o GetAviatrixGatewayImageResultOutput) ToGetAviatrixGatewayImageResultOutputWithContext(ctx context.Context) GetAviatrixGatewayImageResultOutput {
	return o
}

func (o GetAviatrixGatewayImageResultOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v GetAviatrixGatewayImageResult) int { return v.CloudType }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAviatrixGatewayImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAviatrixGatewayImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Image version that is compatible with the given cloudType and software_version.
func (o GetAviatrixGatewayImageResultOutput) ImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAviatrixGatewayImageResult) string { return v.ImageVersion }).(pulumi.StringOutput)
}

func (o GetAviatrixGatewayImageResultOutput) SoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAviatrixGatewayImageResult) string { return v.SoftwareVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAviatrixGatewayImageResultOutput{})
}

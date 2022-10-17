// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAviatrixVpcTracker(ctx *pulumi.Context, args *GetAviatrixVpcTrackerArgs, opts ...pulumi.InvokeOption) (*GetAviatrixVpcTrackerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAviatrixVpcTrackerResult
	err := ctx.Invoke("aviatrix:index/getAviatrixVpcTracker:getAviatrixVpcTracker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAviatrixVpcTracker.
type GetAviatrixVpcTrackerArgs struct {
	AccountName *string `pulumi:"accountName"`
	Cidr        *string `pulumi:"cidr"`
	CloudType   *int    `pulumi:"cloudType"`
	Region      *string `pulumi:"region"`
}

// A collection of values returned by getAviatrixVpcTracker.
type GetAviatrixVpcTrackerResult struct {
	AccountName *string `pulumi:"accountName"`
	Cidr        *string `pulumi:"cidr"`
	CloudType   *int    `pulumi:"cloudType"`
	// The provider-assigned unique ID for this managed resource.
	Id       string                         `pulumi:"id"`
	Region   *string                        `pulumi:"region"`
	VpcLists []GetAviatrixVpcTrackerVpcList `pulumi:"vpcLists"`
}

func GetAviatrixVpcTrackerOutput(ctx *pulumi.Context, args GetAviatrixVpcTrackerOutputArgs, opts ...pulumi.InvokeOption) GetAviatrixVpcTrackerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAviatrixVpcTrackerResult, error) {
			args := v.(GetAviatrixVpcTrackerArgs)
			r, err := GetAviatrixVpcTracker(ctx, &args, opts...)
			var s GetAviatrixVpcTrackerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAviatrixVpcTrackerResultOutput)
}

// A collection of arguments for invoking getAviatrixVpcTracker.
type GetAviatrixVpcTrackerOutputArgs struct {
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
	Cidr        pulumi.StringPtrInput `pulumi:"cidr"`
	CloudType   pulumi.IntPtrInput    `pulumi:"cloudType"`
	Region      pulumi.StringPtrInput `pulumi:"region"`
}

func (GetAviatrixVpcTrackerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAviatrixVpcTrackerArgs)(nil)).Elem()
}

// A collection of values returned by getAviatrixVpcTracker.
type GetAviatrixVpcTrackerResultOutput struct{ *pulumi.OutputState }

func (GetAviatrixVpcTrackerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAviatrixVpcTrackerResult)(nil)).Elem()
}

func (o GetAviatrixVpcTrackerResultOutput) ToGetAviatrixVpcTrackerResultOutput() GetAviatrixVpcTrackerResultOutput {
	return o
}

func (o GetAviatrixVpcTrackerResultOutput) ToGetAviatrixVpcTrackerResultOutputWithContext(ctx context.Context) GetAviatrixVpcTrackerResultOutput {
	return o
}

func (o GetAviatrixVpcTrackerResultOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAviatrixVpcTrackerResult) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o GetAviatrixVpcTrackerResultOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAviatrixVpcTrackerResult) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

func (o GetAviatrixVpcTrackerResultOutput) CloudType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAviatrixVpcTrackerResult) *int { return v.CloudType }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAviatrixVpcTrackerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAviatrixVpcTrackerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAviatrixVpcTrackerResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAviatrixVpcTrackerResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetAviatrixVpcTrackerResultOutput) VpcLists() GetAviatrixVpcTrackerVpcListArrayOutput {
	return o.ApplyT(func(v GetAviatrixVpcTrackerResult) []GetAviatrixVpcTrackerVpcList { return v.VpcLists }).(GetAviatrixVpcTrackerVpcListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAviatrixVpcTrackerResultOutput{})
}
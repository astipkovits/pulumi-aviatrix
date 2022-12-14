// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_transit_gateways** data source provides details about all transit gateways created by the Aviatrix Controller.
func GetAviatrixTransitGateways(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetAviatrixTransitGatewaysResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAviatrixTransitGatewaysResult
	err := ctx.Invoke("aviatrix:index/getAviatrixTransitGateways:getAviatrixTransitGateways", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAviatrixTransitGateways.
type GetAviatrixTransitGatewaysResult struct {
	// The list of all transit gateways
	GatewayLists []GetAviatrixTransitGatewaysGatewayList `pulumi:"gatewayLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

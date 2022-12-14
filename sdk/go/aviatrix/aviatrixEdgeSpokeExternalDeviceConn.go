// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_edge_spoke_external_device_conn** resource creates and manages the connection between Edge as a Spoke and an External Device. This resource is available as of provider version R2.23+.
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
//			_, err := aviatrix.NewAviatrixEdgeSpokeExternalDeviceConn(ctx, "test", &aviatrix.AviatrixEdgeSpokeExternalDeviceConnArgs{
//				BgpLocalAsNum:  pulumi.String("123"),
//				BgpRemoteAsNum: pulumi.String("345"),
//				ConnectionName: pulumi.String("conn"),
//				GwName:         pulumi.String("eaas"),
//				LocalLanIp:     pulumi.String("10.230.3.23"),
//				RemoteLanIp:    pulumi.String("10.0.60.1"),
//				SiteId:         pulumi.String("site-abcd1234"),
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
// **edge_spoke_external_device_conn** can be imported using the `connection_name` and `site_id`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn test connection_name~site_id
//
// ```
type AviatrixEdgeSpokeExternalDeviceConn struct {
	pulumi.CustomResourceState

	// BGP local AS number.
	BgpLocalAsNum pulumi.StringOutput `pulumi:"bgpLocalAsNum"`
	// BGP remote AS number.
	BgpRemoteAsNum pulumi.StringOutput `pulumi:"bgpRemoteAsNum"`
	// Connection name.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// Connection type. Valid value: 'bgp'. Default value: 'bgp'.
	ConnectionType pulumi.StringPtrOutput `pulumi:"connectionType"`
	// Edge as a Spoke name.
	GwName pulumi.StringOutput `pulumi:"gwName"`
	// Local LAN IP.
	LocalLanIp pulumi.StringOutput `pulumi:"localLanIp"`
	// Remote LAN IP.
	RemoteLanIp pulumi.StringOutput `pulumi:"remoteLanIp"`
	// Edge as a Spoke site iD.
	SiteId pulumi.StringOutput `pulumi:"siteId"`
	// Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
	TunnelProtocol pulumi.StringPtrOutput `pulumi:"tunnelProtocol"`
}

// NewAviatrixEdgeSpokeExternalDeviceConn registers a new resource with the given unique name, arguments, and options.
func NewAviatrixEdgeSpokeExternalDeviceConn(ctx *pulumi.Context,
	name string, args *AviatrixEdgeSpokeExternalDeviceConnArgs, opts ...pulumi.ResourceOption) (*AviatrixEdgeSpokeExternalDeviceConn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpLocalAsNum == nil {
		return nil, errors.New("invalid value for required argument 'BgpLocalAsNum'")
	}
	if args.BgpRemoteAsNum == nil {
		return nil, errors.New("invalid value for required argument 'BgpRemoteAsNum'")
	}
	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.GwName == nil {
		return nil, errors.New("invalid value for required argument 'GwName'")
	}
	if args.LocalLanIp == nil {
		return nil, errors.New("invalid value for required argument 'LocalLanIp'")
	}
	if args.RemoteLanIp == nil {
		return nil, errors.New("invalid value for required argument 'RemoteLanIp'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixEdgeSpokeExternalDeviceConn
	err := ctx.RegisterResource("aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixEdgeSpokeExternalDeviceConn gets an existing AviatrixEdgeSpokeExternalDeviceConn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixEdgeSpokeExternalDeviceConn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixEdgeSpokeExternalDeviceConnState, opts ...pulumi.ResourceOption) (*AviatrixEdgeSpokeExternalDeviceConn, error) {
	var resource AviatrixEdgeSpokeExternalDeviceConn
	err := ctx.ReadResource("aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixEdgeSpokeExternalDeviceConn resources.
type aviatrixEdgeSpokeExternalDeviceConnState struct {
	// BGP local AS number.
	BgpLocalAsNum *string `pulumi:"bgpLocalAsNum"`
	// BGP remote AS number.
	BgpRemoteAsNum *string `pulumi:"bgpRemoteAsNum"`
	// Connection name.
	ConnectionName *string `pulumi:"connectionName"`
	// Connection type. Valid value: 'bgp'. Default value: 'bgp'.
	ConnectionType *string `pulumi:"connectionType"`
	// Edge as a Spoke name.
	GwName *string `pulumi:"gwName"`
	// Local LAN IP.
	LocalLanIp *string `pulumi:"localLanIp"`
	// Remote LAN IP.
	RemoteLanIp *string `pulumi:"remoteLanIp"`
	// Edge as a Spoke site iD.
	SiteId *string `pulumi:"siteId"`
	// Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
	TunnelProtocol *string `pulumi:"tunnelProtocol"`
}

type AviatrixEdgeSpokeExternalDeviceConnState struct {
	// BGP local AS number.
	BgpLocalAsNum pulumi.StringPtrInput
	// BGP remote AS number.
	BgpRemoteAsNum pulumi.StringPtrInput
	// Connection name.
	ConnectionName pulumi.StringPtrInput
	// Connection type. Valid value: 'bgp'. Default value: 'bgp'.
	ConnectionType pulumi.StringPtrInput
	// Edge as a Spoke name.
	GwName pulumi.StringPtrInput
	// Local LAN IP.
	LocalLanIp pulumi.StringPtrInput
	// Remote LAN IP.
	RemoteLanIp pulumi.StringPtrInput
	// Edge as a Spoke site iD.
	SiteId pulumi.StringPtrInput
	// Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
	TunnelProtocol pulumi.StringPtrInput
}

func (AviatrixEdgeSpokeExternalDeviceConnState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixEdgeSpokeExternalDeviceConnState)(nil)).Elem()
}

type aviatrixEdgeSpokeExternalDeviceConnArgs struct {
	// BGP local AS number.
	BgpLocalAsNum string `pulumi:"bgpLocalAsNum"`
	// BGP remote AS number.
	BgpRemoteAsNum string `pulumi:"bgpRemoteAsNum"`
	// Connection name.
	ConnectionName string `pulumi:"connectionName"`
	// Connection type. Valid value: 'bgp'. Default value: 'bgp'.
	ConnectionType *string `pulumi:"connectionType"`
	// Edge as a Spoke name.
	GwName string `pulumi:"gwName"`
	// Local LAN IP.
	LocalLanIp string `pulumi:"localLanIp"`
	// Remote LAN IP.
	RemoteLanIp string `pulumi:"remoteLanIp"`
	// Edge as a Spoke site iD.
	SiteId string `pulumi:"siteId"`
	// Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
	TunnelProtocol *string `pulumi:"tunnelProtocol"`
}

// The set of arguments for constructing a AviatrixEdgeSpokeExternalDeviceConn resource.
type AviatrixEdgeSpokeExternalDeviceConnArgs struct {
	// BGP local AS number.
	BgpLocalAsNum pulumi.StringInput
	// BGP remote AS number.
	BgpRemoteAsNum pulumi.StringInput
	// Connection name.
	ConnectionName pulumi.StringInput
	// Connection type. Valid value: 'bgp'. Default value: 'bgp'.
	ConnectionType pulumi.StringPtrInput
	// Edge as a Spoke name.
	GwName pulumi.StringInput
	// Local LAN IP.
	LocalLanIp pulumi.StringInput
	// Remote LAN IP.
	RemoteLanIp pulumi.StringInput
	// Edge as a Spoke site iD.
	SiteId pulumi.StringInput
	// Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
	TunnelProtocol pulumi.StringPtrInput
}

func (AviatrixEdgeSpokeExternalDeviceConnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixEdgeSpokeExternalDeviceConnArgs)(nil)).Elem()
}

type AviatrixEdgeSpokeExternalDeviceConnInput interface {
	pulumi.Input

	ToAviatrixEdgeSpokeExternalDeviceConnOutput() AviatrixEdgeSpokeExternalDeviceConnOutput
	ToAviatrixEdgeSpokeExternalDeviceConnOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnOutput
}

func (*AviatrixEdgeSpokeExternalDeviceConn) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixEdgeSpokeExternalDeviceConn)(nil)).Elem()
}

func (i *AviatrixEdgeSpokeExternalDeviceConn) ToAviatrixEdgeSpokeExternalDeviceConnOutput() AviatrixEdgeSpokeExternalDeviceConnOutput {
	return i.ToAviatrixEdgeSpokeExternalDeviceConnOutputWithContext(context.Background())
}

func (i *AviatrixEdgeSpokeExternalDeviceConn) ToAviatrixEdgeSpokeExternalDeviceConnOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixEdgeSpokeExternalDeviceConnOutput)
}

// AviatrixEdgeSpokeExternalDeviceConnArrayInput is an input type that accepts AviatrixEdgeSpokeExternalDeviceConnArray and AviatrixEdgeSpokeExternalDeviceConnArrayOutput values.
// You can construct a concrete instance of `AviatrixEdgeSpokeExternalDeviceConnArrayInput` via:
//
//	AviatrixEdgeSpokeExternalDeviceConnArray{ AviatrixEdgeSpokeExternalDeviceConnArgs{...} }
type AviatrixEdgeSpokeExternalDeviceConnArrayInput interface {
	pulumi.Input

	ToAviatrixEdgeSpokeExternalDeviceConnArrayOutput() AviatrixEdgeSpokeExternalDeviceConnArrayOutput
	ToAviatrixEdgeSpokeExternalDeviceConnArrayOutputWithContext(context.Context) AviatrixEdgeSpokeExternalDeviceConnArrayOutput
}

type AviatrixEdgeSpokeExternalDeviceConnArray []AviatrixEdgeSpokeExternalDeviceConnInput

func (AviatrixEdgeSpokeExternalDeviceConnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixEdgeSpokeExternalDeviceConn)(nil)).Elem()
}

func (i AviatrixEdgeSpokeExternalDeviceConnArray) ToAviatrixEdgeSpokeExternalDeviceConnArrayOutput() AviatrixEdgeSpokeExternalDeviceConnArrayOutput {
	return i.ToAviatrixEdgeSpokeExternalDeviceConnArrayOutputWithContext(context.Background())
}

func (i AviatrixEdgeSpokeExternalDeviceConnArray) ToAviatrixEdgeSpokeExternalDeviceConnArrayOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixEdgeSpokeExternalDeviceConnArrayOutput)
}

// AviatrixEdgeSpokeExternalDeviceConnMapInput is an input type that accepts AviatrixEdgeSpokeExternalDeviceConnMap and AviatrixEdgeSpokeExternalDeviceConnMapOutput values.
// You can construct a concrete instance of `AviatrixEdgeSpokeExternalDeviceConnMapInput` via:
//
//	AviatrixEdgeSpokeExternalDeviceConnMap{ "key": AviatrixEdgeSpokeExternalDeviceConnArgs{...} }
type AviatrixEdgeSpokeExternalDeviceConnMapInput interface {
	pulumi.Input

	ToAviatrixEdgeSpokeExternalDeviceConnMapOutput() AviatrixEdgeSpokeExternalDeviceConnMapOutput
	ToAviatrixEdgeSpokeExternalDeviceConnMapOutputWithContext(context.Context) AviatrixEdgeSpokeExternalDeviceConnMapOutput
}

type AviatrixEdgeSpokeExternalDeviceConnMap map[string]AviatrixEdgeSpokeExternalDeviceConnInput

func (AviatrixEdgeSpokeExternalDeviceConnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixEdgeSpokeExternalDeviceConn)(nil)).Elem()
}

func (i AviatrixEdgeSpokeExternalDeviceConnMap) ToAviatrixEdgeSpokeExternalDeviceConnMapOutput() AviatrixEdgeSpokeExternalDeviceConnMapOutput {
	return i.ToAviatrixEdgeSpokeExternalDeviceConnMapOutputWithContext(context.Background())
}

func (i AviatrixEdgeSpokeExternalDeviceConnMap) ToAviatrixEdgeSpokeExternalDeviceConnMapOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixEdgeSpokeExternalDeviceConnMapOutput)
}

type AviatrixEdgeSpokeExternalDeviceConnOutput struct{ *pulumi.OutputState }

func (AviatrixEdgeSpokeExternalDeviceConnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixEdgeSpokeExternalDeviceConn)(nil)).Elem()
}

func (o AviatrixEdgeSpokeExternalDeviceConnOutput) ToAviatrixEdgeSpokeExternalDeviceConnOutput() AviatrixEdgeSpokeExternalDeviceConnOutput {
	return o
}

func (o AviatrixEdgeSpokeExternalDeviceConnOutput) ToAviatrixEdgeSpokeExternalDeviceConnOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnOutput {
	return o
}

// BGP local AS number.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) BgpLocalAsNum() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.BgpLocalAsNum }).(pulumi.StringOutput)
}

// BGP remote AS number.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) BgpRemoteAsNum() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.BgpRemoteAsNum }).(pulumi.StringOutput)
}

// Connection name.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.ConnectionName }).(pulumi.StringOutput)
}

// Connection type. Valid value: 'bgp'. Default value: 'bgp'.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringPtrOutput { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

// Edge as a Spoke name.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) GwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.GwName }).(pulumi.StringOutput)
}

// Local LAN IP.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) LocalLanIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.LocalLanIp }).(pulumi.StringOutput)
}

// Remote LAN IP.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) RemoteLanIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.RemoteLanIp }).(pulumi.StringOutput)
}

// Edge as a Spoke site iD.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringOutput { return v.SiteId }).(pulumi.StringOutput)
}

// Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
func (o AviatrixEdgeSpokeExternalDeviceConnOutput) TunnelProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixEdgeSpokeExternalDeviceConn) pulumi.StringPtrOutput { return v.TunnelProtocol }).(pulumi.StringPtrOutput)
}

type AviatrixEdgeSpokeExternalDeviceConnArrayOutput struct{ *pulumi.OutputState }

func (AviatrixEdgeSpokeExternalDeviceConnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixEdgeSpokeExternalDeviceConn)(nil)).Elem()
}

func (o AviatrixEdgeSpokeExternalDeviceConnArrayOutput) ToAviatrixEdgeSpokeExternalDeviceConnArrayOutput() AviatrixEdgeSpokeExternalDeviceConnArrayOutput {
	return o
}

func (o AviatrixEdgeSpokeExternalDeviceConnArrayOutput) ToAviatrixEdgeSpokeExternalDeviceConnArrayOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnArrayOutput {
	return o
}

func (o AviatrixEdgeSpokeExternalDeviceConnArrayOutput) Index(i pulumi.IntInput) AviatrixEdgeSpokeExternalDeviceConnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixEdgeSpokeExternalDeviceConn {
		return vs[0].([]*AviatrixEdgeSpokeExternalDeviceConn)[vs[1].(int)]
	}).(AviatrixEdgeSpokeExternalDeviceConnOutput)
}

type AviatrixEdgeSpokeExternalDeviceConnMapOutput struct{ *pulumi.OutputState }

func (AviatrixEdgeSpokeExternalDeviceConnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixEdgeSpokeExternalDeviceConn)(nil)).Elem()
}

func (o AviatrixEdgeSpokeExternalDeviceConnMapOutput) ToAviatrixEdgeSpokeExternalDeviceConnMapOutput() AviatrixEdgeSpokeExternalDeviceConnMapOutput {
	return o
}

func (o AviatrixEdgeSpokeExternalDeviceConnMapOutput) ToAviatrixEdgeSpokeExternalDeviceConnMapOutputWithContext(ctx context.Context) AviatrixEdgeSpokeExternalDeviceConnMapOutput {
	return o
}

func (o AviatrixEdgeSpokeExternalDeviceConnMapOutput) MapIndex(k pulumi.StringInput) AviatrixEdgeSpokeExternalDeviceConnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixEdgeSpokeExternalDeviceConn {
		return vs[0].(map[string]*AviatrixEdgeSpokeExternalDeviceConn)[vs[1].(string)]
	}).(AviatrixEdgeSpokeExternalDeviceConnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixEdgeSpokeExternalDeviceConnInput)(nil)).Elem(), &AviatrixEdgeSpokeExternalDeviceConn{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixEdgeSpokeExternalDeviceConnArrayInput)(nil)).Elem(), AviatrixEdgeSpokeExternalDeviceConnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixEdgeSpokeExternalDeviceConnMapInput)(nil)).Elem(), AviatrixEdgeSpokeExternalDeviceConnMap{})
	pulumi.RegisterOutputType(AviatrixEdgeSpokeExternalDeviceConnOutput{})
	pulumi.RegisterOutputType(AviatrixEdgeSpokeExternalDeviceConnArrayOutput{})
	pulumi.RegisterOutputType(AviatrixEdgeSpokeExternalDeviceConnMapOutput{})
}

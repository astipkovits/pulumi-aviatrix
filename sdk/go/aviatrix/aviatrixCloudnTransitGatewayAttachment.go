// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixCloudnTransitGatewayAttachment struct {
	pulumi.CustomResourceState

	// CloudN BGP AS Number.
	CloudnBgpAsn pulumi.StringOutput `pulumi:"cloudnBgpAsn"`
	// CloudN LAN Interface Neighbor's BGP AS Number.
	CloudnLanInterfaceNeighborBgpAsn pulumi.StringOutput `pulumi:"cloudnLanInterfaceNeighborBgpAsn"`
	// CloudN LAN Interface Neighbor's IP.
	CloudnLanInterfaceNeighborIp pulumi.StringOutput `pulumi:"cloudnLanInterfaceNeighborIp"`
	// Connection name.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// Device name.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// Enable jumbo frame.
	EnableJumboFrame pulumi.BoolPtrOutput `pulumi:"enableJumboFrame"`
	// Enable over private network.
	EnableOverPrivateNetwork pulumi.BoolPtrOutput `pulumi:"enableOverPrivateNetwork"`
	// AS path prepend.
	PrependAsPaths pulumi.StringArrayOutput `pulumi:"prependAsPaths"`
	// Transit Gateway BGP AS Number.
	TransitGatewayBgpAsn pulumi.StringOutput `pulumi:"transitGatewayBgpAsn"`
	// Transit Gateway name.
	TransitGatewayName pulumi.StringOutput `pulumi:"transitGatewayName"`
}

// NewAviatrixCloudnTransitGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewAviatrixCloudnTransitGatewayAttachment(ctx *pulumi.Context,
	name string, args *AviatrixCloudnTransitGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*AviatrixCloudnTransitGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudnBgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'CloudnBgpAsn'")
	}
	if args.CloudnLanInterfaceNeighborBgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'CloudnLanInterfaceNeighborBgpAsn'")
	}
	if args.CloudnLanInterfaceNeighborIp == nil {
		return nil, errors.New("invalid value for required argument 'CloudnLanInterfaceNeighborIp'")
	}
	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.TransitGatewayBgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayBgpAsn'")
	}
	if args.TransitGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixCloudnTransitGatewayAttachment
	err := ctx.RegisterResource("aviatrix:index/aviatrixCloudnTransitGatewayAttachment:AviatrixCloudnTransitGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixCloudnTransitGatewayAttachment gets an existing AviatrixCloudnTransitGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixCloudnTransitGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixCloudnTransitGatewayAttachmentState, opts ...pulumi.ResourceOption) (*AviatrixCloudnTransitGatewayAttachment, error) {
	var resource AviatrixCloudnTransitGatewayAttachment
	err := ctx.ReadResource("aviatrix:index/aviatrixCloudnTransitGatewayAttachment:AviatrixCloudnTransitGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixCloudnTransitGatewayAttachment resources.
type aviatrixCloudnTransitGatewayAttachmentState struct {
	// CloudN BGP AS Number.
	CloudnBgpAsn *string `pulumi:"cloudnBgpAsn"`
	// CloudN LAN Interface Neighbor's BGP AS Number.
	CloudnLanInterfaceNeighborBgpAsn *string `pulumi:"cloudnLanInterfaceNeighborBgpAsn"`
	// CloudN LAN Interface Neighbor's IP.
	CloudnLanInterfaceNeighborIp *string `pulumi:"cloudnLanInterfaceNeighborIp"`
	// Connection name.
	ConnectionName *string `pulumi:"connectionName"`
	// Device name.
	DeviceName *string `pulumi:"deviceName"`
	// Enable jumbo frame.
	EnableJumboFrame *bool `pulumi:"enableJumboFrame"`
	// Enable over private network.
	EnableOverPrivateNetwork *bool `pulumi:"enableOverPrivateNetwork"`
	// AS path prepend.
	PrependAsPaths []string `pulumi:"prependAsPaths"`
	// Transit Gateway BGP AS Number.
	TransitGatewayBgpAsn *string `pulumi:"transitGatewayBgpAsn"`
	// Transit Gateway name.
	TransitGatewayName *string `pulumi:"transitGatewayName"`
}

type AviatrixCloudnTransitGatewayAttachmentState struct {
	// CloudN BGP AS Number.
	CloudnBgpAsn pulumi.StringPtrInput
	// CloudN LAN Interface Neighbor's BGP AS Number.
	CloudnLanInterfaceNeighborBgpAsn pulumi.StringPtrInput
	// CloudN LAN Interface Neighbor's IP.
	CloudnLanInterfaceNeighborIp pulumi.StringPtrInput
	// Connection name.
	ConnectionName pulumi.StringPtrInput
	// Device name.
	DeviceName pulumi.StringPtrInput
	// Enable jumbo frame.
	EnableJumboFrame pulumi.BoolPtrInput
	// Enable over private network.
	EnableOverPrivateNetwork pulumi.BoolPtrInput
	// AS path prepend.
	PrependAsPaths pulumi.StringArrayInput
	// Transit Gateway BGP AS Number.
	TransitGatewayBgpAsn pulumi.StringPtrInput
	// Transit Gateway name.
	TransitGatewayName pulumi.StringPtrInput
}

func (AviatrixCloudnTransitGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixCloudnTransitGatewayAttachmentState)(nil)).Elem()
}

type aviatrixCloudnTransitGatewayAttachmentArgs struct {
	// CloudN BGP AS Number.
	CloudnBgpAsn string `pulumi:"cloudnBgpAsn"`
	// CloudN LAN Interface Neighbor's BGP AS Number.
	CloudnLanInterfaceNeighborBgpAsn string `pulumi:"cloudnLanInterfaceNeighborBgpAsn"`
	// CloudN LAN Interface Neighbor's IP.
	CloudnLanInterfaceNeighborIp string `pulumi:"cloudnLanInterfaceNeighborIp"`
	// Connection name.
	ConnectionName string `pulumi:"connectionName"`
	// Device name.
	DeviceName string `pulumi:"deviceName"`
	// Enable jumbo frame.
	EnableJumboFrame *bool `pulumi:"enableJumboFrame"`
	// Enable over private network.
	EnableOverPrivateNetwork *bool `pulumi:"enableOverPrivateNetwork"`
	// AS path prepend.
	PrependAsPaths []string `pulumi:"prependAsPaths"`
	// Transit Gateway BGP AS Number.
	TransitGatewayBgpAsn string `pulumi:"transitGatewayBgpAsn"`
	// Transit Gateway name.
	TransitGatewayName string `pulumi:"transitGatewayName"`
}

// The set of arguments for constructing a AviatrixCloudnTransitGatewayAttachment resource.
type AviatrixCloudnTransitGatewayAttachmentArgs struct {
	// CloudN BGP AS Number.
	CloudnBgpAsn pulumi.StringInput
	// CloudN LAN Interface Neighbor's BGP AS Number.
	CloudnLanInterfaceNeighborBgpAsn pulumi.StringInput
	// CloudN LAN Interface Neighbor's IP.
	CloudnLanInterfaceNeighborIp pulumi.StringInput
	// Connection name.
	ConnectionName pulumi.StringInput
	// Device name.
	DeviceName pulumi.StringInput
	// Enable jumbo frame.
	EnableJumboFrame pulumi.BoolPtrInput
	// Enable over private network.
	EnableOverPrivateNetwork pulumi.BoolPtrInput
	// AS path prepend.
	PrependAsPaths pulumi.StringArrayInput
	// Transit Gateway BGP AS Number.
	TransitGatewayBgpAsn pulumi.StringInput
	// Transit Gateway name.
	TransitGatewayName pulumi.StringInput
}

func (AviatrixCloudnTransitGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixCloudnTransitGatewayAttachmentArgs)(nil)).Elem()
}

type AviatrixCloudnTransitGatewayAttachmentInput interface {
	pulumi.Input

	ToAviatrixCloudnTransitGatewayAttachmentOutput() AviatrixCloudnTransitGatewayAttachmentOutput
	ToAviatrixCloudnTransitGatewayAttachmentOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentOutput
}

func (*AviatrixCloudnTransitGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixCloudnTransitGatewayAttachment)(nil)).Elem()
}

func (i *AviatrixCloudnTransitGatewayAttachment) ToAviatrixCloudnTransitGatewayAttachmentOutput() AviatrixCloudnTransitGatewayAttachmentOutput {
	return i.ToAviatrixCloudnTransitGatewayAttachmentOutputWithContext(context.Background())
}

func (i *AviatrixCloudnTransitGatewayAttachment) ToAviatrixCloudnTransitGatewayAttachmentOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixCloudnTransitGatewayAttachmentOutput)
}

// AviatrixCloudnTransitGatewayAttachmentArrayInput is an input type that accepts AviatrixCloudnTransitGatewayAttachmentArray and AviatrixCloudnTransitGatewayAttachmentArrayOutput values.
// You can construct a concrete instance of `AviatrixCloudnTransitGatewayAttachmentArrayInput` via:
//
//	AviatrixCloudnTransitGatewayAttachmentArray{ AviatrixCloudnTransitGatewayAttachmentArgs{...} }
type AviatrixCloudnTransitGatewayAttachmentArrayInput interface {
	pulumi.Input

	ToAviatrixCloudnTransitGatewayAttachmentArrayOutput() AviatrixCloudnTransitGatewayAttachmentArrayOutput
	ToAviatrixCloudnTransitGatewayAttachmentArrayOutputWithContext(context.Context) AviatrixCloudnTransitGatewayAttachmentArrayOutput
}

type AviatrixCloudnTransitGatewayAttachmentArray []AviatrixCloudnTransitGatewayAttachmentInput

func (AviatrixCloudnTransitGatewayAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixCloudnTransitGatewayAttachment)(nil)).Elem()
}

func (i AviatrixCloudnTransitGatewayAttachmentArray) ToAviatrixCloudnTransitGatewayAttachmentArrayOutput() AviatrixCloudnTransitGatewayAttachmentArrayOutput {
	return i.ToAviatrixCloudnTransitGatewayAttachmentArrayOutputWithContext(context.Background())
}

func (i AviatrixCloudnTransitGatewayAttachmentArray) ToAviatrixCloudnTransitGatewayAttachmentArrayOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixCloudnTransitGatewayAttachmentArrayOutput)
}

// AviatrixCloudnTransitGatewayAttachmentMapInput is an input type that accepts AviatrixCloudnTransitGatewayAttachmentMap and AviatrixCloudnTransitGatewayAttachmentMapOutput values.
// You can construct a concrete instance of `AviatrixCloudnTransitGatewayAttachmentMapInput` via:
//
//	AviatrixCloudnTransitGatewayAttachmentMap{ "key": AviatrixCloudnTransitGatewayAttachmentArgs{...} }
type AviatrixCloudnTransitGatewayAttachmentMapInput interface {
	pulumi.Input

	ToAviatrixCloudnTransitGatewayAttachmentMapOutput() AviatrixCloudnTransitGatewayAttachmentMapOutput
	ToAviatrixCloudnTransitGatewayAttachmentMapOutputWithContext(context.Context) AviatrixCloudnTransitGatewayAttachmentMapOutput
}

type AviatrixCloudnTransitGatewayAttachmentMap map[string]AviatrixCloudnTransitGatewayAttachmentInput

func (AviatrixCloudnTransitGatewayAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixCloudnTransitGatewayAttachment)(nil)).Elem()
}

func (i AviatrixCloudnTransitGatewayAttachmentMap) ToAviatrixCloudnTransitGatewayAttachmentMapOutput() AviatrixCloudnTransitGatewayAttachmentMapOutput {
	return i.ToAviatrixCloudnTransitGatewayAttachmentMapOutputWithContext(context.Background())
}

func (i AviatrixCloudnTransitGatewayAttachmentMap) ToAviatrixCloudnTransitGatewayAttachmentMapOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixCloudnTransitGatewayAttachmentMapOutput)
}

type AviatrixCloudnTransitGatewayAttachmentOutput struct{ *pulumi.OutputState }

func (AviatrixCloudnTransitGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixCloudnTransitGatewayAttachment)(nil)).Elem()
}

func (o AviatrixCloudnTransitGatewayAttachmentOutput) ToAviatrixCloudnTransitGatewayAttachmentOutput() AviatrixCloudnTransitGatewayAttachmentOutput {
	return o
}

func (o AviatrixCloudnTransitGatewayAttachmentOutput) ToAviatrixCloudnTransitGatewayAttachmentOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentOutput {
	return o
}

// CloudN BGP AS Number.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) CloudnBgpAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput { return v.CloudnBgpAsn }).(pulumi.StringOutput)
}

// CloudN LAN Interface Neighbor's BGP AS Number.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) CloudnLanInterfaceNeighborBgpAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput {
		return v.CloudnLanInterfaceNeighborBgpAsn
	}).(pulumi.StringOutput)
}

// CloudN LAN Interface Neighbor's IP.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) CloudnLanInterfaceNeighborIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput {
		return v.CloudnLanInterfaceNeighborIp
	}).(pulumi.StringOutput)
}

// Connection name.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput { return v.ConnectionName }).(pulumi.StringOutput)
}

// Device name.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) DeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput { return v.DeviceName }).(pulumi.StringOutput)
}

// Enable jumbo frame.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) EnableJumboFrame() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.BoolPtrOutput { return v.EnableJumboFrame }).(pulumi.BoolPtrOutput)
}

// Enable over private network.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) EnableOverPrivateNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.BoolPtrOutput {
		return v.EnableOverPrivateNetwork
	}).(pulumi.BoolPtrOutput)
}

// AS path prepend.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) PrependAsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringArrayOutput { return v.PrependAsPaths }).(pulumi.StringArrayOutput)
}

// Transit Gateway BGP AS Number.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) TransitGatewayBgpAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput { return v.TransitGatewayBgpAsn }).(pulumi.StringOutput)
}

// Transit Gateway name.
func (o AviatrixCloudnTransitGatewayAttachmentOutput) TransitGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixCloudnTransitGatewayAttachment) pulumi.StringOutput { return v.TransitGatewayName }).(pulumi.StringOutput)
}

type AviatrixCloudnTransitGatewayAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AviatrixCloudnTransitGatewayAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixCloudnTransitGatewayAttachment)(nil)).Elem()
}

func (o AviatrixCloudnTransitGatewayAttachmentArrayOutput) ToAviatrixCloudnTransitGatewayAttachmentArrayOutput() AviatrixCloudnTransitGatewayAttachmentArrayOutput {
	return o
}

func (o AviatrixCloudnTransitGatewayAttachmentArrayOutput) ToAviatrixCloudnTransitGatewayAttachmentArrayOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentArrayOutput {
	return o
}

func (o AviatrixCloudnTransitGatewayAttachmentArrayOutput) Index(i pulumi.IntInput) AviatrixCloudnTransitGatewayAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixCloudnTransitGatewayAttachment {
		return vs[0].([]*AviatrixCloudnTransitGatewayAttachment)[vs[1].(int)]
	}).(AviatrixCloudnTransitGatewayAttachmentOutput)
}

type AviatrixCloudnTransitGatewayAttachmentMapOutput struct{ *pulumi.OutputState }

func (AviatrixCloudnTransitGatewayAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixCloudnTransitGatewayAttachment)(nil)).Elem()
}

func (o AviatrixCloudnTransitGatewayAttachmentMapOutput) ToAviatrixCloudnTransitGatewayAttachmentMapOutput() AviatrixCloudnTransitGatewayAttachmentMapOutput {
	return o
}

func (o AviatrixCloudnTransitGatewayAttachmentMapOutput) ToAviatrixCloudnTransitGatewayAttachmentMapOutputWithContext(ctx context.Context) AviatrixCloudnTransitGatewayAttachmentMapOutput {
	return o
}

func (o AviatrixCloudnTransitGatewayAttachmentMapOutput) MapIndex(k pulumi.StringInput) AviatrixCloudnTransitGatewayAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixCloudnTransitGatewayAttachment {
		return vs[0].(map[string]*AviatrixCloudnTransitGatewayAttachment)[vs[1].(string)]
	}).(AviatrixCloudnTransitGatewayAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixCloudnTransitGatewayAttachmentInput)(nil)).Elem(), &AviatrixCloudnTransitGatewayAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixCloudnTransitGatewayAttachmentArrayInput)(nil)).Elem(), AviatrixCloudnTransitGatewayAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixCloudnTransitGatewayAttachmentMapInput)(nil)).Elem(), AviatrixCloudnTransitGatewayAttachmentMap{})
	pulumi.RegisterOutputType(AviatrixCloudnTransitGatewayAttachmentOutput{})
	pulumi.RegisterOutputType(AviatrixCloudnTransitGatewayAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AviatrixCloudnTransitGatewayAttachmentMapOutput{})
}
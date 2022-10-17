// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixSamlEndpoint struct {
	pulumi.CustomResourceState

	// Access type.
	AccessSetBy pulumi.StringPtrOutput `pulumi:"accessSetBy"`
	// Switch to differentiate if it is for controller login.
	ControllerLogin pulumi.BoolPtrOutput `pulumi:"controllerLogin"`
	// Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
	CustomEntityId pulumi.StringPtrOutput `pulumi:"customEntityId"`
	// Custom SAML Request Template.
	CustomSamlRequestTemplate pulumi.StringPtrOutput `pulumi:"customSamlRequestTemplate"`
	// SAML Endpoint Name.
	EndpointName pulumi.StringOutput `pulumi:"endpointName"`
	// IDP Metadata.
	IdpMetadata pulumi.StringPtrOutput `pulumi:"idpMetadata"`
	// Type of IDP Metadata.
	IdpMetadataType pulumi.StringOutput `pulumi:"idpMetadataType"`
	// IDP Metadata.
	IdpMetadataUrl pulumi.StringPtrOutput `pulumi:"idpMetadataUrl"`
	// List of RBAC groups.
	RbacGroups pulumi.StringArrayOutput `pulumi:"rbacGroups"`
	// Whether to sign SAML AuthnRequests
	SignAuthnRequests pulumi.BoolPtrOutput `pulumi:"signAuthnRequests"`
}

// NewAviatrixSamlEndpoint registers a new resource with the given unique name, arguments, and options.
func NewAviatrixSamlEndpoint(ctx *pulumi.Context,
	name string, args *AviatrixSamlEndpointArgs, opts ...pulumi.ResourceOption) (*AviatrixSamlEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.IdpMetadataType == nil {
		return nil, errors.New("invalid value for required argument 'IdpMetadataType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixSamlEndpoint
	err := ctx.RegisterResource("aviatrix:index/aviatrixSamlEndpoint:AviatrixSamlEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixSamlEndpoint gets an existing AviatrixSamlEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixSamlEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixSamlEndpointState, opts ...pulumi.ResourceOption) (*AviatrixSamlEndpoint, error) {
	var resource AviatrixSamlEndpoint
	err := ctx.ReadResource("aviatrix:index/aviatrixSamlEndpoint:AviatrixSamlEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixSamlEndpoint resources.
type aviatrixSamlEndpointState struct {
	// Access type.
	AccessSetBy *string `pulumi:"accessSetBy"`
	// Switch to differentiate if it is for controller login.
	ControllerLogin *bool `pulumi:"controllerLogin"`
	// Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
	CustomEntityId *string `pulumi:"customEntityId"`
	// Custom SAML Request Template.
	CustomSamlRequestTemplate *string `pulumi:"customSamlRequestTemplate"`
	// SAML Endpoint Name.
	EndpointName *string `pulumi:"endpointName"`
	// IDP Metadata.
	IdpMetadata *string `pulumi:"idpMetadata"`
	// Type of IDP Metadata.
	IdpMetadataType *string `pulumi:"idpMetadataType"`
	// IDP Metadata.
	IdpMetadataUrl *string `pulumi:"idpMetadataUrl"`
	// List of RBAC groups.
	RbacGroups []string `pulumi:"rbacGroups"`
	// Whether to sign SAML AuthnRequests
	SignAuthnRequests *bool `pulumi:"signAuthnRequests"`
}

type AviatrixSamlEndpointState struct {
	// Access type.
	AccessSetBy pulumi.StringPtrInput
	// Switch to differentiate if it is for controller login.
	ControllerLogin pulumi.BoolPtrInput
	// Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
	CustomEntityId pulumi.StringPtrInput
	// Custom SAML Request Template.
	CustomSamlRequestTemplate pulumi.StringPtrInput
	// SAML Endpoint Name.
	EndpointName pulumi.StringPtrInput
	// IDP Metadata.
	IdpMetadata pulumi.StringPtrInput
	// Type of IDP Metadata.
	IdpMetadataType pulumi.StringPtrInput
	// IDP Metadata.
	IdpMetadataUrl pulumi.StringPtrInput
	// List of RBAC groups.
	RbacGroups pulumi.StringArrayInput
	// Whether to sign SAML AuthnRequests
	SignAuthnRequests pulumi.BoolPtrInput
}

func (AviatrixSamlEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSamlEndpointState)(nil)).Elem()
}

type aviatrixSamlEndpointArgs struct {
	// Access type.
	AccessSetBy *string `pulumi:"accessSetBy"`
	// Switch to differentiate if it is for controller login.
	ControllerLogin *bool `pulumi:"controllerLogin"`
	// Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
	CustomEntityId *string `pulumi:"customEntityId"`
	// Custom SAML Request Template.
	CustomSamlRequestTemplate *string `pulumi:"customSamlRequestTemplate"`
	// SAML Endpoint Name.
	EndpointName string `pulumi:"endpointName"`
	// IDP Metadata.
	IdpMetadata *string `pulumi:"idpMetadata"`
	// Type of IDP Metadata.
	IdpMetadataType string `pulumi:"idpMetadataType"`
	// IDP Metadata.
	IdpMetadataUrl *string `pulumi:"idpMetadataUrl"`
	// List of RBAC groups.
	RbacGroups []string `pulumi:"rbacGroups"`
	// Whether to sign SAML AuthnRequests
	SignAuthnRequests *bool `pulumi:"signAuthnRequests"`
}

// The set of arguments for constructing a AviatrixSamlEndpoint resource.
type AviatrixSamlEndpointArgs struct {
	// Access type.
	AccessSetBy pulumi.StringPtrInput
	// Switch to differentiate if it is for controller login.
	ControllerLogin pulumi.BoolPtrInput
	// Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
	CustomEntityId pulumi.StringPtrInput
	// Custom SAML Request Template.
	CustomSamlRequestTemplate pulumi.StringPtrInput
	// SAML Endpoint Name.
	EndpointName pulumi.StringInput
	// IDP Metadata.
	IdpMetadata pulumi.StringPtrInput
	// Type of IDP Metadata.
	IdpMetadataType pulumi.StringInput
	// IDP Metadata.
	IdpMetadataUrl pulumi.StringPtrInput
	// List of RBAC groups.
	RbacGroups pulumi.StringArrayInput
	// Whether to sign SAML AuthnRequests
	SignAuthnRequests pulumi.BoolPtrInput
}

func (AviatrixSamlEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSamlEndpointArgs)(nil)).Elem()
}

type AviatrixSamlEndpointInput interface {
	pulumi.Input

	ToAviatrixSamlEndpointOutput() AviatrixSamlEndpointOutput
	ToAviatrixSamlEndpointOutputWithContext(ctx context.Context) AviatrixSamlEndpointOutput
}

func (*AviatrixSamlEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSamlEndpoint)(nil)).Elem()
}

func (i *AviatrixSamlEndpoint) ToAviatrixSamlEndpointOutput() AviatrixSamlEndpointOutput {
	return i.ToAviatrixSamlEndpointOutputWithContext(context.Background())
}

func (i *AviatrixSamlEndpoint) ToAviatrixSamlEndpointOutputWithContext(ctx context.Context) AviatrixSamlEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSamlEndpointOutput)
}

// AviatrixSamlEndpointArrayInput is an input type that accepts AviatrixSamlEndpointArray and AviatrixSamlEndpointArrayOutput values.
// You can construct a concrete instance of `AviatrixSamlEndpointArrayInput` via:
//
//	AviatrixSamlEndpointArray{ AviatrixSamlEndpointArgs{...} }
type AviatrixSamlEndpointArrayInput interface {
	pulumi.Input

	ToAviatrixSamlEndpointArrayOutput() AviatrixSamlEndpointArrayOutput
	ToAviatrixSamlEndpointArrayOutputWithContext(context.Context) AviatrixSamlEndpointArrayOutput
}

type AviatrixSamlEndpointArray []AviatrixSamlEndpointInput

func (AviatrixSamlEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSamlEndpoint)(nil)).Elem()
}

func (i AviatrixSamlEndpointArray) ToAviatrixSamlEndpointArrayOutput() AviatrixSamlEndpointArrayOutput {
	return i.ToAviatrixSamlEndpointArrayOutputWithContext(context.Background())
}

func (i AviatrixSamlEndpointArray) ToAviatrixSamlEndpointArrayOutputWithContext(ctx context.Context) AviatrixSamlEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSamlEndpointArrayOutput)
}

// AviatrixSamlEndpointMapInput is an input type that accepts AviatrixSamlEndpointMap and AviatrixSamlEndpointMapOutput values.
// You can construct a concrete instance of `AviatrixSamlEndpointMapInput` via:
//
//	AviatrixSamlEndpointMap{ "key": AviatrixSamlEndpointArgs{...} }
type AviatrixSamlEndpointMapInput interface {
	pulumi.Input

	ToAviatrixSamlEndpointMapOutput() AviatrixSamlEndpointMapOutput
	ToAviatrixSamlEndpointMapOutputWithContext(context.Context) AviatrixSamlEndpointMapOutput
}

type AviatrixSamlEndpointMap map[string]AviatrixSamlEndpointInput

func (AviatrixSamlEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSamlEndpoint)(nil)).Elem()
}

func (i AviatrixSamlEndpointMap) ToAviatrixSamlEndpointMapOutput() AviatrixSamlEndpointMapOutput {
	return i.ToAviatrixSamlEndpointMapOutputWithContext(context.Background())
}

func (i AviatrixSamlEndpointMap) ToAviatrixSamlEndpointMapOutputWithContext(ctx context.Context) AviatrixSamlEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSamlEndpointMapOutput)
}

type AviatrixSamlEndpointOutput struct{ *pulumi.OutputState }

func (AviatrixSamlEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSamlEndpoint)(nil)).Elem()
}

func (o AviatrixSamlEndpointOutput) ToAviatrixSamlEndpointOutput() AviatrixSamlEndpointOutput {
	return o
}

func (o AviatrixSamlEndpointOutput) ToAviatrixSamlEndpointOutputWithContext(ctx context.Context) AviatrixSamlEndpointOutput {
	return o
}

// Access type.
func (o AviatrixSamlEndpointOutput) AccessSetBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringPtrOutput { return v.AccessSetBy }).(pulumi.StringPtrOutput)
}

// Switch to differentiate if it is for controller login.
func (o AviatrixSamlEndpointOutput) ControllerLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.BoolPtrOutput { return v.ControllerLogin }).(pulumi.BoolPtrOutput)
}

// Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
func (o AviatrixSamlEndpointOutput) CustomEntityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringPtrOutput { return v.CustomEntityId }).(pulumi.StringPtrOutput)
}

// Custom SAML Request Template.
func (o AviatrixSamlEndpointOutput) CustomSamlRequestTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringPtrOutput { return v.CustomSamlRequestTemplate }).(pulumi.StringPtrOutput)
}

// SAML Endpoint Name.
func (o AviatrixSamlEndpointOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringOutput { return v.EndpointName }).(pulumi.StringOutput)
}

// IDP Metadata.
func (o AviatrixSamlEndpointOutput) IdpMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringPtrOutput { return v.IdpMetadata }).(pulumi.StringPtrOutput)
}

// Type of IDP Metadata.
func (o AviatrixSamlEndpointOutput) IdpMetadataType() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringOutput { return v.IdpMetadataType }).(pulumi.StringOutput)
}

// IDP Metadata.
func (o AviatrixSamlEndpointOutput) IdpMetadataUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringPtrOutput { return v.IdpMetadataUrl }).(pulumi.StringPtrOutput)
}

// List of RBAC groups.
func (o AviatrixSamlEndpointOutput) RbacGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.StringArrayOutput { return v.RbacGroups }).(pulumi.StringArrayOutput)
}

// Whether to sign SAML AuthnRequests
func (o AviatrixSamlEndpointOutput) SignAuthnRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixSamlEndpoint) pulumi.BoolPtrOutput { return v.SignAuthnRequests }).(pulumi.BoolPtrOutput)
}

type AviatrixSamlEndpointArrayOutput struct{ *pulumi.OutputState }

func (AviatrixSamlEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSamlEndpoint)(nil)).Elem()
}

func (o AviatrixSamlEndpointArrayOutput) ToAviatrixSamlEndpointArrayOutput() AviatrixSamlEndpointArrayOutput {
	return o
}

func (o AviatrixSamlEndpointArrayOutput) ToAviatrixSamlEndpointArrayOutputWithContext(ctx context.Context) AviatrixSamlEndpointArrayOutput {
	return o
}

func (o AviatrixSamlEndpointArrayOutput) Index(i pulumi.IntInput) AviatrixSamlEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixSamlEndpoint {
		return vs[0].([]*AviatrixSamlEndpoint)[vs[1].(int)]
	}).(AviatrixSamlEndpointOutput)
}

type AviatrixSamlEndpointMapOutput struct{ *pulumi.OutputState }

func (AviatrixSamlEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSamlEndpoint)(nil)).Elem()
}

func (o AviatrixSamlEndpointMapOutput) ToAviatrixSamlEndpointMapOutput() AviatrixSamlEndpointMapOutput {
	return o
}

func (o AviatrixSamlEndpointMapOutput) ToAviatrixSamlEndpointMapOutputWithContext(ctx context.Context) AviatrixSamlEndpointMapOutput {
	return o
}

func (o AviatrixSamlEndpointMapOutput) MapIndex(k pulumi.StringInput) AviatrixSamlEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixSamlEndpoint {
		return vs[0].(map[string]*AviatrixSamlEndpoint)[vs[1].(string)]
	}).(AviatrixSamlEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSamlEndpointInput)(nil)).Elem(), &AviatrixSamlEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSamlEndpointArrayInput)(nil)).Elem(), AviatrixSamlEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSamlEndpointMapInput)(nil)).Elem(), AviatrixSamlEndpointMap{})
	pulumi.RegisterOutputType(AviatrixSamlEndpointOutput{})
	pulumi.RegisterOutputType(AviatrixSamlEndpointArrayOutput{})
	pulumi.RegisterOutputType(AviatrixSamlEndpointMapOutput{})
}
// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_rbac_group_permission_attachment** resource allows the creation and management of permission attachments to Aviatrix (Role-Based Access Control) RBAC groups.
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
//			_, err := aviatrix.NewAviatrixRbacGroupPermissionAttachment(ctx, "testAttachment", &aviatrix.AviatrixRbacGroupPermissionAttachmentArgs{
//				GroupName:      pulumi.String("write_only"),
//				PermissionName: pulumi.String("all_write"),
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
// **rbac_group_permission_attachment** can be imported using the `group_name` and `permission_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment test group_name~permission_name
//
// ```
type AviatrixRbacGroupPermissionAttachment struct {
	pulumi.CustomResourceState

	// This parameter represents the name of a RBAC group.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// This parameter represents the permission to attach to the RBAC group.
	PermissionName pulumi.StringOutput `pulumi:"permissionName"`
}

// NewAviatrixRbacGroupPermissionAttachment registers a new resource with the given unique name, arguments, and options.
func NewAviatrixRbacGroupPermissionAttachment(ctx *pulumi.Context,
	name string, args *AviatrixRbacGroupPermissionAttachmentArgs, opts ...pulumi.ResourceOption) (*AviatrixRbacGroupPermissionAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.PermissionName == nil {
		return nil, errors.New("invalid value for required argument 'PermissionName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixRbacGroupPermissionAttachment
	err := ctx.RegisterResource("aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixRbacGroupPermissionAttachment gets an existing AviatrixRbacGroupPermissionAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixRbacGroupPermissionAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixRbacGroupPermissionAttachmentState, opts ...pulumi.ResourceOption) (*AviatrixRbacGroupPermissionAttachment, error) {
	var resource AviatrixRbacGroupPermissionAttachment
	err := ctx.ReadResource("aviatrix:index/aviatrixRbacGroupPermissionAttachment:AviatrixRbacGroupPermissionAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixRbacGroupPermissionAttachment resources.
type aviatrixRbacGroupPermissionAttachmentState struct {
	// This parameter represents the name of a RBAC group.
	GroupName *string `pulumi:"groupName"`
	// This parameter represents the permission to attach to the RBAC group.
	PermissionName *string `pulumi:"permissionName"`
}

type AviatrixRbacGroupPermissionAttachmentState struct {
	// This parameter represents the name of a RBAC group.
	GroupName pulumi.StringPtrInput
	// This parameter represents the permission to attach to the RBAC group.
	PermissionName pulumi.StringPtrInput
}

func (AviatrixRbacGroupPermissionAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixRbacGroupPermissionAttachmentState)(nil)).Elem()
}

type aviatrixRbacGroupPermissionAttachmentArgs struct {
	// This parameter represents the name of a RBAC group.
	GroupName string `pulumi:"groupName"`
	// This parameter represents the permission to attach to the RBAC group.
	PermissionName string `pulumi:"permissionName"`
}

// The set of arguments for constructing a AviatrixRbacGroupPermissionAttachment resource.
type AviatrixRbacGroupPermissionAttachmentArgs struct {
	// This parameter represents the name of a RBAC group.
	GroupName pulumi.StringInput
	// This parameter represents the permission to attach to the RBAC group.
	PermissionName pulumi.StringInput
}

func (AviatrixRbacGroupPermissionAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixRbacGroupPermissionAttachmentArgs)(nil)).Elem()
}

type AviatrixRbacGroupPermissionAttachmentInput interface {
	pulumi.Input

	ToAviatrixRbacGroupPermissionAttachmentOutput() AviatrixRbacGroupPermissionAttachmentOutput
	ToAviatrixRbacGroupPermissionAttachmentOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentOutput
}

func (*AviatrixRbacGroupPermissionAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixRbacGroupPermissionAttachment)(nil)).Elem()
}

func (i *AviatrixRbacGroupPermissionAttachment) ToAviatrixRbacGroupPermissionAttachmentOutput() AviatrixRbacGroupPermissionAttachmentOutput {
	return i.ToAviatrixRbacGroupPermissionAttachmentOutputWithContext(context.Background())
}

func (i *AviatrixRbacGroupPermissionAttachment) ToAviatrixRbacGroupPermissionAttachmentOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixRbacGroupPermissionAttachmentOutput)
}

// AviatrixRbacGroupPermissionAttachmentArrayInput is an input type that accepts AviatrixRbacGroupPermissionAttachmentArray and AviatrixRbacGroupPermissionAttachmentArrayOutput values.
// You can construct a concrete instance of `AviatrixRbacGroupPermissionAttachmentArrayInput` via:
//
//	AviatrixRbacGroupPermissionAttachmentArray{ AviatrixRbacGroupPermissionAttachmentArgs{...} }
type AviatrixRbacGroupPermissionAttachmentArrayInput interface {
	pulumi.Input

	ToAviatrixRbacGroupPermissionAttachmentArrayOutput() AviatrixRbacGroupPermissionAttachmentArrayOutput
	ToAviatrixRbacGroupPermissionAttachmentArrayOutputWithContext(context.Context) AviatrixRbacGroupPermissionAttachmentArrayOutput
}

type AviatrixRbacGroupPermissionAttachmentArray []AviatrixRbacGroupPermissionAttachmentInput

func (AviatrixRbacGroupPermissionAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixRbacGroupPermissionAttachment)(nil)).Elem()
}

func (i AviatrixRbacGroupPermissionAttachmentArray) ToAviatrixRbacGroupPermissionAttachmentArrayOutput() AviatrixRbacGroupPermissionAttachmentArrayOutput {
	return i.ToAviatrixRbacGroupPermissionAttachmentArrayOutputWithContext(context.Background())
}

func (i AviatrixRbacGroupPermissionAttachmentArray) ToAviatrixRbacGroupPermissionAttachmentArrayOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixRbacGroupPermissionAttachmentArrayOutput)
}

// AviatrixRbacGroupPermissionAttachmentMapInput is an input type that accepts AviatrixRbacGroupPermissionAttachmentMap and AviatrixRbacGroupPermissionAttachmentMapOutput values.
// You can construct a concrete instance of `AviatrixRbacGroupPermissionAttachmentMapInput` via:
//
//	AviatrixRbacGroupPermissionAttachmentMap{ "key": AviatrixRbacGroupPermissionAttachmentArgs{...} }
type AviatrixRbacGroupPermissionAttachmentMapInput interface {
	pulumi.Input

	ToAviatrixRbacGroupPermissionAttachmentMapOutput() AviatrixRbacGroupPermissionAttachmentMapOutput
	ToAviatrixRbacGroupPermissionAttachmentMapOutputWithContext(context.Context) AviatrixRbacGroupPermissionAttachmentMapOutput
}

type AviatrixRbacGroupPermissionAttachmentMap map[string]AviatrixRbacGroupPermissionAttachmentInput

func (AviatrixRbacGroupPermissionAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixRbacGroupPermissionAttachment)(nil)).Elem()
}

func (i AviatrixRbacGroupPermissionAttachmentMap) ToAviatrixRbacGroupPermissionAttachmentMapOutput() AviatrixRbacGroupPermissionAttachmentMapOutput {
	return i.ToAviatrixRbacGroupPermissionAttachmentMapOutputWithContext(context.Background())
}

func (i AviatrixRbacGroupPermissionAttachmentMap) ToAviatrixRbacGroupPermissionAttachmentMapOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixRbacGroupPermissionAttachmentMapOutput)
}

type AviatrixRbacGroupPermissionAttachmentOutput struct{ *pulumi.OutputState }

func (AviatrixRbacGroupPermissionAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixRbacGroupPermissionAttachment)(nil)).Elem()
}

func (o AviatrixRbacGroupPermissionAttachmentOutput) ToAviatrixRbacGroupPermissionAttachmentOutput() AviatrixRbacGroupPermissionAttachmentOutput {
	return o
}

func (o AviatrixRbacGroupPermissionAttachmentOutput) ToAviatrixRbacGroupPermissionAttachmentOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentOutput {
	return o
}

// This parameter represents the name of a RBAC group.
func (o AviatrixRbacGroupPermissionAttachmentOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixRbacGroupPermissionAttachment) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// This parameter represents the permission to attach to the RBAC group.
func (o AviatrixRbacGroupPermissionAttachmentOutput) PermissionName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixRbacGroupPermissionAttachment) pulumi.StringOutput { return v.PermissionName }).(pulumi.StringOutput)
}

type AviatrixRbacGroupPermissionAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AviatrixRbacGroupPermissionAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixRbacGroupPermissionAttachment)(nil)).Elem()
}

func (o AviatrixRbacGroupPermissionAttachmentArrayOutput) ToAviatrixRbacGroupPermissionAttachmentArrayOutput() AviatrixRbacGroupPermissionAttachmentArrayOutput {
	return o
}

func (o AviatrixRbacGroupPermissionAttachmentArrayOutput) ToAviatrixRbacGroupPermissionAttachmentArrayOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentArrayOutput {
	return o
}

func (o AviatrixRbacGroupPermissionAttachmentArrayOutput) Index(i pulumi.IntInput) AviatrixRbacGroupPermissionAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixRbacGroupPermissionAttachment {
		return vs[0].([]*AviatrixRbacGroupPermissionAttachment)[vs[1].(int)]
	}).(AviatrixRbacGroupPermissionAttachmentOutput)
}

type AviatrixRbacGroupPermissionAttachmentMapOutput struct{ *pulumi.OutputState }

func (AviatrixRbacGroupPermissionAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixRbacGroupPermissionAttachment)(nil)).Elem()
}

func (o AviatrixRbacGroupPermissionAttachmentMapOutput) ToAviatrixRbacGroupPermissionAttachmentMapOutput() AviatrixRbacGroupPermissionAttachmentMapOutput {
	return o
}

func (o AviatrixRbacGroupPermissionAttachmentMapOutput) ToAviatrixRbacGroupPermissionAttachmentMapOutputWithContext(ctx context.Context) AviatrixRbacGroupPermissionAttachmentMapOutput {
	return o
}

func (o AviatrixRbacGroupPermissionAttachmentMapOutput) MapIndex(k pulumi.StringInput) AviatrixRbacGroupPermissionAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixRbacGroupPermissionAttachment {
		return vs[0].(map[string]*AviatrixRbacGroupPermissionAttachment)[vs[1].(string)]
	}).(AviatrixRbacGroupPermissionAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixRbacGroupPermissionAttachmentInput)(nil)).Elem(), &AviatrixRbacGroupPermissionAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixRbacGroupPermissionAttachmentArrayInput)(nil)).Elem(), AviatrixRbacGroupPermissionAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixRbacGroupPermissionAttachmentMapInput)(nil)).Elem(), AviatrixRbacGroupPermissionAttachmentMap{})
	pulumi.RegisterOutputType(AviatrixRbacGroupPermissionAttachmentOutput{})
	pulumi.RegisterOutputType(AviatrixRbacGroupPermissionAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AviatrixRbacGroupPermissionAttachmentMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_account_user** resource allows the creation and management of Aviatrix user accounts.
//
// > **NOTE:** With the release of Controller 5.4 (compatible with Aviatrix provider R2.13), Role-Based Access Control (RBAC) is now integrated into the Accounts workflow. Any **aviatrix_account_user** created in 5.3 by default will have admin privileges (attached to the 'admin' RBAC permission group). In 5.4, any new account users created will no longer have the option to specify an `accountName`, but rather have the option to attach the user to specific RBAC groups through the **aviatrix_rbac_group_user_attachment** resource for more granular security control. Account users created in 5.4 will have minimal access (read_only) unless otherwise specified in the RBAC group permissions that the users are attached to.
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
//			_, err := aviatrix.NewAviatrixAccountUser(ctx, "testAccountuser", &aviatrix.AviatrixAccountUserArgs{
//				Email:    pulumi.String("username1@testdomain.com"),
//				Password: pulumi.String("passwordforuser1-1234"),
//				Username: pulumi.String("username1"),
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
// **account_user** can be imported using the `username` (when doing import, need to leave `password` argument blank), e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixAccountUser:AviatrixAccountUser test username
//
// ```
type AviatrixAccountUser struct {
	pulumi.CustomResourceState

	// Email of address of account user to be created.
	Email pulumi.StringOutput `pulumi:"email"`
	// Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created.
	Password pulumi.StringOutput `pulumi:"password"`
	// Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewAviatrixAccountUser registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAccountUser(ctx *pulumi.Context,
	name string, args *AviatrixAccountUserArgs, opts ...pulumi.ResourceOption) (*AviatrixAccountUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAccountUser
	err := ctx.RegisterResource("aviatrix:index/aviatrixAccountUser:AviatrixAccountUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAccountUser gets an existing AviatrixAccountUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAccountUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAccountUserState, opts ...pulumi.ResourceOption) (*AviatrixAccountUser, error) {
	var resource AviatrixAccountUser
	err := ctx.ReadResource("aviatrix:index/aviatrixAccountUser:AviatrixAccountUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAccountUser resources.
type aviatrixAccountUserState struct {
	// Email of address of account user to be created.
	Email *string `pulumi:"email"`
	// Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created.
	Password *string `pulumi:"password"`
	// Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.
	Username *string `pulumi:"username"`
}

type AviatrixAccountUserState struct {
	// Email of address of account user to be created.
	Email pulumi.StringPtrInput
	// Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created.
	Password pulumi.StringPtrInput
	// Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.
	Username pulumi.StringPtrInput
}

func (AviatrixAccountUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAccountUserState)(nil)).Elem()
}

type aviatrixAccountUserArgs struct {
	// Email of address of account user to be created.
	Email string `pulumi:"email"`
	// Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created.
	Password string `pulumi:"password"`
	// Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a AviatrixAccountUser resource.
type AviatrixAccountUserArgs struct {
	// Email of address of account user to be created.
	Email pulumi.StringInput
	// Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created.
	Password pulumi.StringInput
	// Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.
	Username pulumi.StringInput
}

func (AviatrixAccountUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAccountUserArgs)(nil)).Elem()
}

type AviatrixAccountUserInput interface {
	pulumi.Input

	ToAviatrixAccountUserOutput() AviatrixAccountUserOutput
	ToAviatrixAccountUserOutputWithContext(ctx context.Context) AviatrixAccountUserOutput
}

func (*AviatrixAccountUser) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAccountUser)(nil)).Elem()
}

func (i *AviatrixAccountUser) ToAviatrixAccountUserOutput() AviatrixAccountUserOutput {
	return i.ToAviatrixAccountUserOutputWithContext(context.Background())
}

func (i *AviatrixAccountUser) ToAviatrixAccountUserOutputWithContext(ctx context.Context) AviatrixAccountUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAccountUserOutput)
}

// AviatrixAccountUserArrayInput is an input type that accepts AviatrixAccountUserArray and AviatrixAccountUserArrayOutput values.
// You can construct a concrete instance of `AviatrixAccountUserArrayInput` via:
//
//	AviatrixAccountUserArray{ AviatrixAccountUserArgs{...} }
type AviatrixAccountUserArrayInput interface {
	pulumi.Input

	ToAviatrixAccountUserArrayOutput() AviatrixAccountUserArrayOutput
	ToAviatrixAccountUserArrayOutputWithContext(context.Context) AviatrixAccountUserArrayOutput
}

type AviatrixAccountUserArray []AviatrixAccountUserInput

func (AviatrixAccountUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAccountUser)(nil)).Elem()
}

func (i AviatrixAccountUserArray) ToAviatrixAccountUserArrayOutput() AviatrixAccountUserArrayOutput {
	return i.ToAviatrixAccountUserArrayOutputWithContext(context.Background())
}

func (i AviatrixAccountUserArray) ToAviatrixAccountUserArrayOutputWithContext(ctx context.Context) AviatrixAccountUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAccountUserArrayOutput)
}

// AviatrixAccountUserMapInput is an input type that accepts AviatrixAccountUserMap and AviatrixAccountUserMapOutput values.
// You can construct a concrete instance of `AviatrixAccountUserMapInput` via:
//
//	AviatrixAccountUserMap{ "key": AviatrixAccountUserArgs{...} }
type AviatrixAccountUserMapInput interface {
	pulumi.Input

	ToAviatrixAccountUserMapOutput() AviatrixAccountUserMapOutput
	ToAviatrixAccountUserMapOutputWithContext(context.Context) AviatrixAccountUserMapOutput
}

type AviatrixAccountUserMap map[string]AviatrixAccountUserInput

func (AviatrixAccountUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAccountUser)(nil)).Elem()
}

func (i AviatrixAccountUserMap) ToAviatrixAccountUserMapOutput() AviatrixAccountUserMapOutput {
	return i.ToAviatrixAccountUserMapOutputWithContext(context.Background())
}

func (i AviatrixAccountUserMap) ToAviatrixAccountUserMapOutputWithContext(ctx context.Context) AviatrixAccountUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAccountUserMapOutput)
}

type AviatrixAccountUserOutput struct{ *pulumi.OutputState }

func (AviatrixAccountUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAccountUser)(nil)).Elem()
}

func (o AviatrixAccountUserOutput) ToAviatrixAccountUserOutput() AviatrixAccountUserOutput {
	return o
}

func (o AviatrixAccountUserOutput) ToAviatrixAccountUserOutputWithContext(ctx context.Context) AviatrixAccountUserOutput {
	return o
}

// Email of address of account user to be created.
func (o AviatrixAccountUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAccountUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created.
func (o AviatrixAccountUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAccountUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.
func (o AviatrixAccountUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAccountUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type AviatrixAccountUserArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAccountUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAccountUser)(nil)).Elem()
}

func (o AviatrixAccountUserArrayOutput) ToAviatrixAccountUserArrayOutput() AviatrixAccountUserArrayOutput {
	return o
}

func (o AviatrixAccountUserArrayOutput) ToAviatrixAccountUserArrayOutputWithContext(ctx context.Context) AviatrixAccountUserArrayOutput {
	return o
}

func (o AviatrixAccountUserArrayOutput) Index(i pulumi.IntInput) AviatrixAccountUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAccountUser {
		return vs[0].([]*AviatrixAccountUser)[vs[1].(int)]
	}).(AviatrixAccountUserOutput)
}

type AviatrixAccountUserMapOutput struct{ *pulumi.OutputState }

func (AviatrixAccountUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAccountUser)(nil)).Elem()
}

func (o AviatrixAccountUserMapOutput) ToAviatrixAccountUserMapOutput() AviatrixAccountUserMapOutput {
	return o
}

func (o AviatrixAccountUserMapOutput) ToAviatrixAccountUserMapOutputWithContext(ctx context.Context) AviatrixAccountUserMapOutput {
	return o
}

func (o AviatrixAccountUserMapOutput) MapIndex(k pulumi.StringInput) AviatrixAccountUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAccountUser {
		return vs[0].(map[string]*AviatrixAccountUser)[vs[1].(string)]
	}).(AviatrixAccountUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAccountUserInput)(nil)).Elem(), &AviatrixAccountUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAccountUserArrayInput)(nil)).Elem(), AviatrixAccountUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAccountUserMapInput)(nil)).Elem(), AviatrixAccountUserMap{})
	pulumi.RegisterOutputType(AviatrixAccountUserOutput{})
	pulumi.RegisterOutputType(AviatrixAccountUserArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAccountUserMapOutput{})
}

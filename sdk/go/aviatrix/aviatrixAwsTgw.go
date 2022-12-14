// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **aviatrix_aws_tgw** resource allows the creation and management of Aviatrix-created AWS TGWs.
//
// > **NOTE:** If you are planning to attach VPCs to the **aviatrix_aws_tgw** resource and anticipate updating it often and/or using advanced options such as customized route advertisement, we highly recommend managing those VPCs outside this resource by setting `manageVpcAttachment` to false and using the **aviatrix_aws_tgw_vpc_attachment** resource instead of the in-line `attachedVpc {}` block.
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
//			_, err := aviatrix.NewAviatrixAwsTgw(ctx, "testAwsTgw", &aviatrix.AviatrixAwsTgwArgs{
//				AccountName:                    pulumi.String("devops"),
//				AwsSideAsNumber:                pulumi.String("64512"),
//				ManageTransitGatewayAttachment: pulumi.Bool(false),
//				ManageVpcAttachment:            pulumi.Bool(false),
//				Region:                         pulumi.String("us-east-1"),
//				SecurityDomains: AviatrixAwsTgwSecurityDomainTypeArray{
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Default_Domain"),
//							pulumi.String("Shared_Service_Domain"),
//							pulumi.String("mysdn1"),
//						},
//						SecurityDomainName: pulumi.String("Aviatrix_Edge_Domain"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Aviatrix_Edge_Domain"),
//							pulumi.String("Shared_Service_Domain"),
//						},
//						SecurityDomainName: pulumi.String("Default_Domain"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Aviatrix_Edge_Domain"),
//							pulumi.String("Default_Domain"),
//						},
//						SecurityDomainName: pulumi.String("Shared_Service_Domain"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Aviatrix_Edge_Domain"),
//						},
//						SecurityDomainName: pulumi.String("SDN1"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						SecurityDomainName: pulumi.String("mysdn2"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						AviatrixFirewall:   pulumi.Bool(true),
//						SecurityDomainName: pulumi.String("firewall-domain"),
//					},
//				},
//				TgwName: pulumi.String("test-AWS-TGW"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := aviatrix.NewAviatrixAwsTgw(ctx, "testAwsGovTgw", &aviatrix.AviatrixAwsTgwArgs{
//				AccountName:                    pulumi.String("devops"),
//				AwsSideAsNumber:                pulumi.String("64512"),
//				CloudType:                      pulumi.Int(256),
//				ManageTransitGatewayAttachment: pulumi.Bool(false),
//				ManageVpcAttachment:            pulumi.Bool(false),
//				Region:                         pulumi.String("us-gov-east-1"),
//				SecurityDomains: AviatrixAwsTgwSecurityDomainTypeArray{
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Default_Domain"),
//							pulumi.String("Shared_Service_Domain"),
//							pulumi.String("mysdn1"),
//						},
//						SecurityDomainName: pulumi.String("Aviatrix_Edge_Domain"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Aviatrix_Edge_Domain"),
//							pulumi.String("Shared_Service_Domain"),
//						},
//						SecurityDomainName: pulumi.String("Default_Domain"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Aviatrix_Edge_Domain"),
//							pulumi.String("Default_Domain"),
//						},
//						SecurityDomainName: pulumi.String("Shared_Service_Domain"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						ConnectedDomains: pulumi.StringArray{
//							pulumi.String("Aviatrix_Edge_Domain"),
//						},
//						SecurityDomainName: pulumi.String("SDN1"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						SecurityDomainName: pulumi.String("mysdn2"),
//					},
//					&AviatrixAwsTgwSecurityDomainTypeArgs{
//						AviatrixFirewall:   pulumi.Bool(true),
//						SecurityDomainName: pulumi.String("firewall-domain"),
//					},
//				},
//				TgwName: pulumi.String("test-AWSGov-TGW"),
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
// **aws_tgw** can be imported using the `tgw_name`, e.g.
//
// ```sh
//
//	$ pulumi import aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw test tgw_name
//
// ```
type AviatrixAwsTgw struct {
	pulumi.CustomResourceState

	// Name of the cloud account in the Aviatrix controller.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
	//
	// Deprecated: Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
	AttachedAviatrixTransitGateways pulumi.StringArrayOutput `pulumi:"attachedAviatrixTransitGateways"`
	// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
	AwsSideAsNumber pulumi.StringOutput `pulumi:"awsSideAsNumber"`
	// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
	Cidrs pulumi.StringArrayOutput `pulumi:"cidrs"`
	// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
	CloudType pulumi.IntPtrOutput `pulumi:"cloudType"`
	// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
	EnableMulticast pulumi.BoolPtrOutput `pulumi:"enableMulticast"`
	// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
	InspectionMode pulumi.StringPtrOutput `pulumi:"inspectionMode"`
	// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
	ManageSecurityDomain pulumi.BoolPtrOutput `pulumi:"manageSecurityDomain"`
	// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
	ManageTransitGatewayAttachment pulumi.BoolPtrOutput `pulumi:"manageTransitGatewayAttachment"`
	// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
	ManageVpcAttachment pulumi.BoolPtrOutput `pulumi:"manageVpcAttachment"`
	// AWS region of AWS TGW to be created in
	Region pulumi.StringOutput `pulumi:"region"`
	// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
	//
	// Deprecated: Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
	SecurityDomains AviatrixAwsTgwSecurityDomainTypeArrayOutput `pulumi:"securityDomains"`
	// TGW ID. Available as of provider version R2.19+.
	TgwId pulumi.StringOutput `pulumi:"tgwId"`
	// Name of the AWS TGW to be created
	TgwName pulumi.StringOutput `pulumi:"tgwName"`
}

// NewAviatrixAwsTgw registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAwsTgw(ctx *pulumi.Context,
	name string, args *AviatrixAwsTgwArgs, opts ...pulumi.ResourceOption) (*AviatrixAwsTgw, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AwsSideAsNumber == nil {
		return nil, errors.New("invalid value for required argument 'AwsSideAsNumber'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.TgwName == nil {
		return nil, errors.New("invalid value for required argument 'TgwName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAwsTgw
	err := ctx.RegisterResource("aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAwsTgw gets an existing AviatrixAwsTgw resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAwsTgw(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAwsTgwState, opts ...pulumi.ResourceOption) (*AviatrixAwsTgw, error) {
	var resource AviatrixAwsTgw
	err := ctx.ReadResource("aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAwsTgw resources.
type aviatrixAwsTgwState struct {
	// Name of the cloud account in the Aviatrix controller.
	AccountName *string `pulumi:"accountName"`
	// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
	//
	// Deprecated: Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
	AttachedAviatrixTransitGateways []string `pulumi:"attachedAviatrixTransitGateways"`
	// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
	AwsSideAsNumber *string `pulumi:"awsSideAsNumber"`
	// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
	Cidrs []string `pulumi:"cidrs"`
	// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
	CloudType *int `pulumi:"cloudType"`
	// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
	EnableMulticast *bool `pulumi:"enableMulticast"`
	// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
	InspectionMode *string `pulumi:"inspectionMode"`
	// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
	ManageSecurityDomain *bool `pulumi:"manageSecurityDomain"`
	// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
	ManageTransitGatewayAttachment *bool `pulumi:"manageTransitGatewayAttachment"`
	// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
	ManageVpcAttachment *bool `pulumi:"manageVpcAttachment"`
	// AWS region of AWS TGW to be created in
	Region *string `pulumi:"region"`
	// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
	//
	// Deprecated: Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
	SecurityDomains []AviatrixAwsTgwSecurityDomainType `pulumi:"securityDomains"`
	// TGW ID. Available as of provider version R2.19+.
	TgwId *string `pulumi:"tgwId"`
	// Name of the AWS TGW to be created
	TgwName *string `pulumi:"tgwName"`
}

type AviatrixAwsTgwState struct {
	// Name of the cloud account in the Aviatrix controller.
	AccountName pulumi.StringPtrInput
	// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
	//
	// Deprecated: Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
	AttachedAviatrixTransitGateways pulumi.StringArrayInput
	// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
	AwsSideAsNumber pulumi.StringPtrInput
	// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
	Cidrs pulumi.StringArrayInput
	// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
	CloudType pulumi.IntPtrInput
	// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
	EnableMulticast pulumi.BoolPtrInput
	// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
	InspectionMode pulumi.StringPtrInput
	// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
	ManageSecurityDomain pulumi.BoolPtrInput
	// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
	ManageTransitGatewayAttachment pulumi.BoolPtrInput
	// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
	ManageVpcAttachment pulumi.BoolPtrInput
	// AWS region of AWS TGW to be created in
	Region pulumi.StringPtrInput
	// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
	//
	// Deprecated: Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
	SecurityDomains AviatrixAwsTgwSecurityDomainTypeArrayInput
	// TGW ID. Available as of provider version R2.19+.
	TgwId pulumi.StringPtrInput
	// Name of the AWS TGW to be created
	TgwName pulumi.StringPtrInput
}

func (AviatrixAwsTgwState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwState)(nil)).Elem()
}

type aviatrixAwsTgwArgs struct {
	// Name of the cloud account in the Aviatrix controller.
	AccountName string `pulumi:"accountName"`
	// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
	//
	// Deprecated: Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
	AttachedAviatrixTransitGateways []string `pulumi:"attachedAviatrixTransitGateways"`
	// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
	AwsSideAsNumber string `pulumi:"awsSideAsNumber"`
	// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
	Cidrs []string `pulumi:"cidrs"`
	// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
	CloudType *int `pulumi:"cloudType"`
	// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
	EnableMulticast *bool `pulumi:"enableMulticast"`
	// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
	InspectionMode *string `pulumi:"inspectionMode"`
	// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
	ManageSecurityDomain *bool `pulumi:"manageSecurityDomain"`
	// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
	ManageTransitGatewayAttachment *bool `pulumi:"manageTransitGatewayAttachment"`
	// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
	ManageVpcAttachment *bool `pulumi:"manageVpcAttachment"`
	// AWS region of AWS TGW to be created in
	Region string `pulumi:"region"`
	// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
	//
	// Deprecated: Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
	SecurityDomains []AviatrixAwsTgwSecurityDomainType `pulumi:"securityDomains"`
	// Name of the AWS TGW to be created
	TgwName string `pulumi:"tgwName"`
}

// The set of arguments for constructing a AviatrixAwsTgw resource.
type AviatrixAwsTgwArgs struct {
	// Name of the cloud account in the Aviatrix controller.
	AccountName pulumi.StringInput
	// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
	//
	// Deprecated: Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
	AttachedAviatrixTransitGateways pulumi.StringArrayInput
	// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
	AwsSideAsNumber pulumi.StringInput
	// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
	Cidrs pulumi.StringArrayInput
	// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
	CloudType pulumi.IntPtrInput
	// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
	EnableMulticast pulumi.BoolPtrInput
	// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
	InspectionMode pulumi.StringPtrInput
	// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
	ManageSecurityDomain pulumi.BoolPtrInput
	// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
	ManageTransitGatewayAttachment pulumi.BoolPtrInput
	// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
	ManageVpcAttachment pulumi.BoolPtrInput
	// AWS region of AWS TGW to be created in
	Region pulumi.StringInput
	// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
	//
	// Deprecated: Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
	SecurityDomains AviatrixAwsTgwSecurityDomainTypeArrayInput
	// Name of the AWS TGW to be created
	TgwName pulumi.StringInput
}

func (AviatrixAwsTgwArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwArgs)(nil)).Elem()
}

type AviatrixAwsTgwInput interface {
	pulumi.Input

	ToAviatrixAwsTgwOutput() AviatrixAwsTgwOutput
	ToAviatrixAwsTgwOutputWithContext(ctx context.Context) AviatrixAwsTgwOutput
}

func (*AviatrixAwsTgw) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgw)(nil)).Elem()
}

func (i *AviatrixAwsTgw) ToAviatrixAwsTgwOutput() AviatrixAwsTgwOutput {
	return i.ToAviatrixAwsTgwOutputWithContext(context.Background())
}

func (i *AviatrixAwsTgw) ToAviatrixAwsTgwOutputWithContext(ctx context.Context) AviatrixAwsTgwOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwOutput)
}

// AviatrixAwsTgwArrayInput is an input type that accepts AviatrixAwsTgwArray and AviatrixAwsTgwArrayOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwArrayInput` via:
//
//	AviatrixAwsTgwArray{ AviatrixAwsTgwArgs{...} }
type AviatrixAwsTgwArrayInput interface {
	pulumi.Input

	ToAviatrixAwsTgwArrayOutput() AviatrixAwsTgwArrayOutput
	ToAviatrixAwsTgwArrayOutputWithContext(context.Context) AviatrixAwsTgwArrayOutput
}

type AviatrixAwsTgwArray []AviatrixAwsTgwInput

func (AviatrixAwsTgwArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgw)(nil)).Elem()
}

func (i AviatrixAwsTgwArray) ToAviatrixAwsTgwArrayOutput() AviatrixAwsTgwArrayOutput {
	return i.ToAviatrixAwsTgwArrayOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwArray) ToAviatrixAwsTgwArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwArrayOutput)
}

// AviatrixAwsTgwMapInput is an input type that accepts AviatrixAwsTgwMap and AviatrixAwsTgwMapOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwMapInput` via:
//
//	AviatrixAwsTgwMap{ "key": AviatrixAwsTgwArgs{...} }
type AviatrixAwsTgwMapInput interface {
	pulumi.Input

	ToAviatrixAwsTgwMapOutput() AviatrixAwsTgwMapOutput
	ToAviatrixAwsTgwMapOutputWithContext(context.Context) AviatrixAwsTgwMapOutput
}

type AviatrixAwsTgwMap map[string]AviatrixAwsTgwInput

func (AviatrixAwsTgwMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgw)(nil)).Elem()
}

func (i AviatrixAwsTgwMap) ToAviatrixAwsTgwMapOutput() AviatrixAwsTgwMapOutput {
	return i.ToAviatrixAwsTgwMapOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwMap) ToAviatrixAwsTgwMapOutputWithContext(ctx context.Context) AviatrixAwsTgwMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwMapOutput)
}

type AviatrixAwsTgwOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgw)(nil)).Elem()
}

func (o AviatrixAwsTgwOutput) ToAviatrixAwsTgwOutput() AviatrixAwsTgwOutput {
	return o
}

func (o AviatrixAwsTgwOutput) ToAviatrixAwsTgwOutputWithContext(ctx context.Context) AviatrixAwsTgwOutput {
	return o
}

// Name of the cloud account in the Aviatrix controller.
func (o AviatrixAwsTgwOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.
//
// Deprecated: Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
func (o AviatrixAwsTgwOutput) AttachedAviatrixTransitGateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringArrayOutput { return v.AttachedAviatrixTransitGateways }).(pulumi.StringArrayOutput)
}

// BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001".
func (o AviatrixAwsTgwOutput) AwsSideAsNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringOutput { return v.AwsSideAsNumber }).(pulumi.StringOutput)
}

// Set of TGW CIDRs. For example, `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]`. Available as of provider version R2.18.1+.
func (o AviatrixAwsTgwOutput) Cidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringArrayOutput { return v.Cidrs }).(pulumi.StringArrayOutput)
}

// Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.
func (o AviatrixAwsTgwOutput) CloudType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.IntPtrOutput { return v.CloudType }).(pulumi.IntPtrOutput)
}

// Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.
func (o AviatrixAwsTgwOutput) EnableMulticast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.BoolPtrOutput { return v.EnableMulticast }).(pulumi.BoolPtrOutput)
}

// Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+.
func (o AviatrixAwsTgwOutput) InspectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringPtrOutput { return v.InspectionMode }).(pulumi.StringPtrOutput)
}

// This parameter is a switch used to determine whether or not to manage security domains using the **aviatrix_aws_tgw** resource. If this is set to false, creation and management of security domains must be done using the **aviatrix_aws_tgw_security_domain** resource. Valid values: true, false. Default value: true.
func (o AviatrixAwsTgwOutput) ManageSecurityDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.BoolPtrOutput { return v.ManageSecurityDomain }).(pulumi.BoolPtrOutput)
}

// This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of transit gateways must be done using the **aviatrix_aws_tgw_transit_gateway_attachment** resource. Valid values: true, false. Default value: true.
func (o AviatrixAwsTgwOutput) ManageTransitGatewayAttachment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.BoolPtrOutput { return v.ManageTransitGatewayAttachment }).(pulumi.BoolPtrOutput)
}

// This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the **aviatrix_aws_tgw** resource. If this is set to false, attachment of VPCs must be done using the **aviatrix_aws_tgw_vpc_attachment** resource. Valid values: true, false. Default value: true.
func (o AviatrixAwsTgwOutput) ManageVpcAttachment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.BoolPtrOutput { return v.ManageVpcAttachment }).(pulumi.BoolPtrOutput)
}

// AWS region of AWS TGW to be created in
func (o AviatrixAwsTgwOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.
//
// Deprecated: Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
func (o AviatrixAwsTgwOutput) SecurityDomains() AviatrixAwsTgwSecurityDomainTypeArrayOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) AviatrixAwsTgwSecurityDomainTypeArrayOutput { return v.SecurityDomains }).(AviatrixAwsTgwSecurityDomainTypeArrayOutput)
}

// TGW ID. Available as of provider version R2.19+.
func (o AviatrixAwsTgwOutput) TgwId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringOutput { return v.TgwId }).(pulumi.StringOutput)
}

// Name of the AWS TGW to be created
func (o AviatrixAwsTgwOutput) TgwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgw) pulumi.StringOutput { return v.TgwName }).(pulumi.StringOutput)
}

type AviatrixAwsTgwArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgw)(nil)).Elem()
}

func (o AviatrixAwsTgwArrayOutput) ToAviatrixAwsTgwArrayOutput() AviatrixAwsTgwArrayOutput {
	return o
}

func (o AviatrixAwsTgwArrayOutput) ToAviatrixAwsTgwArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwArrayOutput {
	return o
}

func (o AviatrixAwsTgwArrayOutput) Index(i pulumi.IntInput) AviatrixAwsTgwOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAwsTgw {
		return vs[0].([]*AviatrixAwsTgw)[vs[1].(int)]
	}).(AviatrixAwsTgwOutput)
}

type AviatrixAwsTgwMapOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgw)(nil)).Elem()
}

func (o AviatrixAwsTgwMapOutput) ToAviatrixAwsTgwMapOutput() AviatrixAwsTgwMapOutput {
	return o
}

func (o AviatrixAwsTgwMapOutput) ToAviatrixAwsTgwMapOutputWithContext(ctx context.Context) AviatrixAwsTgwMapOutput {
	return o
}

func (o AviatrixAwsTgwMapOutput) MapIndex(k pulumi.StringInput) AviatrixAwsTgwOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAwsTgw {
		return vs[0].(map[string]*AviatrixAwsTgw)[vs[1].(string)]
	}).(AviatrixAwsTgwOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwInput)(nil)).Elem(), &AviatrixAwsTgw{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwArrayInput)(nil)).Elem(), AviatrixAwsTgwArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwMapInput)(nil)).Elem(), AviatrixAwsTgwMap{})
	pulumi.RegisterOutputType(AviatrixAwsTgwOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwMapOutput{})
}

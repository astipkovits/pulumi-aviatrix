// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The aviatrix.AviatrixSpokeVpc resource allows to create and manage Aviatrix Spoke Gateways.
 *
 * !> **WARNING:** The `aviatrix.AviatrixSpokeVpc` resource is deprecated as of **Release 2.0**. It is currently kept for backward-compatibility and will be removed in the future. Please use the spoke gateway resource instead. If this is already in the state, please remove it from the state file and import as `aviatrix.AviatrixSpokeGateway`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Set Aviatrix aws spoke_vpc
 * const testSpokeVpcAws = new aviatrix.AviatrixSpokeVpc("test_spoke_vpc_aws", {
 *     accountName: "my-aws",
 *     cloudType: 1,
 *     dnsServer: "8.8.8.8",
 *     enableNat: "no",
 *     gwName: "spoke-gw-aws",
 *     subnet: "10.11.0.0/24~~us-west-1b~~spoke-vpc-01-pubsub",
 *     tagLists: [
 *         "k1:v1",
 *         "k2:v2",
 *     ],
 *     vpcId: "vpc-abcd123~~spoke-vpc-01",
 *     vpcReg: "us-west-1",
 *     vpcSize: "t2.micro",
 * });
 * // Set Aviatrix gcp spoke_vpc
 * const testSpokeVpcGcp = new aviatrix.AviatrixSpokeVpc("test_spoke_vpc_gcp", {
 *     accountName: "my-gcp",
 *     cloudType: 4,
 *     enableNat: "no",
 *     gwName: "spoke-gw-gcp",
 *     subnet: "10.12.0.0/24",
 *     vpcId: "gcp-spoke-vpc",
 *     vpcReg: "us-west1-b",
 *     vpcSize: "t2.micro",
 * });
 * // Set Aviatrix arm spoke_vpc
 * const testSpokeVpcArm = new aviatrix.AviatrixSpokeVpc("test_spoke_vpc_arm", {
 *     accountName: "my-arm",
 *     cloudType: 8,
 *     enableNat: "no",
 *     gwName: "spoke-gw-01",
 *     subnet: "10.13.0.0/24",
 *     vpcId: "spoke:test-spoke-gw-123",
 *     vpcReg: "West US",
 *     vpcSize: "t2.micro",
 * });
 * ```
 *
 * ## Import
 *
 * Instance spoke_vpc can be imported using the gw_name, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc test gw_name
 * ```
 */
export class AviatrixSpokeVpc extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixSpokeVpc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixSpokeVpcState, opts?: pulumi.CustomResourceOptions): AviatrixSpokeVpc {
        return new AviatrixSpokeVpc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc';

    /**
     * Returns true if the given object is an instance of AviatrixSpokeVpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixSpokeVpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixSpokeVpc.__pulumiType;
    }

    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Cloud instance ID.
     */
    public /*out*/ readonly cloudInstanceId!: pulumi.Output<string>;
    /**
     * Type of cloud service provider. AWS=1, GCP=4, ARM=8.
     */
    public readonly cloudType!: pulumi.Output<number>;
    /**
     * Specify whether enabling NAT feature on the gateway or not.
     */
    public readonly enableNat!: pulumi.Output<string | undefined>;
    /**
     * Name of the gateway which is going to be created.
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
     */
    public readonly haGwSize!: pulumi.Output<string | undefined>;
    /**
     * HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.
     */
    public readonly haSubnet!: pulumi.Output<string | undefined>;
    /**
     * HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".
     */
    public readonly haZone!: pulumi.Output<string | undefined>;
    /**
     * Set to "enabled" if this feature is desired.
     */
    public readonly singleAzHa!: pulumi.Output<string | undefined>;
    /**
     * Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported
     */
    public readonly tagLists!: pulumi.Output<string[] | undefined>;
    /**
     * Specify the transit Gateway.
     */
    public readonly transitGw!: pulumi.Output<string | undefined>;
    /**
     * VPC-ID/VNet-Name of cloud provider. Required if cloudType is "1" or "4". Example: AWS: "vpc-abcd1234", etc...
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...
     */
    public readonly vpcReg!: pulumi.Output<string>;
    /**
     * Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...
     */
    public readonly vpcSize!: pulumi.Output<string>;

    /**
     * Create a AviatrixSpokeVpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixSpokeVpcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixSpokeVpcArgs | AviatrixSpokeVpcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixSpokeVpcState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["cloudInstanceId"] = state ? state.cloudInstanceId : undefined;
            resourceInputs["cloudType"] = state ? state.cloudType : undefined;
            resourceInputs["enableNat"] = state ? state.enableNat : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["haGwSize"] = state ? state.haGwSize : undefined;
            resourceInputs["haSubnet"] = state ? state.haSubnet : undefined;
            resourceInputs["haZone"] = state ? state.haZone : undefined;
            resourceInputs["singleAzHa"] = state ? state.singleAzHa : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["tagLists"] = state ? state.tagLists : undefined;
            resourceInputs["transitGw"] = state ? state.transitGw : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcReg"] = state ? state.vpcReg : undefined;
            resourceInputs["vpcSize"] = state ? state.vpcSize : undefined;
        } else {
            const args = argsOrState as AviatrixSpokeVpcArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.cloudType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudType'");
            }
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            if ((!args || args.subnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnet'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vpcReg === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcReg'");
            }
            if ((!args || args.vpcSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcSize'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["cloudType"] = args ? args.cloudType : undefined;
            resourceInputs["enableNat"] = args ? args.enableNat : undefined;
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["haGwSize"] = args ? args.haGwSize : undefined;
            resourceInputs["haSubnet"] = args ? args.haSubnet : undefined;
            resourceInputs["haZone"] = args ? args.haZone : undefined;
            resourceInputs["singleAzHa"] = args ? args.singleAzHa : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["tagLists"] = args ? args.tagLists : undefined;
            resourceInputs["transitGw"] = args ? args.transitGw : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpcReg"] = args ? args.vpcReg : undefined;
            resourceInputs["vpcSize"] = args ? args.vpcSize : undefined;
            resourceInputs["cloudInstanceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixSpokeVpc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixSpokeVpc resources.
 */
export interface AviatrixSpokeVpcState {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Cloud instance ID.
     */
    cloudInstanceId?: pulumi.Input<string>;
    /**
     * Type of cloud service provider. AWS=1, GCP=4, ARM=8.
     */
    cloudType?: pulumi.Input<number>;
    /**
     * Specify whether enabling NAT feature on the gateway or not.
     */
    enableNat?: pulumi.Input<string>;
    /**
     * Name of the gateway which is going to be created.
     */
    gwName?: pulumi.Input<string>;
    /**
     * HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
     */
    haGwSize?: pulumi.Input<string>;
    /**
     * HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.
     */
    haSubnet?: pulumi.Input<string>;
    /**
     * HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".
     */
    haZone?: pulumi.Input<string>;
    /**
     * Set to "enabled" if this feature is desired.
     */
    singleAzHa?: pulumi.Input<string>;
    /**
     * Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...
     */
    subnet?: pulumi.Input<string>;
    /**
     * Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported
     */
    tagLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify the transit Gateway.
     */
    transitGw?: pulumi.Input<string>;
    /**
     * VPC-ID/VNet-Name of cloud provider. Required if cloudType is "1" or "4". Example: AWS: "vpc-abcd1234", etc...
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...
     */
    vpcReg?: pulumi.Input<string>;
    /**
     * Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...
     */
    vpcSize?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixSpokeVpc resource.
 */
export interface AviatrixSpokeVpcArgs {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName: pulumi.Input<string>;
    /**
     * Type of cloud service provider. AWS=1, GCP=4, ARM=8.
     */
    cloudType: pulumi.Input<number>;
    /**
     * Specify whether enabling NAT feature on the gateway or not.
     */
    enableNat?: pulumi.Input<string>;
    /**
     * Name of the gateway which is going to be created.
     */
    gwName: pulumi.Input<string>;
    /**
     * HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".
     */
    haGwSize?: pulumi.Input<string>;
    /**
     * HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.
     */
    haSubnet?: pulumi.Input<string>;
    /**
     * HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".
     */
    haZone?: pulumi.Input<string>;
    /**
     * Set to "enabled" if this feature is desired.
     */
    singleAzHa?: pulumi.Input<string>;
    /**
     * Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...
     */
    subnet: pulumi.Input<string>;
    /**
     * Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported
     */
    tagLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify the transit Gateway.
     */
    transitGw?: pulumi.Input<string>;
    /**
     * VPC-ID/VNet-Name of cloud provider. Required if cloudType is "1" or "4". Example: AWS: "vpc-abcd1234", etc...
     */
    vpcId: pulumi.Input<string>;
    /**
     * Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...
     */
    vpcReg: pulumi.Input<string>;
    /**
     * Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...
     */
    vpcSize: pulumi.Input<string>;
}

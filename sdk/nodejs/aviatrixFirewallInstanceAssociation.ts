// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_firewall_instance_association** resource allows for the creation and management of a firewall instance association. To use this resource you must also have an `aviatrix.AviatrixFirenet` resource with it's `manageFirewallInstanceAssociation` attribute set to false.
 *
 * Available in provider version R2.17.1+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@astipkovits/aviatrix";
 *
 * // Associate an Aviatrix FireNet Gateway with a Firewall Instance
 * const firewallInstanceAssociation1 = new aviatrix.AviatrixFirewallInstanceAssociation("firewallInstanceAssociation1", {
 *     vpcId: aviatrix_firewall_instance.firewall_instance_1.vpc_id,
 *     firenetGwName: aviatrix_transit_gateway.transit_gateway_1.gw_name,
 *     instanceId: aviatrix_firewall_instance.firewall_instance_1.instance_id,
 *     firewallName: aviatrix_firewall_instance.firewall_instance_1.firewall_name,
 *     lanInterface: aviatrix_firewall_instance.firewall_instance_1.lan_interface,
 *     managementInterface: aviatrix_firewall_instance.firewall_instance_1.management_interface,
 *     egressInterface: aviatrix_firewall_instance.firewall_instance_1.egress_interface,
 *     attached: true,
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@astipkovits/aviatrix";
 *
 * // Associate an GCP Aviatrix FireNet Gateway with a Firewall Instance
 * const firewallInstanceAssociation1 = new aviatrix.AviatrixFirewallInstanceAssociation("firewallInstanceAssociation1", {
 *     vpcId: aviatrix_firewall_instance.firewall_instance_1.vpc_id,
 *     firenetGwName: aviatrix_transit_gateway.transit_gateway_1.gw_name,
 *     instanceId: aviatrix_firewall_instance.firewall_instance_1.instance_id,
 *     lanInterface: aviatrix_firewall_instance.firewall_instance_1.lan_interface,
 *     managementInterface: aviatrix_firewall_instance.firewall_instance_1.management_interface,
 *     egressInterface: aviatrix_firewall_instance.firewall_instance_1.egress_interface,
 *     attached: true,
 * });
 * ```
 *
 * ## Import
 *
 * **firewall_instance_association** can be imported using the `vpc_id`, `firenet_gw_name` and `instance_id` in the form `vpc_id~~firenet_gw_name~~instance_id` e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation test vpc_id~~firenet_gw_name~~instance_id
 * ```
 *
 *  When using a Native GWLB VPC where there is no `firenet_gw_name` but the ID is in the same form e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation test vpc_id~~~~instance_id
 * ```
 */
export class AviatrixFirewallInstanceAssociation extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixFirewallInstanceAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixFirewallInstanceAssociationState, opts?: pulumi.CustomResourceOptions): AviatrixFirewallInstanceAssociation {
        return new AviatrixFirewallInstanceAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixFirewallInstanceAssociation:AviatrixFirewallInstanceAssociation';

    /**
     * Returns true if the given object is an instance of AviatrixFirewallInstanceAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixFirewallInstanceAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixFirewallInstanceAssociation.__pulumiType;
    }

    /**
     * Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
     */
    public readonly attached!: pulumi.Output<boolean | undefined>;
    /**
     * Egress interface ID. **Required if it is a firewall instance.**
     */
    public readonly egressInterface!: pulumi.Output<string | undefined>;
    /**
     * Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
     */
    public readonly firenetGwName!: pulumi.Output<string | undefined>;
    /**
     * Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
     */
    public readonly firewallName!: pulumi.Output<string | undefined>;
    /**
     * ID of Firewall instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Lan interface ID. **Required if it is a firewall instance.**
     */
    public readonly lanInterface!: pulumi.Output<string | undefined>;
    /**
     * Management interface ID. **Required if it is a firewall instance.**
     */
    public readonly managementInterface!: pulumi.Output<string | undefined>;
    /**
     * Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
     */
    public readonly vendorType!: pulumi.Output<string | undefined>;
    /**
     * VPC ID of the Security VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixFirewallInstanceAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixFirewallInstanceAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixFirewallInstanceAssociationArgs | AviatrixFirewallInstanceAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixFirewallInstanceAssociationState | undefined;
            resourceInputs["attached"] = state ? state.attached : undefined;
            resourceInputs["egressInterface"] = state ? state.egressInterface : undefined;
            resourceInputs["firenetGwName"] = state ? state.firenetGwName : undefined;
            resourceInputs["firewallName"] = state ? state.firewallName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["lanInterface"] = state ? state.lanInterface : undefined;
            resourceInputs["managementInterface"] = state ? state.managementInterface : undefined;
            resourceInputs["vendorType"] = state ? state.vendorType : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixFirewallInstanceAssociationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["attached"] = args ? args.attached : undefined;
            resourceInputs["egressInterface"] = args ? args.egressInterface : undefined;
            resourceInputs["firenetGwName"] = args ? args.firenetGwName : undefined;
            resourceInputs["firewallName"] = args ? args.firewallName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["lanInterface"] = args ? args.lanInterface : undefined;
            resourceInputs["managementInterface"] = args ? args.managementInterface : undefined;
            resourceInputs["vendorType"] = args ? args.vendorType : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixFirewallInstanceAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixFirewallInstanceAssociation resources.
 */
export interface AviatrixFirewallInstanceAssociationState {
    /**
     * Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * Egress interface ID. **Required if it is a firewall instance.**
     */
    egressInterface?: pulumi.Input<string>;
    /**
     * Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
     */
    firenetGwName?: pulumi.Input<string>;
    /**
     * Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
     */
    firewallName?: pulumi.Input<string>;
    /**
     * ID of Firewall instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Lan interface ID. **Required if it is a firewall instance.**
     */
    lanInterface?: pulumi.Input<string>;
    /**
     * Management interface ID. **Required if it is a firewall instance.**
     */
    managementInterface?: pulumi.Input<string>;
    /**
     * Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
     */
    vendorType?: pulumi.Input<string>;
    /**
     * VPC ID of the Security VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixFirewallInstanceAssociation resource.
 */
export interface AviatrixFirewallInstanceAssociationArgs {
    /**
     * Switch to attach/detach firewall instance to/from FireNet. Valid values: true, false. Default value: false.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * Egress interface ID. **Required if it is a firewall instance.**
     */
    egressInterface?: pulumi.Input<string>;
    /**
     * Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.
     */
    firenetGwName?: pulumi.Input<string>;
    /**
     * Firewall instance name. **Required for non-GCP firewall instance. For GCP, this field should not be set.**
     */
    firewallName?: pulumi.Input<string>;
    /**
     * ID of Firewall instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Lan interface ID. **Required if it is a firewall instance.**
     */
    lanInterface?: pulumi.Input<string>;
    /**
     * Management interface ID. **Required if it is a firewall instance.**
     */
    managementInterface?: pulumi.Input<string>;
    /**
     * Type of firewall. Valid values: "Generic", "fqdnGateway". Default value: "Generic". Value "fqdnGateway" is required for FQDN gateway.
     */
    vendorType?: pulumi.Input<string>;
    /**
     * VPC ID of the Security VPC.
     */
    vpcId: pulumi.Input<string>;
}

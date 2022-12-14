// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The **aviatrix_firewall_tag** resource allows the creation and management of [Aviatrix Stateful Firewall tags](https://docs.aviatrix.com/HowTos/tag_firewall.html) for tag-based security for gateways.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Firewall Tag
 * const testFirewallTag = new aviatrix.AviatrixFirewallTag("test_firewall_tag", {
 *     cidrLists: [
 *         {
 *             cidr: "10.1.0.0/24",
 *             cidrTagName: "a1",
 *         },
 *         {
 *             cidr: "10.2.0.0/24",
 *             cidrTagName: "b1",
 *         },
 *     ],
 *     firewallTag: "test-firewall-tag",
 * });
 * ```
 *
 * ## Import
 *
 * **firewall_tag** can be imported using the `firewall_tag`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixFirewallTag:AviatrixFirewallTag test firewall_tag
 * ```
 */
export class AviatrixFirewallTag extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixFirewallTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixFirewallTagState, opts?: pulumi.CustomResourceOptions): AviatrixFirewallTag {
        return new AviatrixFirewallTag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixFirewallTag:AviatrixFirewallTag';

    /**
     * Returns true if the given object is an instance of AviatrixFirewallTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixFirewallTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixFirewallTag.__pulumiType;
    }

    /**
     * Dynamic block representing a CIDR to filter, and a name to identify it:
     */
    public readonly cidrLists!: pulumi.Output<outputs.AviatrixFirewallTagCidrList[] | undefined>;
    /**
     * Name of the stateful firewall tag to be created.
     */
    public readonly firewallTag!: pulumi.Output<string>;

    /**
     * Create a AviatrixFirewallTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixFirewallTagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixFirewallTagArgs | AviatrixFirewallTagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixFirewallTagState | undefined;
            resourceInputs["cidrLists"] = state ? state.cidrLists : undefined;
            resourceInputs["firewallTag"] = state ? state.firewallTag : undefined;
        } else {
            const args = argsOrState as AviatrixFirewallTagArgs | undefined;
            if ((!args || args.firewallTag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewallTag'");
            }
            resourceInputs["cidrLists"] = args ? args.cidrLists : undefined;
            resourceInputs["firewallTag"] = args ? args.firewallTag : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixFirewallTag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixFirewallTag resources.
 */
export interface AviatrixFirewallTagState {
    /**
     * Dynamic block representing a CIDR to filter, and a name to identify it:
     */
    cidrLists?: pulumi.Input<pulumi.Input<inputs.AviatrixFirewallTagCidrList>[]>;
    /**
     * Name of the stateful firewall tag to be created.
     */
    firewallTag?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixFirewallTag resource.
 */
export interface AviatrixFirewallTagArgs {
    /**
     * Dynamic block representing a CIDR to filter, and a name to identify it:
     */
    cidrLists?: pulumi.Input<pulumi.Input<inputs.AviatrixFirewallTagCidrList>[]>;
    /**
     * Name of the stateful firewall tag to be created.
     */
    firewallTag: pulumi.Input<string>;
}

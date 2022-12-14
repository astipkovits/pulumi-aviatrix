// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_rbac_group** resource allows the creation and management of [Aviatrix (Role-Based Access Control) RBAC groups](https://docs.aviatrix.com/HowTos/rbac_faq.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix RBAC Group
 * const testGroup = new aviatrix.AviatrixRbacGroup("test_group", {
 *     groupName: "write_only",
 * });
 * ```
 *
 * ## Import
 *
 * **rbac_group** can be imported using the `group_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixRbacGroup:AviatrixRbacGroup test group_name
 * ```
 */
export class AviatrixRbacGroup extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixRbacGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixRbacGroupState, opts?: pulumi.CustomResourceOptions): AviatrixRbacGroup {
        return new AviatrixRbacGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixRbacGroup:AviatrixRbacGroup';

    /**
     * Returns true if the given object is an instance of AviatrixRbacGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixRbacGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixRbacGroup.__pulumiType;
    }

    /**
     * This parameter represents the name of a RBAC group to be created.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * Whether to allow members of an RBAC group to bypass LDAP/MFA for Duo login . Supported values: true, false. Default value: false. Available in provider version R2.17.1+.
     */
    public readonly localLogin!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AviatrixRbacGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixRbacGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixRbacGroupArgs | AviatrixRbacGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixRbacGroupState | undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["localLogin"] = state ? state.localLogin : undefined;
        } else {
            const args = argsOrState as AviatrixRbacGroupArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["localLogin"] = args ? args.localLogin : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixRbacGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixRbacGroup resources.
 */
export interface AviatrixRbacGroupState {
    /**
     * This parameter represents the name of a RBAC group to be created.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Whether to allow members of an RBAC group to bypass LDAP/MFA for Duo login . Supported values: true, false. Default value: false. Available in provider version R2.17.1+.
     */
    localLogin?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AviatrixRbacGroup resource.
 */
export interface AviatrixRbacGroupArgs {
    /**
     * This parameter represents the name of a RBAC group to be created.
     */
    groupName: pulumi.Input<string>;
    /**
     * Whether to allow members of an RBAC group to bypass LDAP/MFA for Duo login . Supported values: true, false. Default value: false. Available in provider version R2.17.1+.
     */
    localLogin?: pulumi.Input<boolean>;
}

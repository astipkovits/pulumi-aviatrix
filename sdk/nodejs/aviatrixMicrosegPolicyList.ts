// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class AviatrixMicrosegPolicyList extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixMicrosegPolicyList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixMicrosegPolicyListState, opts?: pulumi.CustomResourceOptions): AviatrixMicrosegPolicyList {
        return new AviatrixMicrosegPolicyList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixMicrosegPolicyList:AviatrixMicrosegPolicyList';

    /**
     * Returns true if the given object is an instance of AviatrixMicrosegPolicyList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixMicrosegPolicyList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixMicrosegPolicyList.__pulumiType;
    }

    /**
     * List of micro-segmentation policies.
     */
    public readonly policies!: pulumi.Output<outputs.AviatrixMicrosegPolicyListPolicy[]>;

    /**
     * Create a AviatrixMicrosegPolicyList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixMicrosegPolicyListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixMicrosegPolicyListArgs | AviatrixMicrosegPolicyListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixMicrosegPolicyListState | undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as AviatrixMicrosegPolicyListArgs | undefined;
            if ((!args || args.policies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policies'");
            }
            resourceInputs["policies"] = args ? args.policies : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixMicrosegPolicyList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixMicrosegPolicyList resources.
 */
export interface AviatrixMicrosegPolicyListState {
    /**
     * List of micro-segmentation policies.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.AviatrixMicrosegPolicyListPolicy>[]>;
}

/**
 * The set of arguments for constructing a AviatrixMicrosegPolicyList resource.
 */
export interface AviatrixMicrosegPolicyListArgs {
    /**
     * List of micro-segmentation policies.
     */
    policies: pulumi.Input<pulumi.Input<inputs.AviatrixMicrosegPolicyListPolicy>[]>;
}
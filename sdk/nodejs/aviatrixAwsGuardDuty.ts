// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixAwsGuardDuty extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsGuardDuty resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsGuardDutyState, opts?: pulumi.CustomResourceOptions): AviatrixAwsGuardDuty {
        return new AviatrixAwsGuardDuty(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsGuardDuty:AviatrixAwsGuardDuty';

    /**
     * Returns true if the given object is an instance of AviatrixAwsGuardDuty.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsGuardDuty {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsGuardDuty.__pulumiType;
    }

    /**
     * Account name
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Excluded IPs.
     */
    public readonly excludedIps!: pulumi.Output<string[] | undefined>;
    /**
     * Region.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsGuardDuty resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsGuardDutyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsGuardDutyArgs | AviatrixAwsGuardDutyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsGuardDutyState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["excludedIps"] = state ? state.excludedIps : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as AviatrixAwsGuardDutyArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["excludedIps"] = args ? args.excludedIps : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsGuardDuty.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsGuardDuty resources.
 */
export interface AviatrixAwsGuardDutyState {
    /**
     * Account name
     */
    accountName?: pulumi.Input<string>;
    /**
     * Excluded IPs.
     */
    excludedIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Region.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsGuardDuty resource.
 */
export interface AviatrixAwsGuardDutyArgs {
    /**
     * Account name
     */
    accountName: pulumi.Input<string>;
    /**
     * Excluded IPs.
     */
    excludedIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Region.
     */
    region: pulumi.Input<string>;
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixNetflowAgent extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixNetflowAgent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixNetflowAgentState, opts?: pulumi.CustomResourceOptions): AviatrixNetflowAgent {
        return new AviatrixNetflowAgent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixNetflowAgent:AviatrixNetflowAgent';

    /**
     * Returns true if the given object is an instance of AviatrixNetflowAgent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixNetflowAgent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixNetflowAgent.__pulumiType;
    }

    /**
     * List of excluded gateways.
     */
    public readonly excludedGateways!: pulumi.Output<string[] | undefined>;
    /**
     * Netflow server port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Netflow server IP address.
     */
    public readonly serverIp!: pulumi.Output<string>;
    /**
     * Enabled or not.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Netflow version.
     */
    public readonly version!: pulumi.Output<number | undefined>;

    /**
     * Create a AviatrixNetflowAgent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixNetflowAgentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixNetflowAgentArgs | AviatrixNetflowAgentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixNetflowAgentState | undefined;
            resourceInputs["excludedGateways"] = state ? state.excludedGateways : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["serverIp"] = state ? state.serverIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AviatrixNetflowAgentArgs | undefined;
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.serverIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverIp'");
            }
            resourceInputs["excludedGateways"] = args ? args.excludedGateways : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["serverIp"] = args ? args.serverIp : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixNetflowAgent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixNetflowAgent resources.
 */
export interface AviatrixNetflowAgentState {
    /**
     * List of excluded gateways.
     */
    excludedGateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Netflow server port.
     */
    port?: pulumi.Input<number>;
    /**
     * Netflow server IP address.
     */
    serverIp?: pulumi.Input<string>;
    /**
     * Enabled or not.
     */
    status?: pulumi.Input<string>;
    /**
     * Netflow version.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AviatrixNetflowAgent resource.
 */
export interface AviatrixNetflowAgentArgs {
    /**
     * List of excluded gateways.
     */
    excludedGateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Netflow server port.
     */
    port: pulumi.Input<number>;
    /**
     * Netflow server IP address.
     */
    serverIp: pulumi.Input<string>;
    /**
     * Netflow version.
     */
    version?: pulumi.Input<number>;
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_azure_vng_conn** resource allows the creation and management of the connection between Aviatrix Transit Gateway and Azure VNG.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Attach an Azure VNG to an Aviatrix Transit Gateway
 * const test = new aviatrix.AviatrixAzureVngConn("test", {
 *     connectionName: "connection",
 *     primaryGatewayName: "primary-gateway",
 * });
 * ```
 *
 * ## Import
 *
 * **aviatrix_azure_vng_conn** can be imported using the `connection_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixAzureVngConn:AviatrixAzureVngConn test connection
 * ```
 */
export class AviatrixAzureVngConn extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAzureVngConn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAzureVngConnState, opts?: pulumi.CustomResourceOptions): AviatrixAzureVngConn {
        return new AviatrixAzureVngConn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAzureVngConn:AviatrixAzureVngConn';

    /**
     * Returns true if the given object is an instance of AviatrixAzureVngConn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAzureVngConn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAzureVngConn.__pulumiType;
    }

    /**
     * The status of the connection.
     */
    public /*out*/ readonly attached!: pulumi.Output<boolean>;
    /**
     * Connection name.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * Primary Aviatrix transit gateway name.
     */
    public readonly primaryGatewayName!: pulumi.Output<string>;
    /**
     * Name of Azure VNG.
     */
    public /*out*/ readonly vngName!: pulumi.Output<string>;
    /**
     * VPC ID.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixAzureVngConn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAzureVngConnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAzureVngConnArgs | AviatrixAzureVngConnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAzureVngConnState | undefined;
            resourceInputs["attached"] = state ? state.attached : undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["primaryGatewayName"] = state ? state.primaryGatewayName : undefined;
            resourceInputs["vngName"] = state ? state.vngName : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixAzureVngConnArgs | undefined;
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.primaryGatewayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primaryGatewayName'");
            }
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["primaryGatewayName"] = args ? args.primaryGatewayName : undefined;
            resourceInputs["attached"] = undefined /*out*/;
            resourceInputs["vngName"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAzureVngConn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAzureVngConn resources.
 */
export interface AviatrixAzureVngConnState {
    /**
     * The status of the connection.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * Connection name.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * Primary Aviatrix transit gateway name.
     */
    primaryGatewayName?: pulumi.Input<string>;
    /**
     * Name of Azure VNG.
     */
    vngName?: pulumi.Input<string>;
    /**
     * VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAzureVngConn resource.
 */
export interface AviatrixAzureVngConnArgs {
    /**
     * Connection name.
     */
    connectionName: pulumi.Input<string>;
    /**
     * Primary Aviatrix transit gateway name.
     */
    primaryGatewayName: pulumi.Input<string>;
}

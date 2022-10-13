// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixEdgeSpokeExternalDeviceConn extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixEdgeSpokeExternalDeviceConn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixEdgeSpokeExternalDeviceConnState, opts?: pulumi.CustomResourceOptions): AviatrixEdgeSpokeExternalDeviceConn {
        return new AviatrixEdgeSpokeExternalDeviceConn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixEdgeSpokeExternalDeviceConn:AviatrixEdgeSpokeExternalDeviceConn';

    /**
     * Returns true if the given object is an instance of AviatrixEdgeSpokeExternalDeviceConn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixEdgeSpokeExternalDeviceConn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixEdgeSpokeExternalDeviceConn.__pulumiType;
    }

    /**
     * BGP local AS number.
     */
    public readonly bgpLocalAsNum!: pulumi.Output<string>;
    /**
     * BGP remote AS number.
     */
    public readonly bgpRemoteAsNum!: pulumi.Output<string>;
    /**
     * The name of the spoke external device connection which is going to be created.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * Connection type. Valid values: 'bgp'. Default value: 'bgp'.
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * Name of the BGP Spoke Gateway.
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * Local LAN IP.
     */
    public readonly localLanIp!: pulumi.Output<string>;
    /**
     * Remote LAN IP.
     */
    public readonly remoteLanIp!: pulumi.Output<string>;
    /**
     * ID of the VPC where the BGP Spoke Gateway is located.
     */
    public readonly siteId!: pulumi.Output<string>;
    /**
     * Tunnel Protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
     */
    public readonly tunnelProtocol!: pulumi.Output<string | undefined>;

    /**
     * Create a AviatrixEdgeSpokeExternalDeviceConn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixEdgeSpokeExternalDeviceConnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixEdgeSpokeExternalDeviceConnArgs | AviatrixEdgeSpokeExternalDeviceConnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixEdgeSpokeExternalDeviceConnState | undefined;
            resourceInputs["bgpLocalAsNum"] = state ? state.bgpLocalAsNum : undefined;
            resourceInputs["bgpRemoteAsNum"] = state ? state.bgpRemoteAsNum : undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["connectionType"] = state ? state.connectionType : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["localLanIp"] = state ? state.localLanIp : undefined;
            resourceInputs["remoteLanIp"] = state ? state.remoteLanIp : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["tunnelProtocol"] = state ? state.tunnelProtocol : undefined;
        } else {
            const args = argsOrState as AviatrixEdgeSpokeExternalDeviceConnArgs | undefined;
            if ((!args || args.bgpLocalAsNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpLocalAsNum'");
            }
            if ((!args || args.bgpRemoteAsNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpRemoteAsNum'");
            }
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            if ((!args || args.localLanIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localLanIp'");
            }
            if ((!args || args.remoteLanIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteLanIp'");
            }
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            resourceInputs["bgpLocalAsNum"] = args ? args.bgpLocalAsNum : undefined;
            resourceInputs["bgpRemoteAsNum"] = args ? args.bgpRemoteAsNum : undefined;
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["connectionType"] = args ? args.connectionType : undefined;
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["localLanIp"] = args ? args.localLanIp : undefined;
            resourceInputs["remoteLanIp"] = args ? args.remoteLanIp : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["tunnelProtocol"] = args ? args.tunnelProtocol : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixEdgeSpokeExternalDeviceConn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixEdgeSpokeExternalDeviceConn resources.
 */
export interface AviatrixEdgeSpokeExternalDeviceConnState {
    /**
     * BGP local AS number.
     */
    bgpLocalAsNum?: pulumi.Input<string>;
    /**
     * BGP remote AS number.
     */
    bgpRemoteAsNum?: pulumi.Input<string>;
    /**
     * The name of the spoke external device connection which is going to be created.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * Connection type. Valid values: 'bgp'. Default value: 'bgp'.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * Name of the BGP Spoke Gateway.
     */
    gwName?: pulumi.Input<string>;
    /**
     * Local LAN IP.
     */
    localLanIp?: pulumi.Input<string>;
    /**
     * Remote LAN IP.
     */
    remoteLanIp?: pulumi.Input<string>;
    /**
     * ID of the VPC where the BGP Spoke Gateway is located.
     */
    siteId?: pulumi.Input<string>;
    /**
     * Tunnel Protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
     */
    tunnelProtocol?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixEdgeSpokeExternalDeviceConn resource.
 */
export interface AviatrixEdgeSpokeExternalDeviceConnArgs {
    /**
     * BGP local AS number.
     */
    bgpLocalAsNum: pulumi.Input<string>;
    /**
     * BGP remote AS number.
     */
    bgpRemoteAsNum: pulumi.Input<string>;
    /**
     * The name of the spoke external device connection which is going to be created.
     */
    connectionName: pulumi.Input<string>;
    /**
     * Connection type. Valid values: 'bgp'. Default value: 'bgp'.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * Name of the BGP Spoke Gateway.
     */
    gwName: pulumi.Input<string>;
    /**
     * Local LAN IP.
     */
    localLanIp: pulumi.Input<string>;
    /**
     * Remote LAN IP.
     */
    remoteLanIp: pulumi.Input<string>;
    /**
     * ID of the VPC where the BGP Spoke Gateway is located.
     */
    siteId: pulumi.Input<string>;
    /**
     * Tunnel Protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.
     */
    tunnelProtocol?: pulumi.Input<string>;
}
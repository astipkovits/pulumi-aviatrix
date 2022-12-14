// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get the list of device WAN interfaces for use in other resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Aviatrix Device Interfaces Data Source
 * const test = pulumi.output(aviatrix.getAviatrixDeviceInterfaces({
 *     deviceName: "test-device",
 * }));
 * ```
 */
export function getAviatrixDeviceInterfaces(args: GetAviatrixDeviceInterfacesArgs, opts?: pulumi.InvokeOptions): Promise<GetAviatrixDeviceInterfacesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aviatrix:index/getAviatrixDeviceInterfaces:getAviatrixDeviceInterfaces", {
        "deviceName": args.deviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAviatrixDeviceInterfaces.
 */
export interface GetAviatrixDeviceInterfacesArgs {
    /**
     * Device name.
     */
    deviceName: string;
}

/**
 * A collection of values returned by getAviatrixDeviceInterfaces.
 */
export interface GetAviatrixDeviceInterfacesResult {
    readonly deviceName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of WAN interfaces.
     */
    readonly wanInterfaces: outputs.GetAviatrixDeviceInterfacesWanInterface[];
}

export function getAviatrixDeviceInterfacesOutput(args: GetAviatrixDeviceInterfacesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAviatrixDeviceInterfacesResult> {
    return pulumi.output(args).apply(a => getAviatrixDeviceInterfaces(a, opts))
}

/**
 * A collection of arguments for invoking getAviatrixDeviceInterfaces.
 */
export interface GetAviatrixDeviceInterfacesOutputArgs {
    /**
     * Device name.
     */
    deviceName: pulumi.Input<string>;
}

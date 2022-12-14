// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_geo_vpn** resource enables and manages the [Aviatrix Geo VPN feature](https://docs.aviatrix.com/HowTos/GeoVPN.html).
 *
 * > **NOTE:** If ELBs/gateways are being managed by the Geo VPN, in order to update VPN configurations of the Geo VPN, all the VPN configurations of the ELBs/gateways must be updated simultaneously and share the same values. This can be achieved by managing the VPN configurations through variables and updating their values accordingly.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Geo VPN
 * const testGeoVpn = new aviatrix.AviatrixGeoVpn("test_geo_vpn", {
 *     accountName: "devops-aws",
 *     cloudType: 1,
 *     domainName: "aviatrix.live",
 *     elbDnsNames: [
 *         "elb-test1-497f5e89.elb.us-west-1.amazonaws.com",
 *         "elb-test2-974f895e.elb.us-east-2.amazonaws.com",
 *     ],
 *     serviceName: "vpn",
 * });
 * ```
 *
 * ## Import
 *
 * **geo_vpn** can be imported using the `service_name` and `domain_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn test service_name~domain_name
 * ```
 */
export class AviatrixGeoVpn extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixGeoVpn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixGeoVpnState, opts?: pulumi.CustomResourceOptions): AviatrixGeoVpn {
        return new AviatrixGeoVpn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn';

    /**
     * Returns true if the given object is an instance of AviatrixGeoVpn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixGeoVpn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixGeoVpn.__pulumiType;
    }

    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
     */
    public readonly cloudType!: pulumi.Output<number>;
    /**
     * The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * List of ELB names to attach to this Geo VPN name.
     */
    public readonly elbDnsNames!: pulumi.Output<string[]>;
    /**
     * The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a AviatrixGeoVpn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixGeoVpnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixGeoVpnArgs | AviatrixGeoVpnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixGeoVpnState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["cloudType"] = state ? state.cloudType : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["elbDnsNames"] = state ? state.elbDnsNames : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as AviatrixGeoVpnArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.cloudType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudType'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.elbDnsNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'elbDnsNames'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["cloudType"] = args ? args.cloudType : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["elbDnsNames"] = args ? args.elbDnsNames : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixGeoVpn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixGeoVpn resources.
 */
export interface AviatrixGeoVpnState {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
     */
    cloudType?: pulumi.Input<number>;
    /**
     * The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
     */
    domainName?: pulumi.Input<string>;
    /**
     * List of ELB names to attach to this Geo VPN name.
     */
    elbDnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixGeoVpn resource.
 */
export interface AviatrixGeoVpnArgs {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName: pulumi.Input<string>;
    /**
     * Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
     */
    cloudType: pulumi.Input<number>;
    /**
     * The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
     */
    domainName: pulumi.Input<string>;
    /**
     * List of ELB names to attach to this Geo VPN name.
     */
    elbDnsNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
     */
    serviceName: pulumi.Input<string>;
}

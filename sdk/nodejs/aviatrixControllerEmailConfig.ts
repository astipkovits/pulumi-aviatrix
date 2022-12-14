// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **aviatrix_controller_email_config** resource allows management of an Aviatrix Controller's notification email configurations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Controller Email Config
 * const testEmailConfig = new aviatrix.AviatrixControllerEmailConfig("test_email_config", {
 *     adminAlertEmail: "administrator@mycompany.com",
 *     criticalAlertEmail: "it-support@mycompany.com",
 *     securityEventEmail: "security-admin-group@mycompany.com",
 *     statusChangeEmail: "it-admin-group@mycompany.com",
 * });
 * ```
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aviatrix from "@pulumi/aviatrix";
 *
 * // Create an Aviatrix Controller Email Config and configure the Status Change Notification Interval
 * const testEmailConfig = new aviatrix.AviatrixControllerEmailConfig("test_email_config", {
 *     adminAlertEmail: "administrator@mycompany.com",
 *     criticalAlertEmail: "it-support@mycompany.com",
 *     securityEventEmail: "security-admin-group@mycompany.com",
 *     statusChangeEmail: "it-admin-group@mycompany.com",
 *     statusChangeNotificationInterval: 20,
 * });
 * ```
 *
 * ## Import
 *
 * Instance controller_email_config can be imported using controller IP, e.g. controller IP is 10.11.12.13
 *
 * ```sh
 *  $ pulumi import aviatrix:index/aviatrixControllerEmailConfig:AviatrixControllerEmailConfig test 10-11-12-13
 * ```
 */
export class AviatrixControllerEmailConfig extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixControllerEmailConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixControllerEmailConfigState, opts?: pulumi.CustomResourceOptions): AviatrixControllerEmailConfig {
        return new AviatrixControllerEmailConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixControllerEmailConfig:AviatrixControllerEmailConfig';

    /**
     * Returns true if the given object is an instance of AviatrixControllerEmailConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixControllerEmailConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixControllerEmailConfig.__pulumiType;
    }

    /**
     * Email to receive important account and certification information.
     */
    public readonly adminAlertEmail!: pulumi.Output<string>;
    /**
     * Whether admin alert notification email is verified.
     */
    public /*out*/ readonly adminAlertEmailVerified!: pulumi.Output<boolean>;
    /**
     * Email to receive field notices and critical notices.
     */
    public readonly criticalAlertEmail!: pulumi.Output<string>;
    /**
     * Whether critical alert notification email is verified.
     */
    public /*out*/ readonly criticalAlertEmailVerified!: pulumi.Output<boolean>;
    /**
     * Email to receive security and CVE (Common Vulnerabilities and Exposures) notification emails.
     */
    public readonly securityEventEmail!: pulumi.Output<string>;
    /**
     * Whether security event notification email is verified.
     */
    public /*out*/ readonly securityEventEmailVerified!: pulumi.Output<boolean>;
    /**
     * Email to receive system/tunnel status notification emails.
     */
    public readonly statusChangeEmail!: pulumi.Output<string>;
    /**
     * Whether status change notification email is verified.
     */
    public /*out*/ readonly statusChangeEmailVerified!: pulumi.Output<boolean>;
    /**
     * Status change notification interval in seconds. Default value: 60.
     */
    public readonly statusChangeNotificationInterval!: pulumi.Output<number | undefined>;

    /**
     * Create a AviatrixControllerEmailConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixControllerEmailConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixControllerEmailConfigArgs | AviatrixControllerEmailConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixControllerEmailConfigState | undefined;
            resourceInputs["adminAlertEmail"] = state ? state.adminAlertEmail : undefined;
            resourceInputs["adminAlertEmailVerified"] = state ? state.adminAlertEmailVerified : undefined;
            resourceInputs["criticalAlertEmail"] = state ? state.criticalAlertEmail : undefined;
            resourceInputs["criticalAlertEmailVerified"] = state ? state.criticalAlertEmailVerified : undefined;
            resourceInputs["securityEventEmail"] = state ? state.securityEventEmail : undefined;
            resourceInputs["securityEventEmailVerified"] = state ? state.securityEventEmailVerified : undefined;
            resourceInputs["statusChangeEmail"] = state ? state.statusChangeEmail : undefined;
            resourceInputs["statusChangeEmailVerified"] = state ? state.statusChangeEmailVerified : undefined;
            resourceInputs["statusChangeNotificationInterval"] = state ? state.statusChangeNotificationInterval : undefined;
        } else {
            const args = argsOrState as AviatrixControllerEmailConfigArgs | undefined;
            if ((!args || args.adminAlertEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminAlertEmail'");
            }
            if ((!args || args.criticalAlertEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'criticalAlertEmail'");
            }
            if ((!args || args.securityEventEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityEventEmail'");
            }
            if ((!args || args.statusChangeEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statusChangeEmail'");
            }
            resourceInputs["adminAlertEmail"] = args ? args.adminAlertEmail : undefined;
            resourceInputs["criticalAlertEmail"] = args ? args.criticalAlertEmail : undefined;
            resourceInputs["securityEventEmail"] = args ? args.securityEventEmail : undefined;
            resourceInputs["statusChangeEmail"] = args ? args.statusChangeEmail : undefined;
            resourceInputs["statusChangeNotificationInterval"] = args ? args.statusChangeNotificationInterval : undefined;
            resourceInputs["adminAlertEmailVerified"] = undefined /*out*/;
            resourceInputs["criticalAlertEmailVerified"] = undefined /*out*/;
            resourceInputs["securityEventEmailVerified"] = undefined /*out*/;
            resourceInputs["statusChangeEmailVerified"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixControllerEmailConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixControllerEmailConfig resources.
 */
export interface AviatrixControllerEmailConfigState {
    /**
     * Email to receive important account and certification information.
     */
    adminAlertEmail?: pulumi.Input<string>;
    /**
     * Whether admin alert notification email is verified.
     */
    adminAlertEmailVerified?: pulumi.Input<boolean>;
    /**
     * Email to receive field notices and critical notices.
     */
    criticalAlertEmail?: pulumi.Input<string>;
    /**
     * Whether critical alert notification email is verified.
     */
    criticalAlertEmailVerified?: pulumi.Input<boolean>;
    /**
     * Email to receive security and CVE (Common Vulnerabilities and Exposures) notification emails.
     */
    securityEventEmail?: pulumi.Input<string>;
    /**
     * Whether security event notification email is verified.
     */
    securityEventEmailVerified?: pulumi.Input<boolean>;
    /**
     * Email to receive system/tunnel status notification emails.
     */
    statusChangeEmail?: pulumi.Input<string>;
    /**
     * Whether status change notification email is verified.
     */
    statusChangeEmailVerified?: pulumi.Input<boolean>;
    /**
     * Status change notification interval in seconds. Default value: 60.
     */
    statusChangeNotificationInterval?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AviatrixControllerEmailConfig resource.
 */
export interface AviatrixControllerEmailConfigArgs {
    /**
     * Email to receive important account and certification information.
     */
    adminAlertEmail: pulumi.Input<string>;
    /**
     * Email to receive field notices and critical notices.
     */
    criticalAlertEmail: pulumi.Input<string>;
    /**
     * Email to receive security and CVE (Common Vulnerabilities and Exposures) notification emails.
     */
    securityEventEmail: pulumi.Input<string>;
    /**
     * Email to receive system/tunnel status notification emails.
     */
    statusChangeEmail: pulumi.Input<string>;
    /**
     * Status change notification interval in seconds. Default value: 60.
     */
    statusChangeNotificationInterval?: pulumi.Input<number>;
}

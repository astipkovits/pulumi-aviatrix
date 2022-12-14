// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class GetAviatrixFirenetFirewallInstanceAssociationInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("attached", required: true)]
        public Input<bool> Attached { get; set; } = null!;

        [Input("egressInterface", required: true)]
        public Input<string> EgressInterface { get; set; } = null!;

        /// <summary>
        /// Name of the primary FireNet gateway.
        /// </summary>
        [Input("firenetGwName", required: true)]
        public Input<string> FirenetGwName { get; set; } = null!;

        /// <summary>
        /// Firewall instance name.
        /// * `lan_interface`- Lan interface ID.
        /// </summary>
        [Input("firewallName", required: true)]
        public Input<string> FirewallName { get; set; } = null!;

        /// <summary>
        /// ID of Firewall instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("lanInterface", required: true)]
        public Input<string> LanInterface { get; set; } = null!;

        /// <summary>
        /// Management interface ID.
        /// * `egress_interface`- Egress interface ID.
        /// * `attached`- Switch to attach/detach firewall instance to/from fireNet.
        /// </summary>
        [Input("managementInterface", required: true)]
        public Input<string> ManagementInterface { get; set; } = null!;

        /// <summary>
        /// Type of the firewall.
        /// </summary>
        [Input("vendorType", required: true)]
        public Input<string> VendorType { get; set; } = null!;

        public GetAviatrixFirenetFirewallInstanceAssociationInputArgs()
        {
        }
        public static new GetAviatrixFirenetFirewallInstanceAssociationInputArgs Empty => new GetAviatrixFirenetFirewallInstanceAssociationInputArgs();
    }
}

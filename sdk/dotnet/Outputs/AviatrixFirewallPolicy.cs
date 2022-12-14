// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Outputs
{

    [OutputType]
    public sealed class AviatrixFirewallPolicy
    {
        /// <summary>
        /// Valid values: "allow", "deny" and "force-drop" (in stateful firewall rule to allow immediate packet dropping on established sessions).
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Description of the policy. Example: "This is policy no.1".
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Destination address, a valid IPv4 address or tag name such "HR" or "marketing" etc. Example: "10.30.0.0/16". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        public readonly string DstIp;
        /// <summary>
        /// Valid values: true, false. Default value: false.
        /// </summary>
        public readonly bool? LogEnabled;
        /// <summary>
        /// A single port or a range of port numbers. Example: "25", "25:1024".
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// : Valid values: "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp". Default value: "all".
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Source address, a valid IPv4 address or tag name such "HR" or "marketing" etc. Example: "10.30.0.0/16". The **aviatrix_firewall_tag** resource should be created prior to using the tag name.
        /// </summary>
        public readonly string SrcIp;

        [OutputConstructor]
        private AviatrixFirewallPolicy(
            string action,

            string? description,

            string dstIp,

            bool? logEnabled,

            string port,

            string? protocol,

            string srcIp)
        {
            Action = action;
            Description = description;
            DstIp = dstIp;
            LogEnabled = logEnabled;
            Port = port;
            Protocol = protocol;
            SrcIp = srcIp;
        }
    }
}

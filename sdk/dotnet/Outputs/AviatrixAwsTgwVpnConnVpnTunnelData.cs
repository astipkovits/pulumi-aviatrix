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
    public sealed class AviatrixAwsTgwVpnConnVpnTunnelData
    {
        public readonly string? LastStatusChangeTime;
        public readonly int? RouteCount;
        public readonly string? Status;
        public readonly string? StatusMessage;
        public readonly string? TgwAsn;
        public readonly string? TunnelName;
        public readonly string? VpnInsideAddress;
        public readonly string? VpnOutsideAddress;

        [OutputConstructor]
        private AviatrixAwsTgwVpnConnVpnTunnelData(
            string? lastStatusChangeTime,

            int? routeCount,

            string? status,

            string? statusMessage,

            string? tgwAsn,

            string? tunnelName,

            string? vpnInsideAddress,

            string? vpnOutsideAddress)
        {
            LastStatusChangeTime = lastStatusChangeTime;
            RouteCount = routeCount;
            Status = status;
            StatusMessage = statusMessage;
            TgwAsn = tgwAsn;
            TunnelName = tunnelName;
            VpnInsideAddress = vpnInsideAddress;
            VpnOutsideAddress = vpnOutsideAddress;
        }
    }
}
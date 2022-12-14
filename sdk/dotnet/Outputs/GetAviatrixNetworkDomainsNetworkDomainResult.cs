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
    public sealed class GetAviatrixNetworkDomainsNetworkDomainResult
    {
        /// <summary>
        /// Access Account name.
        /// </summary>
        public readonly string Account;
        /// <summary>
        /// Type of cloud service provider.
        /// </summary>
        public readonly string CloudType;
        /// <summary>
        /// Egress inspection is enable or not.
        /// </summary>
        public readonly bool EgressInspection;
        /// <summary>
        /// Egress inspection name.
        /// </summary>
        public readonly string EgressInspectionName;
        /// <summary>
        /// Inspection policy name.
        /// </summary>
        public readonly string InspectionPolicy;
        /// <summary>
        /// Firewall inspection for traffic within one Security Domain.
        /// </summary>
        public readonly bool IntraDomainInspection;
        /// <summary>
        /// Intra domain inspection name.
        /// </summary>
        public readonly string IntraDomainInspectionName;
        /// <summary>
        /// Network Domain name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Route table's id.
        /// </summary>
        public readonly string RouteTableId;
        /// <summary>
        /// AWS TGW name.
        /// </summary>
        public readonly string TgwName;
        /// <summary>
        /// Type of network domain.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAviatrixNetworkDomainsNetworkDomainResult(
            string account,

            string cloudType,

            bool egressInspection,

            string egressInspectionName,

            string inspectionPolicy,

            bool intraDomainInspection,

            string intraDomainInspectionName,

            string name,

            string region,

            string routeTableId,

            string tgwName,

            string type)
        {
            Account = account;
            CloudType = cloudType;
            EgressInspection = egressInspection;
            EgressInspectionName = egressInspectionName;
            InspectionPolicy = inspectionPolicy;
            IntraDomainInspection = intraDomainInspection;
            IntraDomainInspectionName = intraDomainInspectionName;
            Name = name;
            Region = region;
            RouteTableId = routeTableId;
            TgwName = tgwName;
            Type = type;
        }
    }
}

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
    public sealed class GetAviatrixDeviceInterfacesWanInterfaceResult
    {
        /// <summary>
        /// Name of the WAN primary interface.
        /// </summary>
        public readonly string WanPrimaryInterface;
        /// <summary>
        /// The WAN Primary interface public IP.
        /// </summary>
        public readonly string WanPrimaryInterfacePublicIp;

        [OutputConstructor]
        private GetAviatrixDeviceInterfacesWanInterfaceResult(
            string wanPrimaryInterface,

            string wanPrimaryInterfacePublicIp)
        {
            WanPrimaryInterface = wanPrimaryInterface;
            WanPrimaryInterfacePublicIp = wanPrimaryInterfacePublicIp;
        }
    }
}

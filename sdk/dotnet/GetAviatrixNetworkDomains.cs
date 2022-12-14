// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixNetworkDomains
    {
        /// <summary>
        /// The **aviatrix_network_domains** data source provides details about all Network Domains created by the Aviatrix Controller. Available as of provider version 2.23+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// 
        ///  ```hcl
        ///  # Aviatrix All Network Domains Data Source
        ///  data "aviatrix_network_domains" "foo" {}
        ///  ```
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAviatrixNetworkDomainsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixNetworkDomainsResult>("aviatrix:index/getAviatrixNetworkDomains:getAviatrixNetworkDomains", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAviatrixNetworkDomainsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of all Network Domains
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAviatrixNetworkDomainsNetworkDomainResult> NetworkDomains;

        [OutputConstructor]
        private GetAviatrixNetworkDomainsResult(
            string id,

            ImmutableArray<Outputs.GetAviatrixNetworkDomainsNetworkDomainResult> networkDomains)
        {
            Id = id;
            NetworkDomains = networkDomains;
        }
    }
}

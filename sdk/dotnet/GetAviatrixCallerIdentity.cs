// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixCallerIdentity
    {
        /// <summary>
        /// The **aviatrix_caller_identity** data source provides the Aviatrix CID for use in other resources.
        /// </summary>
        public static Task<GetAviatrixCallerIdentityResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixCallerIdentityResult>("aviatrix:index/getAviatrixCallerIdentity:getAviatrixCallerIdentity", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAviatrixCallerIdentityResult
    {
        /// <summary>
        /// Aviatrix caller identity.
        /// </summary>
        public readonly string Cid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAviatrixCallerIdentityResult(
            string cid,

            string id)
        {
            Cid = cid;
            Id = id;
        }
    }
}

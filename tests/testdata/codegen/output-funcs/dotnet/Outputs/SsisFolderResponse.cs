// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mypkg.Outputs
{

    /// <summary>
    /// Ssis folder.
    /// </summary>
    [OutputType]
    public sealed class SsisFolderResponse
    {
        /// <summary>
        /// Metadata description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Metadata id.
        /// </summary>
        public readonly double? Id;
        /// <summary>
        /// Metadata name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of SSIS object metadata.
        /// Expected value is 'Folder'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SsisFolderResponse(
            string? description,

            double? id,

            string? name,

            string type)
        {
            Description = description;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
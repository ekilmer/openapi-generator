/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// UploadFileWithRequiredFileBody
    /// </summary>
    [DataContract(Name = "uploadFileWithRequiredFileBody")]
    public partial class UploadFileWithRequiredFileBody : IEquatable<UploadFileWithRequiredFileBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="UploadFileWithRequiredFileBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected UploadFileWithRequiredFileBody() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="UploadFileWithRequiredFileBody" /> class.
        /// </summary>
        /// <param name="additionalMetadata">Additional data to pass to server.</param>
        /// <param name="requiredFile">file to upload (required).</param>
        public UploadFileWithRequiredFileBody(string additionalMetadata = default(string), System.IO.Stream requiredFile = default(System.IO.Stream))
        {
            // to ensure "requiredFile" is required (not null)
            this.RequiredFile = requiredFile ?? throw new ArgumentNullException("requiredFile is a required property for UploadFileWithRequiredFileBody and cannot be null");
            this.AdditionalMetadata = additionalMetadata;
        }

        /// <summary>
        /// Additional data to pass to server
        /// </summary>
        /// <value>Additional data to pass to server</value>
        [DataMember(Name = "additionalMetadata", EmitDefaultValue = false)]
        public string AdditionalMetadata { get; set; }

        /// <summary>
        /// file to upload
        /// </summary>
        /// <value>file to upload</value>
        [DataMember(Name = "requiredFile", EmitDefaultValue = false)]
        public System.IO.Stream RequiredFile { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class UploadFileWithRequiredFileBody {\n");
            sb.Append("  AdditionalMetadata: ").Append(AdditionalMetadata).Append("\n");
            sb.Append("  RequiredFile: ").Append(RequiredFile).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return JsonConvert.SerializeObject(this, Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input as UploadFileWithRequiredFileBody).AreEqual;
        }

        /// <summary>
        /// Returns true if UploadFileWithRequiredFileBody instances are equal
        /// </summary>
        /// <param name="input">Instance of UploadFileWithRequiredFileBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(UploadFileWithRequiredFileBody input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input).AreEqual;
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.AdditionalMetadata != null)
                    hashCode = hashCode * 59 + this.AdditionalMetadata.GetHashCode();
                if (this.RequiredFile != null)
                    hashCode = hashCode * 59 + this.RequiredFile.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}

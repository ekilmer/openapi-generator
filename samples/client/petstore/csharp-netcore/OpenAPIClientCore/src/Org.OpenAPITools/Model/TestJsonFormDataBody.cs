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
    /// TestJsonFormDataBody
    /// </summary>
    [DataContract(Name = "testJsonFormDataBody")]
    public partial class TestJsonFormDataBody : IEquatable<TestJsonFormDataBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="TestJsonFormDataBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected TestJsonFormDataBody() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="TestJsonFormDataBody" /> class.
        /// </summary>
        /// <param name="param">field1 (required).</param>
        /// <param name="param2">field2 (required).</param>
        public TestJsonFormDataBody(string param = default(string), string param2 = default(string))
        {
            // to ensure "param" is required (not null)
            this.Param = param ?? throw new ArgumentNullException("param is a required property for TestJsonFormDataBody and cannot be null");
            // to ensure "param2" is required (not null)
            this.Param2 = param2 ?? throw new ArgumentNullException("param2 is a required property for TestJsonFormDataBody and cannot be null");
        }

        /// <summary>
        /// field1
        /// </summary>
        /// <value>field1</value>
        [DataMember(Name = "param", EmitDefaultValue = false)]
        public string Param { get; set; }

        /// <summary>
        /// field2
        /// </summary>
        /// <value>field2</value>
        [DataMember(Name = "param2", EmitDefaultValue = false)]
        public string Param2 { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class TestJsonFormDataBody {\n");
            sb.Append("  Param: ").Append(Param).Append("\n");
            sb.Append("  Param2: ").Append(Param2).Append("\n");
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
            return OpenAPIClientUtils.compareLogic.Compare(this, input as TestJsonFormDataBody).AreEqual;
        }

        /// <summary>
        /// Returns true if TestJsonFormDataBody instances are equal
        /// </summary>
        /// <param name="input">Instance of TestJsonFormDataBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(TestJsonFormDataBody input)
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
                if (this.Param != null)
                    hashCode = hashCode * 59 + this.Param.GetHashCode();
                if (this.Param2 != null)
                    hashCode = hashCode * 59 + this.Param2.GetHashCode();
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

/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.File;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.dataformat.xml.annotation.*;
import javax.xml.bind.annotation.*;

/**
 * UploadFileWithRequiredFileBody
 */
@JsonPropertyOrder({
  UploadFileWithRequiredFileBody.JSON_PROPERTY_ADDITIONAL_METADATA,
  UploadFileWithRequiredFileBody.JSON_PROPERTY_REQUIRED_FILE
})
@JsonTypeName("uploadFileWithRequiredFileBody")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
@XmlRootElement(name = "UploadFileWithRequiredFileBody")
@XmlAccessorType(XmlAccessType.FIELD)
@JacksonXmlRootElement(localName = "UploadFileWithRequiredFileBody")
public class UploadFileWithRequiredFileBody {
  public static final String JSON_PROPERTY_ADDITIONAL_METADATA = "additionalMetadata";
  @XmlElement(name = "additionalMetadata")
  private String additionalMetadata;

  public static final String JSON_PROPERTY_REQUIRED_FILE = "requiredFile";
  @XmlElement(name = "requiredFile")
  private File requiredFile;


  public UploadFileWithRequiredFileBody additionalMetadata(String additionalMetadata) {
    
    this.additionalMetadata = additionalMetadata;
    return this;
  }

   /**
   * Additional data to pass to server
   * @return additionalMetadata
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Additional data to pass to server")
  @JsonProperty(JSON_PROPERTY_ADDITIONAL_METADATA)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  @JacksonXmlProperty(localName = "additionalMetadata")

  public String getAdditionalMetadata() {
    return additionalMetadata;
  }


  public void setAdditionalMetadata(String additionalMetadata) {
    this.additionalMetadata = additionalMetadata;
  }


  public UploadFileWithRequiredFileBody requiredFile(File requiredFile) {
    
    this.requiredFile = requiredFile;
    return this;
  }

   /**
   * file to upload
   * @return requiredFile
  **/
  @ApiModelProperty(required = true, value = "file to upload")
  @JsonProperty(JSON_PROPERTY_REQUIRED_FILE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  @JacksonXmlProperty(localName = "requiredFile")

  public File getRequiredFile() {
    return requiredFile;
  }


  public void setRequiredFile(File requiredFile) {
    this.requiredFile = requiredFile;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UploadFileWithRequiredFileBody uploadFileWithRequiredFileBody = (UploadFileWithRequiredFileBody) o;
    return Objects.equals(this.additionalMetadata, uploadFileWithRequiredFileBody.additionalMetadata) &&
        Objects.equals(this.requiredFile, uploadFileWithRequiredFileBody.requiredFile);
  }

  @Override
  public int hashCode() {
    return Objects.hash(additionalMetadata, requiredFile);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UploadFileWithRequiredFileBody {\n");
    sb.append("    additionalMetadata: ").append(toIndentedString(additionalMetadata)).append("\n");
    sb.append("    requiredFile: ").append(toIndentedString(requiredFile)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

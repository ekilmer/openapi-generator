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


package org.openapitools.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.File;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import javax.validation.constraints.*;
import javax.validation.Valid;

/**
 * UploadFileWithRequiredFileBody
 */
@JsonPropertyOrder({
  UploadFileWithRequiredFileBody.JSON_PROPERTY_ADDITIONAL_METADATA,
  UploadFileWithRequiredFileBody.JSON_PROPERTY_REQUIRED_FILE
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaJerseyServerCodegen")
public class UploadFileWithRequiredFileBody   {
  public static final String JSON_PROPERTY_ADDITIONAL_METADATA = "additionalMetadata";
  @JsonProperty(JSON_PROPERTY_ADDITIONAL_METADATA)
  private String additionalMetadata;

  public static final String JSON_PROPERTY_REQUIRED_FILE = "requiredFile";
  @JsonProperty(JSON_PROPERTY_REQUIRED_FILE)
  private File requiredFile;

  public UploadFileWithRequiredFileBody additionalMetadata(String additionalMetadata) {
    this.additionalMetadata = additionalMetadata;
    return this;
  }

  /**
   * Additional data to pass to server
   * @return additionalMetadata
   **/
  @JsonProperty("additionalMetadata")
  @ApiModelProperty(value = "Additional data to pass to server")
  
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
  @JsonProperty("requiredFile")
  @ApiModelProperty(required = true, value = "file to upload")
  @NotNull 
  public File getRequiredFile() {
    return requiredFile;
  }

  public void setRequiredFile(File requiredFile) {
    this.requiredFile = requiredFile;
  }


  @Override
  public boolean equals(Object o) {
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
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}


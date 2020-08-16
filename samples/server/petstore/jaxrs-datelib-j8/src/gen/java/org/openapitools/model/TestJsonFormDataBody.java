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
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import java.io.Serializable;
import javax.validation.constraints.*;
import javax.validation.Valid;

/**
 * TestJsonFormDataBody
 */
@JsonPropertyOrder({
  TestJsonFormDataBody.JSON_PROPERTY_PARAM,
  TestJsonFormDataBody.JSON_PROPERTY_PARAM2
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaJerseyServerCodegen")
public class TestJsonFormDataBody  implements Serializable {
  public static final String JSON_PROPERTY_PARAM = "param";
  @JsonProperty(JSON_PROPERTY_PARAM)
  private String param;

  public static final String JSON_PROPERTY_PARAM2 = "param2";
  @JsonProperty(JSON_PROPERTY_PARAM2)
  private String param2;

  public TestJsonFormDataBody param(String param) {
    this.param = param;
    return this;
  }

  /**
   * field1
   * @return param
   **/
  @JsonProperty("param")
  @ApiModelProperty(required = true, value = "field1")
  @NotNull 
  public String getParam() {
    return param;
  }

  public void setParam(String param) {
    this.param = param;
  }

  public TestJsonFormDataBody param2(String param2) {
    this.param2 = param2;
    return this;
  }

  /**
   * field2
   * @return param2
   **/
  @JsonProperty("param2")
  @ApiModelProperty(required = true, value = "field2")
  @NotNull 
  public String getParam2() {
    return param2;
  }

  public void setParam2(String param2) {
    this.param2 = param2;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TestJsonFormDataBody testJsonFormDataBody = (TestJsonFormDataBody) o;
    return Objects.equals(this.param, testJsonFormDataBody.param) &&
        Objects.equals(this.param2, testJsonFormDataBody.param2);
  }

  @Override
  public int hashCode() {
    return Objects.hash(param, param2);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TestJsonFormDataBody {\n");
    
    sb.append("    param: ").append(toIndentedString(param)).append("\n");
    sb.append("    param2: ").append(toIndentedString(param2)).append("\n");
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


/*
 * test
 * test
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
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import org.openapitools.client.JSON;


/**
 * ChildSchemaAllOf
 */
@JsonPropertyOrder({
  ChildSchemaAllOf.JSON_PROPERTY_PROP1
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ChildSchemaAllOf {
  public static final String JSON_PROPERTY_PROP1 = "prop1";
  private String prop1;


  public ChildSchemaAllOf prop1(String prop1) {
    this.prop1 = prop1;
    return this;
  }

   /**
   * Get prop1
   * @return prop1
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_PROP1)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getProp1() {
    return prop1;
  }


  public void setProp1(String prop1) {
    this.prop1 = prop1;
  }


  /**
   * Return true if this ChildSchemaAllOf object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ChildSchemaAllOf childSchemaAllOf = (ChildSchemaAllOf) o;
    return Objects.equals(this.prop1, childSchemaAllOf.prop1);
  }

  @Override
  public int hashCode() {
    return Objects.hash(prop1);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ChildSchemaAllOf {\n");
    sb.append("    prop1: ").append(toIndentedString(prop1)).append("\n");
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


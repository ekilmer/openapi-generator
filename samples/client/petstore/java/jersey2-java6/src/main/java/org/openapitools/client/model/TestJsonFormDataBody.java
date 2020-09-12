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

import org.apache.commons.lang3.ObjectUtils;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * TestJsonFormDataBody
 */

public class TestJsonFormDataBody {
  public static final String SERIALIZED_NAME_PARAM = "param";
  @SerializedName(SERIALIZED_NAME_PARAM)
  private String param;

  public static final String SERIALIZED_NAME_PARAM2 = "param2";
  @SerializedName(SERIALIZED_NAME_PARAM2)
  private String param2;


  public TestJsonFormDataBody param(String param) {
    
    this.param = param;
    return this;
  }

   /**
   * field1
   * @return param
  **/
  @ApiModelProperty(required = true, value = "field1")

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
  @ApiModelProperty(required = true, value = "field2")

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
    return ObjectUtils.equals(this.param, testJsonFormDataBody.param) &&
    ObjectUtils.equals(this.param2, testJsonFormDataBody.param2);
  }

  @Override
  public int hashCode() {
    return ObjectUtils.hashCodeMulti(param, param2);
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

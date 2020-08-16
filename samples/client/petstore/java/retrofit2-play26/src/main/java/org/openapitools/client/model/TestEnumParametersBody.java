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
import java.util.ArrayList;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import javax.validation.constraints.*;
import javax.validation.Valid;

/**
 * TestEnumParametersBody
 */
@JsonPropertyOrder({
  TestEnumParametersBody.JSON_PROPERTY_ENUM_FORM_STRING_ARRAY,
  TestEnumParametersBody.JSON_PROPERTY_ENUM_FORM_STRING
})
@JsonTypeName("testEnumParametersBody")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TestEnumParametersBody {
  /**
   * Gets or Sets enumFormStringArray
   */
  public enum EnumFormStringArrayEnum {
    GREATER_THAN(">"),
    
    DOLLAR("$");

    private String value;

    EnumFormStringArrayEnum(String value) {
      this.value = value;
    }

    @JsonValue
    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    @JsonCreator
    public static EnumFormStringArrayEnum fromValue(String value) {
      for (EnumFormStringArrayEnum b : EnumFormStringArrayEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_ENUM_FORM_STRING_ARRAY = "enum_form_string_array";
  private List<EnumFormStringArrayEnum> enumFormStringArray = null;

  /**
   * Form parameter enum test (string)
   */
  public enum EnumFormStringEnum {
    _ABC("_abc"),
    
    _EFG("-efg"),
    
    _XYZ_("(xyz)");

    private String value;

    EnumFormStringEnum(String value) {
      this.value = value;
    }

    @JsonValue
    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    @JsonCreator
    public static EnumFormStringEnum fromValue(String value) {
      for (EnumFormStringEnum b : EnumFormStringEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_ENUM_FORM_STRING = "enum_form_string";
  private EnumFormStringEnum enumFormString = EnumFormStringEnum._EFG;


  public TestEnumParametersBody enumFormStringArray(List<EnumFormStringArrayEnum> enumFormStringArray) {
    
    this.enumFormStringArray = enumFormStringArray;
    return this;
  }

  public TestEnumParametersBody addEnumFormStringArrayItem(EnumFormStringArrayEnum enumFormStringArrayItem) {
    if (this.enumFormStringArray == null) {
      this.enumFormStringArray = new ArrayList<>();
    }
    this.enumFormStringArray.add(enumFormStringArrayItem);
    return this;
  }

   /**
   * Form parameter enum test (string array)
   * @return enumFormStringArray
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Form parameter enum test (string array)")
  @JsonProperty(JSON_PROPERTY_ENUM_FORM_STRING_ARRAY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<EnumFormStringArrayEnum> getEnumFormStringArray() {
    return enumFormStringArray;
  }


  public void setEnumFormStringArray(List<EnumFormStringArrayEnum> enumFormStringArray) {
    this.enumFormStringArray = enumFormStringArray;
  }


  public TestEnumParametersBody enumFormString(EnumFormStringEnum enumFormString) {
    
    this.enumFormString = enumFormString;
    return this;
  }

   /**
   * Form parameter enum test (string)
   * @return enumFormString
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Form parameter enum test (string)")
  @JsonProperty(JSON_PROPERTY_ENUM_FORM_STRING)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public EnumFormStringEnum getEnumFormString() {
    return enumFormString;
  }


  public void setEnumFormString(EnumFormStringEnum enumFormString) {
    this.enumFormString = enumFormString;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TestEnumParametersBody testEnumParametersBody = (TestEnumParametersBody) o;
    return Objects.equals(this.enumFormStringArray, testEnumParametersBody.enumFormStringArray) &&
        Objects.equals(this.enumFormString, testEnumParametersBody.enumFormString);
  }

  @Override
  public int hashCode() {
    return Objects.hash(enumFormStringArray, enumFormString);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TestEnumParametersBody {\n");
    sb.append("    enumFormStringArray: ").append(toIndentedString(enumFormStringArray)).append("\n");
    sb.append("    enumFormString: ").append(toIndentedString(enumFormString)).append("\n");
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


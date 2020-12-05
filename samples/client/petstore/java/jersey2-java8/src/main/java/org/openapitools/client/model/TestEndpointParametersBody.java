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
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.File;
import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.OffsetDateTime;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import org.openapitools.client.JSON;


/**
 * TestEndpointParametersBody
 */
@JsonPropertyOrder({
  TestEndpointParametersBody.JSON_PROPERTY_INTEGER,
  TestEndpointParametersBody.JSON_PROPERTY_INT32,
  TestEndpointParametersBody.JSON_PROPERTY_INT64,
  TestEndpointParametersBody.JSON_PROPERTY_NUMBER,
  TestEndpointParametersBody.JSON_PROPERTY_FLOAT,
  TestEndpointParametersBody.JSON_PROPERTY_DOUBLE,
  TestEndpointParametersBody.JSON_PROPERTY_STRING,
  TestEndpointParametersBody.JSON_PROPERTY_PATTERN_WITHOUT_DELIMITER,
  TestEndpointParametersBody.JSON_PROPERTY_BYTE,
  TestEndpointParametersBody.JSON_PROPERTY_BINARY,
  TestEndpointParametersBody.JSON_PROPERTY_DATE,
  TestEndpointParametersBody.JSON_PROPERTY_DATE_TIME,
  TestEndpointParametersBody.JSON_PROPERTY_PASSWORD,
  TestEndpointParametersBody.JSON_PROPERTY_CALLBACK
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class TestEndpointParametersBody {
  public static final String JSON_PROPERTY_INTEGER = "integer";
  private Integer integer;

  public static final String JSON_PROPERTY_INT32 = "int32";
  private Integer int32;

  public static final String JSON_PROPERTY_INT64 = "int64";
  private Long int64;

  public static final String JSON_PROPERTY_NUMBER = "number";
  private BigDecimal number;

  public static final String JSON_PROPERTY_FLOAT = "float";
  private Float _float;

  public static final String JSON_PROPERTY_DOUBLE = "double";
  private Double _double;

  public static final String JSON_PROPERTY_STRING = "string";
  private String string;

  public static final String JSON_PROPERTY_PATTERN_WITHOUT_DELIMITER = "pattern_without_delimiter";
  private String patternWithoutDelimiter;

  public static final String JSON_PROPERTY_BYTE = "byte";
  private byte[] _byte;

  public static final String JSON_PROPERTY_BINARY = "binary";
  private File binary;

  public static final String JSON_PROPERTY_DATE = "date";
  private LocalDate date;

  public static final String JSON_PROPERTY_DATE_TIME = "dateTime";
  private OffsetDateTime dateTime;

  public static final String JSON_PROPERTY_PASSWORD = "password";
  private String password;

  public static final String JSON_PROPERTY_CALLBACK = "callback";
  private String callback;


  public TestEndpointParametersBody integer(Integer integer) {
    this.integer = integer;
    return this;
  }

   /**
   * None
   * minimum: 10
   * maximum: 100
   * @return integer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_INTEGER)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getInteger() {
    return integer;
  }


  public void setInteger(Integer integer) {
    this.integer = integer;
  }


  public TestEndpointParametersBody int32(Integer int32) {
    this.int32 = int32;
    return this;
  }

   /**
   * None
   * minimum: 20
   * maximum: 200
   * @return int32
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_INT32)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getInt32() {
    return int32;
  }


  public void setInt32(Integer int32) {
    this.int32 = int32;
  }


  public TestEndpointParametersBody int64(Long int64) {
    this.int64 = int64;
    return this;
  }

   /**
   * None
   * @return int64
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_INT64)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Long getInt64() {
    return int64;
  }


  public void setInt64(Long int64) {
    this.int64 = int64;
  }


  public TestEndpointParametersBody number(BigDecimal number) {
    this.number = number;
    return this;
  }

   /**
   * None
   * minimum: 32.1
   * maximum: 543.2
   * @return number
  **/
  @ApiModelProperty(required = true, value = "None")
  @JsonProperty(JSON_PROPERTY_NUMBER)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public BigDecimal getNumber() {
    return number;
  }


  public void setNumber(BigDecimal number) {
    this.number = number;
  }


  public TestEndpointParametersBody _float(Float _float) {
    this._float = _float;
    return this;
  }

   /**
   * None
   * maximum: 987.6
   * @return _float
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_FLOAT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Float getFloat() {
    return _float;
  }


  public void setFloat(Float _float) {
    this._float = _float;
  }


  public TestEndpointParametersBody _double(Double _double) {
    this._double = _double;
    return this;
  }

   /**
   * None
   * minimum: 67.8
   * maximum: 123.4
   * @return _double
  **/
  @ApiModelProperty(required = true, value = "None")
  @JsonProperty(JSON_PROPERTY_DOUBLE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public Double getDouble() {
    return _double;
  }


  public void setDouble(Double _double) {
    this._double = _double;
  }


  public TestEndpointParametersBody string(String string) {
    this.string = string;
    return this;
  }

   /**
   * None
   * @return string
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_STRING)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getString() {
    return string;
  }


  public void setString(String string) {
    this.string = string;
  }


  public TestEndpointParametersBody patternWithoutDelimiter(String patternWithoutDelimiter) {
    this.patternWithoutDelimiter = patternWithoutDelimiter;
    return this;
  }

   /**
   * None
   * @return patternWithoutDelimiter
  **/
  @ApiModelProperty(required = true, value = "None")
  @JsonProperty(JSON_PROPERTY_PATTERN_WITHOUT_DELIMITER)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getPatternWithoutDelimiter() {
    return patternWithoutDelimiter;
  }


  public void setPatternWithoutDelimiter(String patternWithoutDelimiter) {
    this.patternWithoutDelimiter = patternWithoutDelimiter;
  }


  public TestEndpointParametersBody _byte(byte[] _byte) {
    this._byte = _byte;
    return this;
  }

   /**
   * None
   * @return _byte
  **/
  @ApiModelProperty(required = true, value = "None")
  @JsonProperty(JSON_PROPERTY_BYTE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public byte[] getByte() {
    return _byte;
  }


  public void setByte(byte[] _byte) {
    this._byte = _byte;
  }


  public TestEndpointParametersBody binary(File binary) {
    this.binary = binary;
    return this;
  }

   /**
   * None
   * @return binary
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_BINARY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public File getBinary() {
    return binary;
  }


  public void setBinary(File binary) {
    this.binary = binary;
  }


  public TestEndpointParametersBody date(LocalDate date) {
    this.date = date;
    return this;
  }

   /**
   * None
   * @return date
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_DATE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public LocalDate getDate() {
    return date;
  }


  public void setDate(LocalDate date) {
    this.date = date;
  }


  public TestEndpointParametersBody dateTime(OffsetDateTime dateTime) {
    this.dateTime = dateTime;
    return this;
  }

   /**
   * None
   * @return dateTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_DATE_TIME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public OffsetDateTime getDateTime() {
    return dateTime;
  }


  public void setDateTime(OffsetDateTime dateTime) {
    this.dateTime = dateTime;
  }


  public TestEndpointParametersBody password(String password) {
    this.password = password;
    return this;
  }

   /**
   * None
   * @return password
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_PASSWORD)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public TestEndpointParametersBody callback(String callback) {
    this.callback = callback;
    return this;
  }

   /**
   * None
   * @return callback
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "None")
  @JsonProperty(JSON_PROPERTY_CALLBACK)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getCallback() {
    return callback;
  }


  public void setCallback(String callback) {
    this.callback = callback;
  }


  /**
   * Return true if this testEndpointParametersBody object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TestEndpointParametersBody testEndpointParametersBody = (TestEndpointParametersBody) o;
    return Objects.equals(this.integer, testEndpointParametersBody.integer) &&
        Objects.equals(this.int32, testEndpointParametersBody.int32) &&
        Objects.equals(this.int64, testEndpointParametersBody.int64) &&
        Objects.equals(this.number, testEndpointParametersBody.number) &&
        Objects.equals(this._float, testEndpointParametersBody._float) &&
        Objects.equals(this._double, testEndpointParametersBody._double) &&
        Objects.equals(this.string, testEndpointParametersBody.string) &&
        Objects.equals(this.patternWithoutDelimiter, testEndpointParametersBody.patternWithoutDelimiter) &&
        Arrays.equals(this._byte, testEndpointParametersBody._byte) &&
        Objects.equals(this.binary, testEndpointParametersBody.binary) &&
        Objects.equals(this.date, testEndpointParametersBody.date) &&
        Objects.equals(this.dateTime, testEndpointParametersBody.dateTime) &&
        Objects.equals(this.password, testEndpointParametersBody.password) &&
        Objects.equals(this.callback, testEndpointParametersBody.callback);
  }

  @Override
  public int hashCode() {
    return Objects.hash(integer, int32, int64, number, _float, _double, string, patternWithoutDelimiter, Arrays.hashCode(_byte), binary, date, dateTime, password, callback);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TestEndpointParametersBody {\n");
    sb.append("    integer: ").append(toIndentedString(integer)).append("\n");
    sb.append("    int32: ").append(toIndentedString(int32)).append("\n");
    sb.append("    int64: ").append(toIndentedString(int64)).append("\n");
    sb.append("    number: ").append(toIndentedString(number)).append("\n");
    sb.append("    _float: ").append(toIndentedString(_float)).append("\n");
    sb.append("    _double: ").append(toIndentedString(_double)).append("\n");
    sb.append("    string: ").append(toIndentedString(string)).append("\n");
    sb.append("    patternWithoutDelimiter: ").append(toIndentedString(patternWithoutDelimiter)).append("\n");
    sb.append("    _byte: ").append(toIndentedString(_byte)).append("\n");
    sb.append("    binary: ").append(toIndentedString(binary)).append("\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    dateTime: ").append(toIndentedString(dateTime)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    callback: ").append(toIndentedString(callback)).append("\n");
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


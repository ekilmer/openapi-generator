package org.openapitools.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import javax.validation.Valid;
import javax.validation.constraints.*;

/**
 * TestJsonFormDataBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.SpringCodegen")@com.fasterxml.jackson.annotation.JsonFilter(value = "filter-name")
@com.fasterxml.jackson.annotation.JsonIgnoreProperties(value = "id")

public class TestJsonFormDataBody   {
  @JsonProperty("param")
  private String param;

  @JsonProperty("param2")
  private String param2;

  public TestJsonFormDataBody param(String param) {
    this.param = param;
    return this;
  }

  /**
   * field1
   * @return param
  */
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
  */
  @ApiModelProperty(required = true, value = "field2")
  @NotNull


  public String getParam2() {
    return param2;
  }

  public void setParam2(String param2) {
    this.param2 = param2;
  }


  @Override
  public boolean equals(Object o) {
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
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}

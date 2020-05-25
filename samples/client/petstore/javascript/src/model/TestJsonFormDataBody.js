/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 *
 * OpenAPI Generator version: 5.0.0-SNAPSHOT
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OpenApiPetstore) {
      root.OpenApiPetstore = {};
    }
    root.OpenApiPetstore.TestJsonFormDataBody = factory(root.OpenApiPetstore.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';



  /**
   * The TestJsonFormDataBody model module.
   * @module model/TestJsonFormDataBody
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>TestJsonFormDataBody</code>.
   * @alias module:model/TestJsonFormDataBody
   * @class
   * @param param {String} field1
   * @param param2 {String} field2
   */
  var exports = function(param, param2) {
    var _this = this;

    _this['param'] = param;
    _this['param2'] = param2;
  };

  /**
   * Constructs a <code>TestJsonFormDataBody</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TestJsonFormDataBody} obj Optional instance to populate.
   * @return {module:model/TestJsonFormDataBody} The populated <code>TestJsonFormDataBody</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('param')) {
        obj['param'] = ApiClient.convertToType(data['param'], 'String');
      }
      if (data.hasOwnProperty('param2')) {
        obj['param2'] = ApiClient.convertToType(data['param2'], 'String');
      }
    }
    return obj;
  }

  /**
   * field1
   * @member {String} param
   */
  exports.prototype['param'] = undefined;
  /**
   * field2
   * @member {String} param2
   */
  exports.prototype['param2'] = undefined;



  return exports;
}));



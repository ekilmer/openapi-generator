/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.OpenApiPetstore);
  }
}(this, function(expect, OpenApiPetstore) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new OpenApiPetstore.TestEnumParametersBody();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('TestEnumParametersBody', function() {
    it('should create an instance of TestEnumParametersBody', function() {
      // uncomment below and update the code to test TestEnumParametersBody
      //var instane = new OpenApiPetstore.TestEnumParametersBody();
      //expect(instance).to.be.a(OpenApiPetstore.TestEnumParametersBody);
    });

    it('should have the property enumFormStringArray (base name: "enum_form_string_array")', function() {
      // uncomment below and update the code to test the property enumFormStringArray
      //var instane = new OpenApiPetstore.TestEnumParametersBody();
      //expect(instance).to.be();
    });

    it('should have the property enumFormString (base name: "enum_form_string")', function() {
      // uncomment below and update the code to test the property enumFormString
      //var instane = new OpenApiPetstore.TestEnumParametersBody();
      //expect(instance).to.be();
    });

  });

}));

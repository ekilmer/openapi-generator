# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from petstore_api.configuration import Configuration


class TestEnumParametersBody(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'enum_form_string_array': 'list[str]',
        'enum_form_string': 'str'
    }

    attribute_map = {
        'enum_form_string_array': 'enum_form_string_array',
        'enum_form_string': 'enum_form_string'
    }

    def __init__(self, enum_form_string_array=None, enum_form_string='-efg', local_vars_configuration=None):  # noqa: E501
        """TestEnumParametersBody - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._enum_form_string_array = None
        self._enum_form_string = None
        self.discriminator = None

        if enum_form_string_array is not None:
            self.enum_form_string_array = enum_form_string_array
        if enum_form_string is not None:
            self.enum_form_string = enum_form_string

    @property
    def enum_form_string_array(self):
        """Gets the enum_form_string_array of this TestEnumParametersBody.  # noqa: E501

        Form parameter enum test (string array)  # noqa: E501

        :return: The enum_form_string_array of this TestEnumParametersBody.  # noqa: E501
        :rtype: list[str]
        """
        return self._enum_form_string_array

    @enum_form_string_array.setter
    def enum_form_string_array(self, enum_form_string_array):
        """Sets the enum_form_string_array of this TestEnumParametersBody.

        Form parameter enum test (string array)  # noqa: E501

        :param enum_form_string_array: The enum_form_string_array of this TestEnumParametersBody.  # noqa: E501
        :type: list[str]
        """
        allowed_values = [">", "$"]  # noqa: E501
        if (self.local_vars_configuration.client_side_validation and
                not set(enum_form_string_array).issubset(set(allowed_values))):  # noqa: E501
            raise ValueError(
                "Invalid values for `enum_form_string_array` [{0}], must be a subset of [{1}]"  # noqa: E501
                .format(", ".join(map(str, set(enum_form_string_array) - set(allowed_values))),  # noqa: E501
                        ", ".join(map(str, allowed_values)))
            )

        self._enum_form_string_array = enum_form_string_array

    @property
    def enum_form_string(self):
        """Gets the enum_form_string of this TestEnumParametersBody.  # noqa: E501

        Form parameter enum test (string)  # noqa: E501

        :return: The enum_form_string of this TestEnumParametersBody.  # noqa: E501
        :rtype: str
        """
        return self._enum_form_string

    @enum_form_string.setter
    def enum_form_string(self, enum_form_string):
        """Sets the enum_form_string of this TestEnumParametersBody.

        Form parameter enum test (string)  # noqa: E501

        :param enum_form_string: The enum_form_string of this TestEnumParametersBody.  # noqa: E501
        :type: str
        """
        allowed_values = ["_abc", "-efg", "(xyz)"]  # noqa: E501
        if self.local_vars_configuration.client_side_validation and enum_form_string not in allowed_values:  # noqa: E501
            raise ValueError(
                "Invalid value for `enum_form_string` ({0}), must be one of {1}"  # noqa: E501
                .format(enum_form_string, allowed_values)
            )

        self._enum_form_string = enum_form_string

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, TestEnumParametersBody):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, TestEnumParametersBody):
            return True

        return self.to_dict() != other.to_dict()

# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from openapi_server.models.base_model_ import Model
from openapi_server import util


class UpdatePetWithFormBody(Model):
    """NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

    Do not edit the class manually.
    """

    def __init__(self, name=None, status=None):  # noqa: E501
        """UpdatePetWithFormBody - a model defined in OpenAPI

        :param name: The name of this UpdatePetWithFormBody.  # noqa: E501
        :type name: str
        :param status: The status of this UpdatePetWithFormBody.  # noqa: E501
        :type status: str
        """
        self.openapi_types = {
            'name': str,
            'status': str
        }

        self.attribute_map = {
            'name': 'name',
            'status': 'status'
        }

        self._name = name
        self._status = status

    @classmethod
    def from_dict(cls, dikt) -> 'UpdatePetWithFormBody':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The updatePetWithFormBody of this UpdatePetWithFormBody.  # noqa: E501
        :rtype: UpdatePetWithFormBody
        """
        return util.deserialize_model(dikt, cls)

    @property
    def name(self):
        """Gets the name of this UpdatePetWithFormBody.

        Updated name of the pet  # noqa: E501

        :return: The name of this UpdatePetWithFormBody.
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this UpdatePetWithFormBody.

        Updated name of the pet  # noqa: E501

        :param name: The name of this UpdatePetWithFormBody.
        :type name: str
        """

        self._name = name

    @property
    def status(self):
        """Gets the status of this UpdatePetWithFormBody.

        Updated status of the pet  # noqa: E501

        :return: The status of this UpdatePetWithFormBody.
        :rtype: str
        """
        return self._status

    @status.setter
    def status(self, status):
        """Sets the status of this UpdatePetWithFormBody.

        Updated status of the pet  # noqa: E501

        :param status: The status of this UpdatePetWithFormBody.
        :type status: str
        """

        self._status = status
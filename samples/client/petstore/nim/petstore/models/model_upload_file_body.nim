#
# OpenAPI Petstore
# 
# This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
# The version of the OpenAPI document: 1.0.0
# 
# Generated by: https://openapi-generator.tech
#

import json
import tables


type UploadFileBody* = object
  ## 
  additionalMetadata*: string ## Additional data to pass to server
  file*: string ## file to upload
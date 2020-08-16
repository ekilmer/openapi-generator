/**
* OpenAPI Petstore
* This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
*
* The version of the OpenAPI document: 1.0.0
* 
*
* NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
* https://openapi-generator.tech
* Do not edit the class manually.
*/
package org.openapitools.client.models


import kotlinx.serialization.*
import kotlinx.serialization.internal.CommonEnumSerializer

/**
 * 
 * @param additionalMetadata Additional data to pass to server
 * @param file file to upload
 */
@Serializable
data class UploadFileBody (
    /* Additional data to pass to server */
    @SerialName(value = "additionalMetadata") val additionalMetadata: kotlin.String? = null,
    /* file to upload */
    @SerialName(value = "file") val file: org.openapitools.client.infrastructure.OctetByteArray? = null
)


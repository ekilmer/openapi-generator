<?php
/**
 * UploadFileWithRequiredFileBody
 *
 * PHP version 7
 *
 * @package OpenAPIServer\Model
 * @author  OpenAPI Generator team
 * @link    https://github.com/openapitools/openapi-generator
 */

/**
 * NOTE: This class is auto generated by the openapi generator program.
 * https://github.com/openapitools/openapi-generator
 */
namespace OpenAPIServer\Model;

/**
 * UploadFileWithRequiredFileBody
 *
 * @package OpenAPIServer\Model
 * @author  OpenAPI Generator team
 * @link    https://github.com/openapitools/openapi-generator
 */
class UploadFileWithRequiredFileBody
{
    
    /** @var string $additionalMetadata Additional data to pass to server*/
    private $additionalMetadata;
    
    /** @var \SplFileObject $requiredFile file to upload*/
    private $requiredFile;
}
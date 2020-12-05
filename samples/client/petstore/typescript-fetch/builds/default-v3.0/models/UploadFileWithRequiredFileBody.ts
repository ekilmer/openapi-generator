/* tslint:disable */
/* eslint-disable */
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
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UploadFileWithRequiredFileBody
 */
export interface UploadFileWithRequiredFileBody {
    /**
     * Additional data to pass to server
     * @type {string}
     * @memberof UploadFileWithRequiredFileBody
     */
    additionalMetadata?: string;
    /**
     * file to upload
     * @type {Blob}
     * @memberof UploadFileWithRequiredFileBody
     */
    requiredFile: Blob;
}

export function UploadFileWithRequiredFileBodyFromJSON(json: any): UploadFileWithRequiredFileBody {
    return UploadFileWithRequiredFileBodyFromJSONTyped(json, false);
}

export function UploadFileWithRequiredFileBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): UploadFileWithRequiredFileBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'additionalMetadata': !exists(json, 'additionalMetadata') ? undefined : json['additionalMetadata'],
        'requiredFile': json['requiredFile'],
    };
}

export function UploadFileWithRequiredFileBodyToJSON(value?: UploadFileWithRequiredFileBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'additionalMetadata': value.additionalMetadata,
        'requiredFile': value.requiredFile,
    };
}



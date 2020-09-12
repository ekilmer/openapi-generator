// tslint:disable
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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UploadFileBody
 */
export interface UploadFileBody  {
    /**
     * Additional data to pass to server
     * @type {string}
     * @memberof UploadFileBody
     */
    additionalMetadata?: string;
    /**
     * file to upload
     * @type {Blob}
     * @memberof UploadFileBody
     */
    file?: Blob;
}

export function UploadFileBodyFromJSON(json: any): UploadFileBody {
    return {
        'additionalMetadata': !exists(json, 'additionalMetadata') ? undefined : json['additionalMetadata'],
        'file': !exists(json, 'file') ? undefined : json['file'],
    };
}

export function UploadFileBodyToJSON(value?: UploadFileBody): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'additionalMetadata': value.additionalMetadata,
        'file': value.file,
    };
}


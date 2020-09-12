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
 * @interface UpdatePetWithFormBody
 */
export interface UpdatePetWithFormBody  {
    /**
     * Updated name of the pet
     * @type {string}
     * @memberof UpdatePetWithFormBody
     */
    name?: string;
    /**
     * Updated status of the pet
     * @type {string}
     * @memberof UpdatePetWithFormBody
     */
    status?: string;
}

export function UpdatePetWithFormBodyFromJSON(json: any): UpdatePetWithFormBody {
    return {
        'name': !exists(json, 'name') ? undefined : json['name'],
        'status': !exists(json, 'status') ? undefined : json['status'],
    };
}

export function UpdatePetWithFormBodyToJSON(value?: UpdatePetWithFormBody): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'name': value.name,
        'status': value.status,
    };
}


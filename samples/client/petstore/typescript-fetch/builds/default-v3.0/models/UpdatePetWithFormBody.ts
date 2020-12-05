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
 * @interface UpdatePetWithFormBody
 */
export interface UpdatePetWithFormBody {
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
    return UpdatePetWithFormBodyFromJSONTyped(json, false);
}

export function UpdatePetWithFormBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdatePetWithFormBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'status': !exists(json, 'status') ? undefined : json['status'],
    };
}

export function UpdatePetWithFormBodyToJSON(value?: UpdatePetWithFormBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'status': value.status,
    };
}



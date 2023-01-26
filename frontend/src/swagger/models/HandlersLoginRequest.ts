/* tslint:disable */
/* eslint-disable */
/**
 * Distance Challenge API
 * Distance Challenge API.
 *
 * The version of the OpenAPI document: 1.0
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
 * @interface HandlersLoginRequest
 */
export interface HandlersLoginRequest {
    /**
     * 
     * @type {string}
     * @memberof HandlersLoginRequest
     */
    password?: string;
    /**
     * 
     * @type {string}
     * @memberof HandlersLoginRequest
     */
    username?: string;
}

export function HandlersLoginRequestFromJSON(json: any): HandlersLoginRequest {
    return HandlersLoginRequestFromJSONTyped(json, false);
}

export function HandlersLoginRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): HandlersLoginRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'password': !exists(json, 'password') ? undefined : json['password'],
        'username': !exists(json, 'username') ? undefined : json['username'],
    };
}

export function HandlersLoginRequestToJSON(value?: HandlersLoginRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'password': value.password,
        'username': value.username,
    };
}

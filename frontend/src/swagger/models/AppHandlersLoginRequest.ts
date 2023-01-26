/* tslint:disable */
/* eslint-disable */
/**
 * Echo Swagger Example API
 * This is a sample server server.
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface AppHandlersLoginRequest
 */
export interface AppHandlersLoginRequest {
    /**
     * 
     * @type {string}
     * @memberof AppHandlersLoginRequest
     */
    password?: string;
    /**
     * 
     * @type {string}
     * @memberof AppHandlersLoginRequest
     */
    username?: string;
}

export function AppHandlersLoginRequestFromJSON(json: any): AppHandlersLoginRequest {
    return AppHandlersLoginRequestFromJSONTyped(json, false);
}

export function AppHandlersLoginRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AppHandlersLoginRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'password': !exists(json, 'password') ? undefined : json['password'],
        'username': !exists(json, 'username') ? undefined : json['username'],
    };
}

export function AppHandlersLoginRequestToJSON(value?: AppHandlersLoginRequest | null): any {
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


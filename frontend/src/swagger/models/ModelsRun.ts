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
 * @interface ModelsRun
 */
export interface ModelsRun {
    /**
     * 
     * @type {string}
     * @memberof ModelsRun
     */
    createdAt?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelsRun
     */
    distance?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsRun
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsRun
     */
    time?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelsRun
     */
    updatedAt?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelsRun
     */
    userId?: number;
}

export function ModelsRunFromJSON(json: any): ModelsRun {
    return ModelsRunFromJSONTyped(json, false);
}

export function ModelsRunFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelsRun {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createdAt': !exists(json, 'created_at') ? undefined : json['created_at'],
        'distance': !exists(json, 'distance') ? undefined : json['distance'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'time': !exists(json, 'time') ? undefined : json['time'],
        'updatedAt': !exists(json, 'updated_at') ? undefined : json['updated_at'],
        'userId': !exists(json, 'user_id') ? undefined : json['user_id'],
    };
}

export function ModelsRunToJSON(value?: ModelsRun | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'created_at': value.createdAt,
        'distance': value.distance,
        'id': value.id,
        'time': value.time,
        'updated_at': value.updatedAt,
        'user_id': value.userId,
    };
}

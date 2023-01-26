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
 * @interface GithubComSerbrodaDistanceChallengeModelsRun
 */
export interface GithubComSerbrodaDistanceChallengeModelsRun {
    /**
     * 
     * @type {string}
     * @memberof GithubComSerbrodaDistanceChallengeModelsRun
     */
    createdAt?: string;
    /**
     * 
     * @type {number}
     * @memberof GithubComSerbrodaDistanceChallengeModelsRun
     */
    distance?: number;
    /**
     * 
     * @type {number}
     * @memberof GithubComSerbrodaDistanceChallengeModelsRun
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof GithubComSerbrodaDistanceChallengeModelsRun
     */
    time?: number;
    /**
     * 
     * @type {string}
     * @memberof GithubComSerbrodaDistanceChallengeModelsRun
     */
    updatedAt?: string;
    /**
     * 
     * @type {number}
     * @memberof GithubComSerbrodaDistanceChallengeModelsRun
     */
    userId?: number;
}

export function GithubComSerbrodaDistanceChallengeModelsRunFromJSON(json: any): GithubComSerbrodaDistanceChallengeModelsRun {
    return GithubComSerbrodaDistanceChallengeModelsRunFromJSONTyped(json, false);
}

export function GithubComSerbrodaDistanceChallengeModelsRunFromJSONTyped(json: any, ignoreDiscriminator: boolean): GithubComSerbrodaDistanceChallengeModelsRun {
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

export function GithubComSerbrodaDistanceChallengeModelsRunToJSON(value?: GithubComSerbrodaDistanceChallengeModelsRun | null): any {
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


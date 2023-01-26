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


import * as runtime from '../runtime';
import {
    GithubComSerbrodaDistanceChallengeModelsRun,
    GithubComSerbrodaDistanceChallengeModelsRunFromJSON,
    GithubComSerbrodaDistanceChallengeModelsRunToJSON,
    GithubComSerbrodaDistanceChallengeModelsUser,
    GithubComSerbrodaDistanceChallengeModelsUserFromJSON,
    GithubComSerbrodaDistanceChallengeModelsUserToJSON,
} from '../models';

export interface RunsIdDeleteRequest {
    id: string;
}

export interface RunsIdGetRequest {
    id: string;
}

export interface RunsIdPutRequest {
    id: string;
    user: GithubComSerbrodaDistanceChallengeModelsRun;
}

export interface UsersIdRunsPostRequest {
    id: string;
    user: GithubComSerbrodaDistanceChallengeModelsRun;
}

export interface UsersUserIdActivatePutRequest {
    userId: string;
    user: GithubComSerbrodaDistanceChallengeModelsUser;
}

export interface UsersUserIdGetRequest {
    userId: string;
}

export interface UsersUserIdPutRequest {
    userId: string;
    user: GithubComSerbrodaDistanceChallengeModelsUser;
}

/**
 * 
 */
export class ApiApi extends runtime.BaseAPI {

    /**
     */
    async meGetRaw(initOverrides?: RequestInit): Promise<runtime.ApiResponse<GithubComSerbrodaDistanceChallengeModelsUser>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/me`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GithubComSerbrodaDistanceChallengeModelsUserFromJSON(jsonValue));
    }

    /**
     */
    async meGet(initOverrides?: RequestInit): Promise<GithubComSerbrodaDistanceChallengeModelsUser> {
        const response = await this.meGetRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async runsGetRaw(initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<GithubComSerbrodaDistanceChallengeModelsRun>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/runs`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(GithubComSerbrodaDistanceChallengeModelsRunFromJSON));
    }

    /**
     */
    async runsGet(initOverrides?: RequestInit): Promise<Array<GithubComSerbrodaDistanceChallengeModelsRun>> {
        const response = await this.runsGetRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async runsIdDeleteRaw(requestParameters: RunsIdDeleteRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling runsIdDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/runs/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async runsIdDelete(requestParameters: RunsIdDeleteRequest, initOverrides?: RequestInit): Promise<void> {
        await this.runsIdDeleteRaw(requestParameters, initOverrides);
    }

    /**
     */
    async runsIdGetRaw(requestParameters: RunsIdGetRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<GithubComSerbrodaDistanceChallengeModelsRun>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling runsIdGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/runs/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GithubComSerbrodaDistanceChallengeModelsRunFromJSON(jsonValue));
    }

    /**
     */
    async runsIdGet(requestParameters: RunsIdGetRequest, initOverrides?: RequestInit): Promise<GithubComSerbrodaDistanceChallengeModelsRun> {
        const response = await this.runsIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async runsIdPutRaw(requestParameters: RunsIdPutRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<GithubComSerbrodaDistanceChallengeModelsRun>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling runsIdPut.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling runsIdPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/runs/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: GithubComSerbrodaDistanceChallengeModelsRunToJSON(requestParameters.user),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GithubComSerbrodaDistanceChallengeModelsRunFromJSON(jsonValue));
    }

    /**
     */
    async runsIdPut(requestParameters: RunsIdPutRequest, initOverrides?: RequestInit): Promise<GithubComSerbrodaDistanceChallengeModelsRun> {
        const response = await this.runsIdPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async usersGetRaw(initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<GithubComSerbrodaDistanceChallengeModelsUser>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(GithubComSerbrodaDistanceChallengeModelsUserFromJSON));
    }

    /**
     */
    async usersGet(initOverrides?: RequestInit): Promise<Array<GithubComSerbrodaDistanceChallengeModelsUser>> {
        const response = await this.usersGetRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async usersIdRunsPostRaw(requestParameters: UsersIdRunsPostRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<GithubComSerbrodaDistanceChallengeModelsRun>>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling usersIdRunsPost.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling usersIdRunsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{id}/runs`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: GithubComSerbrodaDistanceChallengeModelsRunToJSON(requestParameters.user),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(GithubComSerbrodaDistanceChallengeModelsRunFromJSON));
    }

    /**
     */
    async usersIdRunsPost(requestParameters: UsersIdRunsPostRequest, initOverrides?: RequestInit): Promise<Array<GithubComSerbrodaDistanceChallengeModelsRun>> {
        const response = await this.usersIdRunsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async usersUserIdActivatePutRaw(requestParameters: UsersUserIdActivatePutRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<GithubComSerbrodaDistanceChallengeModelsUser>> {
        if (requestParameters.userId === null || requestParameters.userId === undefined) {
            throw new runtime.RequiredError('userId','Required parameter requestParameters.userId was null or undefined when calling usersUserIdActivatePut.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling usersUserIdActivatePut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{userId}/activate`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters.userId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: GithubComSerbrodaDistanceChallengeModelsUserToJSON(requestParameters.user),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GithubComSerbrodaDistanceChallengeModelsUserFromJSON(jsonValue));
    }

    /**
     */
    async usersUserIdActivatePut(requestParameters: UsersUserIdActivatePutRequest, initOverrides?: RequestInit): Promise<GithubComSerbrodaDistanceChallengeModelsUser> {
        const response = await this.usersUserIdActivatePutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async usersUserIdGetRaw(requestParameters: UsersUserIdGetRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<GithubComSerbrodaDistanceChallengeModelsUser>> {
        if (requestParameters.userId === null || requestParameters.userId === undefined) {
            throw new runtime.RequiredError('userId','Required parameter requestParameters.userId was null or undefined when calling usersUserIdGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{userId}`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters.userId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GithubComSerbrodaDistanceChallengeModelsUserFromJSON(jsonValue));
    }

    /**
     */
    async usersUserIdGet(requestParameters: UsersUserIdGetRequest, initOverrides?: RequestInit): Promise<GithubComSerbrodaDistanceChallengeModelsUser> {
        const response = await this.usersUserIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async usersUserIdPutRaw(requestParameters: UsersUserIdPutRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<GithubComSerbrodaDistanceChallengeModelsUser>> {
        if (requestParameters.userId === null || requestParameters.userId === undefined) {
            throw new runtime.RequiredError('userId','Required parameter requestParameters.userId was null or undefined when calling usersUserIdPut.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling usersUserIdPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{userId}`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters.userId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: GithubComSerbrodaDistanceChallengeModelsUserToJSON(requestParameters.user),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GithubComSerbrodaDistanceChallengeModelsUserFromJSON(jsonValue));
    }

    /**
     */
    async usersUserIdPut(requestParameters: UsersUserIdPutRequest, initOverrides?: RequestInit): Promise<GithubComSerbrodaDistanceChallengeModelsUser> {
        const response = await this.usersUserIdPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

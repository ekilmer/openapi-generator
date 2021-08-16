import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { Cat } from '../models/Cat';
import { CatAllOf } from '../models/CatAllOf';
import { Dog } from '../models/Dog';
import { DogAllOf } from '../models/DogAllOf';
import { InlineBody } from '../models/InlineBody';
import { InlineBody1 } from '../models/InlineBody1';
import { InlineBody2 } from '../models/InlineBody2';
import { InlineBody2File } from '../models/InlineBody2File';
import { PetByAge } from '../models/PetByAge';
import { PetByType } from '../models/PetByType';

import { ObservableDefaultApi } from "./ObservableAPI";
import { DefaultApiRequestFactory, DefaultApiResponseProcessor} from "../apis/DefaultApi";

export interface DefaultApiFilePostRequest {
    /**
     * 
     * @type InlineBody2
     * @memberof DefaultApifilePost
     */
    inlineBody2?: InlineBody2
}

export interface DefaultApiPetsFilteredPatchRequest {
    /**
     * 
     * @type InlineBody1
     * @memberof DefaultApipetsFilteredPatch
     */
    inlineBody1?: InlineBody1
}

export interface DefaultApiPetsPatchRequest {
    /**
     * 
     * @type InlineBody
     * @memberof DefaultApipetsPatch
     */
    inlineBody?: InlineBody
}

export class ObjectDefaultApi {
    private api: ObservableDefaultApi

    public constructor(configuration: Configuration, requestFactory?: DefaultApiRequestFactory, responseProcessor?: DefaultApiResponseProcessor) {
        this.api = new ObservableDefaultApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public filePost(param: DefaultApiFilePostRequest, options?: Configuration): Promise<void> {
        return this.api.filePost(param.inlineBody2,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public petsFilteredPatch(param: DefaultApiPetsFilteredPatchRequest, options?: Configuration): Promise<void> {
        return this.api.petsFilteredPatch(param.inlineBody1,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public petsPatch(param: DefaultApiPetsPatchRequest, options?: Configuration): Promise<void> {
        return this.api.petsPatch(param.inlineBody,  options).toPromise();
    }

}

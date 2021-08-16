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
import { ObservableDefaultApi } from './ObservableAPI';

import { DefaultApiRequestFactory, DefaultApiResponseProcessor} from "../apis/DefaultApi";
export class PromiseDefaultApi {
    private api: ObservableDefaultApi

    public constructor(
        configuration: Configuration,
        requestFactory?: DefaultApiRequestFactory,
        responseProcessor?: DefaultApiResponseProcessor
    ) {
        this.api = new ObservableDefaultApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param inlineBody2 
     */
    public filePost(inlineBody2?: InlineBody2, _options?: Configuration): Promise<void> {
        const result = this.api.filePost(inlineBody2, _options);
        return result.toPromise();
    }

    /**
     * @param inlineBody1 
     */
    public petsFilteredPatch(inlineBody1?: InlineBody1, _options?: Configuration): Promise<void> {
        const result = this.api.petsFilteredPatch(inlineBody1, _options);
        return result.toPromise();
    }

    /**
     * @param inlineBody 
     */
    public petsPatch(inlineBody?: InlineBody, _options?: Configuration): Promise<void> {
        const result = this.api.petsPatch(inlineBody, _options);
        return result.toPromise();
    }


}




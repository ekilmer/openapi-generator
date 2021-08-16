package org.openapitools.client.apis

import org.openapitools.client.infrastructure.CollectionFormats.*
import retrofit2.http.*
import okhttp3.RequestBody
import io.reactivex.rxjava3.core.Single;
import io.reactivex.rxjava3.core.Completable;

import org.openapitools.client.models.User

interface UserApi {
    /**
     * Create user
     * This can only be done by the logged in user.
     * Responses:
     *  - 0: successful operation
     * 
     * @param body Created user object 
<<<<<<< HEAD
    * @return [Call]<[Unit]>
=======
     * @return [Call]<[Unit]>
>>>>>>> master
     */
    @POST("user")
    fun createUser(@Body body: User): Completable

    /**
     * Creates list of users with given input array
     * 
     * Responses:
     *  - 0: successful operation
     * 
     * @param body List of user object 
<<<<<<< HEAD
    * @return [Call]<[Unit]>
=======
     * @return [Call]<[Unit]>
>>>>>>> master
     */
    @POST("user/createWithArray")
    fun createUsersWithArrayInput(@Body body: kotlin.collections.List<User>): Completable

    /**
     * Creates list of users with given input array
     * 
     * Responses:
     *  - 0: successful operation
     * 
     * @param body List of user object 
<<<<<<< HEAD
    * @return [Call]<[Unit]>
=======
     * @return [Call]<[Unit]>
>>>>>>> master
     */
    @POST("user/createWithList")
    fun createUsersWithListInput(@Body body: kotlin.collections.List<User>): Completable

    /**
     * Delete user
     * This can only be done by the logged in user.
     * Responses:
     *  - 400: Invalid username supplied
     *  - 404: User not found
     * 
     * @param username The name that needs to be deleted 
<<<<<<< HEAD
    * @return [Call]<[Unit]>
=======
     * @return [Call]<[Unit]>
>>>>>>> master
     */
    @DELETE("user/{username}")
    fun deleteUser(@Path("username") username: kotlin.String): Completable

    /**
     * Get user by user name
     * 
     * Responses:
     *  - 200: successful operation
     *  - 400: Invalid username supplied
     *  - 404: User not found
     * 
     * @param username The name that needs to be fetched. Use user1 for testing. 
<<<<<<< HEAD
    * @return [Call]<[User]>
=======
     * @return [Call]<[User]>
>>>>>>> master
     */
    @GET("user/{username}")
    fun getUserByName(@Path("username") username: kotlin.String): Single<User>

    /**
     * Logs user into the system
     * 
     * Responses:
     *  - 200: successful operation
     *  - 400: Invalid username/password supplied
     * 
     * @param username The user name for login 
     * @param password The password for login in clear text 
<<<<<<< HEAD
    * @return [Call]<[kotlin.String]>
=======
     * @return [Call]<[kotlin.String]>
>>>>>>> master
     */
    @GET("user/login")
    fun loginUser(@Query("username") username: kotlin.String, @Query("password") password: kotlin.String): Single<kotlin.String>

    /**
     * Logs out current logged in user session
     * 
     * Responses:
     *  - 0: successful operation
     * 
<<<<<<< HEAD
    * @return [Call]<[Unit]>
=======
     * @return [Call]<[Unit]>
>>>>>>> master
     */
    @GET("user/logout")
    fun logoutUser(): Completable

    /**
     * Updated user
     * This can only be done by the logged in user.
     * Responses:
     *  - 400: Invalid user supplied
     *  - 404: User not found
     * 
     * @param username name that need to be deleted 
     * @param body Updated user object 
<<<<<<< HEAD
    * @return [Call]<[Unit]>
=======
     * @return [Call]<[Unit]>
>>>>>>> master
     */
    @PUT("user/{username}")
    fun updateUser(@Path("username") username: kotlin.String, @Body body: User): Completable

}

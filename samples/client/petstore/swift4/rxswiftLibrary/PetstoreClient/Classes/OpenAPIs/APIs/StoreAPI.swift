//
// StoreAPI.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
import RxSwift

open class StoreAPI {
    /**
     Delete purchase order by ID
     
     - parameter orderId: (path) ID of the order that needs to be deleted 
     - returns: Observable<Void>
     */
    open class func deleteOrder(orderId: String) -> Observable<Void> {
        return Observable.create { observer -> Disposable in
            deleteOrderWithRequestBuilder(orderId: orderId).execute { (_, error) -> Void in
                if let error = error {
                    observer.onError(error)
                } else {
                    observer.onNext(())
                }
                observer.onCompleted()
            }
            return Disposables.create()
        }
    }

    /**
     Delete purchase order by ID
     - DELETE /store/order/{order_id}
     - For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
     - parameter orderId: (path) ID of the order that needs to be deleted 
     - returns: RequestBuilder<Void> 
     */
    open class func deleteOrderWithRequestBuilder(orderId: String) -> RequestBuilder<Void> {
        var path = "/store/order/{order_id}"
        let orderIdPreEscape = "\(APIHelper.mapValueToPathItem(orderId))"
        let orderIdPostEscape = orderIdPreEscape.addingPercentEncoding(withAllowedCharacters: .urlPathAllowed) ?? ""
        path = path.replacingOccurrences(of: "{order_id}", with: orderIdPostEscape, options: .literal, range: nil)
        let URLString = PetstoreClientAPI.basePath + path
        let parameters: [String: Any]? = nil

        let url = URLComponents(string: URLString)

        let requestBuilder: RequestBuilder<Void>.Type = PetstoreClientAPI.requestBuilderFactory.getNonDecodableBuilder()

        return requestBuilder.init(method: "DELETE", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**
     Returns pet inventories by status
     
     - returns: Observable<[String:Int]>
     */
    open class func getInventory() -> Observable<[String: Int]> {
        return Observable.create { observer -> Disposable in
            getInventoryWithRequestBuilder().execute { (response, error) -> Void in
                if let error = error {
                    observer.onError(error)
                } else if let response = response {
                    observer.onNext(response.body!)
                } else {
                    fatalError()
                }
                observer.onCompleted()
            }
            return Disposables.create()
        }
    }

    /**
     Returns pet inventories by status
     - GET /store/inventory
     - Returns a map of status codes to quantities
     - API Key:
       - type: apiKey api_key 
       - name: api_key
     - returns: RequestBuilder<[String:Int]> 
     */
    open class func getInventoryWithRequestBuilder() -> RequestBuilder<[String: Int]> {
        let path = "/store/inventory"
        let URLString = PetstoreClientAPI.basePath + path
        let parameters: [String: Any]? = nil

        let url = URLComponents(string: URLString)

        let requestBuilder: RequestBuilder<[String: Int]>.Type = PetstoreClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**
     Find purchase order by ID
     
     - parameter orderId: (path) ID of pet that needs to be fetched 
     - returns: Observable<Order>
     */
    open class func getOrderById(orderId: Int64) -> Observable<Order> {
        return Observable.create { observer -> Disposable in
            getOrderByIdWithRequestBuilder(orderId: orderId).execute { (response, error) -> Void in
                if let error = error {
                    observer.onError(error)
                } else if let response = response {
                    observer.onNext(response.body!)
                } else {
                    fatalError()
                }
                observer.onCompleted()
            }
            return Disposables.create()
        }
    }

    /**
     Find purchase order by ID
     - GET /store/order/{order_id}
     - For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
     - parameter orderId: (path) ID of pet that needs to be fetched 
     - returns: RequestBuilder<Order> 
     */
    open class func getOrderByIdWithRequestBuilder(orderId: Int64) -> RequestBuilder<Order> {
        var path = "/store/order/{order_id}"
        let orderIdPreEscape = "\(APIHelper.mapValueToPathItem(orderId))"
        let orderIdPostEscape = orderIdPreEscape.addingPercentEncoding(withAllowedCharacters: .urlPathAllowed) ?? ""
        path = path.replacingOccurrences(of: "{order_id}", with: orderIdPostEscape, options: .literal, range: nil)
        let URLString = PetstoreClientAPI.basePath + path
        let parameters: [String: Any]? = nil

        let url = URLComponents(string: URLString)

        let requestBuilder: RequestBuilder<Order>.Type = PetstoreClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**
     Place an order for a pet
     
     - parameter body: (body) order placed for purchasing the pet 
     - returns: Observable<Order>
     */
    open class func placeOrder(body: Order) -> Observable<Order> {
        return Observable.create { observer -> Disposable in
            placeOrderWithRequestBuilder(body: body).execute { (response, error) -> Void in
                if let error = error {
                    observer.onError(error)
                } else if let response = response {
                    observer.onNext(response.body!)
                } else {
                    fatalError()
                }
                observer.onCompleted()
            }
            return Disposables.create()
        }
    }

    /**
     Place an order for a pet
     - POST /store/order
     - parameter body: (body) order placed for purchasing the pet 
     - returns: RequestBuilder<Order> 
     */
    open class func placeOrderWithRequestBuilder(body: Order) -> RequestBuilder<Order> {
        let path = "/store/order"
        let URLString = PetstoreClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = URLComponents(string: URLString)

        let requestBuilder: RequestBuilder<Order>.Type = PetstoreClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

}

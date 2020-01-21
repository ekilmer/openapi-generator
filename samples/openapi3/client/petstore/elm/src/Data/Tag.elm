{-
   OpenAPI Petstore
   This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

   The version of the OpenAPI document: 1.0.0

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.Tag exposing (Tag, decoder, encode, encodeWithTag, toString)

import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


{-| A tag for a pet
-}
type alias Tag =
    { id : Maybe (Int)
    , name : Maybe (String)
    }


decoder : Decoder Tag
decoder =
    Decode.succeed Tag
        |> optional "id" (Decode.nullable Decode.int) Nothing
        |> optional "name" (Decode.nullable Decode.string) Nothing



encode : Tag -> Encode.Value
encode =
    Encode.object << encodePairs


encodeWithTag : ( String, String ) -> Tag -> Encode.Value
encodeWithTag (tagField, tag) model =
    Encode.object <| encodePairs model ++ [ ( tagField, Encode.string tag ) ]


encodePairs : Tag -> List (String, Encode.Value)
encodePairs model =
    [ ( "id", Maybe.withDefault Encode.null (Maybe.map Encode.int model.id) )
    , ( "name", Maybe.withDefault Encode.null (Maybe.map Encode.string model.name) )
    ]



toString : Tag -> String
toString =
    Encode.encode 0 << encode





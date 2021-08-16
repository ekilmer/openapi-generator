package org.openapitools.client.infrastructure

import com.squareup.moshi.Moshi
<<<<<<< HEAD
import com.squareup.moshi.adapters.Rfc3339DateJsonAdapter
=======
>>>>>>> master
import com.squareup.moshi.kotlin.reflect.KotlinJsonAdapterFactory
import java.util.Date

object Serializer {
    @JvmStatic
    val moshiBuilder: Moshi.Builder = Moshi.Builder()
<<<<<<< HEAD
        .add(Date::class.java, Rfc3339DateJsonAdapter().nullSafe())
=======
>>>>>>> master
        .add(OffsetDateTimeAdapter())
        .add(LocalDateTimeAdapter())
        .add(LocalDateAdapter())
        .add(UUIDAdapter())
        .add(ByteArrayAdapter())
        .add(KotlinJsonAdapterFactory())

    @JvmStatic
    val moshi: Moshi by lazy {
        moshiBuilder.build()
    }
}

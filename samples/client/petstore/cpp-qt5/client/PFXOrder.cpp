/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

#include "PFXOrder.h"

#include <QDebug>
#include <QJsonArray>
#include <QJsonDocument>
#include <QObject>

#include "PFXHelpers.h"

namespace test_namespace {

PFXOrder::PFXOrder(QString json) {
    this->initializeModel();
    this->fromJson(json);
}

PFXOrder::PFXOrder() {
    this->initializeModel();
}

PFXOrder::~PFXOrder() {}

void PFXOrder::initializeModel() {

    m_id_isSet = false;
    m_id_isValid = false;

    m_pet_id_isSet = false;
    m_pet_id_isValid = false;

    m_quantity_isSet = false;
    m_quantity_isValid = false;

    m_ship_date_isSet = false;
    m_ship_date_isValid = false;

    m_status_isSet = false;
    m_status_isValid = false;

    m_complete_isSet = false;
    m_complete_isValid = false;
}

void PFXOrder::fromJson(QString jsonString) {
    QByteArray array(jsonString.toStdString().c_str());
    QJsonDocument doc = QJsonDocument::fromJson(array);
    QJsonObject jsonObject = doc.object();
    this->fromJsonObject(jsonObject);
}

void PFXOrder::fromJsonObject(QJsonObject json) {

    m_id_isValid = ::test_namespace::fromJsonValue(id, json[QString("id")]);
    m_id_isSet = !json[QString("id")].isNull() && m_id_isValid;

    m_pet_id_isValid = ::test_namespace::fromJsonValue(pet_id, json[QString("petId")]);
    m_pet_id_isSet = !json[QString("petId")].isNull() && m_pet_id_isValid;

    m_quantity_isValid = ::test_namespace::fromJsonValue(quantity, json[QString("quantity")]);
    m_quantity_isSet = !json[QString("quantity")].isNull() && m_quantity_isValid;

    m_ship_date_isValid = ::test_namespace::fromJsonValue(ship_date, json[QString("shipDate")]);
    m_ship_date_isSet = !json[QString("shipDate")].isNull() && m_ship_date_isValid;

    m_status_isValid = ::test_namespace::fromJsonValue(status, json[QString("status")]);
    m_status_isSet = !json[QString("status")].isNull() && m_status_isValid;

    m_complete_isValid = ::test_namespace::fromJsonValue(complete, json[QString("complete")]);
    m_complete_isSet = !json[QString("complete")].isNull() && m_complete_isValid;
}

QString PFXOrder::asJson() const {
    QJsonObject obj = this->asJsonObject();
    QJsonDocument doc(obj);
    QByteArray bytes = doc.toJson();
    return QString(bytes);
}

QJsonObject PFXOrder::asJsonObject() const {
    QJsonObject obj;
    if (m_id_isSet) {
        obj.insert(QString("id"), ::test_namespace::toJsonValue(id));
    }
    if (m_pet_id_isSet) {
        obj.insert(QString("petId"), ::test_namespace::toJsonValue(pet_id));
    }
    if (m_quantity_isSet) {
        obj.insert(QString("quantity"), ::test_namespace::toJsonValue(quantity));
    }
    if (m_ship_date_isSet) {
        obj.insert(QString("shipDate"), ::test_namespace::toJsonValue(ship_date));
    }
    if (m_status_isSet) {
        obj.insert(QString("status"), ::test_namespace::toJsonValue(status));
    }
    if (m_complete_isSet) {
        obj.insert(QString("complete"), ::test_namespace::toJsonValue(complete));
    }
    return obj;
}

qint64 PFXOrder::getId() const {
    return id;
}
void PFXOrder::setId(const qint64 &id) {
    this->id = id;
    this->m_id_isSet = true;
}

qint64 PFXOrder::getPetId() const {
    return pet_id;
}
void PFXOrder::setPetId(const qint64 &pet_id) {
    this->pet_id = pet_id;
    this->m_pet_id_isSet = true;
}

qint32 PFXOrder::getQuantity() const {
    return quantity;
}
void PFXOrder::setQuantity(const qint32 &quantity) {
    this->quantity = quantity;
    this->m_quantity_isSet = true;
}

QDateTime PFXOrder::getShipDate() const {
    return ship_date;
}
void PFXOrder::setShipDate(const QDateTime &ship_date) {
    this->ship_date = ship_date;
    this->m_ship_date_isSet = true;
}

QString PFXOrder::getStatus() const {
    return status;
}
void PFXOrder::setStatus(const QString &status) {
    this->status = status;
    this->m_status_isSet = true;
}

bool PFXOrder::isComplete() const {
    return complete;
}
void PFXOrder::setComplete(const bool &complete) {
    this->complete = complete;
    this->m_complete_isSet = true;
}

bool PFXOrder::isSet() const {
    bool isObjectUpdated = false;
    do {
        if (m_id_isSet) {
            isObjectUpdated = true;
            break;
        }

        if (m_pet_id_isSet) {
            isObjectUpdated = true;
            break;
        }

        if (m_quantity_isSet) {
            isObjectUpdated = true;
            break;
        }

        if (m_ship_date_isSet) {
            isObjectUpdated = true;
            break;
        }

        if (m_status_isSet) {
            isObjectUpdated = true;
            break;
        }

        if (m_complete_isSet) {
            isObjectUpdated = true;
            break;
        }
    } while (false);
    return isObjectUpdated;
}

bool PFXOrder::isValid() const {
    // only required properties are required for the object to be considered valid
    return true;
}

} // namespace test_namespace

<?php
/**
 * ArrayTest
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 6.0.0-SNAPSHOT
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * ArrayTest Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class ArrayTest implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'ArrayTest';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'array_of_string' => 'string[]',
        'array_array_of_integer' => 'int[][]',
        'array_array_of_model' => '\OpenAPI\Client\Model\ReadOnlyFirst[][]'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'array_of_string' => null,
        'array_array_of_integer' => 'int64',
        'array_array_of_model' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'array_of_string' => 'array_of_string',
        'array_array_of_integer' => 'array_array_of_integer',
        'array_array_of_model' => 'array_array_of_model'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'array_of_string' => 'setArrayOfString',
        'array_array_of_integer' => 'setArrayArrayOfInteger',
        'array_array_of_model' => 'setArrayArrayOfModel'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'array_of_string' => 'getArrayOfString',
        'array_array_of_integer' => 'getArrayArrayOfInteger',
        'array_array_of_model' => 'getArrayArrayOfModel'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['array_of_string'] = $data['array_of_string'] ?? null;
        $this->container['array_array_of_integer'] = $data['array_array_of_integer'] ?? null;
        $this->container['array_array_of_model'] = $data['array_array_of_model'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if (!is_null($this->container['array_of_string']) && (count($this->container['array_of_string']) > 3)) {
            $invalidProperties[] = "invalid value for 'array_of_string', number of items must be less than or equal to 3.";
        }

        if (!is_null($this->container['array_of_string']) && (count($this->container['array_of_string']) < 0)) {
            $invalidProperties[] = "invalid value for 'array_of_string', number of items must be greater than or equal to 0.";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets array_of_string
     *
     * @return string[]|null
     */
    public function getArrayOfString()
    {
        return $this->container['array_of_string'];
    }

    /**
     * Sets array_of_string
     *
     * @param string[]|null $array_of_string array_of_string
     *
     * @return self
     */
    public function setArrayOfString($array_of_string)
    {

        if (!is_null($array_of_string) && (count($array_of_string) > 3)) {
            throw new \InvalidArgumentException('invalid value for $array_of_string when calling ArrayTest., number of items must be less than or equal to 3.');
        }
        if (!is_null($array_of_string) && (count($array_of_string) < 0)) {
            throw new \InvalidArgumentException('invalid length for $array_of_string when calling ArrayTest., number of items must be greater than or equal to 0.');
        }
        $this->container['array_of_string'] = $array_of_string;

        return $this;
    }

    /**
     * Gets array_array_of_integer
     *
     * @return int[][]|null
     */
    public function getArrayArrayOfInteger()
    {
        return $this->container['array_array_of_integer'];
    }

    /**
     * Sets array_array_of_integer
     *
     * @param int[][]|null $array_array_of_integer array_array_of_integer
     *
     * @return self
     */
    public function setArrayArrayOfInteger($array_array_of_integer)
    {
        $this->container['array_array_of_integer'] = $array_array_of_integer;

        return $this;
    }

    /**
     * Gets array_array_of_model
     *
     * @return \OpenAPI\Client\Model\ReadOnlyFirst[][]|null
     */
    public function getArrayArrayOfModel()
    {
        return $this->container['array_array_of_model'];
    }

    /**
     * Sets array_array_of_model
     *
     * @param \OpenAPI\Client\Model\ReadOnlyFirst[][]|null $array_array_of_model array_array_of_model
     *
     * @return self
     */
    public function setArrayArrayOfModel($array_array_of_model)
    {
        $this->container['array_array_of_model'] = $array_array_of_model;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}



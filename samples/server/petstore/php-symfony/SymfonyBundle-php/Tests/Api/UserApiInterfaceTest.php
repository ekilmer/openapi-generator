<?php
/**
 * UserApiInterfaceTest
 * PHP version 7.1.3
 *
 * @category Class
 * @package  OpenAPI\Server\Tests\Api
 * @author   openapi-generator contributors
 * @link     https://github.com/openapitools/openapi-generator
 */

/**
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 * Generated by: https://github.com/openapitools/openapi-generator.git
 *
 */

/**
 * NOTE: This class is auto generated by the openapi generator program.
 * https://github.com/openapitools/openapi-generator
 * Please update the test case below to test the endpoint.
 */

namespace OpenAPI\Server\Tests\Api;

use OpenAPI\Server\Configuration;
use OpenAPI\Server\ApiClient;
use OpenAPI\Server\ApiException;
use OpenAPI\Server\ObjectSerializer;
use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

/**
 * UserApiInterfaceTest Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Server\Tests\Api
 * @author   openapi-generator contributors
 * @link     https://github.com/openapitools/openapi-generator
 */
class UserApiInterfaceTest extends WebTestCase
{

    /**
     * Setup before running any test cases
     */
    public static function setUpBeforeClass(): void
    {
    }

    /**
     * Setup before running each test case
     */
    public function setUp(): void
    {
    }

    /**
     * Clean up after running each test case
     */
    public function tearDown(): void
    {
    }

    /**
     * Clean up after running all test cases
     */
    public static function tearDownAfterClass(): void
    {
    }

    /**
     * Test case for createUser
     *
     * Create user.
     *
     */
    public function testCreateUser()
    {
        $client = static::createClient();

        $path = '/user';

        $crawler = $client->request('POST', $path, [], [], ['CONTENT_TYPE' => 'application/json']);
    }

    /**
     * Test case for createUsersWithArrayInput
     *
     * Creates list of users with given input array.
     *
     */
    public function testCreateUsersWithArrayInput()
    {
        $client = static::createClient();

        $path = '/user/createWithArray';

        $crawler = $client->request('POST', $path, [], [], ['CONTENT_TYPE' => 'application/json']);
    }

    /**
     * Test case for createUsersWithListInput
     *
     * Creates list of users with given input array.
     *
     */
    public function testCreateUsersWithListInput()
    {
        $client = static::createClient();

        $path = '/user/createWithList';

        $crawler = $client->request('POST', $path, [], [], ['CONTENT_TYPE' => 'application/json']);
    }

    /**
     * Test case for deleteUser
     *
     * Delete user.
     *
     */
    public function testDeleteUser()
    {
        $client = static::createClient();

        $path = '/user/{username}';
        $pattern = '{username}';
        $data = $this->genTestData('[a-z0-9]+');
        $path = str_replace($pattern, $data, $path);

        $crawler = $client->request('DELETE', $path);
    }

    /**
     * Test case for getUserByName
     *
     * Get user by user name.
     *
     */
    public function testGetUserByName()
    {
        $client = static::createClient();

        $path = '/user/{username}';
        $pattern = '{username}';
        $data = $this->genTestData('[a-z0-9]+');
        $path = str_replace($pattern, $data, $path);

        $crawler = $client->request('GET', $path);
    }

    /**
     * Test case for loginUser
     *
     * Logs user into the system.
     *
     */
    public function testLoginUser()
    {
        $client = static::createClient();

        $path = '/user/login';

        $crawler = $client->request('GET', $path);
    }

    /**
     * Test case for logoutUser
     *
     * Logs out current logged in user session.
     *
     */
    public function testLogoutUser()
    {
        $client = static::createClient();

        $path = '/user/logout';

        $crawler = $client->request('GET', $path);
    }

    /**
     * Test case for updateUser
     *
     * Updated user.
     *
     */
    public function testUpdateUser()
    {
        $client = static::createClient();

        $path = '/user/{username}';
        $pattern = '{username}';
        $data = $this->genTestData('[a-z0-9]+');
        $path = str_replace($pattern, $data, $path);

        $crawler = $client->request('PUT', $path, [], [], ['CONTENT_TYPE' => 'application/json']);
    }

    protected function genTestData($regexp)
    {
        $grammar  = new \Hoa\File\Read('hoa://Library/Regex/Grammar.pp');
        $compiler = \Hoa\Compiler\Llk\Llk::load($grammar);
        $ast      = $compiler->parse($regexp);
        $generator = new \Hoa\Regex\Visitor\Isotropic(new \Hoa\Math\Sampler\Random());

        return $generator->visit($ast); 
    }
}

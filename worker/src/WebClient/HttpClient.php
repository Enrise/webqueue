<?php

namespace Enrise\WebQueue\WebClient;

use Enrise\WebQueue\Message;
use Enrise\WebQueue\WebClient;
use GuzzleHttp\Client;

class HttpClient implements WebClient
{
    /**
     * @var Client
     */
    private $httpClient;
    /**
     * @var string
     */
    private $endpointUrl;

    /**
     * @param Client $httpClient
     * @param string $endpointUrl
     */
    public function __construct(Client $httpClient, $endpointUrl)
    {
        $this->httpClient = $httpClient;
        $this->endpointUrl = $endpointUrl;
    }

    public function dispatch(Message $message)
    {
        $this->httpClient->post(
            $this->endpointUrl,
            [
                'body' => $message->getPayload()
            ]
        );
    }
}

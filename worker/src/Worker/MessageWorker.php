<?php

namespace Enrise\WebQueue\Worker;

use Enrise\WebQueue\QueueDriver;
use Enrise\WebQueue\WebClient;
use Enrise\WebQueue\Worker;
use GuzzleHttp\Client as HttpClient;
use Monolog\Logger;

/**
 * @todo log events
 */
class MessageWorker implements Worker
{
    /**
     * @var QueueDriver
     */
    private $driver;

    /**
     * @var bool
     */
    private $stopping = false;

    /**
     * @var WebClient
     */
    private $webClient;
    /**
     * @var Logger
     */
    private $logger;

    public function __construct(QueueDriver $driver, WebClient $webClient, Logger $logger)
    {
        $this->driver = $driver;
        $this->webClient = $webClient;
        $this->logger = $logger;
    }

    public function start()
    {
        $this->logger->addInfo('Starting new message worker');

        $this->stopping = false;
        while (! $this->stopping) {
            $message = $this->driver->get();

            if ($message === null) {
                usleep(10000);
                continue;
            }

            try {
                // Handle message by doing HTTP-POST
                $this->webClient->dispatch($message);

                // ACK the message
                $this->driver->acknowledge($message);
                $this->logger->addDebug(sprintf('Acknowledged message %s', $message->getUniqueId()));
            } catch (\Exception $exception) {
                $this->driver->reject($message);
                $this->logger->addNotice(sprintf('Rejected message %s', $message->getUniqueId()));
            }
            $this->logger->addDebug('Waiting for new messages');
        }
    }

    public function stop()
    {
        $this->logger->addInfo('Shutting down message worker');
        $this->stopping = true;
    }
}

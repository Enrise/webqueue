<?php

namespace Enrise\WebQueue\QueueDriver;

use Enrise\WebQueue\QueueDriver;
use Enrise\WebQueue\Message;
use PhpAmqpLib\Channel\AMQPChannel;
use PhpAmqpLib\Connection\AMQPStreamConnection;

class AmqpQueueDriver implements QueueDriver
{
    /**
     * @var AMQPStreamConnection
     */
    private $connection;

    /**
     * @var AMQPChannel
     */
    private $channel;
    /**
     * @var string
     */
    private $queueName;

    protected function __construct(AMQPStreamConnection $amqpConnection, $queueName)
    {
        $this->connection = $amqpConnection;
        $this->queueName = $queueName;
    }

    public static function create($hostname, $queueName)
    {
        $connection = new AMQPStreamConnection($hostname, '5672', 'guest', 'guest');
        return new self($connection, $queueName);
    }

    /**
     * @inheritDoc
     */
    public function get()
    {
        if ($this->channel === null) {
            $this->channel = $this->connection->channel();
        }

        $message = $this->channel->basic_get($this->queueName);

        if ($message === null) {
            return null;
        }

        return new QueueDriver\Message\AmqpMessage($message);
    }

    /**
     * @inheritDoc
     */
    public function acknowledge(Message $message)
    {
        $this->channel->basic_ack($message->getOriginalMessage()->get('delivery_tag'));
    }

    /**
     * @inheritDoc
     */
    public function reject(Message $message)
    {
        if (! $message instanceof QueueDriver\Message\AmqpMessage) {
            throw new \InvalidArgumentException('$message must be of type AmqpMessage!');
        }

        /** @var QueueDriver\Message\AmqpMessage $message */
        $originalMessage = $message->getOriginalMessage();
        $this->channel->basic_reject($originalMessage->get('delivery_tag'), true);
    }
}

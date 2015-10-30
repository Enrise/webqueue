<?php

namespace Enrise\WebQueue\QueueDriver\Message;

use Enrise\WebQueue\Message;
use PhpAmqpLib\Message\AMQPMessage as OriginalAMQPMessage;

class AmqpMessage implements Message
{
    /**
     * @var OriginalAMQPMessage
     */
    private $message;

    public function __construct(OriginalAMQPMessage $message)
    {
        $this->message = $message;
    }

    public function getPayload()
    {
        return $this->message->body;
    }

    public function getOriginalMessage()
    {
        return $this->message;
    }

    public function getUniqueId()
    {
        return $this->message->get('delivery_tag');
    }
}

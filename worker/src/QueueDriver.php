<?php

namespace Enrise\WebQueue;

interface QueueDriver
{
    /**
     * Get a message from the message queue. This method must be blocking by design.
     *
     * @return Message|null
     */
    public function get();

    public function reject(Message $message);

    public function acknowledge(Message $message);

    public static function create($hostname, $queueName);
}

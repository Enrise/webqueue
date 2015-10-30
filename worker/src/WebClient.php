<?php

namespace Enrise\WebQueue;

interface WebClient
{
    public function dispatch(Message $message);
}

<?php

namespace Enrise\WebQueue;

interface Message
{
    public function getPayload();

    public function getOriginalMessage();

    public function getUniqueId();
}

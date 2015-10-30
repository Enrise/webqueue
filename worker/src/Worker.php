<?php

namespace Enrise\WebQueue;

interface Worker
{
    /**
     * Start a worker
     */
    public function start();

    /**
     * Stop a worker
     */
    public function stop();
}

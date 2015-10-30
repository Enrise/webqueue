<?php

namespace Enrise\WebQueue\QueueDriver;

use Enrise\WebQueue\QueueDriver;

class QueueDriverFactory
{
    /**
     * @param $driverName
     * @param $hostname
     * @param $queueName
     * @return QueueDriver
     */
    public static function createQueueDriver($driverName, $hostname, $queueName)
    {
        $driverFQCN = sprintf('%s%s%sQueueDriver', __NAMESPACE__, '\\', $driverName);
        if (! class_exists($driverFQCN)) {
            throw new \InvalidArgumentException(sprintf('No such driver exists: %s', $driverFQCN));
        }

        $driver = $driverFQCN::create($hostname, $queueName);

        return $driver;
    }
}

<?php

use Enrise\WebQueue\QueueDriver\QueueDriverFactory;
use Enrise\WebQueue\WebClient\HttpClient;
use Enrise\WebQueue\Worker\MessageWorker;
use GuzzleHttp\Client;
use Symfony\Component\Yaml\Yaml;

require_once __DIR__ . '/../vendor/autoload.php';

$logger = new \Monolog\Logger('webqueue');
$logger->pushHandler(new \Monolog\Handler\StreamHandler('php://output'));

$configurationPath = '/etc/webqueue/worker.yml';
$globalConfiguration = [];
if (! is_readable($configurationPath)) {
    $logger->addError(sprintf('Could not read configuration file: %s', $configurationPath));
    exit(2);
}

$workerConfig = file_get_contents($configurationPath);
$globalConfiguration = Yaml::parse($workerConfig);
$configuration = $globalConfiguration['worker'];

$arguments = ['driver', 'hostname', 'queue', 'endpoint'];
foreach ($arguments as $argument) {
    if (! isset($configuration[$argument])) {
        $logger->addError(sprintf('Invalid configuation: "%s" not set.', $argument));
        exit(2);
    }
}

$driver = QueueDriverFactory::createQueueDriver($configuration['driver'], $configuration['hostname'], $configuration['queue']);
$webClient = new HttpClient(new Client(), $configuration['endpoint']);

$worker = new MessageWorker($driver, $webClient, $logger);

pcntl_signal(SIGTERM, [$worker, 'stop']);
pcntl_signal(SIGINT, [$worker, 'stop']);
pcntl_signal(SIGHUP, [$worker, 'stop']);

$worker->start();

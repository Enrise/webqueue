<?php

use Enrise\WebQueue\QueueDriver\QueueDriverFactory;
use Enrise\WebQueue\WebClient\HttpClient;
use Enrise\WebQueue\Worker\MessageWorker;
use GuzzleHttp\Client;
use Symfony\Component\Yaml\Yaml;

require_once __DIR__ . '/../vendor/autoload.php';

$arguments = [
    'queue',
    'hostname',
    'driver',
    'endpoint',
    'username',
    'password',
    'port',
];
$arguments2 = $arguments;
array_walk($arguments2, function(&$value) {
    $value .= ':';
});

$optionValues = getopt('', $arguments2);

$logger = new \Monolog\Logger('webqueue');
$logger->pushHandler(new \Monolog\Handler\StreamHandler('php://output'));

$configurationPath = '/etc/webqueue/worker.yml';
$globalConfiguration = [];
if (is_readable($configurationPath)) {
    $logger->addDebug('Loading system-wide configuration');
    $workerConfig = file_get_contents($configurationPath);
    $globalConfiguration = Yaml::parse($workerConfig);
}

$configuration = array_merge_recursive($globalConfiguration, ['worker' => $optionValues]);
$configuration = $configuration['worker'];

// TODO: add configuration validation

$driver = QueueDriverFactory::createQueueDriver($configuration['driver'], $configuration['hostname'], $configuration['queue']);
$webClient = new HttpClient(new Client(), $configuration['endpoint']);

$worker = new MessageWorker($driver, $webClient, $logger);

pcntl_signal(SIGTERM, [$worker, 'stop']);
pcntl_signal(SIGINT, [$worker, 'stop']);
pcntl_signal(SIGHUP, [$worker, 'stop']);

$worker->start();

angular.module('webqueue.dashboard.queue-graph', ['ngResource', 'nvd3'])
    .directive('webqueueQueueGraph', function () {
        return {
            restrict: 'E',
            templateUrl: 'template/dashboard/queue-graph.html',
            controller: 'queueGraphCtrl'
        }
    })
    .controller('queueGraphCtrl', ['$scope', '$resource', function ($scope, $resource) {
        var colors = ['#C8C9C7', '#75787B'];
        $scope.options = {
            chart: {
                color: function (data, index) {

                    return colors[index];
                },
                type: 'stackedAreaChart',
                x: function (d) {
                    return d.timestamp;
                },
                y: function (d, i) {
                    return d.sample;
                },
                showValues: true,
                transitionDuration: 100,
                xAxis: {
                    showMaxMin: false,
                    tickFormat: function (d) {
                        return d3.time.format('%X')(new Date(d))
                    }
                },
                yAxis: {
                    axisLabel: 'Y Axis',
                    axisLabelDistance: 30
                }
            },
            caption: {
                enable: true,
                html: 'Messages in the last 10 minutes'
            }
        };

        var RabbitInfo = $resource('http://127.0.0.1:7809/api/queue-info', {}, {get: {isArray: false}});

        var fetchRabbitMqStatistics = function () {
            RabbitInfo.get().$promise.then(function (info) {
                $scope.message_data = [
                    {
                        key: "Processing",
                        values: info.messages_unacknowledged_details.samples
                    },
                    {
                        key: "Waiting",
                        values: info.messages_ready_details.samples
                    }
                ];
            });
        };

        fetchRabbitMqStatistics();

        setInterval(function () {
            fetchRabbitMqStatistics();
        }, 2000);

    }])
;

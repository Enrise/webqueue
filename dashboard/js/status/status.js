'use strict';

angular.module('webqueue.status', ['ngRoute', 'ngResource'])

    .config(['$routeProvider', function ($routeProvider) {
    }])
    .directive('webqueueStatus', function () {
        return {
            restrict: 'E',
            templateUrl: 'template/status/status.html',
            controller: 'statusCtrl'
        }
    })

    .controller('statusCtrl', ['$scope', '$resource', function ($scope, $resource) {
        $scope.allgood = true;

        var Status = $resource('/api/status', {}, {get: {isArray: true}});
        var getStatus = function() {
            Status.get().$promise.then(function (status) {
                var allgood = true;
                angular.forEach(status, function (stat) {
                    if (!allgood) {
                        return;
                    }
                    allgood = stat.Healthy;
                });
                $scope.allgood = allgood;

                $scope.status = status;
            });
        }

        setInterval(function() { getStatus(); }, 1000);
    }])
;

angular.module('webqueue.dashboard.latest-jobs', ['ngResource'])
    .directive('webqueueLatestJobs', function () {
        return {
            restrict: 'E',
            templateUrl: 'template/dashboard/latest-jobs.html',
            controller: 'latestJobsCtrl'
        }
    })
    .controller('latestJobsCtrl', ['$scope', '$resource', function ($scope, $resource) {
        var LatestJobs = $resource('/api/latest-jobs', {}, {get: {isArray: true}});

        var fetchLatestJobs = function () {
            LatestJobs.get().$promise.then(function (jobs) {
                $scope.jobs = jobs;
            })
        };

        fetchLatestJobs();

        setInterval(function () {
            fetchLatestJobs();
        }, 1000);
    }])
;

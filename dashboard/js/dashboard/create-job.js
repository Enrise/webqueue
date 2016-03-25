angular.module('webqueue.dashboard.createjob', ['ngResource'])
    .directive('webqueueCreateJob', function () {
        return {
            restrict: 'E',
            templateUrl: 'template/dashboard/create-job.html',
            controller: 'createJobCtrl'
        }
    })
    .controller('createJobCtrl', ['$scope', '$resource', function ($scope, $resource) {
        $scope.showForm = false;
        $scope.jobContent = "";
        $scope.errorMessage = null;
        $scope.successMessage = null;
        $scope.errorCountdown = 5;
        $scope.successCountdown = 5;

        var Job = $resource('/api/create-job', {}, {post: {}});

        $scope.openForm = function () {
            $scope.showForm = !$scope.showForm;
        };

        $scope.createJob = function () {
            var job = new Job();
            job.payload = $scope.jobContent;
            job.$save().then(function (data) {
                $scope.jobContent = "";
                $scope.successMessage = "Job created!";
                $scope.successCountdown = 5;
                setTimeout(function () {
                    $scope.successMessage = null;
                }, 5000);

                var successInterval = setInterval(function () {
                    $scope.successCountdown -= 1;
                    if ($scope.successCountdown == 0) {
                        clearInterval(successInterval);
                    }
                }, 1000);
            }, function (data) {
                $scope.errorMessage = data.statusText;
                $scope.errorCountdown = 5;
                setTimeout(function () {
                    $scope.errorMessage = null;
                }, 5000);

                var errorInterval = setInterval(function () {
                    $scope.errorCountdown -= 1;
                    if ($scope.errorCountdown == 0) {
                        clearInterval(errorInterval);
                    }
                }, 1000);
            });
        };
    }])
;

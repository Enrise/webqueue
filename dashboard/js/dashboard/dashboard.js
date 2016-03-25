angular.module('webqueue.dashboard', [
    'ngRoute',
    'webqueue.dashboard.latest-jobs',
    'webqueue.dashboard.queue-graph',
    'webqueue.dashboard.createjob'
])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/dashboard', {
            templateUrl: '/template/dashboard.html',
            controller: 'dashboardCtrl'
        })
    }])
    .controller('dashboardCtrl', [function() {

    }])
;

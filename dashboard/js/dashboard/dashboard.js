angular.module('webqueue.dashboard', [
    'ngRoute',
    'webqueue.dashboard.latestmessages',
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

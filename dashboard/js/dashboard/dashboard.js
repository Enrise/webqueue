angular.module('webqueue.dashboard', ['ngRoute', 'webqueue.dashboard.latestmessages'])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/dashboard', {
            templateUrl: '/template/dashboard.html',
            controller: 'dashboardCtrl'
        })
    }])
    .controller('dashboardCtrl', [function() {

    }])
;

'use strict';

angular.module('webqueue', [
    'ngRoute',
    'webqueue.status',
    'webqueue.dashboard'
])
    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.otherwise({redirectTo: '/dashboard'})
    }])
;

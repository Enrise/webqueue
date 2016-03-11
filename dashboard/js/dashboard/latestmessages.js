angular.module('webqueue.dashboard.latestmessages', ['ngResource'])
    .directive('webqueueLatestMessages', function() {
        return {
            restrict: 'E',
            templateUrl: 'template/dashboard/latest-messages.html',
            controller: 'latestMessageCtrl'
        }
    })
    .controller('latestMessageCtrl', ['$scope', '$resource',function($scope, $resource) {
        var LatestMessages = $resource('/api/latest-messages', {}, {get: {isArray:true}});

        LatestMessages.get().$promise.then(function(messages) {
            $scope.messages = messages;
        });
    }])
;

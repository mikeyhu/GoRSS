var app = angular.module('gorssApp',[]);

app.controller('feedController', function($scope, $http) {

    $scope.feeds = [];

    $http.get('../feeds/').
        success(function(data) {
            $scope.feeds = data;
        });

    $scope.addFeed = function() {
      feed = {}
      feed.Url = $scope.feedUrl;
      feed.Tags = [$scope.feedTags];

      $http.post('../feeds/',feed).
        success(function() {
          $scope.feeds.push(feed);
          $scope.feedUrl = '';
          $scope.feedTags = '';
        });
    };
  });

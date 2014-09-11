app.controller('feedController', function($scope, $http) {

    $scope.feedCount = function() {return $scope.feeds.length;};
    $scope.feeds = [];

    $http.get('../feeds/').
      success(function(data) {
          $scope.feeds = data;
      });

    $scope.addFeed = function() {
      var feed = {};
      feed.Url = $scope.feedUrl;
      feed.Tags = $scope.feedTags.split(" ");

      $http.post('../feeds/',feed).
        success(function() {
          $scope.feeds.push(feed);
          $scope.feedUrl = '';
          $scope.feedTags = '';
        });
    };

    $scope.removeFeed = function(id) {
      var feed = $scope.feeds[id];
      var mongoId = feed.id;
      $scope.feeds.splice(id,1);

    }
  });

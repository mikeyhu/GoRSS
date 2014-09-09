app.controller('storyController', function($scope, $http) {

  $scope.stories = [];

  $http.get('../stories/latest').
    success(function(data) { $scope.stories = data; });

});

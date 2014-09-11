app.controller('storyController', function ($scope, $http) {
  "use strict";

  $scope.stories = [];

  $http.get('../stories/latest').
    success(function (data) {
      $scope.stories = data;
    });
});

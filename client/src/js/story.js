app.controller('storyController', function ($scope, $http) {
  "use strict";

  $scope.stories = [];
  $scope.tags = ["select a tag"];

  $http.get('../stories/latest').
    success(function (data) {
      $scope.stories = data;
    });

  $http.get('../stories/tags').
    success(function (data) {
      data.unshift("select a tag");
      $scope.tags = data;
    });
});

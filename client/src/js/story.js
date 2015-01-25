var STATE_FAVORITE = "favorite";
var STATE_NEW = "new";

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

  $scope.toggleFavorite = function (index) {
    var story = $scope.stories[index];
    story.State = (story.State === STATE_FAVORITE) ? STATE_NEW : STATE_FAVORITE;
  };

  $scope.isFavorite = function (index) {
    var story = $scope.stories[index];
    return (story.State === STATE_FAVORITE);
  };

  $scope.starType = function (favorite) {
    return favorite ? "glyphicon-star" : "glyphicon-star-empty";
  };
});

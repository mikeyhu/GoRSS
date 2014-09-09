var app = angular.module('gorssApp',['ngRoute']);

app.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/feeds', {
        templateUrl: 'partials/feeds.html',
        controller: 'feedController'
      }).
      when('/stories', {
        templateUrl: 'partials/stories.html',
        controller: 'storyController'
      }).
      otherwise({
        redirectTo: '/stories'
      });
  }]);

describe("Stories can be displayed",function() {

  beforeEach(module("gorssApp"));

  var ctrl, scope, httpBackend;

  beforeEach(inject(function($controller, $rootScope, $httpBackend){
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../stories/latest").respond([
      {Title:"A story", Link:"A link", Date: Date.now()}
      ]);

    ctrl = $controller("storyController", {
      $scope : scope
    });
  }));



  it("should be able to display feeds", function(){

    httpBackend.flush();
    expect(scope.stories.length).toBe(1);
  });

});

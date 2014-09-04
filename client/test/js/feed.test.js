describe("Testing the feed",function() {

  beforeEach(module("gorssApp"));

  var ctrl, scope;

  beforeEach(inject(function($controller, $rootScope){
    scope = $rootScope;
    ctrl = $controller("feedController", {
      $scope : scope
    });
  }));

  it("should say hello", function(){
    expect(scope.feedCount()).toBe(0);
  });
});

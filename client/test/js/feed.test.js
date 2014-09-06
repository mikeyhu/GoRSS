describe("Feeds can be displayed",function() {

  beforeEach(module("gorssApp"));

  var ctrl, scope, httpBackend;

  beforeEach(inject(function($controller, $rootScope, $httpBackend){
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../feeds/").respond([
      {Url:"A", Tags:["tag"]},
      {Url:"B", Tags:["tag"]}
      ]);

    ctrl = $controller("feedController", {
      $scope : scope
    });
  }));


  it("should be able to display feeds", function(){

    httpBackend.flush();
    expect(scope.feedCount()).toBe(2);
  });

  it("should be able to post content", function(){

    httpBackend.expectPOST("../feeds/").respond();

    scope.feedUrl = "http://mike.was.here";
    scope.feedTags = "tag";
    scope.addFeed();

    httpBackend.flush();

    expect(scope.feedUrl).toBe("");
    expect(scope.feedTags).toBe("");
    expect(scope.feedCount()).toBe(3);
    expect(scope.feeds[2].Url).toBe("http://mike.was.here")
  });
});

describe("Feeds can have tags",function() {

  beforeEach(module("gorssApp"));

  var ctrl, scope, httpBackend;

  beforeEach(inject(function($controller, $rootScope, $httpBackend){
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../feeds/").respond([]);

    ctrl = $controller("feedController", {
      $scope : scope
    });
  }));

  it("should be able to save multiple tags", function(){

    httpBackend.expectPOST("../feeds/").respond();

    scope.feedUrl = "http://mike.was.here";
    scope.feedTags = "tag1 tag2";
    scope.addFeed();

    httpBackend.flush();

    expect(scope.feeds[0].Tags).toEqual(["tag1","tag2"])
  });
});

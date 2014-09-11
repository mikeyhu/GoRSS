describe("Feeds can be displayed", function () {
  "use strict";

  beforeEach(module("gorssApp"));

  var scope, httpBackend;

  beforeEach(inject(function ($controller, $rootScope, $httpBackend) {
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../feeds/").respond([
      {Url: "A", Tags: ["tag"], id: "aaa"},
      {Url: "B", Tags: ["tag"], id: "bbb"}
    ]);

    $controller("feedController", {
      $scope : scope
    });
  }));

  it("should be able to display feeds", function () {
    httpBackend.flush();
    expect(scope.feedCount()).toBe(2);
  });

  it("should be able to delete feeds", function () {
    httpBackend.flush();
    scope.removeFeed("aaa");
    expect(scope.feedCount()).toBe(1);
  });

  it("should be able to post content", function () {
    httpBackend.expectPOST("../feeds/").respond();

    scope.feedUrl = "http://mike.was.here";
    scope.feedTags = "tag";
    scope.addFeed();

    httpBackend.flush();

    expect(scope.feedUrl).toBe("");
    expect(scope.feedTags).toBe("");
    expect(scope.feedCount()).toBe(3);
    expect(scope.feeds[2].Url).toBe("http://mike.was.here");
  });
});

describe("Feeds can have tags", function () {
  "use strict";

  beforeEach(module("gorssApp"));

  var scope, httpBackend;

  beforeEach(inject(function ($controller, $rootScope, $httpBackend) {
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../feeds/").respond([]);

    $controller("feedController", {
      $scope : scope
    });
  }));

  it("should be able to save multiple tags", function () {

    httpBackend.expectPOST("../feeds/").respond();

    scope.feedUrl = "http://mike.was.here";
    scope.feedTags = "tag1 tag2";
    scope.addFeed();

    httpBackend.flush();

    expect(scope.feeds[0].Tags).toEqual(["tag1", "tag2"]);
  });
});

describe("Stories can be displayed", function () {
  "use strict";

  beforeEach(module("gorssApp"));

  var scope, httpBackend;

  beforeEach(inject(function ($controller, $rootScope, $httpBackend) {
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../stories/latest").respond([
      {Title: "A story", Link: "A link", Date: Date.now()}
    ]);

    httpBackend.whenGET("../stories/tags").respond(["news","technology"]);

    $controller("storyController", {
      $scope : scope
    });
  }));

  it("should be able to display feeds", function () {
    httpBackend.flush();
    expect(scope.stories.length).toBe(1);
  });

  it("should be able to display feeds", function () {
    httpBackend.flush();
    expect(scope.tags.length).toBe(3);
    expect(scope.tags[0]).toBe("select a tag");
    expect(scope.tags[1]).toBe("news");
    expect(scope.tags[2]).toBe("technology");
  });

});

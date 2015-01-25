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

    httpBackend.whenGET("../stories/tags").respond(["news", "technology"]);

    $controller("storyController", {
      $scope: scope
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

describe("Stories with state", function () {
  "use strict";

  beforeEach(module("gorssApp"));

  var scope, httpBackend;

  beforeEach(inject(function ($controller, $rootScope, $httpBackend) {
    scope = $rootScope;
    httpBackend = $httpBackend;

    httpBackend.whenGET("../stories/latest").respond([
      {Title: "A story", Link: "A link", Date: Date.now(), State: "new"},
      {Title: "Another story", Link: "A link", Date: Date.now(), State: "favorite"}
    ]);

    httpBackend.whenGET("../stories/tags").respond(["news", "technology"]);

    $controller("storyController", {
      $scope: scope
    });
  }));


  it("should mark new stories as not favorited", function () {
    httpBackend.flush();
    expect(scope.stories.length).toBe(2);
    expect(scope.isFavorite(0)).toBe(false);
    expect(scope.isFavorite(1)).toBe(true);
  });

  it("should mark favorite stories as favorited", function () {
    httpBackend.flush();
    expect(scope.isFavorite(1)).toBe(true);
  });

  it("should be able to toggle on a new story", function () {
    httpBackend.flush();
    scope.toggleFavorite(0);
    expect(scope.isFavorite(0)).toBe(true);
  });

  it("should be able to toggle off a favorite story", function () {
    httpBackend.flush();
    scope.toggleFavorite(1);
    expect(scope.isFavorite(1)).toBe(false);
  });

  it("should show new stories with empty star", function () {
    httpBackend.flush();
    expect(scope.starType(false)).toBe("glyphicon-star-empty");
  });

  it("should show favorite stories with filled star", function () {
    httpBackend.flush();
    expect(scope.starType(true)).toBe("glyphicon-star");
  });
});

'use strict';

angular.module('userApp')
  .controller('MainCtrl', function($scope, $http) {
    $scope.resp = [];
    $scope.user = null;
    $http.get('collector/'+$scope.userNum).success(function(data) {
      $scope.user = data;
    })
  	$scope.$on( 'login', function(event, name) {
  		$scope.user = name;
  	})
  	$scope.$on( 'logout', function(event) {
  		$scope.user = null;
  	})
  })
  .controller('RestCtrl', function($scope, $http) {
    $scope.userNum = null;
  })
  .controller('LoginCtrl', function($scope) {
  	$scope.name = '';
  	$scope.login = function() {
  		$scope.$emit( 'login', $scope.name );
  	}
  	$scope.logout = function() {
  		$scope.$emit( 'logout', null );
  	}
  })

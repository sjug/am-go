'use strict';

angular.module('userApp')
  .controller('MainCtrl', function($scope, $http) {
    $scope.resp = [];
    $scope.user = null;
    $scope.update = function() {
      $http.get('collector/'+$scope.user).success(function(data) {
        $scope.resp = data;
      })
  	}
  	$scope.$on( 'login', function(event, name) {
  		$scope.user = name;
  	})
  	$scope.$on( 'logout', function(event) {
  		$scope.user = null;
  	})
    $scope.$watch( 'user', $scope.update);
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

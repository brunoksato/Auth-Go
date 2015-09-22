(function() {
  'use strict';

  angular
    .module('app.users.service', [])
    .factory('UsersService', UsersService);

  function UsersService($http, $q, API) {

    function get(){

      var defer = $q.defer();

      $http.post(API.URL + 'api/users').success(function(result){
        defer.resolve(result.status.Value);
      }).error(function(error){
        defer.reject(error);
      });

      return defer.promise;
    }

    return{
      get: get
    }

  }

})();

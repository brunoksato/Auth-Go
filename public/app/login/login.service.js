(function() {
  'use strict';

  angular
    .module('app.login.service', [])
    .factory('LoginService', LoginService);

  function LoginService($http, $q, API) {

    function login(model){

      var defer = $q.defer();

      $http.post(API.URL + 'auth/login', model).success(function(result){
        defer.resolve(result);
      }).error(function(error){
        defer.reject(error);
      });

      return defer.promise;
    }

    return{
      login: login
    }

  }

  LoginService.$inject = ['$http', '$q', 'API']

})();

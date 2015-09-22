(function() {
  'use strict';

  angular
    .module('app.signup.service', [])
    .factory('SignupService', SignupService);

  /** @ngInject */
  function SignupService($http, $q, API) {

    function register(model){

      var defer = $q.defer();

      $http.post(API.URL + 'api/register', model).success(function(result){
        defer.resolve(result);
      }).error(function(error){
        defer.reject(error);
      });

      return defer.promise;
    }

    return{
      register: register
    }

  }

})();

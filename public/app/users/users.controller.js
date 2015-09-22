(function() {
  'use strict';

  angular
    .module('app.users', ['app.users.service'])
    .controller('UsersController', UsersController);

  function UsersController($state, $timeout, UsersService) {
    var vm = this;
    vm.users = [];

    UsersService.get().then(function(response){
      vm.users = response;
    })

  }

})();

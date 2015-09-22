(function() {
  'use strict';

  angular
    .module('app.login', ['app.login.service'])
    .controller('LoginController', LoginController);

  function LoginController($state, $timeout, LoginService) {
    var vm = this;
    vm.model = {};

    vm.logar = function(isValid){
        if (isValid) {
          LoginService.login(vm.model).then(function(result){
            console.log(result);
          });
        } else {
          alert('Preencha todos os campos!');
        }

    };

  }

})();

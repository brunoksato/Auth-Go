(function() {
  'use strict';

  angular
    .module('app.signup', ['app.signup.service'])
    .controller('SignupController', SignupController);

  function SignupController($scope, $timeout, $state, SignupService) {
    var vm = this;
    vm.model = {};

    vm.create = function(isValid){

        console.log(isValid, 'entro');

        if (isValid) {

            if(vm.model.password.toLowerCase().trim() !== vm.model.confirmPassword.toLowerCase().trim()){
                alert('senha iguais');
            }

            var model = angular.copy(vm.model);
            delete model["confirmPassword"];

            SignupService.register(model).then(function(response){
                console.log(response);
            })

        } else {
           alert('Erro ao salvar');
       }

    };

  }

})();

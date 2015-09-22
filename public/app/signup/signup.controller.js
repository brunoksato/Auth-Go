(function() {
  'use strict';

  angular
    .module('app.signup', ['app.signup.service'])
    .controller('SignupController', SignupController);

  function SignupController($scope, $timeout, $state, SignupService) {
    var vm = this;
    vm.model = {};
    vm.alertDanger = false;
    vm.alertSuccess = false;
    vm.alertPassword = false;

    vm.create = function(isValid){

        if (isValid) {

            if(vm.model.password.toLowerCase().trim() !== vm.model.confirmPassword.toLowerCase().trim()){
                vm.alertPassword = true;

                $timeout(function(){
                    vm.alertPassword = false;
                },5000);
                return;
            }

        } else {
            vm.alertDanger = true;
            $timeout(function(){
                vm.alertDanger = false;
            },3000);
        }

    };

  }

})();

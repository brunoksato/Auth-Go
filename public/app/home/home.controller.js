(function() {
  'use strict';

  angular
    .module('app.home', [])
    .controller('HomeController', HomeController);

  /** @ngInject */
  function HomeController($state) {
    var vm = this;

    vm.signin = function(){
        $state.go('signin');
    };

    vm.login = function(){
        $state.go('login');
    };

  }

})();

(function() {
  'use strict';

  describe('controllers', function(){

    beforeEach(module('app.home'));
    beforeEach(module('ui.router'));
    beforeEach(module('stateMock'));

    it('should respond to URL', inject(function($state) {

      $state.expectTransitionTo('home');

    }));

    it('vm.signin', inject(function($state) {

      $state.expectTransitionTo('signup');

    }));

    it('vm.login', inject(function($state) {

      $state.expectTransitionTo('login');

    }));

  });
})();

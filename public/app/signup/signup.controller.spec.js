(function() {
  'use strict';

  describe('controllers', function(){

    beforeEach(module('app.signup'));
    beforeEach(module('ui.router'));
    beforeEach(module('stateMock'));

    it('should respond to URL', inject(function($state) {

      $state.expectTransitionTo('signup');

    }));

    it('vm.create', inject(function($state) {

       $state.expectTransitionTo('users');

    }));

  });
})();

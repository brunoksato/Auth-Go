(function() {
  'use strict';

  describe('controller', function(){

    beforeEach(module('app.login'));
    beforeEach(module('ui.router'));
    beforeEach(module('stateMock'));

    it('should respond to URL', inject(function($state) {

      $state.expectTransitionTo('login');

    }));

    it('vm.login', inject(function($state) {

       $state.expectTransitionTo('users');

    }));
  });

})();

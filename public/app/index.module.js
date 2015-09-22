(function() {
  'use strict';

  angular
    .module('app', [
        'ui.router',
        'app.home',
        'app.login',
        'app.signup',
        'app.users'
    ]);

})();

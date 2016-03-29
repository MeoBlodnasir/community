import Ember from 'ember';

export default Ember.Route.extend({

  setupController: function(controller, model) {
    controller.set('model', model);
    controller.set('editingPassword', false);
    controller.set('editingPasssword', "");
    controller.set('errorMessage', "");
  },
});

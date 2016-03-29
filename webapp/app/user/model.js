import DS from 'ember-data';

export default DS.Model.extend({
  email: DS.attr('string'),
  activated: DS.attr('boolean'),
  isAdmin: DS.attr('boolean'),
  firstName: DS.attr('string'),
  lastName: DS.attr('string'),
  password: DS.attr('string'),

  fullName: function() {
    return `${this.get('firstName')}  ${this.get('lastName')}`;
  }.property('firstName', 'lastName'),
});

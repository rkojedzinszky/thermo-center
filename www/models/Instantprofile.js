import Model from './model';
import meta from './g/Instantprofile';
export default Model.extend('Models.Instantprofile', { // static members
    _meta: meta
}, { // dynamic members
	toggle() {
		var self = this;
		var active = self.attr('active');
		this.attr('active', !this.active);
		this.save().then(null, function() {
			self.attr('active', active);
		});
	}
});

import Model from './model';
import meta from './g/Thsensor';
export default Model.extend('Models.Thsensor', { // static members
    _meta: meta,
    _maxdelta: 30000
}, { // dynamic members
	startRefresh() {
		var self = this;
		var to_sleep;
		var age = this.getAge();

		if (age == null) {
			to_sleep = this.constructor._maxdelta;
		} else if (age * 1000 < this.constructor._maxdelta) {
			to_sleep = this.constructor._maxdelta - age * 1000;
		} else {
			to_sleep = 1000; // sleep 1 seconds
		}

		steal.dev.log('THSensor ' + this.attr('id') + ': sleeping ' + to_sleep + ' mseconds');

		this._th = setTimeout(function() {
			self._detail().then(function() {
				self.startRefresh();
			})
		}, to_sleep);
	},
	stopRefresh() {
		clearTimeout(this._th);
		delete this._th;
	}
});

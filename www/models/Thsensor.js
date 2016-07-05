import Model from './model';
import meta from './g/Thsensor';
export default Model.extend('Models.Thsensor', { // static members
    _meta: meta,
    _maxdelta: 30000
}, { // dynamic members
	startRefresh() {
		var self = this;
		var time_passed = new Date() - new Date(this.attr('last_ts'));
		var to_sleep;

		if (time_passed < this.constructor._maxdelta) {
			to_sleep = this.constructor._maxdelta - time_passed;
		} else {
			to_sleep = 1000;
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
	}
});

import 'can/component/';
import list from './list.stache!';
import sensor from './sensor.stache!';
import './summary.less!';

can.Component.extend({
	tag: 'sensor-sensor',
	template: sensor,
	viewModel(attrs, scope, element) {
		return new can.Map({
			sensor() { return scope._context; },
			expand: false,
			classes() {
				if (this.sensor().getValid() === false) {
					return 'text-warning';
				}

				return '';
			},
			can_resync() {
				return this.sensor().attr('sensor_resync') != null;
			},
			do_resync() {
				this.sensor().getSensor_resync().then(function(o) {
					o.save();
				});
			},
			toggle() {
				this.attr('expand', !this.attr('expand'));
			}
		});
	},
	helpers: {
		format_num(value, fix) {
			value = value();

			if (typeof(value) == 'number') {
				return value.toFixed(fix);
			}

			return value;
		}
	}
});

can.Component.extend({
	tag: 'sensor-list',
	template: list
});

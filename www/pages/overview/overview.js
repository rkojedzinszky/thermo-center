import THSensor from 'models/Thsensor';
import 'can/component/';
import 'can/construct/super/';
import list from './list.stache!';
import sensor from './sensor.stache!';
import './overview.less!';

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
	tag: 'page-overview',
	template: list,
	viewModel(attrs, parentScope, element) {
		return can.Map.extend({
			define: {
				sensors: {
					Value: THSensor.List
				}
			}
		});
	},
	events: {
		inserted() {
			var view = this.viewModel;
			THSensor.findAll({'order_by': 'id'}).then(function(res) {
				can.each(res, function(s) {
					view.sensors.push(s);
					s.startRefresh();
				});
			});
		},
		removed() {
			var view = this.viewModel;
			can.each(view.sensors, (s) => s.stopRefresh());
		}
	}
});

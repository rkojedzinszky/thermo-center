import HeatSensor from 'models/Heatsensor';
import DayType from 'models/Daytype';
import 'can/component/';
import list from './list.stache!';
import sensor from './sensor.stache!';
import './detail';

can.Component.extend({
	tag: 'heatcontrol-sensor',
	template: sensor,
	viewModel: {
		toggle() {
			this.attr('expand', !this.attr('expand'));
		}
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
	tag: 'page-heatcontrol',
	template: list,
	viewModel: {
		sensors: [],
		daytypes: null
	},
	events: {
		inserted() {
			var view = this.viewModel;
			HeatSensor.findAll({'order_by': 'id'}).then(function(res) {
				view.attr('sensors', res);
				can.each(res, function(s) {
					s.startRefresh();
				});
			});
			view.attr('daytypes', DayType.findAll());
		},
		removed() {
			var view = this.viewModel;
			can.each(view.sensors, (s) => s.stopRefresh());
		}
	}
});

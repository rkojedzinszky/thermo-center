import 'can/component/';
import template from './detail.stache!';
import DayType from 'models/Daytype';
import HeatSensor from 'models/Heatsensor';
import HeatSensorTime from 'models/Heatsensortime';
import HeatSensorOverride from 'models/Heatsensoroverride';

can.Component.extend({
	tag: 'page-edit',
	template: template,
	viewModel: {
		days: [],
		overrides: [],
		d: 1,
		add() {
			var self = this;
			var st = new Date();
			var end = new Date(st.getTime() + this.attr('d') * 3600 * 1000);
			var hso = new HeatSensorOverride({
				sensor: self.attr('sensor'),
				start: st,
				end: end,
				target_temp: self.attr('t'),
			});
			hso.save().then(function(hso) {
				self.attr('overrides').push(hso);
			});
		},
		addProfile(day) {
			day.attr('times').push(new HeatSensorTime({daytype: day, sensor: this.sensor, target_temp: 20}));
		}
	},
	events: {
		inserted() {
			var view = this.viewModel;
			var days = view.attr('days');

			can.when(HeatSensor.findOne({id: can.route.attr('id')}).then(function(s) {
				view.attr('sensor', s);
				HeatSensorOverride.findAll({sensor: s.attr('id')}).then(function(overrides) {
					view.attr('overrides', overrides);
				});
				return s;
			}), DayType.findAll()).then(function(s, r) {
				can.each(r, function(d) {
					var dt = new DayType(d);
					dt.attr('times', []);
					HeatSensorTime.findAll({sensor: s.attr('id'), daytype: dt.attr('id')}).then(function(times) {
						dt.attr('times', times);
					});
					days.push(dt);
				});
			});
		}
	}
});

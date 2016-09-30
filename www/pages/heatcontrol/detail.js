import 'can/component/';
import template from './detail.stache!';
import DayType from 'models/Daytype';
import HeatSensorTime from 'models/Heatsensortime';
import HeatSensorOverride from 'models/Heatsensoroverride';

can.Component.extend({
	tag: 'heatcontrol-detail',
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
				target_temp: this.attr('tt')
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
			view.attr('tt', view.attr('t'));
			var days = view.attr('days');
			view.attr('daytypes').then(function(r) {
				can.each(r, function(d) {
					var dt = new DayType(d);
					dt.attr('times', []);
					HeatSensorTime.findAll({sensor: view.sensor.attr('id'), daytype: dt.attr('id')}).then(function(times) {
						dt.attr('times', times);
					});
					days.push(dt);
				});
			});
			HeatSensorOverride.findAll({sensor: view.sensor.attr('id')}).then(function(overrides) {
				view.attr('overrides', overrides);
			});
		}
	}
});

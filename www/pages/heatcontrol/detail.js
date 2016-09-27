import 'can/component/';
import template from './detail.stache!';
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
				target_temp: this.attr('t')
			});
		},
	},
	events: {
		inserted() {
			var view = this.viewModel;
			var days = view.attr('days');
			view.attr('daytypes').then(function(r) {
				can.each(r, function(d) {
					var dt = new can.Map(d);
					dt.attr('times', []);
					HeatSensorTime.findAll({sensor_id: view.sensor.attr('id'), daytype_id: dt.attr('id')}).then(function(times) {
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

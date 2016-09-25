import 'can/component/';
import template from './detail.stache!';
import HeatSensorTime from 'models/Heatsensortime';

can.Component.extend({
	tag: 'heatcontrol-detail',
	template: template,
	viewModel: {
		days: []
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
		}
	}
});

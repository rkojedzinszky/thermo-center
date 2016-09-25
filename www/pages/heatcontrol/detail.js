import DayType from 'models/Daytype';
import DayTime from 'models/Daytime';
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
			DayType.findAll().then(function(r) {
				can.each(r, function(dt) {
					dt.attr('times', []);
					HeatSensorTime.findAll({sensor_id: view.sensor.attr('id'), daytype_id: dt.attr('id')}).then(function(times) {
						dt.attr('times', times);
					});
				});
				view.attr('days', r);
			});
		}
	}
});

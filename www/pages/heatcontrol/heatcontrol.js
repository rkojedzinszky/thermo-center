import HeatSensor from 'models/Heatsensor';
import DayType from 'models/Daytype';
import 'can/component/';
import list from './list.stache!';
import stache from 'can/view/stache/';

can.Component.extend({
	tag: 'page-heatcontrol',
	template: list,
	viewModel: {
		sensors: [],
		daytypes: null
	},
	helpers: {
		edit_link(sensor) {
			return stache.safeString(can.route.url({'page': 'edit', 'id': sensor.getId()}));
		}
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

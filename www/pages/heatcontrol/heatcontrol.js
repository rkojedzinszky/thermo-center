import Control from 'models/Control';
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
			Control.findAll().then(function(res) {
				if (view.attr('sensors')) {
					view.attr('sensors', res);
					can.each(view.sensors, (s) => s.startRefresh());
				}
			});
		},
		removed() {
			var view = this.viewModel;
			can.each(view.sensors, (s) => s.stopRefresh());
			view.attr('sensors', undefined);
		}
	}
});

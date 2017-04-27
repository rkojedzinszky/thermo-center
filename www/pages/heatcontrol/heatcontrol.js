import Control from 'models/Control';
import InstantProfile from 'models/Instantprofile';
import 'can/component/';
import list from './list.stache!';
import stache from 'can/view/stache/';
import './heatcontrol.less!';

can.Component.extend({
	tag: 'page-heatcontrol',
	template: list,
	viewModel: {
		sensors: [],
		instantprofiles: [],
		daytypes: null
	},
	helpers: {
		edit_link(sensor) {
			return stache.safeString(can.route.url({'page': 'edit', 'id': sensor.getId()}));
		},
		ip_classes(iprofile) {
			return iprofile.attr('active') ? 'active' : '';
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
			InstantProfile.findAll().then(function(res) {
				view.attr('instantprofiles', res);
			});
		},
		removed() {
			var view = this.viewModel;
			can.each(view.sensors, (s) => s.stopRefresh());
			view.attr('sensors', undefined);
		}
	}
});

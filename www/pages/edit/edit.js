import 'can/component/';
import template from './detail.stache!';
import DayType from 'models/Daytype';
import Control from 'models/Control';
import HeatControlProfile from 'models/Heatcontrolprofile';
import HeatControlOverride from 'models/Heatcontroloverride';

can.Component.extend({
	tag: 'page-edit',
	template: template,
	viewModel: {
		days: [],
		overrides: [],
		d: 1,
		control: null,
		add() {
			var self = this;
			var st = new Date();
			var end = new Date(st.getTime() + this.attr('d') * 3600 * 1000);
			var hco = new HeatControlOverride({
				control: self.attr('control'),
				start: st,
				end: end,
				target_temp: self.attr('t'),
			});
			hco.save().then(function(hco) {
				self.attr('overrides').push(hco);
			});
		},
		addProfile(day) {
			day.attr('times').push(new HeatControlProfile({daytype: day, control: this.attr('control'), target_temp: 20}));
		},
		hcSave() {
			this.control.save();
		},
	},
	events: {
		inserted() {
			var view = this.viewModel;
			var days = view.attr('days');

			can.when(Control.findOne({id: can.route.attr('id')}).then(function(hc) {
				view.attr('control', hc);
				HeatControlOverride.findAll({control: hc.getId()}).then(function(overrides) {
					view.attr('overrides', overrides);
				});
				return hc;
			}), DayType.findAll()).then(function(hc, r) {
				can.each(r, function(d) {
					var dt = new DayType(d);
					dt.attr('times', []);
					HeatControlProfile.findAll({control: hc.getId(), daytype: dt.getId()}).then(function(times) {
						dt.attr('times', times);
					});
					days.push(dt);
				});
			});
		}
	}
});

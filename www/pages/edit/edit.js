import 'can/component/';
import template from './detail.stache!';
import DayType from 'models/Daytype';
import HeatControl from 'models/Heatcontrol';
import HeatControlProfile from 'models/Heatcontrolprofile';
import HeatControlOverride from 'models/Heatcontroloverride';

can.Component.extend({
	tag: 'page-edit',
	template: template,
	viewModel: {
		days: [],
		overrides: [],
		d: 1,
		heatcontrol: null,
		add() {
			var self = this;
			var st = new Date();
			var end = new Date(st.getTime() + this.attr('d') * 3600 * 1000);
			var hco = new HeatControlOverride({
				heatcontrol: self.attr('heatcontrol'),
				start: st,
				end: end,
				target_temp: self.attr('t'),
			});
			hco.save().then(function(hco) {
				self.attr('overrides').push(hco);
			});
		},
		addProfile(day) {
			day.attr('times').push(new HeatControlProfile({daytype: day, heatcontrol: this.attr('heatcontrol'), target_temp: 20}));
		},
		hcSave() {
			this.heatcontrol.save();
		},
	},
	events: {
		inserted() {
			var view = this.viewModel;
			var days = view.attr('days');

			can.when(HeatControl.findOne({id: can.route.attr('id')}).then(function(hc) {
				view.attr('heatcontrol', hc);
				HeatControlOverride.findAll({heatcontrol: hc.getId()}).then(function(overrides) {
					view.attr('overrides', overrides);
				});
				return hc;
			}), DayType.findAll()).then(function(hc, r) {
				can.each(r, function(d) {
					var dt = new DayType(d);
					dt.attr('times', []);
					HeatControlProfile.findAll({heatcontrol: hc.getId(), daytype: dt.getId()}).then(function(times) {
						dt.attr('times', times);
					});
					days.push(dt);
				});
			});
		}
	}
});

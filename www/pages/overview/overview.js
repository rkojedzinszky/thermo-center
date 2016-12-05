import THSensor from 'models/Thsensor';
import 'can/component/';
import list from './list.stache!';

can.Component.extend({
	tag: 'sensor-resync',
	template: can.stache('{{#if can_resync}}<button ($click)="do_resync()">Resync</button>{{/if}}'),
	viewModel: {
		can_resync() {
			return this.sensor.attr('sensor_resync') != null;
		},
		do_resync() {
			this.sensor.getSensor_resync().then(function(o) {
				o.save();
			});
		}
	}
});

can.Component.extend({
	tag: 'page-overview',
	template: list,
	viewModel: {
		sensors: []
	},
	helpers: {
		classes(args) {
			if (args.context.getValid() === false) {
				return 'text-warning';
			}

			return '';
		},
	},
	events: {
		inserted() {
			var view = this.viewModel;
			THSensor.findAll({'order_by': 'id'}).then(function(res) {
				if (view.attr('sensors')) {
					view.attr('sensors', res);
					can.each(res, function(s) {
						s.startRefresh();
					});
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

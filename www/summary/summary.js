import 'can/component/';
import list from './list.stache!';
import sensor from './sensor.stache!';
import './summary.less!';

can.Component.extend({
	tag: 'sensor-sensor',
	template: sensor,
	viewModel: {
		expand: false,
		toggle() {
			this.attr('expand', !this.attr('expand'));
		}
	},
	helpers: {
		format_num(value, fix) {
			return value().toFixed(fix);
		}
	}
});

can.Component.extend({
	tag: 'sensor-list',
	template: list
});

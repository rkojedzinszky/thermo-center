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
			value = value();

			if (typeof(value) == 'number') {
				return value.toFixed(fix);
			}

			return value;
		}
	}
});

can.Component.extend({
	tag: 'sensor-list',
	template: list
});

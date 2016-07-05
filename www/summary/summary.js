import 'can/component/';
import list from './list.stache!';

can.Component.extend({
	tag: 'sensor-list',
	template: list,
	helpers: {
		format_vcc(ctx) {
			return ctx.context.attr('vcc').toFixed(2);
		},
		format_rssi(ctx) {
			return ctx.context.attr('rssi').toFixed(1);
		},
		format_lqi(ctx) {
			return ctx.context.attr('lqi').toFixed(0);
		},
		format_interval(ctx) {
			return ctx.context.attr('interval').toFixed(1);
		},
		format_temperature(ctx) {
			return ctx.context.attr('temperature').toFixed(1);
		},
		format_humidity(ctx) {
			return ctx.context.attr('humidity').toFixed(1);
		}
	},
});

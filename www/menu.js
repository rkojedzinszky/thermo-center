import 'can/component/';
import 'can/route/';
import stache from 'can/view/stache/';
import view from './menu.stache!';

can.Component.extend({
	tag: 'thermo-menu',
	template: view,
	helpers: {
		menu_link(name) {
			return stache.safeString(can.route.url({'page': name}));
		}
	}
});

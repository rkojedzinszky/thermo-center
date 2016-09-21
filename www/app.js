import 'models/Thsensor';
import 'can/map/';
import 'can/map/define/';
import 'can/route/';

can.route(':page', {'page': 'overview'});
can.route.ready();

export default can.Map.extend({
	define: {
		menu: {
			get() {
				return can.route.attr('page');
			},
			serialize: false
		},
		menus: {
			get() {
				return ['overview', 'graphs'];
			},
			serialize: false
		}
	}
});

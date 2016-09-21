import 'can/map/';
import 'can/map/define/';
import 'can/route/';
import stache from 'can/view/stache/';
import './menu';
import 'bootstrap/dist/css/bootstrap.css!';
import $ from 'jquery';
import './tastypie';

var AppState = can.Map.extend({
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

var appState = new AppState();
var body = $('body');
body.append(stache('<thermo-menu />')(appState));
var cdiv = $('<div>');
body.append(cdiv);

appState.bind('menu', function(ev, newVal, oldVal) {
	var p = 'pages/' + newVal + '/';
	var t = '<page-' + newVal + ' />';
	System.import(p).then(function(m) {
		cdiv.html(stache(t)(appState));
	});
});

can.route(':page', {'page': 'overview'});
can.route.ready();

import 'can/map/';
import 'can/map/define/';
import 'can/route/';
import 'can/construct/super/';
import stache from 'can/view/stache/';
import './menu';
import 'bootstrap/dist/css/bootstrap.css!';
import $ from 'jquery';
import './tastypie';
import Session from 'models/Session';
import 'pages/login/';
import 'pages/logout/';
import './index.less!';

var AppState = can.Map.extend({
	define: {
		menu: {
			get() {
				if (!this.attr('ready')) {
					return null;
				}

				if (this.loggedIn()) {
					return can.route.attr('page');
				} else {
					return 'login';
				}
			},
			serialize: false
		},
		menus: {
			get() {
				if (this.loggedIn()) {
					return ['overview', 'heatcontrol', 'logout'];
				} else {
					return [];
				}
			},
			serialize: false
		},
		ready: {
			serialize: false,
			set() {
				can.route.ready();
			}
		},
		session: {
		}
	},
	loggedIn() {
		return this.attr('session') != null;
	},
	login(user, pass) {
		var self = this;
		var s = new Session({username: user, password: pass});
		s.save().then(function(s) {
			self.attr('session', s);
		});
	},
	logout() {
		if (this.attr('session')) {
			var self = this;

			this.attr('session').destroy().then(function(s) {
				self.removeAttr('session');
				if (can.route.attr('page') == 'logout') {
					can.route.removeAttr('page');
				}
			});
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
	System.import(p).then(function(m) {
		if (m.init && typeof(m.init) === 'function') {
			m.init(appState);
		} else {
			var t = '<page-' + newVal + ' {app}="app"/>';
			cdiv.html(stache(t)({app: appState}));
		}
	});
});

can.route(':page', {'page': 'overview'});

Session.findAll().then(function(res) {
	if (res.length == 1) {
		appState.attr('session', res[0]);
	}
	appState.attr('ready', true);
});

'use strict';
import 'bootstrap/dist/css/bootstrap-reboot.css!';
import 'bootstrap/dist/css/bootstrap.css!';
import Component from 'can-component';
import DefineMap from 'can-define/map/map';
import stache from 'can-stache';
import 'can-stache-bindings';
import route from 'can-route';
import {Session} from 'models/Session';
import './navbar';
import './footer';
import './common.less!';
import '~/common';
import Version from './version';

// Register number formatting helpers
const Precisions = new (DefineMap.extend({
	temperature: { type: 'number', default: 2 },
	humidity: { type: 'number', default: 1 },
	vcc: { type: 'number', default: 2 },
	interval: { type: 'number', default: 2 },
	rssi: { type: 'number', default: 0 },
	lqi: { type: 'number', default: 0 },
	target_temp: { type: 'number', default: 2 },
	pidcontrol: { type: 'number', default: 3 },
}))();

stache.addHelper('format', function(metric, value) {
	if (typeof(value) == 'number') {
		return value.toFixed(Precisions.get(metric));
	}

	return value;
});

const wsurl = function() {
	const protocol = document.location.protocol == 'https:' ? 'wss:' : 'ws:';
	let path = document.location.pathname.split('/');
	path.pop();
	path.push('ws');

	return protocol + '//' + document.location.host + path.join('/') + '/';
}();

const WsHandler = DefineMap.extend({
	'ws': { serialize: false, default: null },
	'app': { serialize: false, default: null },
	'_timeout': { serialize: false, default: null },

	start() {
		this._timeout = null;
		if (this.app == null) {
			return;
		}

		const ws = new WebSocket(wsurl);
		var self = this;
		ws.addEventListener('open', function() {
			self.ws = ws;
		});

		ws.addEventListener('message', function(e) {
			if (self.app && self.app.onmessage) {
				self.app.onmessage(e.data);
			}
		});

		ws.addEventListener('close', function(e) {
			console.log(e);
			self.ws = null;
			if (self.app) {
				self._timeout = setTimeout(function() {
					self.start();
				}, 1000);
			}
		});
	},

	stop() {
		this.app = null;
		if (this.ws) {
			this.ws.close();
		}
		if (this._timeout) {
			clearTimeout(this._timeout);
		}
	},

	is_connected() {
		return this.ws != null;
	}
});

const AppState = DefineMap.extend({
	'session': { serialize: false },
	'element': { serialize: false },
	'visible': { serialize: false, type: "boolean", default: true },
	'loaded': { serialize: false, type: "boolean", default: false },
	'url': { default: () => new DefineMap() },
	'ws': { serialize: false, default: null },
	'onmessage': { serialize: false },
	'current_time': { serialize: false, default: function() { return new Date() } },
	'current_timer': { serialize: false },
	'uiVersion': { default: () => Version },
	'displaypage': {
		get() {
			if (this.session != null)
				return this.url.page;
			return 'login';
		}
	},
	'need_ws': {
		get() {
			return this.session && this.visible;
		}
	},
	'ws_connected': {
		get() {
			return this.ws && this.ws.is_connected();
		}
	},
	setpage() {
		var self = this;
		var page = this.displaypage;

		steal.import('~/pages/' + page + '/').then(function(module) {
			if (typeof(module.default) === 'function') {
				module.default(self);
			} else {
				while (self.element.firstChild) {
					self.element.removeChild(self.element.firstChild);
				}
				self.element.appendChild(stache('<thermo-p-' + page + ' app:bind="."/>')(self));
			}
		});
	},
	connectedCallback(element) {
		var self = this;
		self.element = element.querySelector('.content');

		route.data = this.url;
		route.register('{page}', {'page': 'overview'});
		route.start();

		this.listenTo('need_ws', this._ws.bind(this));
		this.listenTo('displaypage', this.setpage.bind(this));

		Session.getList().then(function(res) {
			if (res.length == 1) {
				self.session = res[0]; // will call setpage
			} else {
				self.setpage();
			}
		});

		window.addEventListener('visibilitychange', self.visibilitychanged.bind(self));

		self.visibilitychanged();
	},

	visibilitychanged() {
		this.visible = document.visibilityState == 'visible';

		if (this.visible) {
			this._startTimer();
		} else {
			this._stopTimer();
		}
	},

	// Start and stop ws handler based on session
	_ws(event) {
		const need_ws = event.value;

		if (this.ws) {
			this.ws.stop();
			this.ws = null;
		}

		if (need_ws) {
			this.ws = new WsHandler();
			this.ws.app = this;
			this.ws.start()
		}
	},

	_startTimer() {
		var self = this;
		this.current_timer = window.setInterval(function() {
			self.current_time = new Date();
		}, 1000);
	},

	_stopTimer() {
		window.clearInterval(this.current_timer);
	},
});

Component.extend({
	tag: 'thermo-main',
	view: `
	<thermo-navbar app:bind="." />
	<div class="content {{#if loaded}}loaded{{/if}}" />
	<thermo-footer app:bind="." />
	`,
	ViewModel: AppState
});

// Register service-worker
if ('serviceWorker' in navigator) {
	navigator.serviceWorker.register("/sw.js");
}

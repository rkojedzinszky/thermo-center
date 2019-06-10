'use strict';
import 'bootstrap/dist/css/bootstrap-reboot.css!';
import 'bootstrap/dist/css/bootstrap.css!';
import Component from 'can-component';
import DefineMap from 'can-define/map/map';
import stache from 'can-stache';
import 'can-stache-bindings';
import route from 'can-route';
import {Session} from 'models/Session';
import $ from 'jquery';
import './navbar';
import './common.less!';

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

const AppState = DefineMap.extend({
	'session': { serialize: false },
	'url': { default: () => new DefineMap() },
	'ws': { serialize: false, default: null },
	'onmessage': { serialize: false },
	'current_time': { serialize: false, default: function() { return new Date() } },
	'current_timer': { serialize: false },
	'displaypage': {
		get() {
			if (this.session != null)
				return this.url.page;
			return 'login';
		}
	},
	setpage(element) {
		var self = this;
		var page = this.displaypage;

		steal.import('~/pages/' + page + '/').then(function(module) {
			if (typeof(module.default) === 'function') {
				module.default(self);
			} else {
				element.html(stache('<thermo-p-' + page + ' app:bind="."/>')(self));
			}
		});
	},
	connectedCallback(element) {
		var self = this;
		var element = $(element).find('.content');

		route.data = this.url;
		route.register('{page}', {'page': 'overview'});
		route.start();

		this.listenTo('displaypage', this.setpage.bind(this, element));
		Session.getList().then(function(res) {
			if (res.length == 1) {
				self.session = res[0]; // will call setpage
			} else {
				self.setpage(element);
			}
		});

		this._openws();
		this._startTimer();
	},
	_openws() {
		var self = this;
		const ws = new WebSocket(wsurl);
		ws.addEventListener('open', function() {
			self.ws = ws;
		});

		ws.addEventListener('message', function(e) {
			// console.log('wsmsg: ' + e.data);
			if (self.onmessage) {
				self.onmessage(e.data);
			}
		});

		ws.addEventListener('close', function(e) {
			console.log(e);
			self.ws = null;
			setTimeout(function() {
				self._openws();
			}, 1000);
		});
	},
	_startTimer() {
		var self = this;
		this.current_timer = window.setInterval(function() {
			self.current_time = new Date();
		}, 1000);
	},
});

Component.extend({
	tag: 'thermo-main',
	view: `
	<thermo-navbar app:bind="." />
	<div class="content" />
	`,
	ViewModel: AppState
});


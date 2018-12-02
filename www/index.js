'use strict';
import {Component, DefineMap, stache, route} from 'can';
import 'bootstrap/dist/css/bootstrap-reboot.css!';
import 'bootstrap/dist/css/bootstrap.css!';
import {Session} from 'models/Session';
import $ from 'jquery';

// Register number formatting helpers
const Precisions = new (DefineMap.extend({
	temperature: { type: 'number', default: 2 },
	humidity: { type: 'number', default: 1 },
	vcc: { type: 'number', default: 2 },
	interval: { type: 'number', default: 2 },
	rssi: { type: 'number', default: 0 },
	lqi: { type: 'number', default: 0 },
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
	'session': { default: null, serialize: false },
	'page': { default: 'overview' },
	'ws': { serialize: false, default: null },
	'onmessage': { serialize: false },
	'displaypage': {
		get() {
			if (this.session != null)
				return this.page;
			return 'login';
		}
	},
	setpage(element) {
		var self = this;
		var page = this.displaypage;
		System.import('pages/' + page + '/').then(function(module) {
			element.html(stache('<thermo-p-' + page + ' appstate:bind="."/>')(self));
		});
	},
	connectedCallback(element) {
		var self = this;
		var element = $(element).find(".content");
		this.listenTo('displaypage', this.setpage.bind(this, element));
		Session.getList().then(function(res) {
			if (res.length == 1) {
				self.session = res[0];
			}
			self.setpage(element);
		});

		this._openws();
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
	}
});

Component.extend({
	tag: 'thermo-main',
	view: '<div><h3>{{#ws}}{{else}}NOT {{/ws}}CONNECTED</h3><div class="content" /></div>',
	ViewModel: AppState
});

import {Component, DefineMap, stache, route} from 'can';
import 'bootstrap/dist/css/bootstrap-reboot.css!';
import 'bootstrap/dist/css/bootstrap.css!';
import {Session} from 'models/Session';
import $ from 'jquery';

const AppState = DefineMap.extend({
	'session': { default: null, serialize: false },
	'page': { default: 'overview' },
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
		var element = $(element);
		this.listenTo('displaypage', this.setpage.bind(this, element));
		Session.getList().then(function(res) {
			if (res.length == 1) {
				self.session = res[0];
			}
			self.setpage(element);
		});
	}
});

Component.extend({
	tag: 'thermo-main',
	view: '<div />',
	ViewModel: AppState
});

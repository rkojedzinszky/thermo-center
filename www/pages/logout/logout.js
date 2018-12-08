'use strict';
// import {Component} from 'can';
import queues from 'can-queues';

/*
Component.extend({
	tag: 'thermo-p-logout',
	view: `Logging out...`,
	ViewModel: {
		connectedCallback(element) {
			var self = this;
			this.app.session.destroy().then(function() {
				queues.batch.start();
				self.app.url.page = '';
				self.app.session = null;
				queues.batch.stop();
			});
		}
	}
});
*/

export default function (app) {
	queues.batch.start();
	app.url.page = '';
	app.session = null;
	queues.batch.stop();
};

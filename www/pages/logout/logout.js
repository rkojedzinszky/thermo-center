'use strict';
import queues from 'can-queues';

export default function (app) {
	app.session.destroy().then(function() {
		queues.batch.start();
		app.url.page = '';
		app.session = null;
		queues.batch.stop();
	});
};

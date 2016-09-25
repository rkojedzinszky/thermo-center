import Session from 'models/Session';
import template from './login.stache!';
import './login.less!';
import 'can/component/';

can.Component.extend({
	tag: 'page-login',
	template: template,
	viewModel: {
		login(scope, el, ev) {
			ev.preventDefault();

			this.app.login(this.attr('username'), this.attr('password'));
		}
	}
});

'use strict';
import Component from 'can-component';
import {Session} from 'models/Session';
import './login.less!';

Component.extend({
	tag: 'thermo-p-login',
	view: `
<form class="loginbox text-center form-horizontal" on:submit="submit(scope)">
	<h3>Thermo Center login</h3>
	<div class="form-group">
		<input class="form-control" placeholder="Username" value:to="username" />
	</div>
	<div class="form-group">
		<input class="form-control" placeholder="Password" type="password" value:to="password" />
	</div>
	<div class="form-group">
		<button type="submit" class="btn btn-default">Login</button>
	</div>
	{{#loginerror}}
	<div class="form-group">
		<div class="bg-danger">Login error</div>
	</div>
	{{/loginerror}}
</form>
	`,
	ViewModel: {
		username: 'string',
		password: 'string',
		loginerror: 'boolean',
		submit(scope) {
			var self = this;
			self.loginerror = false;
			scope.event.preventDefault();
			new Session({username: this.username, password: this.password}).save().then(function(s) {
				self.app.session = s;
			}, function() {
				self.loginerror = true;
			});
		}
	}
});

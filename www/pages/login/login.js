'use strict';
import {Component} from 'can';
import './login.less!';
import {Session} from 'models/Session';

Component.extend({
	tag: 'thermo-p-login',
	view: `
	<div class="loginbox container-fluid">
		<div class="row">
		<div class="col-sm-4"></div>
		<form class="col-sm-4 container-fluid form-horizontal" on:submit="submit(scope)">
			<div class="form-group row">
				<label class="col-sm-2 col-xs-2 control-label">Username</label>
				<div class="col-sm-4">
					<input class="form-control" value:to="username" />
				</div>
			</div>
			<div class="form-group row">
				<label class="col-sm-2 col-xs-2 control-label">Password</label>
				<div class="col-sm-4">
					<input class="form-control" type="password" value:to="password" />
				</div>
			</div>
			<div class="form-group row">
				<div class="col-sm-2"></div>
				<div class="col-sm-4">
					<button type="submit" class="btn btn-default">Login</button>
				</div>
			</div>
			{{#loginerror}}
			<div class="form-group row">
				<div class="col-sm-2"></div>
				<div class="col-sm-2 bg-danger">Login error</div>
				<div class="col-sm-2"></div>
			</div>
			{{/loginerror}}
		</form>
		<div class="col-sm-4"></div>
		</div>
	</div>
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
				self.appstate.session = s;
			}, function() {
				self.loginerror = true;
			});
		}
	}
});

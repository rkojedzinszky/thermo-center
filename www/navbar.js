"use strict";
import Component from "can-component";
import 'can-stache-route-helpers';
import 'bootstrap/js/src/collapse';
import '@fortawesome/fontawesome-free/less/solid.less!';
import '@fortawesome/fontawesome-free/less/fontawesome.less!';

Component.extend({
	tag: 'thermo-navbar',
	view: `
	<nav class="navbar navbar-expand-sm navbar-light bg-light">
	 <a class="navbar-brand" href="#"><i class="fas fa-thermometer-half"></i></a>
         <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
	  <span class="navbar-toggler-icon"></span>
	 </button>
	 <div class="collapse navbar-collapse" id="navbarNav">
	  <ul class="navbar-nav">
	   {{#for (page of this.pages)}}
	    <li class="nav-item {{#if(routeCurrent(page=page.link))}}active{{/if}}">
	     <a class="nav-link" href="{{ routeUrl(page=page.link) }}">{{ page.name }}</a>
	    </li>
	   {{/for}}
	  </ul>
	  <i style="margin-left: auto" class="nav-item nav-link fas fa-asterisk {{#if (app.ws_connected)}}text-success{{else}}text-warning{{/if}}" title="{{#if (app.ws_connected)}}{{else}}not {{/if}}connected"></i>
	 </div>
	</nav>
	`,
	ViewModel: {
		is_logged_in: {
			get() {
				return this.app.session != null;
			}
		},
		is_admin: {
			get() {
				return this.app.session && this.app.session.is_admin;
			}
		},
		pages: {
			get() {
				let list = [];
				if (this.is_logged_in) {
					list = list.concat([
						{name: 'Overview', link: 'overview'},
						{name: 'Heat control', link: 'heatcontrol'},
					]);
					if (this.is_admin) {
						list.push({name: 'Admin', link: 'admin'});
					};
					list.push({name: 'Logout', link: 'logout'});
				}
				return list;
			}
		},
	}
});

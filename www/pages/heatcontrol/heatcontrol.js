'use strict';
import Component from 'can-component';
import route from 'can-route';
import {Control} from 'models/Control';
import {InstantProfile} from 'models/InstantProfile';

let ControlCache = new Control.List();
let InstantProfileCache = new InstantProfile.List();

Component.extend({
	tag: 'thermo-p-heatcontrol',
	view: `
<table class="table table-striped table-bordered table-hover table-sm cell-align-middle">
<thead>
<tr>
	<th>Name</th>
	<th>Temp</th>
	<th>Target temp</th>
	<th>Pid control</th>
	<th>Edit</th>
</tr>
</thead>
<tbody>
{{#for (s of sensors)}}
<tr>
	<td>{{s.name}}</td>
	<td>{{format('temperature', s.temperature)}}</td>
	<td>{{format('target_temp', s.target_temp)}}</td>
	<td>{{format('pidcontrol', s.pidcontrol)}}</td>
	<td><a class="btn btn-primary mx-1" role="button" href="{{{ edit_link(s) }}}">Edit</a></td>
</tr>
{{/for}}
</tbody>
</table>
<ul class="list-unstyled list-inline">
{{#for (i of instantprofiles)}}
<li class="list-inline-item">
	<button class="btn {{#if (i.active)}}btn-primary{{else}}btn-outline-primary{{/if}} m-1" on:click="toggle(i)">{{ i.name }}</button>
</li>
{{/for}}
</ul>
	`,
	ViewModel: {
		sensors: { default: () => ControlCache },
		instantprofiles: { default: () => InstantProfileCache },
		daytypes: { default: null },

		toggle(i) {
			const saved = i.active;
			i.active = !saved;
			i.save().then(() => true, function() {
				i.active = saved;
			});
		},

		connectedCallback(element) {
			var self = this;

			InstantProfile.getList().then(function(res) {
				InstantProfileCache = self.instantprofiles = res;
			});

			self.app.onmessage = function(el) {
				Control.getList({sensor_id: el});
			};

			let listener = self.visible.bind(self);
			self.app.listenTo('visible', listener);

			self.visible();

			return function() {
				self.app.onmessage = null;
				self.app.stopListening('visible', listener);
				self.stopListening();
			};
		},
		visible() {
			var self = this;

			if (self.app.visible) {
				Control.getList().then(function(res) {
					ControlCache.update(res);
					self.app.active = true;
				});
			}
		}
	},
	helpers: {
		edit_link(sensor) {
			return route.url({'page': 'edit', 'id': sensor.id});
		},
	},
});

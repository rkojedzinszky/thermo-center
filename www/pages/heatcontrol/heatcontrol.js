'use strict';
import Component from 'can-component';
import {Control} from 'models/Control';
import {InstantProfile} from 'models/InstantProfile';

Component.extend({
	tag: 'thermo-p-heatcontrol',
	view: `
	<table class="table table-striped table-bordered table-condensed table-hover table-sm">
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
		<td><a class="btn btn-default btn-xs" href="{{ edit_link }}">Edit</a></td>
	</tr>
	{{/for}}
	</tbody>
	</table>
	<ul class="list-unstyled list-inline iprofiles">
	{{#for (i of instantprofiles)}}
	<li>
		<button class="iprofile btn btn-default btn-lg {{ ip_classes . }}" on:click="toggle(i)">{{ i.name }}</button>
	</li>
	{{/for}}
	</ul>
	`,
	ViewModel: {
		sensors: { default: () => [] },
		instantprofiles: { default: () => [] },
		daytypes: { default: null },

		connectedCallback(element) {
			var self = this;
			Control.getList().then(function(res) {
				self.sensors = res;

				self.app.onmessage = function(el) {
					Control.getList({sensor_id: el});
				};
			});

			InstantProfile.getList().then(function(res) {
				self.instantprofiles = res;
			});

			return function() {
				self.app.onmessage = null;
				self.stopListening();
			};
		},
	},
	helpers: {
		edit_link(sensor) {
			return stache.safeString(can.route.url({'page': 'edit', 'id': sensor.getId()}));
		},
	},
});

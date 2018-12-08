'use strict';
import Component from 'can-component';
import {THSensor} from 'models/THSensor';
import {SensorResync} from 'models/SensorResync';

Component.extend({
	tag: 'thermo-p-overview',
	view: `
		<table class="table table-striped table-bordered table-hover table-sm">
		<thead>
		<tr>
			<th scope="col">#</th>
			<th scope="col">Name</th>
			<th scope="col">Temperature</th>
			<th scope="col">Humidity</th>
			{{#advanced}}
			<th scope="col">VCC</th>
			<th scope="col">Interval</th>
			<th scope="col">RSSI</th>
			<th scope="col">LQI</th>
			{{/advanced}}
		</tr>
		</thead>
		<tbody>
		{{#for(s of this.sensors)}}
			<tr class="{{#s.sensor_resync}}table-danger{{/s.sensor_resync}}">
				<th scope="row">{{s.id}}</th>
				<td>{{s.name}}</td>
				{{#s.sensor_resync}}
				<td colspan="2"><thermo-sensor-resync sensor:bind="s" /></td>
				{{else}}
				<td>{{format('temperature', s.temperature)}}</td>
				<td>{{format('humidity', s.humidity)}}</td>
				{{/s.sensor_resync}}
				{{#advanced}}
				<td>{{format('vcc', s.vcc)}}</td>
				<td>{{format('interval', s.interval)}}</td>
				<td>{{format('rssi', s.rssi)}}</td>
				<td>{{format('lqi', s.lqi)}}</td>
				{{/advanced}}
			</tr>
		{{/for}}
		</tbody>
		</table>
		<div class="checkbox">
			<label>
				<input type="checkbox" checked:bind="advanced" />
				Show advanced fields
			</label>
		</div>
	`,
	ViewModel: {
		sensors: { default: () => [] },
		advanced: 'boolean',
		connectedCallback(element) {
			var self = this;
			THSensor.getList({'order_by': 'id'}).then(function(res) {
				self.sensors = res;

				self.app.onmessage = function(el) {
					THSensor.get({id: el});
				};
			});

			return function() {
				self.app.onmessage = null;
				self.stopListening();
			};
		}
	}
});

Component.extend({
	tag: 'thermo-sensor-resync',
	view: `<button class="btn btn-sm btn-secondary" on:click="do_resync(scope.event)">Resync!</button>`,
	ViewModel: {
		do_resync(event) {
			event.preventDefault();

			new SensorResync({sensor: this.sensor.resource_uri}).save().then(function(s) {
				THSensor.getByResourceUri(s.sensor);
			});
		}
	}
});

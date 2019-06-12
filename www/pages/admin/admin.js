'use strict';
import Component from 'can-component';
import {THSensor} from 'models/THSensor';
import {ConfigureSensorTask} from 'models/ConfigureSensorTask';
import prettyMilliseconds from 'pretty-ms';
import stache from 'can-stache';
import $ from 'jquery';
import 'bootstrap/js/src/modal';

const taskFields = [
	{desc: 'Name', field: 'sensor_name', okclass: 'bg-primary'},
	{desc: 'Id', field: 'sensor_id', okclass: 'bg-primary'},
	{desc: 'Created', field: 'created', okclass: 'bg-success'},
	{desc: 'Started', field: 'started', okclass: 'bg-success', notokclass: 'bg-warning'},
	{desc: 'First discovery', field: 'first_discovery', okclass: 'bg-success', notokclass: 'bg-warning'},
	{desc: 'Last discovery', field: 'last_discovery', okclass: 'bg-success', notokclass: 'bg-warning'},
	{desc: 'Finished', field: 'finished', okclass: 'bg-success', notokclass: 'bg-warning'},
	{desc: 'Error', field: 'error', okclass: 'bg-danger', notokclass: 'bg-success'},
];

const taskView = stache(`
<div class="modal fade taskmodal" tabindex="-1" role="dialog">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">
	Configuring {{task.sensor_name}} (#{{task.sensor_id}})
	</h5>
      </div>
      <div class="modal-body">
        <div class="container-fluid">
	{{#for (f of fields)}}
          <div class="row {{#if (task[f.field])}}{{f.okclass}}{{else}}{{f.notokclass}}{{/if}}">
            <div class="col-sm-4">{{f.desc}}</div>
            <div class="col-sm-8">{{task[f.field]}}</div>
          </div>
        {{/for}}
	</div>
      </div>
    </div>
  </div>
</div>
`);

Component.extend({
	tag: 'thermo-p-admin',
	view: `
<table class="table table-striped table-bordered table-hover table-sm cell-align-middle">
<thead>
<tr>
	<th scope="col">#</th>
	<th scope="col">Name</th>
	<th scope="col">VCC</th>
	<th scope="col">RSSI</th>
	<th scope="col">LQI</th>
	<th scope="col" class="overview-age">Age</th>
	<th>op</th>
</tr>
</thead>
<tbody>
{{#for(s of this.sensors)}}
	<tr>
		<th scope="row">{{s.id}}</th>
		<td>{{s.name}}</td>
		<td>{{format('vcc', s.vcc)}}</td>
		<td>{{format('rssi', s.rssi)}}</td>
		<td>{{format('lqi', s.lqi)}}</td>
		<td>{{calculate_age(s)}}</td>
		<td><button class="btn btn-primary" on:click="reprogram(s)">Configure</button><button class="btn btn-danger" on:click="destroy(s)">Delete</button></td>
	</tr>
{{/for}}
</tbody>
</table>
<button class="btn btn-primary" on:click="addSensor()">Add Sensor</button>
<div class="mmodal"></div>
	`,
	ViewModel: {
		sensors: { default: () => [] },
		element: {},
		task: {},
		modal: {},
		connectedCallback(element) {
			let self = this;
			self.element = element;
			self.modal = $(this.element.querySelector('.mmodal'));

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
		},
		_pollTask() {
			let self = this;
			let task = self.task;
			if (!task)
				return;

			ConfigureSensorTask.getByResourceUri(task.resource_uri).then(function(task) {
				if (task.finished == null) {
					window.setTimeout(self._pollTask.bind(self), 1000);
				} else {
					self.task = null;
					let timeout = 1000;
					if (task.error) {
						timeout = 3000;
					}

					window.setTimeout(function() {
						self.modal.find('.taskmodal').modal('hide');
					}, timeout);
				}
			});
		},
		_newTask(params) {
			let self = this;
			let task = new ConfigureSensorTask(params);
			task.save().then(function() {
				self.task = task;
				self.modal.html(taskView({task: task, fields: taskFields}));
				self.modal.find('.taskmodal').modal();
				self._pollTask();
				if (params.sensor_id == null) {
					THSensor.get({id: task.sensor_id}).then(function(s) {
						self.sensors.push(s);
					});
				}
			});
		},
	},
	helpers: {
		calculate_age(s) {
			if (s.last_tsf != null) {
				let elapsed = Math.ceil(this.app.current_time - 1000 * s.last_tsf);
				if (elapsed < 0) {
					elapsed = 0;
				}
				return prettyMilliseconds(elapsed, {compact: true});
			} else {
				return '';
			}
		},
		destroy(s) {
			if (window.confirm('Sensor (#' + s.id + ') ' + s.name + ' will be removed. Are you sure?')) {
				s.destroy();
			}
		},
		reprogram(s) {
			if (this.task) {
				alert('A task is already running.');
				return;
			}

			this._newTask({sensor_id: s.id});
		},
		addSensor() {
			if (this.task) {
				alert('A task is already running.');
				return;
			}

			let sensor_name = window.prompt("Enter sensor name:");
			if (sensor_name == null || sensor_name == "") {
				return;
			}

			this._newTask({sensor_name: sensor_name});
		},
	},
});

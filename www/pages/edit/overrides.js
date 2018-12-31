'use strict';
import Component from 'can-component';
import ScheduledOverride from 'models/ScheduledOverride';

Component.extend({
	tag: 'thermo-p-edit-override',
	view: `
	<td>{{dateformat(o.start)}}
	<td>{{dateformat(o.end)}}
	<td>{{o.target_temp}}</td>
	<td class="align-middle text-center"><i class="btn btn-danger fas fa-trash" on:click="o.destroy()"></i></td>
	`,
	helpers: {
		dateformat(date) {
			return new Date(date).toLocaleString();
		}
	},
});

Component.extend({
	tag: 'thermo-p-edit-overrides',
	view: `
<h3 class="text-center">Quick overrides</h3>
<table class="table table-striped table-bordered table-hover table-sm">
<thead>
<tr>
	<th scope="col">Start</th>
	<th scope="col">End</th>
	<th scope="col">Target temp</th>
	<th scope="col">&nbsp;</th>
</tr>
</thead>
<tbody>
<tr>
	<td colspan="2">
		<div class="input-group">
			<input style="text-align: right" class="form-control" type="number" min="15" max="240" step="15" value:bind="newduration" />
			<div class="input-group-append"><div class="input-group-text">min</div></div>
		</div>
	</td>
	<td>
		<div class="input-group">
			<input style="text-align: right" class="form-control" type="number" min="16" max="24" step="0.5" value:bind="newtemp" />
			<div class="input-group-append"><div class="input-group-text">℃</div></div>
		</div>
	</td>
	<td class="align-middle text-center"><i class="btn btn-primary fas fa-plus" on:click="create()"></i></td>
</tr>
	{{#for (o of overrides)}}
	<thermo-p-edit-override style="display: table-row" o:bind="o" />
	{{/for}}
</tbody>
</table>`,
	ViewModel: {
		slider: {},
		newtemp: { default: 20, type: 'number' },
		newduration: { default: 60, type: 'number' },
		overrides: { default: () => new ScheduledOverride.List() },
		connectedCallback(element) {
			var self = this;
			ScheduledOverride.getList({control: this.control.id}).then(function(res) {
				self.overrides = res;
			});
		},
		create() {
			const start = new Date();
			const end = new Date(start.getTime() + this.newduration * 60 * 1000);
			new ScheduledOverride({
				control: this.control.resource_uri,
				start: start.toISOString(),
				end: end.toISOString(),
				target_temp: this.newtemp
			}).save();
		},
	}
});

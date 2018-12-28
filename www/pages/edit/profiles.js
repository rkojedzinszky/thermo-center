'use strict';
import Component from 'can-component';
import DayType from 'models/DayType';
import Profile from 'models/Profile';

Component.extend({
	tag: 'thermo-p-edit-daytype',
	view: `
<h3 class="text-center">Profiles for {{control.name}} / {{daytype.name}}</h3>
<table class="table table-striped table-bordered table-hover table-sm">
<thead>
<tr>
	<th scope="col">Start</th>
	<th scope="col">Target temp</th>
	<th scope="col">&nbsp;</th>
</tr>
</thead>
<tbody>
	{{#for (profile of profiles)}}
<tr class="form-group">
	<td><input class="form-control" type="time" value:bind="profile.start" on:blur="profile.save()"/></td>
	<td><input class="form-control" type="number" step="0.5" value:bind="profile.target_temp_s" on:blur="profile.save()"/></td>
	<td class="align-middle text-center"><i class="form-control btn btn-sm btn-danger fas fa-trash" on:click="profile.destroy()"></i></td>
</tr>
	{{/for}}
<tr class="form-group">
	<td><input class="form-control" type="time" value:bind="newp.start" /></td>
	<td><input class="form-control" type="number" step="0.5" value:bind="newp.target_temp_s" /></td>
	<td class="align-middle text-center"><i class="form-control btn btn-sm btn-primary fas fa-plus" on:click="createNewP()"></i></td>
</tr>
</tbody>
</table>`,
	ViewModel: {
		newp: { Default: Profile },
		profiles: { Default: Profile.List },
		createNewP() {
			var self = this;

			this.newp.daytype = this.daytype.resource_uri;
			this.newp.control = this.control.resource_uri;

			this.newp.save().then(function(newp) {
				self.newp = new Profile();
			});
		},
		connectedCallback(element) {
			var self = this;
			Profile.getList({filter: {control: this.control.resource_uri, daytype: this.daytype.resource_uri}, tastypiefilter: {control: this.control.id, daytype: this.daytype.id}, sort: 'start'}).then(function(res) {
				self.profiles = res;
			});
		}
	},
});

Component.extend({
	tag: 'thermo-p-edit-profiles',
	view: `
	{{#for (daytype of daytypes)}}
		<thermo-p-edit-daytype daytype:from="daytype" control:from="control" />
	{{/for}}
	`,
	ViewModel: {
		daytypes: { Default: DayType.List },
		connectedCallback(element) {
			var self = this;
			DayType.getList().then(function(res) {
				self.daytypes = res;
			});
		}
	}
});

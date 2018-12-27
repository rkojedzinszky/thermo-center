import Component from 'can-component';
import DayType from 'models/DayType';
import Control from 'models/Control';
import Profile from 'models/Profile';
import './overrides';

Component.extend({
	tag: 'thermo-p-edit',
	view: `
<div class="container-fluid">
{{#control}}
	<h3 class="center-block">
		Settings for {{name}}({{sensor_id}})
	</h3>
{{/control}}
	<div class="row">
	{{#if control}}
		<div class="col-sm-4">
			<legend>Quick overrides</legend>
			<thermo-p-edit-overrides control:bind="control" />
		</div>
		<div class="col-sm-4">
			<fieldset {{#if control.isSaving()}}disabled{{/if}}>
			<legend>Pid control loop parameters</legend>
			<div class="form-group form-inline">
				<div class="input-group col-sm-4">
					<div class="input-group-prepend"><div class="input-group-text">Kp</div></div>
					<input type="number" class="form-control" value:bind="control.kp" on:blur="control.save()"/>
				</div>
				<div class="input-group col-sm-4">
					<div class="input-group-prepend"><div class="input-group-text">Ki</div></div>
					<input type="number" class="form-control" value:bind="control.ki" on:blur="control.save()"/>
				</div>
				<div class="input-group col-sm-4">
					<div class="input-group-prepend"><div class="input-group-text">Kd</div></div>
					<input type="number" class="form-control" value:bind="control.kd" on:blur="control.save()"/>
				</div>
			</div>
			</fieldset>
		</div>
	{{/if}}
	</div>
</div>
	`,
	ViewModel: {
		days: { default: () => [] },
		control: { default: null },
		addProfile(day) {
			day.attr('times').push(new Profile({daytype: day, control: this.attr('control'), target_temp: 20}));
		},
		hcSave() {
			this.control.save();
		},
		connectedCallback(element) {
			var self = this;
			var days = this.days;

			Promise.all([Control.findOne({id: this.app.url.id}).then(function(hc) {
				self.control = hc;
				return hc;
			}), DayType.findAll()]).then(function(values) {
				const hc = values[0];
				const dt = values[1];

				console.log(hc);
				console.log(dt);
			});
		}
	}
});

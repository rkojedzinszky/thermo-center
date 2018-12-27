'use strict';
import Component from 'can-component';

Component.extend({
	tag: 'thermo-p-edit-pidsettings',
	view: `
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
</fieldset>`,
});

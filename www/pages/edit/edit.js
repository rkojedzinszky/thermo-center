'use strict';
import Component from 'can-component';
import Control from 'models/Control';
import './overrides';
import './pidsettings';
import './profiles';

Component.extend({
	tag: 'thermo-p-edit',
	view: `
<div class="container-fluid">
{{#control}}
	<h3 class="center-block">
		Settings for {{name}}({{sensor_id}})
	</h3>
{{/control}}
	{{#if control}}
	<div class="row">
		<div class="col-sm-6">
			<thermo-p-edit-overrides control:from="control" />
		</div>
		<div class="col-sm-6">
			<thermo-p-edit-pidsettings control:from="control" />
		</div>
	</div>
	<div class="row">
		<div class="col-sm-3" />
		<div class="col-sm-6">
			<thermo-p-edit-profiles control:from="control" />
		</div>
		<div class="col-sm-3" />
	</div>
	{{/if}}
</div>
	`,
	ViewModel: {
		control: { default: null },
		connectedCallback(element) {
			var self = this;

			Control.findOne({id: this.app.url.id}).then(function(hc) {
				self.control = hc;
			});
		}
	}
});

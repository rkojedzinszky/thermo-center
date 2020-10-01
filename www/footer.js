"use strict";
import Component from "can-component";

Component.extend({
	tag: 'thermo-footer',
	view: `
	<footer class="container border-top">
		<div class="row justify-content-center">
			<!--
			<div class="col-sm-2 text-muted">
				<div>api: {{ app.apiVersion }}</div>
			</div>
			-->
			<div class="col-sm-2 text-muted">
				<div>ui: {{ app.uiVersion }}</div>
			</div>
		</div>
	</footer>
	`,
});

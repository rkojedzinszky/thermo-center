"use strict";
import Component from "can-component";

Component.extend({
	tag: 'thermo-footer',
	view: `
	<footer class="container border-top">
		<div class="row justify-content-center">
			<!--
			<div class="col-sm-2 text-muted text-center">
				<div>api: {{ app.apiVersion }}</div>
			</div>
			-->
			<div class="col-sm-2 text-muted text-center">
				<div>ui: {{ app.uiVersion }}</div>
			</div>
		</div>
		{{#if app.updateAvailable}}
			<div class="row justify-content-center">
				<div class="col-sm-2 text-muted text-center">
					<a href="#" on:click="app.update()">Reload for new version</a>
				</div>
			</div>
		{{/if}}
	</footer>
	`
});

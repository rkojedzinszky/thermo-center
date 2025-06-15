"use strict";
import Component from "can-component";
import DefineMap from "can-define/map/map";

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
		{{#if app.update}}
			<div class="row justify-content-center">
				<div class="col-sm-2 text-muted text-center">
					<button on:click="doReload()">Reload for new version</button>
				</div>
			</div>
		{{/if}}
	</footer>
	`,
	ViewModel: DefineMap.extend({
		doReload() {
			window.location.reload();
		}
	})
});

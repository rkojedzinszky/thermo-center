'use strict';
import realtimeRestModel from 'can-realtime-rest-model';
import {default as ajax, ajaxSetup} from 'can-ajax';
import {default as Cookies} from 'js-cookie';

function sameOrigin(url) {
	const loc = window.location, a = document.createElement('a');
	a.href = url;

	return a.hostname == loc.hostname &&
		a.port == loc.port &&
		a.protocol == loc.protocol;
}

ajaxSetup({
	beforeSend(xhr, settings) {
		if (sameOrigin(settings.url) && !/^(GET|HEAD|OPTIONS|TRACE)$/.test(settings.type)) {
			xhr.setRequestHeader("X-CSRFToken", Cookies.get('csrftoken'));
		}
	}
});

function tastypieRestModel(options)
{
	var endpoint = options['url'];

	options['url'] = {
		getListData: "GET " + endpoint,
		getData: "GET " + endpoint + "{id}/",
		createData: "POST " + endpoint,
		updateData: "PUT " + endpoint + "{id}/",
		destroyData: "DELETE " + endpoint + "{id}/"
	};
	options['parseListProp'] = 'objects';
	options['updateInstanceWithAssignDeep'] = true;

	const connection = realtimeRestModel(options);

	connection.getByResourceUri = function(uri) {
		var self = this;
		return ajax({
			type: 'get',
			url: uri
		}).then(function(data) {
			return self.hydrateInstance(data);
		});
	};
	connection.Map.getByResourceUri = function(uri) {
		return connection.getByResourceUri(uri);
	};

	return connection;
}

export {tastypieRestModel};
export {tastypieRestModel as default};

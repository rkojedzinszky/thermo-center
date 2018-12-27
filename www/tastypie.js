'use strict';
import realtimeRestModel from 'can-realtime-rest-model';
import {default as ajax, ajaxSetup} from 'can-ajax';
import {default as Cookies} from 'js-cookie';
import assign from 'can-assign';

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

function tastypieajax(options)
{
	if (options.type.toUpperCase() == 'GET' && typeof options.data === 'object') {
		if (options.data.hasOwnProperty('filter') || options.data.hasOwnProperty('sort')) {
			const filter = options.data.filter || {};
			const tastypiefilter = options.data.tastypiefilter || {};
			const order_by = options.data.sort;

			let newfilter = {};
			assign(newfilter, filter);
			assign(newfilter, tastypiefilter);
			if (order_by !== undefined) {
				newfilter['order_by'] = order_by;
			}

			options.data = newfilter;
		}
	}

	return ajax(options);
}

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
	options['ajax'] = tastypieajax;

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

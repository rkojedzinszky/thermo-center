'use strict';
import {restModel, realtimeRestModel} from "can";

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

	return realtimeRestModel(options);
}

export {tastypieRestModel};
export {tastypieRestModel as default};

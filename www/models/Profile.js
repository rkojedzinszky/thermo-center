'use strict';
import meta from './g/Profile';
import DefineMap from 'can-define/map/map';
import DefineList from 'can-define/list/list';
import assign from 'can-assign';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
	target_temp_s: {
		get() {
			return this.target_temp;
		},
		set(value) {
			if (value == '') {
				value = null;
			}

			this.target_temp = value;
		}
	}
};
assign(prototype, meta.d);

const Profile = DefineMap.extend('Profile', staticProps, prototype);
Profile.List = DefineList.extend('ProfileList', {'#': Profile});

Profile.connect = tastypieRestModel({
    Map: Profile,
    List: Profile.List,
    url: meta.e,
});

export {Profile, Profile as default};

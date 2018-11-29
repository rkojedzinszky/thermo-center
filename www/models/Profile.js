'use strict';
import meta from './g/Profile';
import {DefineMap, DefineList, Reflect} from 'can';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
Reflect.assign(prototype, meta.d);

const Profile = DefineMap.extend('Profile', staticProps, prototype);
Profile.List = DefineList.extend('ProfileList', {'#': Profile});

Profile.connect = tastypieRestModel({
    Map: Profile,
    List: Profile.List,
    url: meta.e,
});

export {Profile, Profile as default};

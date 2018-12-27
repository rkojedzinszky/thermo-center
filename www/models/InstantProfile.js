'use strict';
import meta from './g/InstantProfile';
import DefineMap from 'can-define/map/map';
import DefineList from 'can-define/list/list';
import assign from 'can-assign';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
assign(prototype, meta.d);

const InstantProfile = DefineMap.extend('InstantProfile', staticProps, prototype);
InstantProfile.List = DefineList.extend('InstantProfileList', {'#': InstantProfile});

InstantProfile.connect = tastypieRestModel({
    Map: InstantProfile,
    List: InstantProfile.List,
    url: meta.e,
});

export {InstantProfile, InstantProfile as default};

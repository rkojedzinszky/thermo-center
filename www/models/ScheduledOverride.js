'use strict';
import meta from './g/ScheduledOverride';
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

const ScheduledOverride = DefineMap.extend('ScheduledOverride', staticProps, prototype);
ScheduledOverride.List = DefineList.extend('ScheduledOverrideList', {'#': ScheduledOverride});

ScheduledOverride.connect = tastypieRestModel({
    Map: ScheduledOverride,
    List: ScheduledOverride.List,
    url: meta.e,
});

export {ScheduledOverride, ScheduledOverride as default};

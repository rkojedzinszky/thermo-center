'use strict';
import meta from './g/ScheduledOverride';
import {DefineMap, DefineList, Reflect} from 'can';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
Reflect.assign(prototype, meta.d);

const ScheduledOverride = DefineMap.extend('ScheduledOverride', staticProps, prototype);
ScheduledOverride.List = DefineList.extend('ScheduledOverrideList', {'#': ScheduledOverride});

ScheduledOverride.connect = tastypieRestModel({
    Map: ScheduledOverride,
    List: ScheduledOverride.List,
    url: meta.e,
});

export {ScheduledOverride, ScheduledOverride as default};

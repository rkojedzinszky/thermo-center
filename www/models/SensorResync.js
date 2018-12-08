'use strict';
import meta from './g/SensorResync';
import DefineMap from 'can-define/map/map';
import DefineList from 'can-define/list/list';
import Reflect from 'can-reflect';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
Reflect.assign(prototype, meta.d);

const SensorResync = DefineMap.extend('SensorResync', staticProps, prototype);
SensorResync.List = DefineList.extend('SensorResyncList', {'#': SensorResync});

SensorResync.connect = tastypieRestModel({
    Map: SensorResync,
    List: SensorResync.List,
    url: meta.e,
});

export {SensorResync, SensorResync as default};

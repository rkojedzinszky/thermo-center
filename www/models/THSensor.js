'use strict';
import meta from './g/THSensor';
import {DefineMap, DefineList, Reflect} from 'can';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
Reflect.assign(prototype, meta.d);

const THSensor = DefineMap.extend('THSensor', staticProps, prototype);
THSensor.List = DefineList.extend('THSensorList', {'#': THSensor});

THSensor.connect = tastypieRestModel({
    Map: THSensor,
    List: THSensor.List,
    url: meta.e,
});

export {THSensor, THSensor as default};

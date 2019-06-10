'use strict';
import meta from './g/ConfigureSensorTask';
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

const ConfigureSensorTask = DefineMap.extend('ConfigureSensorTask', staticProps, prototype);
ConfigureSensorTask.List = DefineList.extend('ConfigureSensorTaskList', {'#': ConfigureSensorTask});

ConfigureSensorTask.connect = tastypieRestModel({
    Map: ConfigureSensorTask,
    List: ConfigureSensorTask.List,
    url: meta.e,
});

export {ConfigureSensorTask, ConfigureSensorTask as default};

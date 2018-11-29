'use strict';
import meta from './g/Control';
import {DefineMap, DefineList, Reflect} from 'can';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
Reflect.assign(prototype, meta.d);

const Control = DefineMap.extend('Control', staticProps, prototype);
Control.List = DefineList.extend('ControlList', {'#': Control});

Control.connect = tastypieRestModel({
    Map: Control,
    List: Control.List,
    url: meta.e,
});

export {Control, Control as default};

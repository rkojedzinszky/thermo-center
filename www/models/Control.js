'use strict';
import meta from './g/Control';
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

const Control = DefineMap.extend('Control', staticProps, prototype);
Control.List = DefineList.extend('ControlList', {'#': Control});

Control.connect = tastypieRestModel({
    Map: Control,
    List: Control.List,
    url: meta.e,
});

export {Control, Control as default};

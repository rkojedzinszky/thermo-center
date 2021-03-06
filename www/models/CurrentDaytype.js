'use strict';
import meta from './g/CurrentDaytype';
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

const CurrentDaytype = DefineMap.extend('CurrentDaytype', staticProps, prototype);
CurrentDaytype.List = DefineList.extend('CurrentDaytypeList', {'#': CurrentDaytype});

CurrentDaytype.connect = tastypieRestModel({
    Map: CurrentDaytype,
    List: CurrentDaytype.List,
    url: meta.e,
});

export {CurrentDaytype, CurrentDaytype as default};

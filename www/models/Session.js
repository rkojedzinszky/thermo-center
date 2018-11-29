'use strict';
import meta from './g/Session';
import {DefineMap, DefineList, Reflect} from 'can';
import {tastypieRestModel} from '../tastypie';

const staticProps = {
    seal: true,
};
const prototype = {
};
Reflect.assign(prototype, meta.d);

const Session = DefineMap.extend('Session', staticProps, prototype);
Session.List = DefineList.extend('SessionList', {'#': Session});

Session.connect = tastypieRestModel({
    Map: Session,
    List: Session.List,
    url: meta.e,
});

export {Session, Session as default};

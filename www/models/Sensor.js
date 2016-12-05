import Model from './model';
import meta from './g/Sensor';
import refresh from 'models/refresh';
export default Model.extend(refresh.static, refresh.dynamic).extend('Models.Sensor', { // static members
    _meta: meta,
}, { // dynamic members
});

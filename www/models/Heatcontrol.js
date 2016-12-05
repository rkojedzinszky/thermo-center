import Model from './model';
import meta from './g/Heatcontrol';
import refresh from 'models/refresh';
export default Model.extend(refresh.static, refresh.dynamic).extend('Models.Heatcontrol', { // static members
    _meta: meta,
}, { // dynamic members
});

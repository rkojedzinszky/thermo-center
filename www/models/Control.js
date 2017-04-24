import Model from './model';
import meta from './g/Control';
import refresh from 'models/refresh';
export default Model.extend(refresh.static, refresh.dynamic).extend('Models.Control', { // static members
    _meta: meta,
}, { // dynamic members
});

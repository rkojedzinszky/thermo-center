import 'can/model/model';
import 'can/construct/super/super';

function _mapNum(func, value)
{
	var r = func(value);
	return isNaN(r) ? null : r;
}

function _mapValue(field, value)
{
	switch (field.type) {
		case 'integer':
			value = _mapNum(parseInt, value);
			break;
		case 'float':
			value = _mapNum(parseFloat, value);
			break;
		case 'decimal':
			value = _mapNum(parseFloat, value);
			break;
		case 'boolean':
			if (value != null)
				value = !!value;
			break;
	}
	return value;
}

export default can.Model.extend({
	removeAttr: true,
	parseModels: "objects",
	findOneByUri(uri, opts) {
		var def = can.Deferred();

		if (uri == null) {
			def.resolveWith(def, [null]);
		} else {
			var self = this;
			can.ajax({
				type: 'GET',
				url: uri
			}).then(function(res) {
				def.resolveWith(def, [self.model(res)]);
			}, function() {
				def.rejectWith(def, arguments);
			});
		}

		return def.promise();
	},

	setup(base, fullName, staticProps, protoProps) {
		var meta = this._meta;
		if (meta === undefined)
			return;

		var resource_uri = meta.endpoint;
		var object_uri = meta.endpoint + '{id}/';
		var objecti_uri = '{resource_uri}';

		staticProps = can.extend({
			findAll: 'GET ' + resource_uri,
			findOne: 'GET ' + object_uri,
			create : 'POST ' + resource_uri,
			update : 'PATCH ' + objecti_uri,
			destroy: 'DELETE ' + objecti_uri
		}, staticProps || { });

		can.Model.setup.call(this, base, fullName, staticProps, protoProps);

		var _d = { };
		can.each(meta.fields, function(field, fname) {
			var f;
			var capname = can.capitalize(fname);
			var getter;
			if (field.type == 'related') {
				if (field.related_type != 'to_one') { // csak to_one mezok
					return;
				}
				field.relates_to = meta.res_to_class[field.relates_to];
				getter = function(opts) {
					return field.relates_to.findOneByUri(this.attr(fname), opts);
				};
			} else {
				getter = function() {
					return this.attr(fname);
				}
			}
			var setter = function(val) {
					return this.attr(fname, val);
			};
			_d['get' + capname] = getter;
			_d['set' + capname] = setter;
		});
		can.Construct._inherit(_d, this.prototype, this.prototype);
	}
}, {
	__convert(prop, value) {
		var field = this.constructor._meta.fields[prop];
		var name = this.constructor._meta.name;
		if (field) {
			if (field.type == 'related') {
				if (field.related_type == 'to_one') {
					if (value != null) {
						if (typeof(value) !== 'string') {
							if (!(value instanceof field.relates_to)) {
								steal.dev.warn(name + "." + prop + "(): argument is not of " + field.relates_to.fullName);
								throw name + "." + prop + "(): argument is not of " + field.relates_to.fullName;
							}
							value = value.attr('resource_uri');
						}
					}
				} else if (field.related_type == 'to_many') {
					var ret = [];
					$.each(value || [], function(idx, el) {
						if (typeof(el) !== 'string') {
							if (!(el instanceof field.relates_to)) {
								steal.dev.warn(name + "." + prop + "(): argument " + idx + " is not of " + field.relates_to.fullName);
								throw name + "." + prop + "(): argument " + idx + " is not of " + field.relates_to.fullName;
							}
							el = el.attr('resource_uri');
						}
						ret.push(el);
					});
					value = ret;
				}
			} else {
				value = _mapValue(field, value);
			}

			// null check
			if (!field.nullable && value == null) {
				steal.dev.warn(name + '.' + prop + ' cannot be null');
				throw name + '.' + prop + ' cannot be null';
			}
		}

		return value;
	},
	_detail() {
		if (this.isNew()) {
			var def = can.Deferred();
			def.resolveWith(def, [this]);
			return def.promise();
		} else {
			return this.constructor.findOneByUri(this.attr('resource_uri'));
		}
	},
	_eq : function(obj) {
		return obj instanceof this.constructor && this.id === obj.id;
	}
});

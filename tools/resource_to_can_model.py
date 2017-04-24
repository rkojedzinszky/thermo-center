#!/usr/bin/env pyton

import os
import sys
import json
import types
import re

os.environ.setdefault("DJANGO_SETTINGS_MODULE", "application.settings")
sys.path.insert(0, os.getcwd())
import django
django.setup()
from django.utils.module_loading import autodiscover_modules
autodiscover_modules('restapi')

from tastypie.fields import NOT_PROVIDED, ToOneField

apiUrl = 'api/v1/'

modelsdir = os.path.dirname(__file__) + '/../' + 'www/models/'
metadir = modelsdir + 'g/'

def _obj_serialize(obj):
    if isinstance(obj, NOT_PROVIDED):
        return ''
    try:
        return str(obj)
    except:
        raise TypeError(repr(obj) + " is not JSON serializable")

def generate_canjs_model(capName, name, resource):
    endpoint = apiUrl + name + '/'
    schema = resource.build_schema()
    schema['endpoint'] = endpoint
    schema['name'] = capName
    referreds = []
    referreds_dict = {}
    field_to_resource = {}
    for f in resource.fields:
        field = resource.fields[f]
        if isinstance(field, ToOneField):
            if isinstance(field.to, str):
                refd = re.search(r'\.(\w+)Resource$', field.to).group(1).lower().capitalize()
            else:
                refd = field.to._meta.resource_name.capitalize()
            if refd not in referreds_dict:
                referreds_dict[refd] = True
                referreds.append(refd)
            field_to_resource[f] = refd
    for res in referreds:
        yield "import %s from '../%s';\n" % (res, res)
    for f in schema['fields']:
        field = schema['fields'][f]
        if field['type'] == 'related' and f in field_to_resource:
            field['relates_to'] = field_to_resource[f]
    yield 'var meta = %s;' % (json.dumps(schema, default=_obj_serialize))
    yield 'meta.res_to_class = {%s};' % (', '.join(['"'+r+'":'+r for r in referreds]))
    yield 'export default meta;'

def generate_canjs_models():
    import glob
    try:
        os.mkdir(metadir)
    except OSError:
        pass
    for f in glob.glob(metadir + '/*.js'):
        os.unlink(f)
    from application.restapi import RestApi
    for cls in RestApi._registry:
        capName = cls.capitalize()
        jsname = metadir + '/' + capName + '.js'
        fh = open(jsname, 'w')
        fh.write('/* %s */\n' % cls)
        for line in generate_canjs_model(capName, cls, RestApi._registry[cls]):
            fh.write(line + '\n')
        modeljsname = modelsdir + capName + '.js'
        if os.path.exists(modeljsname) == False:
            fh = open(modeljsname, 'w')
            fh.write("import Model from './model';\n")
            fh.write("import meta from './g/" + capName + "';\n")
            fh.write("export default Model.extend('Models." + capName + "', { // static members\n")
            fh.write('    _meta: meta\n')
            fh.write('}, { // dynamic members\n')
            fh.write('});\n')
            print ('{} generated, dont forget to add it to the repository'.format(modeljsname))

if __name__ == '__main__':
    generate_canjs_models()

import os
import json
from django.core.management.base import BaseCommand
from django.utils.module_loading import autodiscover_modules
from django.conf import settings

apiUrl = 'api/v1/'
modelsdir = os.path.join(settings.BASE_DIR, 'www', 'models')
metadir = os.path.join(modelsdir, 'g')

def _prepare_metadir():
    import glob
    try:
        os.mkdir(metadir)
    except OSError:
        pass
    for f in glob.glob(os.path.join(metadir, '*.js')):
        os.unlink(f)

_TYPE_MAP = {
        'float': 'number',
        'integer': 'number',
        'boolean': 'boolean',
        }

def _get_type(typ_):
    return _TYPE_MAP.get(typ_, 'string')

def _gen_meta(name, cls, endpoint):
    fname = os.path.join(metadir, '{}.js'.format(name))
    schema = cls.build_schema()['fields']
    can_dyn = {}
    for name, desc in schema.items():
        typ_ = _get_type(desc['type'])
        can_dyn[name] = {
                'type': typ_,
                }
    meta = {
            'd': can_dyn,
            'e': endpoint
            }

    content = """'use strict';
const meta = {};
export {{meta as default}};
""".format(json.dumps(meta))
    with open(fname, 'w') as fh:
        fh.write(content)

def _gen_model(name, force=False):
    fname = os.path.join(modelsdir, '{}.js'.format(name))
    if os.path.exists(fname) and not force:
        return

    content = """'use strict';
import meta from './g/{name}';
import DefineMap from 'can-define/map/map';
import DefineList from 'can-define/list/list';
import assign from 'can-assign';
import {{tastypieRestModel}} from '../tastypie';

const staticProps = {{
    seal: true,
}};
const prototype = {{
}};
assign(prototype, meta.d);

const {name} = DefineMap.extend('{name}', staticProps, prototype);
{name}.List = DefineList.extend('{name}List', {{'#': {name}}});

{name}.connect = tastypieRestModel({{
    Map: {name},
    List: {name}.List,
    url: meta.e,
}});

export {{{name}, {name} as default}};
""".format(name=name)
    with open(fname, 'w') as fh:
        fh.write(content)

class Command(BaseCommand):
    help = "Generate CanJS models"

    def add_arguments(self, parser):
        parser.add_argument('-f', '--force', action='store_true', help='Force regeneration of all model files')

    def handle(self, *args, **options):
        force = options['force']

        _prepare_metadir()

        autodiscover_modules('restapi')

        from application.restapi import RestApi
        for name, cls in RestApi._registry.items():
            rname = cls.__class__.__name__.replace('Resource', '')
            endpoint = RestApi._build_reverse_url("api_dispatch_list", kwargs={
                'api_name': RestApi.api_name,
                'resource_name': name,
                })
            _gen_meta(rname, cls, endpoint)
            _gen_model(rname, force)


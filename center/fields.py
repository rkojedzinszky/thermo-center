from django.db import models

class RangedIntegerField(models.IntegerField):
    """ A field to represent an integer with min-max ranges """

    class InvalidIdException(RuntimeError):
        pass

    def __init__(self, min_value=None, max_value=None, **kwargs):
        super(RangedIntegerField, self).__init__(**kwargs)
        self._min_value = min_value
        self._max_value = max_value

    def deconstruct(self):
        name, path, args, kwargs = super(RangedIntegerField, self).deconstruct()
        kwargs['min_value'] = self._min_value
        kwargs['max_value'] = self._max_value
        return name, path, args, kwargs

    def _check_ranges(self, value):
        if self._min_value is not None and value is not None and value < self._min_value:
            raise RangedIntegerField.InvalidIdException()

        if self._max_value is not None and value is not None and value > self._max_value:
            raise RangedIntegerField.InvalidIdException()

        return value

    def get_prep_value(self, value):
        return self._check_ranges(super(RangedIntegerField, self).get_prep_value(value))

    def to_python(self, value):
        return self._check_ranges(super(RangedIntegerField, self).to_python(value))

class SensorIdField(RangedIntegerField):
    """ A field to represent a sensor id """

    def __init__(self, **kwargs):
        super(SensorIdField, self).__init__(min_value=1, max_value=127, **kwargs)

    def deconstruct(self):
        name, path, args, kwargs = super(SensorIdField, self).deconstruct()
        del kwargs['min_value']
        del kwargs['max_value']
        return name, path, args, kwargs

import datetime
from django.core.management.base import BaseCommand, CommandError
from heatcontrol import models

class Command(BaseCommand):
    help = 'Regenerate calendar and day types'

    def handle(self, *args, **options):
        models.Calendar.objects.all().delete()
        models.DayType.objects.all().delete()

        w = models.DayType.objects.create(name='workday')
        sat = models.DayType.objects.create(name='Saturday')
        sun = models.DayType.objects.create(name='Sunday')

        st = datetime.date.today()
        for i in range(10000):
            c = models.Calendar(day=st)
            if st.weekday() < 5:
                c.daytype = w
            elif st.weekday() == 5:
                c.daytype = sat
            else:
                c.daytype = sun
            c.save()

            st += datetime.timedelta(days=1)

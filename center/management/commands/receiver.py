from django.core.management.base import BaseCommand

class Command(BaseCommand):
    help = "Start the receiver daemon"

    def add_arguments(self, parser):
        parser.add_argument('-d', '--daemonize', action='store_true', help='Run in the background')

    def handle(self, *args, **options):
        from center.receiver import Main

        Main().run(daemonize=options['daemonize'])

from django.core.management.base import BaseCommand

class Command(BaseCommand):
    help = "Start the receiver daemon"

    def handle(self, *args, **kwargs):
        from center.receiver import Main

        Main().run()

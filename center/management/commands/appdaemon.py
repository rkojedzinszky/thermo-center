import os
import sys
import asyncio
from django.core.management.base import BaseCommand

class Command(BaseCommand):
    help = "Start the application daemon (receiver and websocket)"

    def add_arguments(self, parser):
        parser.add_argument('-d', '--daemonize', action='store_true', help='Run in the background')

    def handle(self, *args, **options):
        if options['daemonize']:
            if os.fork() > 0:
                sys.exit()

            os.chdir('/')
            os.setsid()

            if os.fork() > 0:
                sys.exit()

            with open('/dev/null', 'r') as fh:
                os.dup2(fh.fileno(), sys.stdin.fileno())
            with open('/dev/null', 'a+') as fh:
                os.dup2(fh.fileno(), sys.stdout.fileno())
                os.dup2(fh.fileno(), sys.stderr.fileno())

        loop = asyncio.get_event_loop()

        from center import websocket
        websocketDaemon = websocket.Main(loop)
        wstask = loop.create_task(websocketDaemon.run())

        from center import receiver
        receiverDaemon = receiver.Main(loop)
        recvtask = loop.create_task(receiverDaemon.run())

        tasks = {wstask, recvtask}

        loop.run_until_complete(asyncio.wait(tasks, return_when=asyncio.FIRST_COMPLETED))

        for t in tasks:
            t.cancel()

        loop.run_until_complete(asyncio.wait(tasks))

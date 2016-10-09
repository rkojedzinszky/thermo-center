import ipaddress
from django.contrib.auth import get_user_model
from django.contrib.auth.backends import ModelBackend

class Backend(ModelBackend):
    def authenticate(self, username, password, remote_addr=None):
        if remote_addr is not None:
            user = get_user_model().objects.filter(username=username, is_active=True).first()

            if user is not None:
                for n in user.networkauth_set.all():
                    p = ipaddress.ip_network(unicode(n.address))
                    if remote_addr in p:
                        return user

        return super(Backend, self).authenticate(username=username, password=password)

import ipaddress
from django.contrib.auth import get_user_model
from django.contrib.auth.backends import ModelBackend

class Backend(ModelBackend):
    def authenticate(self, request, username, password):
        remote_addr = request.META.get('REMOTE_ADDR', None)

        if remote_addr:
            remote_addr = ipaddress.ip_address(remote_addr)
            user = get_user_model().objects.filter(username=username, is_active=True).first()

            if user is not None:
                for n in user.networkauth_set.all():
                    p = ipaddress.ip_network(n.address)
                    if remote_addr in p:
                        return user

        return super(Backend, self).authenticate(request=request, username=username, password=password)

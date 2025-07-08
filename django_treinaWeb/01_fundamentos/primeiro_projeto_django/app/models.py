from django.db import models


class Client(models.Model):
    name = models.CharField(max_length=100)
    age = models.IntegerField(null=True, blank=True)
    is_active = models.BooleanField(default=True)

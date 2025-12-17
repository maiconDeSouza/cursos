from django.db import models


# Create your models here.
class EnterpriseModel(models.Model):
    name = models.CharField(max_length=255)
    cnpj = models.CharField(max_length=18, unique=True)
    active = models.BooleanField(default=False)

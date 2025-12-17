from django.contrib import admin

from .models import EnterpriseModel


# Register your models here.
@admin.register(EnterpriseModel)
class EnterpriseAdmin(admin.ModelAdmin):
    list_display = ('name', 'cnpj', 'active')
    list_filter = ('active',)
    search_fields = ('name', 'cnpj')

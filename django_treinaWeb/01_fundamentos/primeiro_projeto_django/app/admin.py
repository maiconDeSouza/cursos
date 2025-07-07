from django.contrib import admin

from .models import Client

# Register your models here.


@admin.register(Client)
class ClientAdmin(admin.ModelAdmin):
    fields = ('name', 'age', 'date_of_birth', 'is_active',)
    search_fields = ('name',)

from django.contrib import admin
from django.urls import path, include


urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('apps.base.urls', namespace='base')),
    path('recipes/', include('apps.recipes.urls', namespace='recipes')),
]

from django.urls import path

from . import views

app_name = 'recipes'
urlpatterns = [
    path('', views.recipes, name='recipes_list'),
]

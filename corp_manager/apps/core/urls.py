from django.urls import path

from . import views

app_name = 'home'

urlpatterns = [
    path('', views.Home.as_view(), name='home'),
    path('empresa/nova/', views.FormView.as_view(), name='enterprise-form'),
]

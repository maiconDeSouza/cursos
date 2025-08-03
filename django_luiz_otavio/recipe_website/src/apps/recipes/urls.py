from django.urls import path

from . import views

app_name = 'recipes'
urlpatterns = [
    path('', views.home, name='home'),
    path('recipes/', views.recipes, name='recipes_list'),
    path(
        'recipe/<slug:slug>', views.RecipeDetail.as_view(), name='recipe_detail'
    ),
]

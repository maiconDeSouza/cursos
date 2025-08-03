from django.urls import path

from . import views

app_name = 'recipes'
urlpatterns = [
    path('', views.Home.as_view(), name='home'),
    path('recipes/', views.Recipes.as_view(), name='recipes_list'),
    path(
        'category/<str:category_name>',
        views.Category.as_view(),
        name='category',
    ),
    path(
        'recipe/<slug:slug>', views.RecipeDetail.as_view(), name='recipe_detail'
    ),
]

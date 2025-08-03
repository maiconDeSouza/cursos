from django.shortcuts import render, redirect
from django.views import View

from .models import Recipe

# Create your views here.


def home(request):
    return redirect('recipes:recipes_list')


def recipes(request):
    recipes_all = Recipe.objects.filter(is_published=True).order_by(
        '-created_at'
    )
    context = {'recipes': recipes_all}
    return render(request, 'recipes/pages/home.html', context)


class RecipeDetail(View):
    def get(self, request, slug):
        return render(request, 'recipes/pages/main-recipe-detail.html')

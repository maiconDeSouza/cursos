from django.shortcuts import (
    render,
    redirect,
    get_object_or_404,
    get_list_or_404,
)
from django.views import View

from .models import Recipe

# Create your views here.


class Home(View):
    def get(self, request):
        return redirect('recipes:recipes_list')


class Recipes(View):
    def get(self, request):
        recipes_published_all = Recipe.objects.filter(
            is_published=True
        ).order_by('-created_at')

        context = {'recipes': recipes_published_all}

        return render(request, 'recipes/pages/home.html', context)


class Category(View):
    def get(self, request, category_name):
        recipes_category = get_list_or_404(
            Recipe.objects.order_by('-created_at'),
            category__name=category_name,
            is_published=True,
        )

        context = {
            'recipes': recipes_category,
            'is_category': True,
            'category': category_name,
        }

        return render(request, 'recipes/pages/home.html', context)


class RecipeDetail(View):
    def get(self, request, slug):
        recipe = get_object_or_404(Recipe, slug=slug)

        context = {'recipe': recipe}

        return render(request, 'recipes/pages/main-recipe-detail.html', context)

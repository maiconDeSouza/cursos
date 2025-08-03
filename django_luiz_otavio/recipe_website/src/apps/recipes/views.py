from django.shortcuts import render, redirect
from django.views import View


# Create your views here.


def home(request):
    return redirect('recipes:recipes_list')


def recipes(request):
    return render(request, 'recipes/pages/home.html')


class RecipeDetail(View):
    def get(self, request, slug):
        return render(request, 'recipes/pages/main-recipe-detail.html')

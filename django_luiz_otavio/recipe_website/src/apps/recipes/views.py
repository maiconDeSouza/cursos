from django.shortcuts import render, redirect


# Create your views here.


def home(request):
    return redirect('recipes:recipes_list')


def recipes(request):
    return render(request, 'recipes/pages/home.html')

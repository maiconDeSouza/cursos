from django.shortcuts import HttpResponse


# Create your views here.
def index(request):
    return HttpResponse('Bem-vindo ao meu primeiro projeto Django!')

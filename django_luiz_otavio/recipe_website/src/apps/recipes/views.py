from django.shortcuts import HttpResponse


# Create your views here.
def recipes(request):
    return HttpResponse('Hello, Django!')

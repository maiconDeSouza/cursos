from django.shortcuts import render, HttpResponse
import datetime

# Create your views here.


def home(request):
    time_current = datetime.datetime.now()
    return HttpResponse(time_current)

from django.shortcuts import render
import datetime

# Create your views here.


def home(request):
    time_current = datetime.datetime.now()
    return render(request, 'app/home.html', {'time_current': time_current})

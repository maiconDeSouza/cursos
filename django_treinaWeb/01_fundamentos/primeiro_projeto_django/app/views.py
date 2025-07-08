from django.shortcuts import render, redirect
import datetime

from . import forms
from .models import Client


def home(request):
    time_current = datetime.datetime.now()
    return render(request, 'app/home.html', {'time_current': time_current})


def register(resquest):
    if resquest.method == 'POST':
        form_data = forms.FormsClient(resquest.POST)
        if form_data.is_valid():
            print('Formulário Válido')
            name = resquest.POST.get('name')
            age = resquest.POST.get('age')
            new_cliente = Client(
                name=name, age=age, is_active=True)
            new_cliente.save()
            return redirect('home')
        else:
            form = forms.FormsClient()
            return render(resquest, 'app/register.html', {'form': form, 'is_invalid': True})

    form = forms.FormsClient()
    return render(resquest, 'app/register.html', {'form': form})

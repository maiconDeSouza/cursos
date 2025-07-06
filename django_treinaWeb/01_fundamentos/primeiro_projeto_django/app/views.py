from django.shortcuts import render, redirect
import datetime

from . import forms as forms_client


def home(request):
    time_current = datetime.datetime.now()
    return render(request, 'app/home.html', {'time_current': time_current})


def register(resquest):
    if resquest.method == 'POST':
        form_data = forms_client.FormsClient(resquest.POST)
        if form_data.is_valid():
            print('Formulário Válido')
            return redirect('home')
        else:
            form = forms_client.FormsClient()
            return render(resquest, 'app/register.html', {'form': form, 'is_invalid': True})

    form = forms_client.FormsClient()
    return render(resquest, 'app/register.html', {'form': form})

from django.shortcuts import render
from django.views import View

from .forms import EnterpriseForm


# Create your views here.
class Home(View):
    def get(self, request):
        context = {
            'status': 'online2',
            'sistema': 'Corporativo',
        }
        return render(request, 'core/pages/index.html', context)


class FormView(View):
    def get(self, request):
        form = EnterpriseForm()

        context = {
            'form': form,
        }

        return render(request, 'core/pages/enterprise-form.html', context)

    def post(self, request):
        form = EnterpriseForm(request.POST)

        if form.is_valid():
            form.save()

        context = {
            'form': form,
        }

        return render(request, 'core/pages/enterprise-form.html', context)

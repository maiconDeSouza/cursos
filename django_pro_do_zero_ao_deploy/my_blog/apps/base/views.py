from django.shortcuts import render
from django.views import View


# Create your views here.
class Home(View):
    def get(self, resquest):
        return render(resquest, 'base/pages/index.html')

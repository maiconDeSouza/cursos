from django.shortcuts import render


# Create your views here.
def home(request):
    context = {
        'status': 'online2',
        'sistema': 'Corporativo',
    }
    return render(request, 'core/pages/index.html', context)

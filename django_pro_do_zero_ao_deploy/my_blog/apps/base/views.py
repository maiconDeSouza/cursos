from django.shortcuts import render
from django.views import View

from .models import Post


# Create your views here.
class Home(View):
    def get(self, resquest):
        posts = Post.objects.order_by('-created_at')[:3]
        context = {'posts': posts}
        return render(resquest, 'base/pages/index.html', context)

from django.contrib import admin

from .models import Post, Comment


# Register your models here.
class PostAdmin(admin.ModelAdmin):
    list_display = ('title', 'content', 'author')
    search_fields = ('title', 'author')


class CommentAdmin(admin.ModelAdmin):
    list_display = ('title', 'author')
    search_fields = ('title', 'author')


admin.site.register(Post, PostAdmin)
admin.site.register(Comment, CommentAdmin)

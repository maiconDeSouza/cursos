from django.db import models
from django.contrib.auth.models import User
from django.utils.text import slugify


class Categorie(models.Model):
    name = models.CharField(max_length=65, unique=True)

    def __str__(self):
        return self.name


class RecipeQuerySet(models.QuerySet):
    def recipes_published_order_by(self):
        return self.filter(is_published=True).order_by('-created_at')

    def category_published_order_by(self, category_name):
        return self.filter(
            category__name=category_name, is_published=True
        ).order_by('-created_at')

    def by_slug_published_order_by(self, slug):
        return self.filter(slug=slug, is_published=True)


class RecipePublishedManager(models.Manager):
    def get_queryset(self):
        return RecipeQuerySet(self.model, using=self._db)

    def recipes(self):
        return self.get_queryset().recipes_published_order_by()

    def category(self, category_name):
        return self.get_queryset().category_published_order_by(category_name)

    def by_slug(self, slug):
        return self.get_queryset().by_slug_published_order_by(slug)


# Create your models here.
class Recipe(models.Model):
    title = models.CharField(max_length=100)
    slug = models.SlugField(unique=True, blank=True)
    description = models.CharField(max_length=165)
    preparetion_time = models.IntegerField()
    preparetion_unit = models.CharField(max_length=65)
    servings = models.IntegerField()
    servings_unit = models.CharField(max_length=65)
    preparetion_steps = models.TextField()
    preparetion_steps_is_html = models.BooleanField(default=False)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    is_published = models.BooleanField(default=False)
    cover = models.ImageField(upload_to='recipes/covers/%Y/%M/%d/')
    author = models.ForeignKey(User, on_delete=models.SET_NULL, null=True)
    category = models.ForeignKey(
        Categorie, on_delete=models.SET_NULL, null=True
    )

    def __str__(self):
        return self.title

    published = RecipePublishedManager()

    def save(self, *args, **kwargs):
        if not self.slug:  # só cria se estiver vazio
            base_slug = slugify(self.title)
            slug = base_slug
            counter = 1
            while Recipe.objects.filter(slug=slug).exists():
                slug = f'{base_slug}-{counter}'
                counter += 1
            self.slug = slug
        super().save(*args, **kwargs)

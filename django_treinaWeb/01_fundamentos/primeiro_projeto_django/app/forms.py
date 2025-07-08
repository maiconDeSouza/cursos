from django import forms


class FormsClient(forms.Form):
    name = forms.CharField(label='Nome do Cliente', max_length=200)
    age = forms.IntegerField(label='Digite sua idade')
    is_active = forms.BooleanField()

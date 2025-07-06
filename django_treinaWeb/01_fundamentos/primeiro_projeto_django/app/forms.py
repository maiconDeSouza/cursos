from django import forms


class FormsClient(forms.Form):
    name = forms.CharField(label='Nome do Cliente', max_length=200)
    age = forms.IntegerField(label='Digite sua idade')
    date_of_birth = forms.DateField(label='Data de nascimento')
    is_active = forms.BooleanField()

from django import forms
from .models import EnterpriseModel


class EnterpriseForm(forms.ModelForm):
    class Meta:
        model = EnterpriseModel
        fields = ['name', 'cnpj', 'active']

    def clean_cnpj(self):
        cnpj = self.cleaned_data['cnpj']

        if len(cnpj) < 14:
            raise forms.ValidationError('CNPJ inválido')

        return cnpj

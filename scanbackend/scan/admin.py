from django.contrib import admin
from .models import *

# Register your models here.

@admin.register(ScanGroup)
class ScanGroupAdmin(admin.ModelAdmin):
	prepopulated_fields = {"slug": ("name",)}



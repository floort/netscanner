from django.db import models

# Create your models here.

class ScanGroup(models.Model):
	name = models.CharField(max_length=256)
	slug = models.SlugField(unique=True)
	description = models.TextField()
	ranges = models.ManyToManyField('CIDRRange')

	def __unicode__(self):
		return name


class CIDRRange(models.Model):
	cidr = models.CharField(max_length=256, indexed=True, Unique=True)

	def __unicode__(self):
		return self.cidr


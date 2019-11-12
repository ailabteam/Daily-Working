from .models import User, Block
from rest_framework import serializers


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('username', 'name', 'public_key', 'private_key', 'age', 'description', 'is_manage', 'password')

    def create(self, validated_data):
        return User.objects.create_user(**validated_data)

    def update(self, instance, validated_data):
        print(validated_data)
        instance.password = validated_data.get('password', instance.password)
        instance.name = validated_data.get('name', instance.name)
        instance.age = validated_data.get('age', instance.name)
        instance.description = validated_data.get('description', instance.description)
        instance.is_manage = validated_data.get('is_manage', instance.is_manage)
        instance.set_password(instance.password)
        print(instance.password, instance.name, instance.age, instance.description)
        instance.save()
        return instance


class BlockSerializer(serializers.ModelSerializer):
    class Meta:
        model = Block
        fields = ('block_id', 'public_key', 'private_key')

    def create(self, validated_data):
        return Block.objects.create(**validated_data)

    def update(self, instance, validated_data):
        instance.block_id = validated_data.get('block_id', instance.block_id)
        instance.public_key = validated_data.get('public_key', instance.public_key)
        instance.private_key = validated_data.get('private_key', instance.private_key)
        return instance

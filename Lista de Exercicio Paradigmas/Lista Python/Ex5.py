class Animal:
    def som(self):
        raise NotImplementedError("Este método deve ser implementado por subclasses.")

class Cachorro(Animal):
    def som(self):
        return "Au Au!"

class Gato(Animal):
    def som(self):
        return "Miau!"

def fazer_som_animais(animais):
    for animal in animais:
        print(f'O animal faz: {animal.som()}')

animais = [Cachorro(), Gato(), Cachorro(), Gato()]

fazer_som_animais(animais)
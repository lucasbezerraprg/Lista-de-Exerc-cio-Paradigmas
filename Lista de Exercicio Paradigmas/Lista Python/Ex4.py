class Animal:
    def som(self):
        raise NotImplementedError("Este método deve ser implementado por subclasses.")

class Cachorro(Animal):
    def som(self):
        return "Au Au!"

class Gato(Animal):
    def som(self):
        return "Miau!"

def fazer_som(animal):
    print(f'O animal faz: {animal.som()}')

cachorro = Cachorro()
gato = Gato()

fazer_som(cachorro)
fazer_som(gato)

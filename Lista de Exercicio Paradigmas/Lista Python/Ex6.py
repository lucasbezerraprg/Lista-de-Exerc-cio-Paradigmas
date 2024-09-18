class Motor:
    def __init__(self, potencia, tipo):
        self.potencia = potencia  
        self.tipo = tipo  

    def detalhes(self):
        return f'Motor: {self.potencia} CV, Tipo: {self.tipo}'

class Carro:
    def __init__(self, marca, modelo, ano, motor):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.motor = motor  

    def detalhes(self):
        return f'Carro: {self.marca} {self.modelo}, Ano: {self.ano}, {self.motor.detalhes()}'

# Exemplo de uso
motor1 = Motor(150, 'Gasolina')
carro1 = Carro('Toyota', 'Corolla', 2020, motor1)

motor2 = Motor(120, 'Etanol')
carro2 = Carro('Honda', 'Civic', 2019, motor2)

print(carro1.detalhes())
print(carro2.detalhes())
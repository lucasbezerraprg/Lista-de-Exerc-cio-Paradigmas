class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0  

    def detalhes(self):
        return f'Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}'

    def acelerar(self, incremento):
        self.velocidade += incremento
        print(f'O carro acelerou para {self.velocidade} km/h.')

    def frear(self, decremento):
        self.velocidade -= decremento
        if self.velocidade < 0:
            self.velocidade = 0  
        print(f'O carro freou para {self.velocidade} km/h.')

    def exibir_velocidade(self):
        print(f'A velocidade atual é {self.velocidade} km/h.')

carro1 = Carro('Toyota', 'Corolla', 2020)
carro2 = Carro('Honda', 'Civic', 2019)
carro3 = Carro('Ford', 'Fiesta', 2021)

print(carro1.detalhes())
print(carro2.detalhes())
print(carro3.detalhes())

carro1.acelerar(50)
carro1.exibir_velocidade()
carro1.frear(20)
carro1.exibir_velocidade()
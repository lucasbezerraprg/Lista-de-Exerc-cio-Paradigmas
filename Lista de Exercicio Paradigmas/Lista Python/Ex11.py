from abc import ABC, abstractmethod

# Classe abstrata Funcionario
class Funcionario(ABC):
    @abstractmethod
    def calcularSalario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, horas_trabalhadas, valor_hora):
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_hora = valor_hora

    def calcularSalario(self):
        return self.horas_trabalhadas * self.valor_hora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, salario_mensal):
        self.salario_mensal = salario_mensal

    def calcularSalario(self):
        return self.salario_mensal

horista = FuncionarioHorista(160, 20) 
assalariado = FuncionarioAssalariado(3000)  
print(f'Salário do Funcionário Horista: R$ {horista.calcularSalario()}')
print(f'Salário do Funcionário Assalariado: R$ {assalariado.calcularSalario()}')

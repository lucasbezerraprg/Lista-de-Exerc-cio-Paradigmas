class SaldoInsuficienteException(Exception):
    def __init__(self, mensagem):
        super().__init__(mensagem)

class ContaBancaria:
    def __init__(self, saldo_inicial=0):
        self.saldo = saldo_inicial

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException("Saldo insuficiente para realizar o saque.")
        self.saldo -= valor
        return self.saldo

    def depositar(self, valor):
        self.saldo += valor
        return self.saldo

# Exemplo de uso
conta = ContaBancaria(100)

try:
    conta.sacar(150)
except SaldoInsuficienteException as e:
    print(e)  
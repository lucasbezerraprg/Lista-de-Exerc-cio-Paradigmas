class ContaBancaria:
    def __init__(self, titular):
        self.titular = titular
        self.__saldo = 0  

    def depositar(self, valor):
        if valor > 0:
            self.__saldo += valor
            print(f'Depósito de R$ {valor:.2f} realizado. Saldo atual: R$ {self.__saldo:.2f}.')
        else:
            print('Valor de depósito inválido.')

    def sacar(self, valor):
        if 0 < valor <= self.__saldo:
            self.__saldo -= valor
            print(f'Saque de R$ {valor:.2f} realizado. Saldo atual: R$ {self.__saldo:.2f}.')
        elif valor > self.__saldo:
            print('Saldo insuficiente para saque.')
        else:
            print('Valor de saque inválido.')

    def exibir_saldo(self):
        print(f'Saldo atual: R$ {self.__saldo:.2f}.')

conta1 = ContaBancaria('João Silva')

conta1.depositar(100)
conta1.exibir_saldo()

conta1.sacar(30)
conta1.exibir_saldo()

conta1.sacar(80)
conta1.exibir_saldo()
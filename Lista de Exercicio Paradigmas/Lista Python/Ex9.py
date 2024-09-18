from abc import ABC, abstractmethod

class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):
    def __init__(self, conteudo):
        self.conteudo = conteudo

    def imprimir(self):
        print(f"Imprimindo Relatório:\n{self.conteudo}")

class Contrato(Imprimivel):
    def __init__(self, partes):
        self.partes = partes

    def imprimir(self):
        print("Imprimindo Contrato:")
        for parte in self.partes:
            print(f"- {parte}")

# Exemplo de uso
relatorio = Relatorio("Este é um relatório de vendas.")
contrato = Contrato(["Parte A: João", "Parte B: Maria"])

relatorio.imprimir()
contrato.imprimir()
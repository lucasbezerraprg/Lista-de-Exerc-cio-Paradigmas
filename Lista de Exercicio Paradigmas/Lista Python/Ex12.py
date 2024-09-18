class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, outro):
        if isinstance(outro, Produto):
            novo_nome = f"{self.nome} + {outro.nome}"
            novo_preco = self.preco + outro.preco
            return Produto(novo_nome, novo_preco)
        return NotImplemented

    def __repr__(self):
        return f"Produto(nome='{self.nome}', preco={self.preco:.2f})"

# Exemplo de uso
produto1 = Produto("Produto A", 20.50)
produto2 = Produto("Produto B", 15.75)

produto_soma = produto1 + produto2
print(produto_soma)
class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []  

    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self)  

    def listar_escolas(self):
        return [escola.nome for escola in self.escolas]

class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []  

    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)

    def listar_professores(self):
        return [professor.nome for professor in self.professores]

escola1 = Escola('Escola Primária')
escola2 = Escola('Escola Secundária')

professor1 = Professor('Carlos Silva')
professor2 = Professor('Maria Oliveira')

professor1.adicionar_escola(escola1)
professor1.adicionar_escola(escola2)
professor2.adicionar_escola(escola1)

print(f'Professores da {escola1.nome}: {escola1.listar_professores()}')
print(f'Escolas que o {professor1.nome} leciona: {professor1.listar_escolas()}')
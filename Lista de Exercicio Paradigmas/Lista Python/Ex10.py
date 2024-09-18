class Calculadora:
    def somar_dois(self, a, b):
        return a + b

    def somar_tres(self, a, b, c):
        return a + b + c

# Exemplo de uso
calc = Calculadora()

resultado1 = calc.somar_dois(5, 10)
resultado2 = calc.somar_tres(5, 10, 15)

print(f'Soma de dois números: {resultado1}')
print(f'Soma de três números: {resultado2}')
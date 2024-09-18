class Configuracao:
    _instancia = None

    def __new__(cls, *args, **kwargs):
        if not cls._instancia:
            cls._instancia = super(Configuracao, cls).__new__(cls)
            cls._instancia.configuracoes = {}
        return cls._instancia

    def set(self, chave, valor):
        self.configuracoes[chave] = valor

    def get(self, chave):
        return self.configuracoes.get(chave, None)

config1 = Configuracao()
config1.set("url", "https://exemplo.com")

config2 = Configuracao()
print(config2.get("url"))  

print(config1 is config2)  
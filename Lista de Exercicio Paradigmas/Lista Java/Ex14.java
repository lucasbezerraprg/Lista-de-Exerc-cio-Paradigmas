import java.util.HashMap;
import java.util.Map;

class Configuracao {
    private static Configuracao instancia = null;
    private Map<String, String> configuracoes;

    private Configuracao() {
        configuracoes = new HashMap<>();
    }

    public static Configuracao getInstancia() {
        if (instancia == null) {
            instancia = new Configuracao();
        }
        return instancia;
    }

    public void set(String chave, String valor) {
        configuracoes.put(chave, valor);
    }

    public String get(String chave) {
        return configuracoes.get(chave);
    }

    public static void main(String[] args) {
        Configuracao config1 = Configuracao.getInstancia();
        config1.set("url", "https://exemplo.com");

        Configuracao config2 = Configuracao.getInstancia();
        System.out.println(config2.get("url"));

        System.out.println(config1 == config2);
    }
}

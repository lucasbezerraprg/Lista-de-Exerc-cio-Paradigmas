import java.util.List;

abstract class Imprimivel {
    abstract void imprimir();
}

class Relatorio extends Imprimivel {
    private String conteudo;

    public Relatorio(String conteudo) {
        this.conteudo = conteudo;
    }

    @Override
    void imprimir() {
        System.out.println("Imprimindo Relatório:\n" + this.conteudo);
    }
}

class Contrato extends Imprimivel {
    private List<String> partes;

    public Contrato(List<String> partes) {
        this.partes = partes;
    }

    @Override
    void imprimir() {
        System.out.println("Imprimindo Contrato:");
        for (String parte : partes) {
            System.out.println("- " + parte);
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Relatorio relatorio = new Relatorio("Este é um relatório de vendas.");
        Contrato contrato = new Contrato(List.of("Parte A: João", "Parte B: Maria"));

        relatorio.imprimir();
        contrato.imprimir();
    }
}

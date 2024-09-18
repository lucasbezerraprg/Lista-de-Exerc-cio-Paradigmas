import java.util.ArrayList;
import java.util.List;

class Professor {
    private String nome;
    private List<Escola> escolas;

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this);
        }
    }

    public List<String> listarEscolas() {
        List<String> nomesEscolas = new ArrayList<>();
        for (Escola escola : escolas) {
            nomesEscolas.add(escola.getNome());
        }
        return nomesEscolas;
    }

    public String getNome() {
        return nome;
    }
}

class Escola {
    private String nome;
    private List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
        }
    }

    public List<String> listarProfessores() {
        List<String> nomesProfessores = new ArrayList<>();
        for (Professor professor : professores) {
            nomesProfessores.add(professor.getNome());
        }
        return nomesProfessores;
    }

    public String getNome() {
        return nome;
    }
}

public class Main {
    public static void main(String[] args) {
        Escola escola1 = new Escola("Escola Primária");
        Escola escola2 = new Escola("Escola Secundária");

        Professor professor1 = new Professor("Carlos Silva");
        Professor professor2 = new Professor("Maria Oliveira");

        professor1.adicionarEscola(escola1);
        professor1.adicionarEscola(escola2);
        professor2.adicionarEscola(escola1);

        System.out.println("Professores da " + escola1.getNome() + ": " + escola1.listarProfessores());
        System.out.println("Escolas que o " + professor1.getNome() + " leciona: " + professor1.listarEscolas());
    }
}

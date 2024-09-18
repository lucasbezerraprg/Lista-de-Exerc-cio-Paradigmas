abstract class Funcionario {
    abstract double calcularSalario();
}

class FuncionarioHorista extends Funcionario {
    private int horasTrabalhadas;
    private double valorHora;

    public FuncionarioHorista(int horasTrabalhadas, double valorHora) {
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorHora = valorHora;
    }

    @Override
    double calcularSalario() {
        return horasTrabalhadas * valorHora;
    }
}

class FuncionarioAssalariado extends Funcionario {
    private double salarioMensal;

    public FuncionarioAssalariado(double salarioMensal) {
        this.salarioMensal = salarioMensal;
    }

    @Override
    double calcularSalario() {
        return salarioMensal;
    }
}

public class Main {
    public static void main(String[] args) {
        FuncionarioHorista horista = new FuncionarioHorista(160, 20);
        FuncionarioAssalariado assalariado = new FuncionarioAssalariado(3000);

        System.out.println("Salário do Funcionário Horista: R$ " + horista.calcularSalario());
        System.out.println("Salário do Funcionário Assalariado: R$ " + assalariado.calcularSalario());
    }
}

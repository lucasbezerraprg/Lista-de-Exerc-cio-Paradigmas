public class ContaBancaria {
    private String titular;
    private double saldo;

    public ContaBancaria(String titular) {
        this.titular = titular;
        this.saldo = 0;
    }

    public void depositar(double valor) {
        if (valor > 0) {
            this.saldo += valor;
            System.out.println("Depósito de R$ " + String.format("%.2f", valor) + " realizado. Saldo atual: R$ " + String.format("%.2f", this.saldo) + ".");
        } else {
            System.out.println("Valor de depósito inválido.");
        }
    }

    public void sacar(double valor) {
        if (0 < valor && valor <= this.saldo) {
            this.saldo -= valor;
            System.out.println("Saque de R$ " + String.format("%.2f", valor) + " realizado. Saldo atual: R$ " + String.format("%.2f", this.saldo) + ".");
        } else if (valor > this.saldo) {
            System.out.println("Saldo insuficiente para saque.");
        } else {
            System.out.println("Valor de saque inválido.");
        }
    }

    public void exibirSaldo() {
        System.out.println("Saldo atual: R$ " + String.format("%.2f", this.saldo) + ".");
    }

    public static void main(String[] args) {
        ContaBancaria conta1 = new ContaBancaria("João Silva");

        conta1.depositar(100);
        conta1.exibirSaldo();

        conta1.sacar(30);
        conta1.exibirSaldo();

        conta1.sacar(80);
        conta1.exibirSaldo();
    }
}

class SaldoInsuficienteException extends Exception {
    public SaldoInsuficienteException(String mensagem) {
        super(mensagem);
    }
}

class ContaBancaria {
    private double saldo;

    public ContaBancaria(double saldoInicial) {
        this.saldo = saldoInicial;
    }

    public double sacar(double valor) throws SaldoInsuficienteException {
        if (valor > saldo) {
            throw new SaldoInsuficienteException("Saldo insuficiente para realizar o saque.");
        }
        saldo -= valor;
        return saldo;
    }

    public double depositar(double valor) {
        saldo += valor;
        return saldo;
    }

    public static void main(String[] args) {
        ContaBancaria conta = new ContaBancaria(100);

        try {
            conta.sacar(150);
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }
    }
}

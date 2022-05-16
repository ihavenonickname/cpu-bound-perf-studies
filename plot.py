import matplotlib.pyplot as plt

def plotar_throughput():
    plt.clf()

    X = [1, 2, 3, 4, 5, 6]

    plt.plot(X, [50, 83, 106, 123, 129, 141], label='.net')
    plt.plot(X, [16, 29, 40, 50, 59, 64], label='Go')
    plt.plot(X, [28, 31, 31, 31, 31, 31], label='node.js')

    plt.legend()
    plt.ylabel('Requisições por minuto')
    plt.xlabel('Usuários simultâneos')
    plt.ylim([0, 150])
    plt.yticks(list(range(10, 141, 10)))
    plt.tight_layout()
    plt.savefig('throughput.png', bbox_inches='tight')

def plotar_tempo_de_resposta():
    plt.clf()

    X = [1, 2, 3, 4, 5, 6]

    plt.plot(X, [1.1, 1.4, 1.6, 1.9, 2.1, 2.4], label='.net')
    plt.plot(X, [3.6, 4.0, 4.4, 4.7, 4.9, 5.2], label='Go')
    plt.plot(X, [2.0, 3.7, 5.5, 7.3, 8.9, 10.6], label='node.js')

    plt.legend()
    plt.ylabel('Segundos')
    plt.xlabel('Usuários simultâneos')
    plt.ylim([0, 12])
    plt.yticks(list(range(1, 12)))
    plt.savefig('tempo-de-resposta.png', bbox_inches='tight')

def main():
    plotar_throughput()
    plotar_tempo_de_resposta()
    print('Ok')

if __name__ == '__main__':
    main()

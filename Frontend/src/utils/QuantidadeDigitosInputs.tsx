export const quantidadeDigitosMax10 = (valor: number) => {
    if (valor < 10) return 4;
    if (valor >= 10) return 2;
    return 3;
}

export const quantidadeDigitosMax100 = (valor: number) => {
    if (valor < 10) return 4;
    if (valor >= 10 && valor < 100) return 5;
    return 3;
}

export const quantidadeDigitosMax1000 = (valor: number) => {
    if (valor < 10) return 4;
    if (valor >= 10 && valor < 100) return 5;
    if (valor >= 100 && valor < 1000) return 6;
    return 4;
}

export const quantidadeDigitosMax10000 = (valor: number) => {
    if (valor < 10) return 4;
    if (valor >= 10 && valor < 100) return 5;
    if (valor >= 100 && valor < 1000) return 6;
    if (valor >= 1000 && valor < 10000) return 7;
    return 5;
}

export const quantidadeDigitosMax100000 = (valor: number) => {
    if (valor < 10) return 4;
    if (valor >= 10 && valor < 100) return 5;
    if (valor >= 100 && valor < 1000) return 6;
    if (valor >= 1000 && valor < 10000) return 7;
    if (valor >= 10000 && valor < 100000) return 8;
    return 6;
}

// Iterative Algorithm

ndondomeko fibonacci(n) {
    ngati(n <= 1)
    {
        bweza n;
    }
    nambala n2 = 0;
    nambala n1 = 1;

    za(nambala i = 2; i <= n; i++) {
        n2, n1 = n1, n1+n2
    }

    return n1;
}

// Recursive Algorithm


// Dynamic Programming Algorithm

console.lemba(fibonacci(9));
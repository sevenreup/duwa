// Iterative Algorithm
console.lemba("Here is the Fibonacci Series: ");
ndondomeko fibonacci(n) {
    nambala a = 0;
    nambala b = 1;
    nambala c = 0;
    nambala i = 0;

    za(i = 2; i <= n; i++) {
        c = a + b;
        a = b;
        b = c;
    }

    bweza b;
}

// Recursive Algorithm


// Dynamic Programming Algorithm
nambala value = fibonacci(9);
console.lemba("\n");
console.lemba(value);
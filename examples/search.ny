// Linear Search
ndondomeko linearSearch(arr, x) {
    za(nambala i = 0; i < arr.length(); i++) {
            lembanzr(arr[i] == x);
        ngati (arr[i] == x) {
            bweza i;
        }
    }
    bweza -1;
}

ndondomeko doLinearSearch() {
    nambala[] arr = [2, 3, 4, 10, 40];
    nambala x = 10;

    nambala result = linearSearch(arr, x);
    ngati (result == -1) { 
        lembanzr("Element is not present in array");
    } kapena {
        lembanzr("Element is present at index "+ result);
    }
}

doLinearSearch();

// Binary Search
// Jump Search
// Interpolation Search
// Exponential Search
// Fibonacci Search
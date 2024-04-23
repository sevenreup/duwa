ndondomeko printArray(arr) {
    nambala n = arr.length();
    za(nambala i=0; i<n; i++) {
        lemba(arr[i]);
        lemba(" ");
    }
    lembanzr();
}

// Selection Sort

ndondomeko selectionSort(arr) {
    nambala n = arr.length();

    za(nambala i = 0; i < n-1; i++) {
        nambala min_idx = i;
        za(nambala j = i+1; j < n; j++) {
            ngati (arr[j] < arr[min_idx]) {
                min_idx = j;
            }
        }

        nambala temp = arr[min_idx];
        arr[min_idx] = arr[i];
        arr[i] = temp;
    }
}

nambala[] arr = [64,25,12,22,11];
selectionSort(arr);
printArray(arr);

// Bubble Sort
// Insertion Sort
// Merge Sort
// Quick Sort
// Heap Sort
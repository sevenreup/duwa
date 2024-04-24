// Linear Search
ndondomeko linearSearch(arr, x) {
    za(nambala i = 0; i < arr.length(); i++) {
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
        lembanzr("Linear Search: Element is not present in array");
    } kapena {
        lembanzr("Linear Search: Element at index "+ result);
    }
}

doLinearSearch();

// Binary Search

ndondomeko binarySearch(arr,x) {
    nambala l = 0;
    nambala right = arr.length() - 1;
    pamene (l <= right) {
        nambala mid = l + (right - l) / 2;
        ngati (arr[mid] == x) {
            bweza mid;
        }
        ngati (arr[mid] < x) {
            l = mid + 1;
        } kapena {
            right = mid - 1;
        }
    }
    bweza -1;
}

ndondomeko doBinarySearch() {
    nambala[] arr = [2, 3, 4, 10, 40];
    nambala x = 10;

    nambala result = binarySearch(arr, x);
    ngati (result == -1) {
        lembanzr("Binary Search: Element is not present in array");
    } kapena {
        lembanzr("Binary Search: Element at index "+ result);
    }
}

doBinarySearch();

// Jump Search

// TODO: needs math library

// Interpolation Search

ndondomeko interpolationSearch(arr, lo, hi, x) {
    nambala pos = 0;
    // Since array is sorted, an element present
    // in array must be in range defined by corner
    ngati (lo <= hi && x >= arr[lo] && x <= arr[hi]) {
        // Probing the position with keeping
        // uniform distribution in mind.
        pos = lo + (((hi - lo) / (arr[hi] - arr[lo])) * (x - arr[lo]));
 
        // Condition of target found
        ngati (arr[pos] == x){
            bweza pos;
        }
 
        // If x is larger, x is in right sub array
        ngati (arr[pos] < x) {
            return interpolationSearch(arr, pos + 1, hi, x);
        }
 
        // If x is smaller, x is in left sub array
        ngati (arr[pos] > x) {
            bweza interpolationSearch(arr, lo, pos - 1, x);
        }
    }
    bweza -1;
}

ndondomeko doInterpolationSearch()
{
    // Array of items on which search will
    // be conducted.
    nambala[] arr = [10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47];
    nambala n = arr.length();
 
    nambala x = 18; // Element to be searched
    nambala index = interpolationSearch(arr, 0, n - 1, x);
 
    // If element was found
    ngati (index != -1) {
        lembanzr("Element found at index" + index);
    } kapena {
        lembanzr("Element not found.");
    }

    bweza 0;
}

doInterpolationSearch();

// Exponential Search
// Fibonacci Search
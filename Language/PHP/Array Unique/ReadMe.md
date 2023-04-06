# array_unique()

This function takes an array as input and returns a new array with all duplicate values removed.
It works by looping through the input array, and adding each value to an associative array with the value as the key and true as the value. If a value already exists in the associative array, it is not added again.
Finally, the function returns the keys of the associative array as the output array, which is guaranteed to contain only unique values.

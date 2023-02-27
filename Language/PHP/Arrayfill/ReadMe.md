The array_fill function in PHP is used to create an array and fill it with a given value. Its syntax is as follows:
array_fill ( int $start_index , int $num , mixed $value ) : array

It takes three arguments:

$start_index: The index of the first element in the array.
$num: The number of elements to add to the array.
$value: The value to fill the array with.
The function returns an array filled with $num elements of the value $value, starting at the index $start_index.

The time complexity of array_fill is O(n), where n is the number of elements to be added to the array. This is because the function needs to loop over the number of elements and set each value to the given value.

The space complexity of array_fill is O(n), as the function creates an array with n elements.
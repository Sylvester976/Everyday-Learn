<?php
$array = array('apple', 'banana', 'orange');
$delimiter = ', ';
$string = implode($delimiter, $array);
echo $string; // Output: "apple, banana, orange"

// #2

$string2 = implode($array);
echo ' ,'.$string2; // Output: "applebananaorange"
?>




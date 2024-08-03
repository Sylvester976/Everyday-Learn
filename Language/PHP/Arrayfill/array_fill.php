<?php

function array_fill_custom($start_index, $num, $value) {
    $arr = array();
    for($i = 0; $i < $num; $i++) {
        $arr[$start_index + $i] = $value;
    }
    return $arr;
}

// Example usage
$start_index = 2;
$num = 5;
$value = "hello";
$result = array_fill_custom($start_index, $num, $value);

print_r($result);

?>

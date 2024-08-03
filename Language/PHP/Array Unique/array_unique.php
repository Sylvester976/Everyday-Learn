<?php
function array_unique_custom($array) {
    $unique_array = array();
    foreach ($array as $value) {
        if (!array_key_exists($value, $unique_array)) {
            $unique_array[$value] = true;
        }
    }
    return array_keys($unique_array);
}

?>

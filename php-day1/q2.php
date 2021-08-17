<?php
    $input= readline('Enter an number: ');
    $output = "";
    if(strlen($input)<10) {
        echo "invalid number\n";
        return;
    }
    $n = strlen($input);
    $mid = "";
    for ($i = 2; $i < $n - 2; $i++){
        $mid .= "*";
    }
    $output .= substr($input,0,2);
    $output .= "$mid";
    $output .= substr($input,$n - 2, 2);
    echo $output;
    echo "\n";
?>
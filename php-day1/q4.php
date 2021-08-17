<?php
    $obj="{\"players\":[{\"name\":\"Ganguly\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Dravid\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Dhoni\",\"age\":37,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Virat\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}}]}";
    $obj1 = json_decode($obj,true);
    $name = array();
    $age = array();
    $city = array();
    foreach($obj1["players"] as $row) {
        array_push($name, $row["name"]);
        array_push($age, $row["age"]);
        array_push($city, $row["address"]["city"]);
        $mx_age = max($mx_age,$row["age"]);
    }

    print_r($name);
    echo "\n";
    print_r($age);
    echo "\n";
    print_r($city);
    echo "\n";
    print_r(array_unique($name));
    echo "\n";
    $max_age = array();
    foreach($obj1["players"] as $row){
        if($row["age"]==$mx_age) {
            array_push ($max_age, $row["name"]);
        }
    }
    print_r($max_age);
?>
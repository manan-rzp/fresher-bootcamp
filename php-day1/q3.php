<?php
    function camel_case ($val){
        $str="";
        for($i=0;$i<strlen($val);$i++){
            if($val[$i]=='_'){
                $i++;
                break;
            }
            $str.=$val[$i];
        }
        $str.=ucfirst(substr($val,$i,strlen($val)-$i));
        return $str;
    }

    $input= array("snake_case", "camel_case");
    foreach($input as $val){
        echo camel_case($val);
        echo "\n";
    }
?>
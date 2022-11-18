<?php


header("Content-Type:text/html;charset=utf-8");

highlight_file(__FILE__);


if(isset($_GET['one']) && $_GET['one'] == 0 && $_GET['one']){
    echo 'pass 1<br>';
}else{
    die('no no no!');
}

$query = $_SERVER['QUERY_STRING'];
if(strpos($query,'one') !== false){
    die('no no no!');
}else{
    echo 'pass 2<br>';
}

$f=$_GET['file'];
@include($f . '.php');
?>
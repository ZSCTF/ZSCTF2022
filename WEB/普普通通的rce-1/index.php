<?php
 // 来个简单的rce
highlight_file(__FILE__);
if(isset($_GET['a'])){
    $a = $_GET['a'];
    if(!preg_match("/flag| |system|`|cat/i", $a)){
        eval($a);
    }
    
}
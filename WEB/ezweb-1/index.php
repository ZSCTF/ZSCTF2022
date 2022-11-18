<?php

highlight_file(__FILE__);

if(md5(time())==$_GET['time']){
    new $_GET['a']($_GET['b']);
}

echo "enjoy yourself in ZSCTF";


?>
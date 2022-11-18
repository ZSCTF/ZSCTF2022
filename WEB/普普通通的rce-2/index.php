<?php
 // 能做出上一题,你已经学会了简单的RCE技巧了,那么我们来稍微加点难度
highlight_file(__FILE__);
if(isset($_GET['a'])){
    $a = $_GET['a'];
    if(!preg_match("/\;|cat|flag| |[0-9]|\*|more|wget|less|head|sort|tail|sed|cut|tac|awk|strings|od|curl|\`|\%|\x0a|\x0b|\x0c|\x0d|\x09|\x26|\>|\<|die|exit/i", $a)){
        eval($a);
        $s = ob_get_contents();
        ob_end_clean();
        echo preg_replace("/[a-z]/i","*",$s);
    }else{
        echo "hacker!";
    }
    
}
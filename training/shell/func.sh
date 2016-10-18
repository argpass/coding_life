#!/usr/bin/env bash

hello(){
    echo "hello world :$*"
    echo "args len:$#"
}
hello shell go

add(){
    echo -n "Input your first num:"
    read a
    echo  -n "Input your second num:"
    read b
    return $(($a + $b))
}
add
ret=$?
echo "return is $ret after add"

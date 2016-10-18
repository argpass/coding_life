for i in if.sh select.sh; do
    echo "goto :$i"
    cat $i
done
unset i
echo $i

echo "test seq ==========>"
for i in `seq 10`; do
    echo "$i"
done
echo "i end is:$i"

echo "test for i++"
for ((i=0; i<10; i++)); do
    echo $i
done

echo "for i in str_var"
list="a b c"
for i in $list; do
    echo $i
done

echo "for i in str"
for i in "a b c"; do
    echo $i
done

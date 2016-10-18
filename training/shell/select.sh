echo "What's your favourite OS ?"

select ch in "Mac" "Win" "Linux";do
    echo $ch
    if [ $ch == "Mac" ]; then
        echo "mac os"
        break
    fi
done

echo "You chose $ch"

for m in `seq 10 12`; do
    echo "doing $m month"
    for d in `seq 1 31`; do
        # 10月只需要15号之后的
        if [ $m -eq 10 ]; then
            if [ $d -lt 15 ]; then
                continue
            fi
        fi
        # 12月只要今天以前的
        if [ $m -eq 12 ]; then
            if [ $d -gt 25 ]; then
                continue
            fi
        fi
        # 日期
        ds="2016.$m.$d"
        echo "fetch $ds"
        url="www.qijigps.com/api/install/yh?date=$ds"
        name="$ds.xlsx"
        wget -O $name $url
    done
done

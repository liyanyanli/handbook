#!/bin/sh

echo "检查路径:"$1
echo "过期时间:"$2
echo "匹配文件:"$3
echo "运行间隔:"$4

while true; do
    find $1 -mtime +$2 -name $3 -exec rm {} \;
    echo `date`;
    sleep $4'd';
done

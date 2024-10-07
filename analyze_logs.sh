#!/bin/bash

# 检查日志文件路径
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <log_file_path>"
    exit 1
fi

log_file="$1"

# 检查日志文件是否存在
if [ ! -f "$log_file" ]; then
    echo "Log file not found: $log_file"
    exit 1
fi

# 提取IP地址并计算每个IP的请求次数
echo "Top 5 IP addresses with the most requests:"
awk '{print $1}' "$log_file" | sort | uniq -c | sort -nr | head -n 5 | while read line; do
    ip=$(echo $line | awk '{print $2}')
    count=$(echo $line | awk '{print $1}')
    echo "$ip - $count requests"
done
#!/bin/bash

echo "=== 测试Go反射示例 ==="

echo -e "\n1. 测试核心知识点示例..."
go run reflect_core.go
if [ $? -eq 0 ]; then
    echo "✓ reflect_core.go 运行成功"
else
    echo "✗ reflect_core.go 运行失败"
    exit 1
fi

echo -e "\n2. 测试完整示例..."
go run reflect_demo.go
if [ $? -eq 0 ]; then
    echo "✓ reflect_demo.go 运行成功"
else
    echo "✗ reflect_demo.go 运行失败"
    exit 1
fi

echo -e "\n=== 所有测试通过 ==="
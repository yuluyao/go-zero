#!/bin/bash

wd=$(pwd)
output="$wd/hello"
api_file="$wd/hello.api"
goctl_root="$(dirname $(dirname $wd))"

# 检查hello.api文件是否存在
if [ ! -f "$api_file" ]; then
    echo "Error: hello.api file not found at $api_file"
    exit 1
fi

# 清理之前的输出目录
echo "🧹 Cleaning output directory: $output"
rm -rf $output

# 使用go run方式生成API代码
echo "🚀 Generating Go code from API file: $api_file"
echo "📁 Output directory: $output"

cd "$goctl_root"
go run goctl.go api go -api "$api_file" -dir "$output" --style go_zero

if [ $? -ne 0 ]; then
    echo "❌ Generate failed"
    exit 1
fi

echo "✅ Code generation completed successfully!"

# 切换到输出目录进行后续操作
cd "$output"

# 初始化go模块（如果不存在）
if [ ! -f "go.mod" ]; then
    echo "📦 Initializing Go module..."
    go mod init hello
fi

# 修复生成代码中的导入路径
echo "🔧 Fixing import paths..."
old_import_path="github.com/zeromicro/go-zero/tools/goctl/example/api/hello"
new_import_path="hello"

# 使用find和sed替换所有Go文件中的导入路径
find . -name "*.go" -type f -exec sed -i '' "s|$old_import_path|$new_import_path|g" {} \;

echo "✅ Import paths fixed!"

# 整理依赖
echo "🔧 Tidying Go modules..."
go mod tidy

if [ $? -ne 0 ]; then
    echo "❌ Tidy failed"
    exit 1
fi

echo "🎉 All done! Generated code is in: $output"
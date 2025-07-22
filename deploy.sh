#!/bin/bash

set -e

# 日志函数
log() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [INFO] $1"
}

error() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [ERROR] $1" >&2
}

log "拉取最新代码..."
git fetch && git reset --hard origin/master

log "清理旧文件..."
rm -rf ./bin/webconsole

log "编译可执行文件..."
if ! go build -o ./bin/webconsole main.go; then
  error "编译失败，请检查代码"
  exit 1
fi

log "检查编译结果..."
if [ ! -f "./bin/webconsole" ]; then
  error "可执行文件不存在，编译可能失败"
  exit 1
fi

log "停止旧进程..."
# 使用 pkill -f 匹配完整命令行，避免误杀
# 使用 || true 防止 pkill 在没有找到进程时导致脚本退出
pkill -f "webconsole" || true

# 可选：等待一段时间确保进程已停止
sleep 2

# 检查进程是否真的已停止
if pgrep -f "webconsole" > /dev/null; then
  error "旧进程未停止，尝试强制终止"
  pkill -9 -f "webconsole" || true
  sleep 2
  
  if pgrep -f "webconsole" > /dev/null; then
    error "无法终止旧进程，请手动处理"
    exit 1
  fi
fi

log "启动新进程..."
nohup ./bin/webconsole &
NEW_PID=$!

# 验证进程是否成功启动
sleep 2
if ! ps -p $NEW_PID > /dev/null; then
  error "新进程启动失败"
  
  # 输出应用日志帮助排查
  if [ -f "nohup.out" ]; then
    error "应用日志内容："
    tail -n 20 nohup.out
  fi
  
  exit 1
fi

log "进程启动成功, PID: $NEW_PID"
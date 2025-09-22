#!/bin/bash

set -e

# 配置变量
SERVICE_NAME="identify_service"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
INSTALL_DIR="/usr/share/${SERVICE_NAME}"
BINARY_NAME="identify_service"
CONFIG_FILE="${INSTALL_DIR}/config.yaml"

log() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [INFO] $1"
}

error() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [ERROR] $1" >&2
}

log "[1] 停止服务"
systemctl stop identify_service

log "[2] 检测服务是否停止"
# 使用 systemctl is-active 检查服务是否真的停止（仅关注状态，不输出详细日志）
if systemctl is-active --quiet identify_service; then
  error "服务未成功停止，无法继续操作"
  exit 1
else
  log "服务已成功停止"
fi

log "[3] 检查目录是否存在"
if [ ! -d "/usr/share/identify_service" ]; then
  log "目录不存在，创建目录"
  mkdir -p /usr/share/identify_service
  mkdir -p /usr/share/identify_service/resource/casbin
  mkdir -p /usr/share/identify_service/resource/log/server
  mkdir -p /usr/share/identify_service/resource/log/sql
  mkdir -p /usr/share/identify_service/resource/data/distTokenDb
  mkdir -p /usr/share/identify_service/resource/session
fi

log "[4] 拷贝可执行文件到/usr/share/identify_service"
/bin/cp -f ./resource/casbin/rbac_model.conf /usr/share/identify_service/resource/casbin/rbac_model.conf
/bin/cp -f ./resource/casbin/rbac_policy.csv /usr/share/identify_service/resource/casbin/rbac_policy.csv
/bin/cp -f ./bin/identify_service /usr/share/identify_service/
/bin/cp -f ./config.yaml /usr/share/identify_service/

log "[5] 重启identify_service服务"
systemctl restart identify_service

log "[6] 查看identify_service服务状态"
systemctl status identify_service

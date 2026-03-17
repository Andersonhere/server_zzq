#!/bin/bash

# 生成自签名 SSL 证书脚本
# 用于本地开发和测试

CERT_DIR="certs"
CERT_FILE="$CERT_DIR/server.crt"
KEY_FILE="$CERT_DIR/server.key"

# 创建证书目录
mkdir -p "$CERT_DIR"

# 检查是否已存在证书
if [ -f "$CERT_FILE" ] && [ -f "$KEY_FILE" ]; then
    echo "证书已存在，是否重新生成？(y/n)"
    read -r response
    if [ "$response" != "y" ]; then
        echo "跳过证书生成"
        exit 0
    fi
fi

# 生成自签名证书
echo "正在生成自签名证书..."
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout "$KEY_FILE" \
    -out "$CERT_FILE" \
    -subj "/C=CN/ST=State/L=City/O=Organization/OU=Unit/CN=localhost" \
    -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"

if [ $? -eq 0 ]; then
    echo "证书生成成功！"
    echo "证书文件：$CERT_FILE"
    echo "私钥文件：$KEY_FILE"
    echo ""
    echo "注意：自签名证书在浏览器中会显示安全警告，这是正常现象。"
    echo "如需信任此证书，请将 $CERT_FILE 添加到系统的信任证书列表中。"
else
    echo "证书生成失败！"
    exit 1
fi

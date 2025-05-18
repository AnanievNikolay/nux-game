#!/bin/bash

set -e

APP_ROOT=~/app/nux-game

echo "📁 dir check..."

mkdir -p "$APP_ROOT/logs"
echo "✅ Created: $APP_ROOT/logs"

mkdir -p "$APP_ROOT/sqlite-db"
echo "✅ Created: $APP_ROOT/sqlite-db"

cp ./config.json "$APP_ROOT/config.json"
echo "✅ Copied: ./config.json -> $APP_ROOT/config.json"

echo "🐳 docker-compose..."
docker-compose up -d

sleep 3

echo "🩺 Checking service health..."
STATUS=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/is_alive || true)

if [ "$STATUS" = "200" ]; then
    echo "✅ Service is alive (200 OK)"
    exit 0
else
    echo "❌ ERROR: Service is not healthy (HTTP $STATUS)"
    exit 1
fi
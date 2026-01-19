#!/bin/sh
set -e

: "${BACKEND_URL:?BACKEND_URL env var not set}"

echo "Replacing placeholders with real URLs..."
ROOT="/usr/share/nginx/html"

find "$ROOT" -type f \( -name '*.html' -o -name '*.js' -o -name '*.css' \) -exec grep -l "__BACKEND_URL__" {} + | \
  xargs -r sed -i "s|__BACKEND_URL__|$BACKEND_URL|g"

echo "Backend URL replaced!"
echo "Starting nginx..."

exec nginx -g 'daemon off;'

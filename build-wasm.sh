#!/usr/bin/env bash
# Build TinyGo WASM, optimize with wasm-opt, brotli-compress, keep only .br for deploy.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")" && pwd)"
cd "$ROOT"

cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./wasm_exec.js

tinygo build -o web/app.wasm -target=wasm -no-debug -opt=z ./src

if command -v wasm-opt >/dev/null 2>&1; then
  wasm-opt --enable-bulk-memory -Oz web/app.wasm -o web/app.wasm
else
  echo "warning: wasm-opt not found; skipping binaryen optimize step" >&2
fi

if ! command -v brotli >/dev/null 2>&1; then
  echo "error: brotli is required (deploy ships only web/app.wasm.br)" >&2
  exit 1
fi

# Uncompressed size for go-app's loading progress (file is removed after compress).
wc -c < web/app.wasm | tr -d ' ' > web/app.wasm.size
brotli -f -9 web/app.wasm
rm -f web/app.wasm

ls -lh web/app.wasm.br web/app.wasm.size wasm_exec.js

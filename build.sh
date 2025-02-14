#!/usr/bin/env bash

rm -rf build
mkdir -p build

os_archs=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "windows/amd64")
for os_arch in "${os_archs[@]}"; do
  IFS="/" read -r os arch <<< "$os_arch"
  output_name="hll-arty-calc-${os}-${arch}"
  if [ "$os" = "windows" ]; then
    output_name+=".exe"
  fi
  echo "Building for $os/$arch..."
  GOOS=$os GOARCH=$arch go build -o "build/$output_name"
done

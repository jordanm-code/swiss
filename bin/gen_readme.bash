#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PKG_DIR=$SCRIPT_DIR/..

if which gomarkdoc > /dev/null; then
  echo "Found gomarkdoc"
else
  echo "Installing gomarkdoc"
  go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
fi

echo "Generating README.md"
gomarkdoc -e -o $PKG_DIR/README.md $PKG_DIR

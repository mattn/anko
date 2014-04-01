#!/bin/bash

DIR=$(cd $(dirname $0);pwd)

ls $DIR/*.t |\
while read f; do
  anko $DIR/tester.ank $f
done

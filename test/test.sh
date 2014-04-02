#!/bin/bash

DIR=$(cd $(dirname $0);pwd)

ls $DIR/*.t |\
while read f; do
  $DIR/../anko $DIR/tester.ank $f
done

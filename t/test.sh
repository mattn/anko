#!/bin/bash

DIR=$(cd $(dirname $0);pwd)

ls $DIR/*.ank |\
while read f; do
  $DIR/../anko $DIR/lib/tester.ank $f
done

#!/bin/bash

set -ev

$GOPATH/bin/godepgraph -horizontal -s -o github.com/dev-drprasad/qt/internal/examples/showcases/wallet github.com/dev-drprasad/qt/internal/examples/showcases/wallet | dot -Tpng -o depgraph.png

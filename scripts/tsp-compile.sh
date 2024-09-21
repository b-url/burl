#!/bin/bash

current_dir=$(pwd)
cd api/typespec || exit
tsp compile . --emit @typespec/openapi3 --option "@typespec/openapi3.emitter-output-dir=${current_dir}/api"
cd "$current_dir" || exit

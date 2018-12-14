#!/bin/bash
cmake -G Ninja -DCMAKE_BUILD_TYPE=Release -B .build_osx .
cmake --build .build_osx

mkdir -p ../../tools/flatbuffers/bin/Darwin
cp .build_osx/flatc ../../tools/flatbuffers/bin/Darwin/

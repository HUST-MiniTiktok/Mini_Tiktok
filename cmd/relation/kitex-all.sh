#! /usr/bin/env bash

kitex -service relationservice -gen-path ../../kitex_gen --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ relation.thrift 
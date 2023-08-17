#! /usr/bin/env bash

kitex -service feedservice -gen-path ../../kitex_gen --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ feed.thrift 
#! /usr/bin/env bash

kitex -service publishservice --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ publish.thrift 
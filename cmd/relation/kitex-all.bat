#! /usr/bin/env bash

kitex -service relationservice --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ relation.thrift 
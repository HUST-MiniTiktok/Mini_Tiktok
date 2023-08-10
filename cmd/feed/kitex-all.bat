#! /usr/bin/env bash

kitex -service feedservice --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ feed.thrift 
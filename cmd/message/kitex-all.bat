#! /usr/bin/env bash

kitex -service messageservice --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ message.thrift 
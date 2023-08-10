#! /usr/bin/env bash

kitex -service userservice --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ user.thrift 
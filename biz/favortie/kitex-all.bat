#! /usr/bin/env bash

kitex -service favoriteservice --record --thrift-plugin validator -module github.com/HUST-MiniTiktok/mini_tiktok -I ../../idl/ favorite.thrift 
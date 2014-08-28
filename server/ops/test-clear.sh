#!/bin/bash
echo "Clearing test mongo"
path=`pwd`
rm -rf test/resources/database/*
rm -rf test/log/database.log

mkdir -p test/log
mkdir -p test/resources
mkdir -p test/resources/database


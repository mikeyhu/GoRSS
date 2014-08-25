#!/bin/bash
mongod --port 27000 --pidfilepath "test/resources/database.pid" --nssize 1 --smallfiles --noprealloc --dbpath test/resources/database/ > test/log/database.log &

#!/bin/bash

###########################################
#Script Name    : bootstrap.sh
#Description    : holds commands for vagrant file to access to setup
#                 enviornment
#Args           : NIL
#Author         : SHREY SONI
#Email          : shrey@sodio.tech
###########################################


echo "CHECK FOR CURL "
# update the system
# sudo apt-get update
# check curl available or not
CurlVersion = $(sudo curl --version) 
# if logic
echo "Here is the version of curl"
echo "${CurlVersion}"
if [ -z "$CurlVerison" ]
then
    echo "\$Curl is already installed"
else 
     sudo apt-get install curl
fi


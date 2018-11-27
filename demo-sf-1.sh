#!/usr/bin/env bash

########################
# include the magic
########################
. ./demo-magic.sh 


########################
# Configure the options
########################

#
# speed at which to simulate typing. bigger num = faster
#
TYPE_SPEED=10

#
# custom prompt
#
# see http://www.tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html for escape sequences
#
DEMO_PROMPT="${GREEN}➜ ${CYAN}\W "

# hide the evidence
clear

echo "# pre-requirements: SAP CP, Service Manager, Service Fabrik"

echo
echo "# Pretty print the catalog"
pe "cf marketplace -s postgresql"

echo
echo "# Order one Service-Fabrik service instance provisioned in ali-cloud with name ali-postgresql"
pe "cf create-service postgresql v9.4-apsaradb ali-postgresql"


echo
count=0
sleep 1
while [[ `cf service ali-postgresql | sed -n '13 p' | awk '{print $3}'` != succeeded ]]; do printf "\r waiting for instance creation...you can view it in the ali cloud portal ......... waiting for $count seconds"; count=$((count+1)); sleep 1; done

echo
echo "# Check the state of the service instance provisioned in ali-cloud with name ali-postgresql"
pe "cf service ali-postgresql"

echo "delete the service"
wait

cf ds ali-postgresql -f

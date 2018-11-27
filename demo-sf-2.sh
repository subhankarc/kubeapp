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
echo "# Check for already provisioned service instance provisioned in ali-cloud"
pe "cf services"


echo
echo "# Push an application on SAP CP and consume the service from it."
pe "cat manifest.yml"
pe "cf push"
echo "# Open the url demo-app.cfapps.dev2.sf.sapcloud.io to see the application on the browser"

#!/bin/bash
export PATH=.:$PATH
read-system-health.sh
echo
read-api-kafka-admin-config.sh
echo
read-api-kafka-admin-cluster.sh
echo
read-api-kafka-admin-topic.sh
echo
read-api-kafka-admin-consumer-group.sh
echo

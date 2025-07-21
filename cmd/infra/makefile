#!/bin/bash

canonical-orders-local:
	kcat -P -b localhost:9092 \
        -t canonical-orders \
        -H ev=1 \
        -H id=MDFIVjczN1NIOFczQzUyVDYzN0UzTk1SRVYtMDFIVjczN1NINzhYNUZFUjI1NFNYN0dFUDE= \
        -H sn=canonical-orders \
        -H sk=01HV737SH8W3C52T637E3NMREV-01HV737SH78X5FER254SX7GEP1 \
        -H si=create-order \
        -H sv=0 \
        -H tg=MDFIVjczN1NIOFczQzUyVDYzN0UzTk1SRVYtMDFIVjczN1NINzhYNUZFUjI1NFNYN0dFUDE= \
        -H tc=MDFIVjczN1NIOFczQzUyVDYzN0UzTk1SRVYtMDFIVjczN1NINzhYNUZFUjI1NFNYN0dFUDE= \
        -H ps=json \
        order.json

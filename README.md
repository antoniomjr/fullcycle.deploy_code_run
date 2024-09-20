# Go Challenge for FullCycle Pós Go Expert

## Overview

This repository contains a Go application developed as a part of the "`fullcyle.deploy-code-run`" (Go Challenge) from the Pós Go Expert.

## Challenge Requirements

Objective: Develop a system in Go that receives a Brazilian postal code (CEP), identifies the city, and returns the current weather (temperature in Celsius, Fahrenheit, and Kelvin). This system will be deployed on Google Cloud Run.

## Instructions

To build and run this application, follow these steps:

#### Running the Go application
This will run the application
```bash
docker-compose up -build
```

#### Request by terminal
```bash
curl http://localhost:8080/weather?cep=seu_cep_aqui
```

#### Request by terminal
```bash
https://deploy-code-run-563666353735.us-central1.run.app/weather?cep=seu_cep_aqui
```

# cf-tunnel-service
CloudFoundry CLI Plugin to forward ports from service instances to your local computer

## Overview
Applications deployed to CloudFoundry can be bound to Service Instances available through the marketplace. 
These services are often deployed within the same private network as the CloudFoundry installation, on a private IP Address. 

As an Application Developer you may need to connect directly to the service instance in order to perform actions such as migrations, data import / export or just generally debugging. 
Unless you ask your CloudFoundry operator to poke you a hole in the firewall (*not a good idea*) then this can be hard to do.

This CF CLI plugin will push an empty application into the space which you are currently targetted to and bind it to the service which you input. It will then use the SSH Functionality provided by Diego to forward the port you specify to your local computer

## Prerequisites

* CloudFoundry running Diego
* Diego SSH functionality enabled

## Installation

## Usage

``

## Removal

## Credit

Credit to [Marco Nicosia](github.com/menicosia)


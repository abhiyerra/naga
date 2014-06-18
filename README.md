# Bootup #

## Use Case ##

The goal of bootup is to provision environments to test gerrit changes
quickly. Whenever server engineers are developing a change they may
need to setup an environment that is used by other members of the team
so that they may test the change. The best way to do this is to
provision an environment in the cloud.

Bootup attempts to make this simple by getting a gerrit change that is
still in testing and provisioning an environment in which it can be
tested.

## How it works ##

- It checks out a revision of code.
- It creates a vagrant environment which it provisions
- It sets up the code to run a default script (./bootup-run.sh).

## Setup ##

### Environment Variables ###

 - BOOTUP_GERRIT_HOST
 - BOOTUP_GERRIT_USERNAME
 - BOOTUP_GERRIT_PASSWORD

## Run it ##

    bootup -change_id 1234 -revision 1

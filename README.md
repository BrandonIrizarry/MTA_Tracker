# MTA Tracker

An app for tracking bus and train arrivals for a registered user's
favorite stops within the NYC Metropolitan Transit Authority (MTA).

# Rationale

While this is effectively a solved problem on many fronts (see the
[Notable Mentions](#notable-mentions) section later in this README), I
feel enthusiastic about crafting my own implementation: something
simple and intuitive. Since I commute daily by bus (or train, if I
feel like it), I want somthing that I can use **myself** on a
day-to-day basis.

# Idea

## Configuration

The user selects either a bus or a train stop. Once a stop is
selected, the user selects the lines they wish to track (or all of
them, if desired.) In this way, the user accrues a structured, nested
set of data, designed for easy browseability.

## Example Scenario

Alice needs to decide whether they should take the bus or train, based
on whichever one gets them to work on time.  She opens the app, and
examines some bus and train stations she has bookmarked (fake names
here are used merely for the sake of illustration):

XYZ Street W Avenue
	Uptown 1: 8:30 AM
	Uptown 2: 8:45 AM
	
Nth Street Hill Drive
	Northbound Q31: 8:15 AM
	Northbound Q32: 8:30 AM
	
The first entry's header would be in red, signifying a train stop, and
the second entry's header would be in blue, signifying a bus stop.

In this case, Alice may well choose to take the bus, since it arrives
a bit sooner than the next train.

# Notable Mentions

## MTA BusTime
The [MTA BusTime Website](https://bustime.mta.info/ "MTA BusTime URL")
provides excellent real-time data on bus arrivals for a given route
among the five boroughs. I constantly use the [mobile
version](https://bustime.mta.info/m/ "MTA BusTime URL mobile version")
of the website.

## realtimerail.nyc
An especially remarkable (and hassle free!) tool for tracking train
arrivals is [realtimerail.nyc](realtimerail.nyc "realtimerail.nyc
URL"). The quality and care put into this app is worthy of envy. I
rely on it for determining whether I'll make a certain train on time.



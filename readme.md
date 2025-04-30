# FooService

Foo Service is a fictional service used to describe how we might experiment
in production with an improvement to our service.

Foo Service is dependent upon Flaggy!, a feature toggle service found [here](https://github.com/timpellison/flaggy).
It is optimized for the Mac M3 chip so if you're running an intel box you will want to alter the
make start command to build the correct GOARCH.

## A little pretend goes a long way!

The Foo Service has been in production for a while and our SLO (Service Level Objective) is 
p(99) < 2s

This service is meeting the SLO but we've had complaints from customers that it takes too long
for their account data to be returned.

We have investigated and don't have the same amount of data in pre-prod so
we need to conduct an experiment.  We believe that with a little tweaking to 
our query, we can produce results putting us in the p(99) < 100ms range.

We aren't allowed access to production for info sec reasons so our recourse is to 
deploy the updated query to production and collect duration statistics to determine 
if it is indeed faster.

Our benefactors have incorporate Flaggy! into our enterprise.  

### Pretend
There is no dashboard however, we are going to pretend we have a dashboard 
and we'll use k6 with a short iteration run to pretend to be the dashboard.


## Running this Demo
[Prerequisites]
1. Go 1.24.2
2. Docker (Desktop, Colima, Rancher, we don't care!)

[Steps]
1. clone github.com/timpellison/flaggy
2. run make start in flaggy.  you may have to change GOARCH if you're not m3 on Mac
3. run make baseline-test in this repo
   a. Observe the p(95) results
4. run make experiment in this repo
   a. Observe the p(95) results


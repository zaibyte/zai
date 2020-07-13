<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents** 

- [Fantastic Figures](#fantastic-figures)
- [What is durability](#what-is-durability)
- [Zai is weak](#zai-is-weak)
- [Zai is strong](#zai-is-strong)
- [FAQ](#faq)
  - [How dare are you! Isn't durability the most important thing?](#how-dare-are-you-isnt-durability-the-most-important-thing)
  - [Is that true any storage system would lose data?](#is-that-true-any-storage-system-would-lose-data)
  - [What is the strongest storage media?](#what-is-the-strongest-storage-media)
  - [Will Zai get stronger in the future?](#will-zai-get-stronger-in-the-future)
  - [If a storage system has a better durability than Zai, what are the things it may have done beyond Zai?](#if-a-storage-system-has-a-better-durability-than-zai-what-are-the-things-it-may-have-done-beyond-zai)
  - [Sounds it's not easy to get "100%" ?](#sounds-its-not-easy-to-get-100%25-)
  - [Ceph has copyset to improve durability, how about Zai?](#ceph-has-copyset-to-improve-durability-how-about-zai)
  - [If I want to make Zai stronger, what can I do?](#if-i-want-to-make-zai-stronger-what-can-i-do)
  - [2 or 3 replications, which one is better for ZBuffer?](#2-or-3-replications-which-one-is-better-for-zbuffer)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

Durability
===`

## Fantastic Figures

I use "Fantastic Figures" here, because the best article about durability I have ever read is Feynman's `<Investigating The Space Shuttle Challenger Disaster>` , and "Fantanstic Figures" is one of the most funny parts.

NASA said the probability of failure was more like 1 in 10^5.

"That means you could fly the shuttle every day for an average of 300 years between accidents ---- every day, one flight, for 300 years ---- which is obviously crazy!"

"Yes, I know", said Mr.Ullian. "I moved my number up to 1 in 1000..."

"1 in 200"

"1 in 300"

"100 percent!"

"I said it was 1 in 300, and I still say it's 1 in 300, but I don't want to tell you how I got my number."

...

Emm... Sound familiar.

## What is durability

I got multi replication and erasure codes, isn't that enough?

Durability is not just about repairing things, it's more about data integrity, that is:

"I put A, I get A."

But in the real world, maybe "I put A, but send B, when B arrived, it becomes C, then changes to D, write E down, return F when there is a read request, and F becomes G..." Wait a minute, maybe after such a series of flipping, we still can get "A"...

## Zai is weak

Zai is not designed for data which needs strict durability. 
It's not a right place to explain why Zai chooses that way here, because I think the users who wanna use Zai would know better about what's the durability they really want.

1. The data source is out of control: most of the end-user are not enterprise devices, 
and can't ensure there is no misdirected-writes or lost-writes during the write process or can't verify the read process either. 
How could we do the right things with an unreliable source?

2. Zai doesn't perform write check after write immediately

## Zai is strong

1. Zai enable E2E checking by default.

2. Zai has 3 replications and Erasure Codes.

## FAQ

### How dare are you! Isn't durability the most important thing?

Because I'm honest.

### Is that true any storage system would lose data?

Yes! Unless you're Dr.Strange who can make time moves back.

### What is the strongest storage media?

I think DNA with Fountain Codes is a good choice, and Superman's DNA is better :D

### Will Zai get stronger in the future?

Maybe not. Zai is not design for that purpose.

### If a storage system has a better durability than Zai, what are the things it may have done beyond Zai?

a. Add location info to each checksum for avoiding misdirected-writes

b. Using write verify or make the checksum elsewhere to avoid lost-writes

c. Do the same things(mentioned above) in end-user side

### Sounds it's not easy to get "100%" ?

I have to say if end-user do nothing, the improvement is almost 0. Because usually end-user and the connection to Zai are the weakest part of the whole data stream.

But we still can catch at least one SDC(Silent Data Corruption) by normal checking (e.g. E2E checking) or 
advanced checking mentioned above, if run long enough.

So it's worth to do all these things? Yes or not! 

### Ceph has copyset to improve durability, how about Zai?

It's easy to implement copyset in Zai, I will make it someday.

And yes, copyset will help the upgrade/restart process too, we can restart all the nodes which are in different copyset
at the same time. If we don't have copyset, it won't be so gracefully restart more than two nodes at the same time,
because any two nodes non-responded may cause a uploading failed.

### If I want to make Zai stronger, what can I do?

a. Follow the rules in [op-guide](../op-guide/durability.md)

b. Use end2end check with all connections

c. Change the data structure on storage node, ensure writing the right data. And check data after read from the disk

### 2 or 3 replications, which one is better for ZBuffer?

In Zai, 3 replications is default setting, and can't be changed it.

Here is a brief explanation:

My way to measure the durability is copy from Kevin Greenan's paper `<Reliability and Power-Efficiency in Erasure-Coded Storage Systems>`,
just modify the codes for replications model(but not erasure codes), and replace some const by other papers or my experience.

The results is below:

For 3 reps/year:

    Average bytes lost per usable TB: 20856.21467
    
For 2 reps/year:

    Average bytes lost per usable TB: 38297882.30313
    
Actually, the results in practice should be better than these numbers, because Zai implements repairing
in parallel, it will be really fast to repair a single disk, but I'm too lazy to fix the model of repairing. :D

We also need to care about the availability of uploading objects. In present,
Zai only support all reps returns ok than regards the uploading as ok one. 
After 3 tries, the 2reps availability is `8e-09`, the 3reps availability is `2.7e-08.
(Assume the single node SLA is 99.9%). 2reps is much better, but 3reps is enough in
most cases.

And for gaining better read performance, more replicas will help.

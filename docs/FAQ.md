<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  

- [Frequently asked questions (FAQs)](#frequently-asked-questions-faqs)
  - [Q: Why call it "Zai"?](#q-why-call-it-zai)
  - [Q: Why call it "4Mile" before?](#q-why-call-it-4mile-before)
  - [Q: What Zai can't do?](#q-what-zai-cant-do)
  - [Q: Why max object size is 8MB?](#q-why-max-object-size-is-8mb)
  - [Q: Up to 2814 trillions objects?](#q-up-to-2814-trillions-objects)
  - [Q: How about more than 2814 trillions?](#q-how-about-more-than-2814-trillions)
  - [Q: Why should we keep box small?](#q-why-should-we-keep-box-small)
  - [Q: Please talk about Load Balance in Zai](#q-please-talk-about-load-balance-in-zai)
  - [Q: Where is the fast GC in ZCold?](#q-where-is-the-fast-gc-in-zcold)
  - [Q: Why Zai can't protect my data 100%?](#q-why-zai-cant-protect-my-data-100%25)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

Frequently asked questions (FAQs)
===

## Q: Why call it "Zai"?

Zai, in Chinese is "载" which means a huge number just like what Zai does -- provides
ability to hold trillion objects.

## Q: Why call it "4Mile" before?

The old name of Zai is 4Mile, and it's a lovely name too.

See [Why 4Mile](why_4mile.md) for details.

## Q: What Zai can't do?

1. Can't persist large objects (default max size is 8MB).
2. Up to 2814 trillions objects in a single namespace.
3. Can't ensure every bit is correct.
4. ...

## Q: Why max object size is 8MB?

8MB is good for load balance, we can gain better performance because there is no
big objects cause disk I/O stall. And Zai is for LOSF, 8MB is a fit number.

One major design purpose of Zai is for SNS. For example, short video clips, 
they are always short and in low bit rate. If it's 30s, and under 2000kbps, the size would be 8MB.
Mostly, they are only few seconds, so 8MB is large enough. 
(e.g. video clips on TikTok are always around 5MB).
And for longer videos, m3u8 is very popular now,
so it's not a big deal.

Anyway, if you want to support much larger objects, you could split the large one to
pieces and link them together by link API.

## Q: Up to 2814 trillions objects?

Zai is built for LOSF, so it's reasonable to keep usable storage capacity in a small size:

If the usable storage capacity is 10PB (not include write buffer):

1024(1024 boxes per namespace is far away for the limitation, but it's already huge in practice) * (10PB/4KB) = 2814 trillions 

## Q: How about more than 2814 trillions? 

Congratulations!

You must be the biggest data company in the world. 

You could deploy more boxes in a namespace or just deploy more namespaces.

## Q: Why should we keep box small?

Actually, you can have a much bigger box. I assume it's small (e.g. 10PB), because I'm a humble boy.

## Q: Please talk about Load Balance in Zai

First load isn't that heavy in a storage system likes Zai, 
and the hottest data will be placed all over the box.

So optimizing the load is not urgent, maybe I'll make a better solution later.

## Q: Where is the fast GC in ZCold?

ZCold is using Erasure Codes to protect data, and it's not so easy to implement GC on a EC stripe.
That's because we combine lots of objects in a single extent, and several extents form a EC stripe.
We need Copy-and-Paste data, make a new EC stripe, then modify the metadata of extents to finish the GC process.

But I didn't implement it, not just because it's difficult to achieve it. 

In statistics, 99% delete ops happen in 90days. 
At that time, files are in ZBuf, it's really easy to reuse space there.
So implementing GC in store may only can gain a little benefits.

Anyway, we may have faster GC in ZCold in the future if I got time.

## Q: Why Zai can't protect my data 100%?

First, none of system can do 100%.

You can find more details about durability discussion in [Durability](conversation-street/durability.md).

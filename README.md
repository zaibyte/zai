Zai: User defined object storage with high performance and thousands trillion capacity
===

Zai is developed and maintained by Temple3x (temple3x@gmail.com) and his cat, Noodles.

<img src="./docs/img/n1.jpeg" width = "200"/>

Zai is a user defined object storage System suited for LOSF (lots of small files) with high performance and thousand trillion capacity.

- __Thousands Trillion Capacity__

    Size of an object can range from [1B to 1TB, up to 2814 trillion objects](./docs/FAQ.md) in a single namespace.
    
- __User Defined__

    Zai is easily upgrade/modify, building your own zai is strongly recommended. 
    
    Feel free to build filesystem or implement S3 protocol upon Zai, it's easy to gain all Zai features in this way. 
    I believe Zai could help you achieve business success.
    
    There is a [document](./docs/dev-guide/build_your_distribution.md) to show you
    how to set up highly customized storage service step by step.
    
- __High Performance__  

    Zai maybe the fastest open source object storage on GitHub. You could find a brief benchmark result [here](README.md#Performance).
    
- __High Availability__  
    
    No single point of failure in Zai. Automatic failure recovery and self-healing are provided.
    
- __High Durability__ 

    Replicas and Erasure Codes protect data together. 

- __Horizontal Scalability__

    Zai scales horizontally by simply adding new nodes. No rebalance issues.
        
- __Data Tiers__  
    
    Saving 90% TOC. Transport data to cheaper devices with lower storage redundancy automatically when data gets cold.
   
- __Generational Garbage Collection__   

    Speed up reclaiming disk space and reduce I/O impact in GC process, suited for applications which would delete files often 
        
- __Efficient Reconstruction__  

    Fast and efficient data reconstruction in [Erasure Codes](https://github.com/templexxx/xrs) (Saving about 30% I/O in reconstruction process).
    And there is no read amplification.
    
## Performance

    TODO
        
## Adopters

View the current list of in-production zai adopters [here](./docs/adopters.md).

## Roadmap

Read the [Roadmap](./docs/roadmap.md).

## Getting Started

Most readers will want to start with the [Getting Started](./docs/getting_started.md) or [FAQ](./docs/FAQ.md)

## Documentation

+ [Docs](./docs) 

## Architecture Overview

![architecture](./docs/img/arch.jpg)

## Contributing

I appreciate your help!

To report bugs, please open an [issue](https://github.com/zaibyte/zai/issues).

To submit patches, please send in the pull request.

## License

zai is under the Apache 2.0 license. See the [LICENSE](./LICENSE) file for details.

Deployment
===

## Basic Steps:

1. Start Keeper

2. Start ZBuf & ZCold(If has)

3. Start Application embed Zai

> **Note:**
>
> - All Zai components can start with this command:
  
    <cmd> -c <config_file_path>

> - Each Disk on write_buffer/storage nodes must be mounted under /zai`, and has a UUID as its name.
    This UUID will be the diskID in Zai.

> - Each ZBuf/ZCold nodes must has an unique nodeID, default is node's MAC address in hex.
        


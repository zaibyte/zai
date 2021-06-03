package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jaypipes/ghw/pkg/block"
)

var dev = flag.String("dev", "", "dev name")

func main() {
	flag.Parse()

	dn := *dev

	printIDs(dn)
}

func printIDs(dn string) {

	blk, err := block.New()
	if err != nil {
		log.Fatal(err)
	}

	if len(blk.Disks) == 0 {
		log.Fatal("cannot find any disk")
	}

	var found bool
	if dn == "" {
		found = true
	}

	for _, d := range blk.Disks {

		if dn != "" && dn != d.Name {
			continue
		}

		fmt.Println("-----------")
		fmt.Printf("dev_name: %s, S/N: %s\n", d.Name, d.SerialNumber)

		if len(d.Partitions) == 0 {
			fmt.Println("no partition for this device")
			continue
		}
		for _, p := range d.Partitions {
			fmt.Printf("part_name: %s, mnt_path: %s, uuid: %s\n",
				p.Name, p.MountPoint, p.UUID)
		}

	}

	if !found {
		fmt.Printf("dev: %s not found\n", dn)
	}
}

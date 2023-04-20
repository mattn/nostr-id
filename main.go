package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nbd-wtf/go-nostr/nip19"
)

func show(s string, err error) {
	if err == nil {
		fmt.Println(s)
	} else {
		log.Println(err)
	}
}

func main() {
	var d bool
	var typ string
	flag.BoolVar(&d, "d", false, "decode")
	flag.StringVar(&typ, "t", "npub", "type (note/nsec/npub)")
	flag.Parse()

	if d {
		for _, v := range flag.Args() {
			if _, vv, err := nip19.Decode(v); err == nil {
				fmt.Println(vv)
			} else {
				log.Println(err)
			}
		}
	} else {
		for _, v := range flag.Args() {
			switch typ {
			case "note":
				show(nip19.EncodeNote(v))
			case "nsec":
				show(nip19.EncodePrivateKey(v))
			case "npub":
				show(nip19.EncodePublicKey(v))
			}
		}
	}
}

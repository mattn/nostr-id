package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/urfave/cli/v2"
)

func showEncode(s string, err error) {
	if err == nil {
		fmt.Println(s)
	} else {
		log.Println(err)
	}
}

func showDecode(s string) {
	if _, vv, err := nip19.Decode(s); err == nil {
		fmt.Println(vv)
	} else {
		log.Println(err)
	}
}

func convertNote(cCtx *cli.Context) error {
	for _, arg := range cCtx.Args().Slice() {
		if strings.HasPrefix(arg, "note") {
			showDecode(arg)
		} else {
			showEncode(nip19.EncodeNote(arg))
		}
	}
	return nil
}

func convertPublicKey(cCtx *cli.Context) error {
	for _, arg := range cCtx.Args().Slice() {
		if strings.HasPrefix(arg, "npub") {
			showDecode(arg)
		} else {
			showEncode(nip19.EncodePublicKey(arg))
		}
	}
	return nil
}

func convertPrivateKey(cCtx *cli.Context) error {
	for _, arg := range cCtx.Args().Slice() {
		if strings.HasPrefix(arg, "nsec") {
			showDecode(arg)
		} else {
			showEncode(nip19.EncodePrivateKey(arg))
		}
	}
	return nil
}

func convertProfile(cCtx *cli.Context) error {
	for _, arg := range cCtx.Args().Slice() {
		if strings.HasPrefix(arg, "nprofile") {
			showDecode(arg)
		} else {
			showEncode(nip19.EncodeProfile(arg, cCtx.StringSlice("relay")))
		}
	}
	return nil
}

func convertEvent(cCtx *cli.Context) error {
	for _, arg := range cCtx.Args().Slice() {
		if strings.HasPrefix(arg, "nevent") {
			showDecode(arg)
		} else {
			showEncode(nip19.EncodeEvent(arg, cCtx.StringSlice("relay"), cCtx.String("author")))
		}
	}
	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "note",
				Usage:  "convert note/hex",
				Action: convertNote,
			},
			{
				Name:   "npub",
				Usage:  "convert npub/hex",
				Action: convertPublicKey,
			},
			{
				Name:   "nsec",
				Usage:  "convert nsec/hex",
				Action: convertPrivateKey,
			},
			{
				Name:   "nprofile",
				Usage:  "convert nprofile/hex",
				Action: convertProfile,
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name: "relay",
					},
				},
			},
			{
				Name:  "nevent",
				Usage: "convert nevent/hex",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name: "relay",
					},
					&cli.StringFlag{
						Name: "author",
					},
				},
				Action: convertEvent,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

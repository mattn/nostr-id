package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/urfave/cli/v2"
)

func decode(out io.Writer, s string) {
	if _, vv, err := nip19.Decode(s); err == nil {
		fmt.Fprintln(out, vv)
	} else {
		log.Println(err)
	}
}

func ConvertNote(cCtx *cli.Context) error {
	return convertNote(os.Stdout, cCtx.Args().Slice())
}

func convertNote(out io.Writer, args []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "note") {
			decode(out, arg)
		} else {
			v, err := nip19.EncodeNote(arg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(out, v)
		}
	}
	return nil
}

func ConvertPublicKey(cCtx *cli.Context) error {
	return convertPublicKey(os.Stdout, cCtx.Args().Slice())
}

func convertPublicKey(out io.Writer, args []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "npub") {
			decode(out, arg)
		} else {
			v, err := nip19.EncodePublicKey(arg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(out, v)
		}
	}
	return nil
}

func ConvertPrivateKey(cCtx *cli.Context) error {
	return convertPrivateKey(os.Stdout, cCtx.Args().Slice())
}

func convertPrivateKey(out io.Writer, args []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "nsec") {
			decode(out, arg)
		} else {
			v, err := nip19.EncodePrivateKey(arg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(out, v)
		}
	}
	return nil
}

func ConvertProfile(cCtx *cli.Context) error {
	return convertProfile(os.Stdout, cCtx.Args().Slice(), cCtx.StringSlice("relay"))
}

func convertProfile(out io.Writer, args []string, relays []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "nprofile") {
			decode(out, arg)
		} else {
			v, err := nip19.EncodeProfile(arg, relays)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(out, v)
		}
	}
	return nil
}

func ConvertEvent(cCtx *cli.Context) error {
	return convertEvent(os.Stdout, cCtx.Args().Slice(), cCtx.StringSlice("relay"), cCtx.String("author"))
}

func convertEvent(out io.Writer, args []string, relays []string, author string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "nevent") {
			decode(out, arg)
		} else {
			v, err := nip19.EncodeEvent(arg, relays, author)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(out, v)
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
				Action: ConvertNote,
			},
			{
				Name:   "npub",
				Usage:  "convert npub/hex",
				Action: ConvertPublicKey,
			},
			{
				Name:   "nsec",
				Usage:  "convert nsec/hex",
				Action: ConvertPrivateKey,
			},
			{
				Name:   "nprofile",
				Usage:  "convert nprofile/hex",
				Action: ConvertProfile,
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
				Action: ConvertEvent,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

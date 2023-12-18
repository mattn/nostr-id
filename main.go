package main

import (
	"bufio"
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
		fmt.Fprintf(out, "%+v\n", vv)
	} else {
		log.Println(err)
	}
}

func argsOrStdinLines(cCtx *cli.Context) []string {
	if cCtx.Args().Present() {
		return cCtx.Args().Slice()
	}
	var ret []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func ConvertNote(cCtx *cli.Context) error {
	return convertNote(os.Stdout, argsOrStdinLines(cCtx))
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
	return convertPublicKey(os.Stdout, argsOrStdinLines(cCtx))
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
	return convertPrivateKey(os.Stdout, argsOrStdinLines(cCtx))
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
	return convertProfile(os.Stdout, argsOrStdinLines(cCtx), cCtx.StringSlice("relay"))
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
	return convertEvent(os.Stdout, argsOrStdinLines(cCtx), cCtx.StringSlice("relay"), cCtx.String("author"))
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

func ConvertAddr(cCtx *cli.Context) error {
	return convertAddr(os.Stdout, argsOrStdinLines(cCtx), cCtx.Int("kind"), cCtx.String("author"), cCtx.StringSlice("relay"))
}

func convertAddr(out io.Writer, args []string, kind int, author string, relays []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "naddr") {
			decode(out, arg)
		} else {
			v, err := nip19.EncodeEntity(author, kind, arg, relays)
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
			{
				Name:  "naddr",
				Usage: "convert naddr/hex",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name: "relay",
					},
					&cli.StringFlag{
						Name: "author",
					},
				},
				Action: ConvertAddr,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestNote(t *testing.T) {
	var buf bytes.Buffer
	err := convertNote(&buf, []string{"note1n7caul29027wd472988h9tf69unff4yrv9dse6pk6zjak88fp68s2yynsg"})
	if err != nil {
		t.Fatal(err)
	}
	want := "9fb1de7d457abce6d7ca29cf72ad3a2f2694d483615b0ce836d0a5db1ce90e8f\n"
	if s := buf.String(); s != want {
		t.Fatalf("want %q but got %q", want, s)
	}
}

func TestPublicKey(t *testing.T) {
	var buf bytes.Buffer
	err := convertPublicKey(&buf, []string{"npub1937vv2nf06360qn9y8el6d8sevnndy7tuh5nzre4gj05xc32tnwqauhaj6"})
	if err != nil {
		t.Fatal(err)
	}
	want := "2c7cc62a697ea3a7826521f3fd34f0cb273693cbe5e9310f35449f43622a5cdc\n"
	if s := buf.String(); s != want {
		t.Fatalf("want %q but got %q", want, s)
	}
}

func TestPrivateKey(t *testing.T) {
	var buf bytes.Buffer
	err := convertPrivateKey(&buf, []string{"nsec180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsgyumg0"})
	if err != nil {
		t.Fatal(err)
	}
	want := "3bf0c63fcb93463407af97a5e5ee64fa883d107ef9e558472c4eb9aaaefa459d\n"
	if s := buf.String(); s != want {
		t.Fatalf("want %q but got %q", want, s)
	}
}

func TestProfile(t *testing.T) {
	var buf bytes.Buffer
	err := convertProfile(&buf, []string{"nprofile1qqsrhuxx8l9ex335q7he0f09aej04zpazpl0ne2cgukyawd24mayt8gpp4mhxue69uhhytnc9e3k7mgpz4mhxue69uhkg6nzv9ejuumpv34kytnrdaksjlyr9p"}, []string{"wss://example.com"})
	if err != nil {
		t.Fatal(err)
	}
	want := "{3bf0c63fcb93463407af97a5e5ee64fa883d107ef9e558472c4eb9aaaefa459d [wss://r.x.com wss://djbas.sadkb.com]}\n"
	if s := buf.String(); s != want {
		t.Fatalf("want %q but got %q", want, s)
	}
}

func TestEvent(t *testing.T) {
	var buf bytes.Buffer
	err := convertEvent(&buf, []string{"nevent1qqs8p8522pzp8hc2paa2yxqypu6jxpqxdhxywtelfcmvqrx703k5v4qprpmhxue69uhkummnw3ezu6r0d3ukyetp9e3k7mf0qgs2808qjhxsaq6gkz2vjgqm6rpjwr404au0waj30aqszvnd4awt0lcrqsqqqqqpu0ur69"}, []string{"wss://example.com"}, "npub1937vv2nf06360qn9y8el6d8sevnndy7tuh5nzre4gj05xc32tnwqauhaj6")
	if err != nil {
		t.Fatal(err)
	}
	want := "{709e8a504413df0a0f7aa218040f352304066dcc472f3f4e36c00cde7c6d4654 [wss://nostr.holybea.com/] a3bce095cd0e8348b094c9201bd0c3270eafaf78f776517f4101326daf5cb7ff}\n"
	if s := buf.String(); s != want {
		t.Fatalf("want %q but got %q", want, s)
	}
}

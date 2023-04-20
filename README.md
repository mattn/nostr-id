# nostr-id

ID converter for nostr

## Usage

Decode npub to hex public key

```
$ nostr-id -d npub1a3pvwe2p3v7mnjz6hle63r628wl9w567aw7u23fzqs062v5vqcqqu3sgh3
ec42c765418b3db9c85abff3a88f4a3bbe57535eebbdc54522041fa5328c0600
```

Encode hex public key to npub

```
$ nostr-id ec42c765418b3db9c85abff3a88f4a3bbe57535eebbdc54522041fa5328c0600
npub1a3pvwe2p3v7mnjz6hle63r628wl9w567aw7u23fzqs062v5vqcqqu3sgh3
```

Encode hex note id to note

```
$ nostr-id -t note fbf26c1c5690773b3ae6b9abff884f65969f1be27d8978bcb8584ee9ac3855e0
note1l0exc8zkjpmnkwhxhx4llzz0vktf7xlz0kyh309ctp8wntpc2hsqqhn0mt
```

Decode note to hex note id

```
$ nostr-id -d note1l0exc8zkjpmnkwhxhx4llzz0vktf7xlz0kyh309ctp8wntpc2hsqqhn0mt
fbf26c1c5690773b3ae6b9abff884f65969f1be27d8978bcb8584ee9ac3855e0
```

## Installation

```
go install github.com/mattn/nostr-id@latest
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)

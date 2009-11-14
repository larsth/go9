package main

import fmt "fmt"

const (
	VERSION = "9P2000";
	NOTAG = 0xffff;
	NOFID = 0xffffffff;
	
	MAX_VERSION = 32;
	MAX_MSG = 8192;
	MAX_ERROR = 128;
	MAX_CACHE = 32;
	MAX_FLEN = 128;
	MAX_ULEN = 32;
	MAX_WELEM = 16;

	// From libc.h in p9p
	OREAD = 0;
	OWRITE = 1;
	ORDWR = 2;
	OEXEC = 3;
	OTRUNC = 16;
	OCEXEC = 32;
	ORCLOSE = 64;
	ODIRECT = 128;
	ONONBLOCK = 256;
	OEXCL = 0x1000;
	OLOCK = 0x2000;
	OAPPEND = 0x4000;

	// Bits for Qid
	QTDIR = 0x80;
	QTAPPEND = 0x40;
	QTEXCL = 0x20;
	QTMOUNT = 0x10;
	QTAUTH = 0x08;
	QTTMP = 0x04;
	QTSYMLINK = 0x02;
	QTFILE = 0x00;

	// Bits for Dir
	DMDIR = 0x80000000;
	DMAPPEND = 0x40000000;	
	DMEXCL = 0x20000000;	
	DMMOUNT = 0x10000000;
	DMAUTH = 0x08000000;	
	DMTMP = 0x04000000;
)

type Message struct {
	name string;
	data []byte;
}

// Initialize message formats
func prepareMessages() (msg map[string] int, msgFmt map[int] string) {
	j := 0;
	num := 100;
	msgs := []string{"version", "auth", "attach", "error", "flush", "walk",
		"open", "create", "read", "write", "clunk", "remove", "stat", "wstat"};
	fmts := []string{"4S", "4S", "4SS", "Q", "44SS", "Q", "", "S", "2", "",
		"{Twalk}", "{Rwalk}", "41", "Q4", "4S41", "Q4", "484", "D", "48D",
		"4", "4", "", "4", "", "4", "{Stat}", "4{Stat}", ""};
	
	message := make(map[string] int);
	messageFmt := make(map[int] string);
	for i := 0; i < len(msgs); i++ {
		message["T" + msgs[i]] = num;
		message["R" + msgs[i]] = num + 1;
		messageFmt[num] = fmts[j];
		messageFmt[num + 1] = fmts[j + 1];
		
		j += 2;
		num += 2;
	}
	
	return message, messageFmt;
}

func main() {
	message, messageFmt := prepareMessages();
	fmt.Printf("%v\n%v\n", message, messageFmt);
}

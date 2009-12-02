package main

import (
	"io";
	"fmt";
	"bufio";
	"reflect";
	"strings";
	"encoding/base64";
	"encoding/binary";
)

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
)

const (
	// From libc.h in p9p
	OREAD = iota;
	OWRITE;
	ORDWR;
	OEXEC;
	OTRUNC = 16;
	OCEXEC = 32;
	ORCLOSE = 64;
	ODIRECT = 128;
	ONONBLOCK = 256;
	OEXCL = 0x1000;
	OLOCK = 0x2000;
	OAPPEND = 0x4000;
)

const (
	// Bits for Qid
	QTFILE = 1 << iota;
	QTSYMLINK;
	QTTMP;
	QTAUTH;
	QTMOUNT;
	QTEXCL;
	QTAPPEND;
	QTDIR;
)

const (
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

func sendTwalk(buf io.Writer, args ...) {
	
}

func sendRwalk(buf io.Writer, args ...) {
	
}

func sendStat(buf io.Writer, args ...) {
	
}

func sendMessage(socket io.Writer, tag uint16, typ uint8, args ...) {
	en = binary.BigEndian;
	buf := bufio.NewWriter();
	binary.Write(buf, en, typ);
	binary.Write(buf, en, tag);
	
	argnum := 0;
	format := strings.Bytes(messageFmt[typ]);
	fields := reflect.NewValue(args).(*reflect.StructValue);
	for i := 0; i < len(format); i++ {
		switch format[i] {
		case '1':
			// 1 byte number
			binary.Write(buf, en, int8(fields.Field(argnum));
		case '2':
			// 2 byte number
			binary.Write(buf, en, int16(fields.Field(argnum));
		case '4':
			// 4 byte number
			binary.Write(buf, en, int32(fields.Field(argnum));
		case '8':
			// 8 byte number
			binary.Write(buf, en, int64(fields.Field(argnum));
		case 'S':
			// Variable string with 2 byte length
			str := fields.Field(argnum).Get();
			binary.Write(buf, en, int16(len(str)));
			buf.Write(strings.Bytes(str));
		case 'D':
			// Variable string with 4 byte length
			str := fields.Field(argnum).Get();
			binary.Write(buf, en, int32(len(str)));
			buf.Write(strings.Bytes(str));
		case 'Q':
			// QID 13byte value = [type, version, path]
			binary.Write(buf, en, int8(fields.Field(argnum));
			binary.Write(buf, en, int16(fields.Field(argnum + 1));
			binary.Write(buf, en, int64(fields.Field(argnum + 2));
		case '{':
			k := i + 1;
			ftype := "";
			while format[k] != '}' {
				ftype += format[k];
				k += 1;
			}
			switch format {
				case "Twalk":
					sendTwalk(buf, args);
				case "Rwalk":
					sendRwalk(buf, args);
				case "Stat":
					sendStat(buf, args);
			}
		default:
			// Invalid format type!
		}
		j += 1;
	}	
}

func main() {
	message, messageFmt := prepareMessages();
	fmt.Printf("%v\n%v\n", message, messageFmt);
}

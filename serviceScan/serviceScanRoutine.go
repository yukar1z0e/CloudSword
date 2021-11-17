package serviceScan

import (
	"CloudSword/structure"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// assume_http 检测HTTPS和HTTP
func assume_http(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	ret := structure.HostInfoStruct{Host: host, Port: port}
	t := []byte{0x16, 0x03, 0x01, 0x00, 0xb5, 0x01, 0x00, 0x00, 0xb1, 0x03, 0x03, 0xb2, 0xd3, 0x4d, 0xfd, 0x63, 0xbe, 0x89, 0xdb, 0xe5, 0x46, 0xcc, 0xaf, 0x39, 0x6e, 0xba, 0x63, 0x63, 0x75, 0xce, 0x30, 0xda, 0xe0, 0x4f, 0xab, 0xa2, 0x3e, 0x50, 0xea, 0x41, 0x20, 0x10, 0xc4, 0x00, 0x00, 0x18, 0xc0, 0x2b, 0xc0, 0x2f, 0xc0, 0x2c, 0xc0, 0x30, 0xc0, 0x13, 0xc0, 0x14, 0x00, 0x9c, 0x00, 0x9d, 0x00, 0x2f, 0x00, 0x35, 0x00, 0x0a, 0x00, 0xff, 0x01, 0x00, 0x00, 0x70, 0x00, 0x00, 0x00, 0x15, 0x00, 0x13, 0x00, 0x00, 0x10, 0x77, 0x77, 0x77, 0x2e, 0x73, 0x6f, 0x2d, 0x63, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x0b, 0x00, 0x04, 0x03, 0x00, 0x01, 0x02, 0x00, 0x0a, 0x00, 0x06, 0x00, 0x04, 0x00, 0x17, 0x00, 0x18, 0x00, 0x23, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x20, 0x00, 0x1e, 0x06, 0x01, 0x06, 0x02, 0x06, 0x03, 0x05, 0x01, 0x05, 0x02, 0x05, 0x03, 0x04, 0x01, 0x04, 0x02, 0x04, 0x03, 0x03, 0x01, 0x03, 0x02, 0x03, 0x03, 0x02, 0x01, 0x02, 0x02, 0x02, 0x03, 0x00, 0x05, 0x00, 0x05, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x01, 0x01, 0x00, 0x10, 0x00, 0x0b, 0x00, 0x09, 0x08, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x31, 0x2e, 0x31}

	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 1024)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		if err != nil {
			ret.Service = ""
		}
		if string(recvBuf[0:4]) == string([]byte{22, 3, 3, 0}) || port == 443 {
			ret.Service = "https"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: HTTPS")
		} else if strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("HTTP")) || port == 80 {
			ret.Service = "http"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: HTTP")
		}
		result <- ret

	}
	conn.Close()
}

// ssh
func assume_ssh(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte("test ssh")
	ret := structure.HostInfoStruct{Host: host, Port: port}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 1024)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(string(recvBuf))
		if err != nil {
			ret.Service = ""
		}
		if strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("ssh")) || port == 22 {
			ret.Service = "ssh"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: SSH")
		}
		result <- ret

	}
	conn.Close()
}

// ftp
func assume_ftp(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte("test ftp")
	ret := structure.HostInfoStruct{Host: host, Port: port}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 1024)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(string(recvBuf))
		if err != nil {
			ret.Service = ""
		}
		if strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("ftp")) || strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("filezilla")) || strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("220")) || port == 21 {
			ret.Service = "ftp"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: FTP")
		}

		result <- ret

	}
	conn.Close()
}

// mysql
func assume_mysql(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte("test mysql")
	ret := structure.HostInfoStruct{Host: host, Port: port}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 1024)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(string(recvBuf))
		if err != nil {
			ret.Service = ""
		}
		if strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("mysql")) || port == 3306 {
			ret.Service = "mysql"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: MYSQL")
		}
		result <- ret

	}
	conn.Close()
}

// mssql
func assume_mssql(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte{0x12, 0x01, 0x00, 0x2F, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x1A, 0x00, 0x06, 0x01, 0x00, 0x20, 0x00, 0x01, 0x02, 0x00, 0x21, 0x00, 0x01, 0x03, 0x00, 0x22, 0x00, 0x04, 0x04, 0x00, 0x26, 0x00, 0x01, 0xFF, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0xB8, 0x0D, 0x00, 0x00, 0x01}
	ret := structure.HostInfoStruct{Host: host, Port: port}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 2048)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(recvBuf)
		if err != nil {
			ret.Service = ""
		}
		if string(recvBuf[0:4]) == string([]byte{4, 1, 0, 43}) || port == 1433 {
			ret.Service = "mssql"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: MSSQL")
		}
		result <- ret

	}
	conn.Close()
}

// mongodb
func assume_mongodb(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte{0x3f, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xd4, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x24, 0x63, 0x6d, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00}
	ret := structure.HostInfoStruct{Host: host, Port: port}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 2048)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(string(recvBuf))
		if err != nil {
			ret.Service = ""
		}
		if strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("command")) || strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("Unauthorized")) || strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("listDatabases")) || port == 27017 {
			ret.Service = "mongodb"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: Mongodb")
		}
		result <- ret

	}
	conn.Close()
}

// redis
func assume_redis(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte{0x2A, 0x32, 0x0D, 0x24, 0x34, 0x0D, 0x41, 0x55, 0x54, 0x48, 0x0D, 0x24, 0x35, 0x0D, 0x31, 0x32, 0x33, 0x34, 0x35}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	ret := structure.HostInfoStruct{Host: host, Port: port}
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 1024)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(string(recvBuf))
		if err != nil {
			ret.Service = ""
		}
		if strings.Contains(strings.ToLower(string(recvBuf)), strings.ToLower("-err")) || port == 6379 {
			ret.Service = "redis"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: Redis")
		}
		result <- ret

	}
	conn.Close()
}

// telnet
func assume_telnet(host string, port int, timeout time.Duration, result chan structure.HostInfoStruct) {
	t := []byte("test telnet")
	ret := structure.HostInfoStruct{Host: host, Port: port}
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), timeout)
	//defer conn.Close()
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		ret.Service = ""
	} else {
		conn.Write(t)
		recvBuf := make([]byte, 1024)
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			ret.Service = ""
		}
		_, err = conn.Read(recvBuf[:])
		err = conn.SetReadDeadline(time.Time{})
		//fmt.Println(recvBuf[0:3])
		if err != nil {
			ret.Service = ""
		}
		if bytes.Contains(recvBuf[0:3], []byte{255}) || bytes.Contains(recvBuf[0:3], []byte{253}) || port == 23 {
			ret.Service = "telnet"
			fmt.Println("Host: " + host + " Port: " + strconv.Itoa(port) + " Service: Telnet")
		}
		result <- ret

	}
	conn.Close()
}

package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/hashicorp/memberlist"
	"net"
	"net/rpc"
	"strconv"
)

type Status struct {
	Role  string
	State string
=======
	"net"
	"net/rpc"
	"strconv"
	"time"

	"github.com/nanobox-core/scribble"
)

//
type(

	//
	Status struct {
		CRole    string
		DBRole   string
		State  	 string
		UpdateAt time.Time
	}
)

var(
	client *rpc.Client
	status *Status
	store  *scribble.Driver
)

//
func StatusStart() error {

	//
	port := strconv.FormatInt(int64(conf.ClusterPort+1), 10)

	//
	s := Status{CRole: conf.Role, State: "booting"}

	//
	store = scribble.New("./status", log)
	t := scribble.Transaction{Operation: "write", Collection: "cluster", RecordID: "node", Container: &s}
	if err := store.Transact(t); err != nil {
		return err
	}

	//
	rpc.Register(s)

	// RPC SERVER
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	//
	go func(l net.Listener) {
		for {
			if conn, err := l.Accept(); err != nil {
				fmt.Println("accept error: " + err.Error())
			} else {
				fmt.Printf("new connection established\n")
				go rpc.ServeConn(conn)
			}
		}
	}(l)

	// RPC CLIENT
	client, err = rpc.Dial("tcp", port)
 	if err != nil {
		return err
	}

	return nil
>>>>>>> working on rpc server and client
}

// 'public' methods

<<<<<<< HEAD
func (self *Status) save() {

=======
//
func (s *Status) SetCRole(role string) {
	s.CRole = role
	if err := save(s); err != nil {
		log.Fatal("BONK!", err)
		panic("Unable to set set cluster role! " + err.Error())
	}
>>>>>>> working on rpc server and client
}

//
func (s *Status) SetDBRole(role string) {
	s.DBRole = role
	if err := save(s); err != nil {
		log.Fatal("BONK!", err)
		panic("Unable to set db role! " + err.Error())
	}
}

//
func (s *Status) SetState(state string) {
	s.State = state
	if err := save(s); err != nil {
		log.Fatal("BONK!", err)
		panic("Unable to set state! " + err.Error())
	}
}

<<<<<<< HEAD
func (self *Status) Get(who string, reply *Status) error {
	fmt.Println("whos asking: " + who)
	self.retrieve()
	reply = self
	return nil
=======
// 'public' (RPC mapped) function

//
func Whoami() (*Status, error) {
	fmt.Println("WHAT IS THIS", list.LocalNode())

	s := &Status{}

	if err := client.Call("Status.whoami", nil, s); err != nil {
		return nil, err
	}

	return s, nil
}

//
func Whois(role string) (*Status, error) {
	s := &Status{}

	if err := client.Call("Status.whois", role, s); err != nil {
		return nil, err
	}

	return s, nil
>>>>>>> working on rpc server and client
}

//
func Cluster() ([]*Status, error) {
	var members = []*Status{}

	if err := client.Call("Status.cluster", nil, &members); err != nil {
		return nil, err
	}

	return members, nil
}

//
func Demote() error {
	if err := client.Call("Status.demote", nil, nil); err != nil {
		return err
	}

	return nil
}

// 'private' (RPC) methods

//
func (s *Status) whoami(v interface{}) error {

	//
	if err := get(s.CRole, v); err != nil {
		return err
	}

	return nil
}

//
func (s *Status) whois(role string, v interface{}) error {

	//
	for _, m := range list.Members() {
		if m.Name == role {
			if err := get(role, v); err != nil {
				return err
			}
		}
	}

	return nil
}

//
func (s *Status) cluster(v []*Status) error {

	//
	for range list.Members() {

		//
		status, err := Whoami()
		if err != nil {
			return err
		}

		//
		v = append(v, status)
	}

	return nil
}

//
func (s *Status) demote() (error) {
	return nil
}

<<<<<<< HEAD
var status *Status

func RpcStart() error {
	status = new(Status)
	rpc.Register(status)
	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(conf.ClusterPort+1), 10))
	if err != nil {
		return err
	}

	go func(listener net.Listener) {
		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Println("accept error: " + err.Error())
			} else {
				fmt.Printf("new connection established\n")
				go rpc.ServeConn(conn)
			}
		}
	}(listener)
=======
// 'private' functions

//
func get(role string, v interface{}) error {
	t := scribble.Transaction{Operation: "read", Collection: "cluster", RecordID: role, Container: &v}
	if err := store.Transact(t); err != nil {
		return err
	}

	return nil
}

//
func save(v interface{}) error {
	t := scribble.Transaction{Operation: "write", Collection: "cluster", RecordID: "node", Container: &v}
	if err := store.Transact(t); err != nil {
		return err
	}

>>>>>>> working on rpc server and client
	return nil
}

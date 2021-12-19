package core

import (
	"log"

	"github.com/gocql/gocql"
)

type Session struct {
	cqlSession *gocql.Session
}

func (session *Session) Get() *gocql.Session {
	return session.cqlSession
}

func (session *Session) Instance() *Session {
	return session
}

func (session *Session) IsActive() bool { return session.cqlSession != nil }

func (session *Session) New() (*Session, error) {
	if session.IsActive() {
		return session, nil
	}

	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	cluster.Keyspace = "microgo"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	_session, err := cluster.CreateSession()

	if err != nil {
		log.Fatalln("Error initializing new session")
		return nil, err
	}

	log.Println("Initialized new session")
	session.cqlSession = _session
	return session, nil
}

func (session *Session) Close() {
	session.cqlSession.Close()
	session.cqlSession = nil
	log.Println("Closed active session")
}

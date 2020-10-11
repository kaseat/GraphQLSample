package main

type User struct {
	FirstName string
	LastName  string
}

const schema = `
type Query {
	user: User!
}
type User {
	firstName: String!
	lastName: String!
}
`

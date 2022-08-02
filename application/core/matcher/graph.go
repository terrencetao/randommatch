// Package graph creates a ItemGraph data structure for the Item type
package matcher

import (
	"fmt"
	//"github.com/cheekybits/genny/generic"
)

// Item the type of the binary search tree
//type Item generic.Type

// Node a single node that composes the tree
type User struct {
	UserId string `json:"userId"`
}

func (n *User) String() string {
	return n.UserId
}

// UserGraph the Items graph
type UserGraph struct {
	users []*User
	edges map[User][]*User
	// lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *UserGraph) AddUser(n *User) {
	if find, _ := Search(g.users, n); !find {
		g.users = append(g.users, n)
	}

}

// AddEdge adds an edge to the graph
func (g *UserGraph) AddEdge(n1, n2 *User) {
	if g.edges == nil {
		g.edges = make(map[User][]*User)
	}
	if find, _ := Search(g.edges[*n1], n2); !find {
		g.edges[*n1] = append(g.edges[*n1], n2)
	}
	if find, _ := Search(g.edges[*n2], n1); !find {
		g.edges[*n2] = append(g.edges[*n2], n1)
	}
}

// search a user in a list of user
func Search(users []*User, n *User) (bool, int) {
	index := -1
	find := false
	for i, user := range users {
		if user.UserId == n.UserId {
			find = true
			index = i
			break
		}
	}

	return find, index
}

// SearchNode findout a specifique node in a graph

func (g *UserGraph) SearchUser(n *User) (bool, int) {
	index := -1
	find := false
	if g.users != nil {
		find, index = Search(g.users, n)
	}

	return find, index

}

// remove a user in a list of users

func Remove(s []*User, i int) []*User {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// RemoveEdge remove an edge from the graph
func (g *UserGraph) RemoveEdge(n *User) {
	for _, user := range g.users {
		find, index := Search(g.edges[*user], n)
		if find {
			g.edges[*user] = Remove(g.edges[*user], index)
		}
		delete(g.edges, *n)
	}
}

// RemoveUser remove a user from the graph
func (g *UserGraph) RemoveUser(n *User) {

	g.RemoveEdge(n)
	find, index := g.SearchUser(n) // find out the index of this node
	if find {
		g.users = Remove(g.users, index)

	}

}

// print the graph
func (g *UserGraph) String() {
	s := ""
	for _, usernode := range g.users {
		s += usernode.String() + " -> "
		near := g.edges[*usernode]
		for _, user := range near {
			s += user.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)

}

// UsersToGraph create a graph of some users and connections

func UsersToGraph(users []User, connections [][]User) *UserGraph {

	/* input :
	         users: table users for matching
			 connections: Matrix for connection in the graph;
			               for each line:
						      - first element is connected to all the rest
	   output: return the graph
	*/
	var graph UserGraph
	for _, user := range users {
		user := user
		graph.AddUser(&user)
	}
	for _, usersAlreadyMatch := range connections {
		if len(usersAlreadyMatch) > 0 {
			node := usersAlreadyMatch[0]
			for _, user := range usersAlreadyMatch[1:] {
				user := user
				graph.AddEdge(&node, &user)
			}
		}
	}
	return &graph
}

// Subgraph extract a subgraph G' from a graph G
func (g *UserGraph) Subgraph(users []*User) *UserGraph {
	var subG UserGraph
	subG.edges = make(map[User][]*User)
	subG.users = users

	for _, user := range users {
		subG.edges[*user] = g.edges[*user]
	}
	return &subG
}

package matcher

import (
	"fmt"
	"testing"
)

func TestRandomChoices(t *testing.T) {
	//id := []string{"2", "5", "6", "8", "10", "12", "24", "25"}
	g.String()
	var forbiddenConnections [][]User
	constraint := []Constraint{Dejavu}
	matching := RandomChoices(&g, 2, constraint, forbiddenConnections)
	fmt.Printf("Match of %d: [", len(matching.Users))
	for _, user := range matching.Users {
		fmt.Printf("%s,", user.String())
	}

	fmt.Printf("]")
}

func TestRanSubGroup(t *testing.T) {
	g.String()
	var forbiddenConnections [][]User
	interConstraint := []Constraint{Dejavu}
	A := []*User{&User{"5"}}
	subA := g.Subgraph(A)
	fmt.Println("Sous groupe A")
	subA.String()
	B := []*User{&User{"3"}}
	fmt.Println("Sous groupe B")
	subB := g.Subgraph(B)
	subB.String()
	matching := RandSubGroup(subA, subB, 1, 1, interConstraint, []Constraint{Dejavu}, forbiddenConnections)
	fmt.Printf("Match of %d: [", len(matching.Users))
	for _, user := range matching.Users {
		fmt.Printf("%s,", user.String())
	}

	fmt.Printf("]")
}
func TestMatcher1(t *testing.T) {

	g.String()
	var forbiddenConnections [][]User
	A := []User{User{"2"}, User{"1"}}
	forbiddenConnections = append(forbiddenConnections, A)
	constraint := []Constraint{Dejavu}
	SELECTOR := Basic
	matching := Matcher(&g, 2, constraint, SELECTOR, forbiddenConnections, []*User{}, []*User{}, 0, 0, []Constraint{}, []Constraint{})

	for _, match := range matching {
		fmt.Printf("Match : [")
		for _, user := range match.Users {
			fmt.Printf("%s,", user.String())

		}

		fmt.Printf("]")
		fmt.Println("")

	}
	g.String()
}
func TestMatcher2(t *testing.T) {
	var G UserGraph
	nA := User{"1"}
	nB := User{"2"}
	nC := User{"3"}
	nD := User{"4"}
	nE := User{"5"}
	nF := User{"6"}
	G.AddUser(&nA)
	G.AddUser(&nB)
	G.AddUser(&nC)
	G.AddUser(&nD)
	G.AddUser(&nE)
	G.AddUser(&nF)

	G.AddEdge(&nA, &nB)
	G.AddEdge(&nA, &nC)
	G.AddEdge(&nB, &nE)
	G.AddEdge(&nF, &nE)
	G.AddEdge(&nA, &nE)

	G.String()
	var forbiddenConnections [][]User
	A := []*User{&User{"4"}, &User{"5"}, &User{"6"}}
	B := []*User{&User{"1"}, &User{"2"}, &User{"3"}}
	interConstraint := []Constraint{Dejavu}
	SELECTOR := Group
	matching := Matcher(&G, 0, []Constraint{}, SELECTOR, forbiddenConnections, A, B, 1, 1, interConstraint, []Constraint{})

	for _, match := range matching {
		fmt.Printf("Match : [")
		for _, user := range match.Users {
			fmt.Printf("%s,", user.String())

		}

		fmt.Printf("]")
		fmt.Println("")

	}

}

func TestGenTuple(t *testing.T) {

	users := []User{User{"1"}, User{"2"}, User{"3"}, User{"4"}, User{"5"}, User{"6"}}
	var connections, forbiddenConnections [][]User
	A := []User{User{"4"}, User{"5"}, User{"6"}}
	B := []User{User{"1"}, User{"2"}, User{"3"}}

	matching := GenerateTuple(users, connections, Group, forbiddenConnections, 0, A, B, 1, 1)

	for _, match := range matching {
		fmt.Printf("Match : [")
		for _, user := range match.Users {
			fmt.Printf("%s,", user.String())
		}

		fmt.Printf("]")
		fmt.Println("")
	}
}

package main

import "fmt"

// Takes graph with names, returns name of the nearest one
// who meets the condition
func search(name string, graph map[string][]string) bool {

	// Stores persons in queue
	queue := []string{}

	// Stores persons that are already searched
	searched := []string{}
	queue = append(queue, graph[name]...)
	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]
		fmt.Println(person)
		if !person_in_searched(person, searched) {
			if person_is_seller(person) {
				fmt.Println(person + " is a mango seller")
				return true
			}
			queue = append(queue, graph[person]...)
			searched = append(searched, person)
		}

	}
	return false
}

func person_in_searched(person string, searched []string) bool {
	for _, n := range searched {
		if n == person {
			return true
		}
	}
	return false
}

func person_is_seller(name string) bool {
	return name[len(name)-1] == 'j'
}

func main() {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	search("you", graph)
}

// given two employees, find the closest shared manager
package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	id        int
	name      string
	manager   *Employee
	reportees []*Employee
}

func findNearestBoss(e1, e2 *Employee) (*Employee, error) {
	bossMap := map[int]bool{}

	// check whether e1 or e2 is ceo
	if e1.manager == nil {
		return nil, errors.New("e1 is the ceo")
	} else if e2.manager == nil {
		return nil, errors.New("e2 is the ceo")
	}

	// iterate through e1's boss chain until the CEO (where manager == nil)
	// write the id's of each boss in the chain to bosses map
	for n := e1.manager; n != nil; n = n.manager {
		bossMap[n.id] = true
	}

	// iterate through e2's boss chain, check if each employee is contained in
	// e1's boss chain. The first one which is contained will be the closest shared boss
	for n := e2.manager; n != nil; n = n.manager {
		if _, ok := bossMap[n.id]; ok {
			return n, nil
		}
	}

	return nil, errors.New("no shared boss found")
}

func main() {
	// employee hierarchy
	// 		  ceo: joe
	//		/			\
	//	e2: bob		e3: jake
	//				/		\
	//			e4: jess	e5: nick
	//							\
	//							e6: dan

	ceo := &Employee{id: 1, name: "joe"}
	e2 := &Employee{id: 2, name: "bob"}
	e3 := &Employee{id: 3, name: "jake"}
	e4 := &Employee{id: 4, name: "jess"}
	e5 := &Employee{id: 5, name: "nick"}
	e6 := &Employee{id: 6, name: "dan"}

	ceo.reportees = []*Employee{e2, e3}
	e2.manager = ceo
	e3.manager = ceo
	e3.reportees = []*Employee{e4, e5}
	e4.manager = e3
	e5.manager = e3
	e5.reportees = []*Employee{e6}
	e6.manager = e5

	nearestBoss, err := findNearestBoss(e4, e6)
	if err != nil {
		fmt.Printf("Err: %v", err.Error())
	} else {
		// prints jake
		fmt.Printf("Nearest boss of %v and %v is %v\n", e4.name, e6.name, nearestBoss.name)
	}

	nearestBoss, err = findNearestBoss(e2, e6)
	if err != nil {
		fmt.Printf("Err: %v", err.Error())
	} else {
		// prints joe
		fmt.Printf("Nearest boss of %v and %v is %v\n", e2.name, e6.name, nearestBoss.name)
	}

	_, err = findNearestBoss(ceo, e6)
	if err != nil {
		// should print error that one of the employees is the ceo
		fmt.Printf("Err: %v", err.Error())
	} else {
		fmt.Println("Something is wrong, this should not happen")
	}
}

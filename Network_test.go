package main

import (
	"testing"
)

func TestGetNetworkInstance(t *testing.T) {
	n := GetNetworkInstance(5, 5, 0)

	if n.numConnections < 10 || n.id < 9 {
		t.Errorf("we got %d and %d", n.numConnections, n.id)
	}
}

func TestCreateNode(t *testing.T) {
	n := GetNetworkInstance(5, 5, 0)

	if n.createNode().id != 10 || n.createNode().id != 11 {
		t.Errorf("we got %d and %d", n.createNode().id, n.createNode().id)
	}
}
func TestGetNode(t *testing.T) {
	n := GetNetworkInstance(5, 5, 0)

	if n.getNode(0).id != 0 || n.getNode(9).id != 9 {
		t.Errorf("we got", n.getNode(0).id)
	}
}

func TestMutateNode(t *testing.T) {
	n := GetNetworkInstance(5, 5, 0)

	n.mutateNode(0, 8, 10, 11)

	ans := true
	ansA := false
	ansB := false
	//the node to is nil (because this has default initial
	for i := 0; i < len(n.getNode(0).send); i++ {
		if n.getNode(0).getSendCon(i).nodeTo != nil && n.getNode(0).getSendCon(i).nodeTo.id == 8 {
			ans = false
		} else if n.getNode(0).getSendCon(i).nodeTo != nil && n.getNode(0).getSendCon(i).nodeTo.id == 11 {
			ansA = true

			for a := 0; a < len(n.getNode(0).getSendCon(i).nodeTo.send); a++ {
				if n.getNode(0).getSendCon(i).nodeTo.getSendCon(a).nodeTo.id == 8 {
					ansB = true
				}
			}
		}
	}

	if !ans && ansA && ansB {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", ans, ansA)
	}
}

func TestMutateConnection(t *testing.T) {

}

func TestAddInnovation(t *testing.T) {

}

func TestRemoveInnovation(t *testing.T) {

}

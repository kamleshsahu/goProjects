package main

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashRing struct {
	nodes         []Node
	vnodesPerNode int
}

type Node struct {
	hash   int
	server string
}

func NewHashRing(vnodesPerNode int) *HashRing {
	return &HashRing{
		vnodesPerNode: vnodesPerNode,
	}
}

func (h *HashRing) AddServer(server string) {
	for i := 0; i < h.vnodesPerNode; i++ {
		vnode := server + "#" + strconv.Itoa(i)
		hash := int(crc32.ChecksumIEEE([]byte(vnode)))
		h.nodes = append(h.nodes, Node{hash: hash, server: server})
	}
	sort.Slice(h.nodes, func(i, j int) bool {
		return h.nodes[i].hash < h.nodes[j].hash
	})
}

func (h *HashRing) RemoveServer(server string) {
	var newNodes []Node
	for _, node := range h.nodes {
		if node.server != server {
			newNodes = append(newNodes, node)
		}
	}
	h.nodes = newNodes
}

func (h *HashRing) GetServer(key string) string {
	if len(h.nodes) == 0 {
		return ""
	}
	hash := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i].hash >= hash
	})
	if idx == len(h.nodes) {
		idx = 0
	}
	return h.nodes[idx].server
}

func main() {
	// Example usage
	hashRing := NewHashRing(10) // 10 virtual nodes per server
	hashRing.AddServer("Server1")
	hashRing.AddServer("Server2")
	hashRing.AddServer("Server3")
	hashRing.AddServer("Server4")
	hashRing.AddServer("Server5")

	// Add more servers or remove them as needed
	// Get the server for a given key
	server1 := hashRing.GetServer("some_key2")
	server3 := hashRing.GetServer("some_key3")
	server4 := hashRing.GetServer("some_key4")
	server := hashRing.GetServer("some_key")

	println(server1)
	println(server3)
	println(server4)
	println(server)
	hashRing.RemoveServer("Server2")
	hashRing.RemoveServer("Server3")

	println()
	server1 = hashRing.GetServer("some_key2")
	server3 = hashRing.GetServer("some_key3")
	server4 = hashRing.GetServer("some_key4")
	server = hashRing.GetServer("some_key")

	println(server1)
	println(server3)
	println(server4)
	println(server)
}

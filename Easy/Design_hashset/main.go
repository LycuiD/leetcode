// https://leetcode.com/problems/design-hashset/
package main

type MyHashSet struct {
	set []bool
}

func Constructor() MyHashSet {
	return MyHashSet{set: make([]bool, 1e6+1)}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.set[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.set[key]
}

func main() {}

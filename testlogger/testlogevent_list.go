// An encapsulated []*TestLogEvent.
// Thread-safe.
//
// Generated from threadsafe/list.tpl with Type=*TestLogEvent
// options: Comparable:<no value> Numeric:<no value> Integer:<no value> Ordered:<no value>
//          StringLike:<no value> StringParser:<no value> Stringer:<no value>
// GobEncode:<no value> Mutable:always ToList:always ToSet:<no value> MapTo:<no value>
// by runtemplate v3.10.0
// See https://github.com/rickb777/runtemplate/blob/master/BUILTIN.md

package testlogger

import (
	"math/rand"
	"sort"
	"sync"
)

// TestLogEventList contains a slice of type *TestLogEvent.
// It encapsulates the slice and provides methods to access or mutate it.
//
// List values follow a similar pattern to Scala Lists and LinearSeqs in particular.
// For comparison with Scala, see e.g. http://www.scala-lang.org/api/2.11.7/#scala.collection.LinearSeq
type TestLogEventList struct {
	s *sync.RWMutex
	m []*TestLogEvent
}

//-------------------------------------------------------------------------------------------------

// MakeTestLogEventList makes an empty list with both length and capacity initialised.
func MakeTestLogEventList(length, capacity int) *TestLogEventList {
	return &TestLogEventList{
		s: &sync.RWMutex{},
		m: make([]*TestLogEvent, length, capacity),
	}
}

// NewTestLogEventList constructs a new list containing the supplied values, if any.
func NewTestLogEventList(values ...*TestLogEvent) *TestLogEventList {
	list := MakeTestLogEventList(len(values), len(values))
	copy(list.m, values)
	return list
}

// ConvertTestLogEventList constructs a new list containing the supplied values, if any.
// The returned boolean will be false if any of the values could not be converted correctly.
// The returned list will contain all the values that were correctly converted.
func ConvertTestLogEventList(values ...interface{}) (*TestLogEventList, bool) {
	list := MakeTestLogEventList(0, len(values))

	for _, i := range values {
		switch j := i.(type) {
		case TestLogEvent:
			list.m = append(list.m, &j)
		case *TestLogEvent:
			list.m = append(list.m, j)
		}
	}

	return list, len(list.m) == len(values)
}

// BuildTestLogEventListFromChan constructs a new TestLogEventList from a channel that supplies
// a sequence of values until it is closed. The function doesn't return until then.
func BuildTestLogEventListFromChan(source <-chan *TestLogEvent) *TestLogEventList {
	list := MakeTestLogEventList(0, 0)
	for v := range source {
		list.m = append(list.m, v)
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// IsSequence returns true for lists and queues.
func (list *TestLogEventList) IsSequence() bool {
	return true
}

// IsSet returns false for lists or queues.
func (list *TestLogEventList) IsSet() bool {
	return false
}

// slice returns the internal elements of the current list. This is a seam for testing etc.
func (list *TestLogEventList) slice() []*TestLogEvent {
	if list == nil {
		return nil
	}
	return list.m
}

// ToList returns the elements of the list as a list, which is an identity operation in this case.
func (list *TestLogEventList) ToList() *TestLogEventList {
	return list
}

// ToSlice returns the elements of the current list as a slice.
func (list *TestLogEventList) ToSlice() []*TestLogEvent {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]*TestLogEvent, len(list.m), len(list.m))
	copy(s, list.m)
	return s
}

// ToInterfaceSlice returns the elements of the current list as a slice of arbitrary type.
func (list *TestLogEventList) ToInterfaceSlice() []interface{} {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	s := make([]interface{}, 0, len(list.m))
	for _, v := range list.m {
		s = append(s, v)
	}
	return s
}

// Clone returns a shallow copy of the list. It does not clone the underlying elements.
func (list *TestLogEventList) Clone() *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return NewTestLogEventList(list.m...)
}

//-------------------------------------------------------------------------------------------------

// Get gets the specified element in the list.
// Panics if the index is out of range or the list is nil.
func (list *TestLogEventList) Get(i int) *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[i]
}

// Head gets the first element in the list. Head plus Tail include the whole list. Head is the opposite of Last.
// Panics if list is empty or nil.
func (list *TestLogEventList) Head() *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	return list.m[0]
}

// HeadOption gets the first element in the list, if possible.
// Otherwise returns nil.
func (list *TestLogEventList) HeadOption() (*TestLogEvent, bool) {
	if list == nil {
		return nil, false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return nil, false
	}
	return list.m[0], true
}

// LastOption gets the last element in the list, if possible.
// Otherwise returns nil.
func (list *TestLogEventList) LastOption() (*TestLogEvent, bool) {
	if list == nil {
		return nil, false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return nil, false
	}
	return list.m[len(list.m)-1], true
}

// Tail gets everything except the head. Head plus Tail include the whole list. Tail is the opposite of Init.
// Panics if list is empty or nil.
func (list *TestLogEventList) Tail() *TestLogEventList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeTestLogEventList(0, 0)
	result.m = list.m[1:]
	return result
}

// Init gets everything except the last. Init plus Last include the whole list. Init is the opposite of Tail.
// Panics if list is empty or nil.
func (list *TestLogEventList) Init() *TestLogEventList {
	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeTestLogEventList(0, 0)
	result.m = list.m[:len(list.m)-1]
	return result
}

// IsEmpty tests whether TestLogEventList is empty.
func (list *TestLogEventList) IsEmpty() bool {
	return list.Size() == 0
}

// NonEmpty tests whether TestLogEventList is empty.
func (list *TestLogEventList) NonEmpty() bool {
	return list.Size() > 0
}

//-------------------------------------------------------------------------------------------------

// Size returns the number of items in the list - an alias of Len().
func (list *TestLogEventList) Size() int {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	return len(list.m)
}

// Len returns the number of items in the list - an alias of Size().
// This is one of the three methods in the standard sort.Interface.
func (list *TestLogEventList) Len() int {
	return list.Size()
}

// Swap exchanges two elements, which is necessary during sorting etc.
// This is one of the three methods in the standard sort.Interface.
func (list *TestLogEventList) Swap(i, j int) {
	list.s.Lock()
	defer list.s.Unlock()

	list.m[i], list.m[j] = list.m[j], list.m[i]
}

//-------------------------------------------------------------------------------------------------

// Exists verifies that one or more elements of TestLogEventList return true for the predicate p.
func (list *TestLogEventList) Exists(p func(*TestLogEvent) bool) bool {
	if list == nil {
		return false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return true
		}
	}
	return false
}

// Forall verifies that all elements of TestLogEventList return true for the predicate p.
func (list *TestLogEventList) Forall(p func(*TestLogEvent) bool) bool {
	if list == nil {
		return true
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if !p(v) {
			return false
		}
	}
	return true
}

// Foreach iterates over TestLogEventList and executes function f against each element.
// The function can safely alter the values via side-effects.
func (list *TestLogEventList) Foreach(f func(*TestLogEvent)) {
	if list == nil {
		return
	}

	list.s.Lock()
	defer list.s.Unlock()

	for _, v := range list.m {
		f(v)
	}
}

// Send returns a channel that will send all the elements in order.
// A goroutine is created to send the elements; this only terminates when all the elements
// have been consumed. The channel will be closed when all the elements have been sent.
func (list *TestLogEventList) Send() <-chan *TestLogEvent {
	ch := make(chan *TestLogEvent)
	go func() {
		if list != nil {
			list.s.RLock()
			defer list.s.RUnlock()

			for _, v := range list.m {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

//-------------------------------------------------------------------------------------------------

// Reverse returns a copy of TestLogEventList with all elements in the reverse order.
//
// The original list is not modified.
func (list *TestLogEventList) Reverse() *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	n := len(list.m)
	result := MakeTestLogEventList(n, n)
	last := n - 1
	for i, v := range list.m {
		result.m[last-i] = v
	}
	return result
}

// DoReverse alters a TestLogEventList with all elements in the reverse order.
// Unlike Reverse, it does not allocate new memory.
//
// The list is modified and the modified list is returned.
func (list *TestLogEventList) DoReverse() *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	mid := (len(list.m) + 1) / 2
	last := len(list.m) - 1
	for i := 0; i < mid; i++ {
		r := last - i
		if i != r {
			list.m[i], list.m[r] = list.m[r], list.m[i]
		}
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Shuffle returns a shuffled copy of TestLogEventList, using a version of the Fisher-Yates shuffle.
//
// The original list is not modified.
func (list *TestLogEventList) Shuffle() *TestLogEventList {
	if list == nil {
		return nil
	}

	return list.Clone().doShuffle()
}

// DoShuffle returns a shuffled TestLogEventList, using a version of the Fisher-Yates shuffle.
//
// The list is modified and the modified list is returned.
func (list *TestLogEventList) DoShuffle() *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doShuffle()
}

func (list *TestLogEventList) doShuffle() *TestLogEventList {
	n := len(list.m)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		list.m[i], list.m[r] = list.m[r], list.m[i]
	}
	return list
}

//-------------------------------------------------------------------------------------------------

// Clear the entire collection.
func (list *TestLogEventList) Clear() {
	if list != nil {
		list.s.Lock()
		defer list.s.Unlock()
		list.m = list.m[:0]
	}
}

// Add adds items to the current list. This is a synonym for Append.
func (list *TestLogEventList) Add(more ...*TestLogEvent) {
	list.Append(more...)
}

// Append adds items to the current list.
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
func (list *TestLogEventList) Append(more ...*TestLogEvent) *TestLogEventList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeTestLogEventList(0, len(more))
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doAppend(more...)
}

func (list *TestLogEventList) doAppend(more ...*TestLogEvent) *TestLogEventList {
	list.m = append(list.m, more...)
	return list
}

// DoInsertAt modifies a TestLogEventList by inserting elements at a given index.
// This is a generalised version of Append.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if the index is out of range.
func (list *TestLogEventList) DoInsertAt(index int, more ...*TestLogEvent) *TestLogEventList {
	if list == nil {
		if len(more) == 0 {
			return nil
		}
		list = MakeTestLogEventList(0, len(more))
		return list.doInsertAt(index, more...)
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doInsertAt(index, more...)
}

func (list *TestLogEventList) doInsertAt(index int, more ...*TestLogEvent) *TestLogEventList {
	if len(more) == 0 {
		return list
	}

	if index == len(list.m) {
		// appending is an easy special case
		return list.doAppend(more...)
	}

	newlist := make([]*TestLogEvent, 0, len(list.m)+len(more))

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	newlist = append(newlist, more...)

	newlist = append(newlist, list.m[index:]...)

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoDeleteFirst modifies a TestLogEventList by deleting n elements from the start of
// the list.
//
// If the list is nil, a new list is allocated and returned. Otherwise the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *TestLogEventList) DoDeleteFirst(n int) *TestLogEventList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(0, n)
}

// DoDeleteLast modifies a TestLogEventList by deleting n elements from the end of
// the list.
//
// The list is modified and the modified list is returned.
// Panics if n is large enough to take the index out of range.
func (list *TestLogEventList) DoDeleteLast(n int) *TestLogEventList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(len(list.m)-n, n)
}

// DoDeleteAt modifies a TestLogEventList by deleting n elements from a given index.
//
// The list is modified and the modified list is returned.
// Panics if the index is out of range or n is large enough to take the index out of range.
func (list *TestLogEventList) DoDeleteAt(index, n int) *TestLogEventList {
	list.s.Lock()
	defer list.s.Unlock()
	return list.doDeleteAt(index, n)
}

func (list *TestLogEventList) doDeleteAt(index, n int) *TestLogEventList {
	if n == 0 {
		return list
	}

	newlist := make([]*TestLogEvent, 0, len(list.m)-n)

	if index != 0 {
		newlist = append(newlist, list.m[:index]...)
	}

	index += n

	if index != len(list.m) {
		newlist = append(newlist, list.m[index:]...)
	}

	list.m = newlist
	return list
}

//-------------------------------------------------------------------------------------------------

// DoKeepWhere modifies a TestLogEventList by retaining only those elements that match
// the predicate p. This is very similar to Filter but alters the list in place.
//
// The list is modified and the modified list is returned.
func (list *TestLogEventList) DoKeepWhere(p func(*TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()
	return list.doKeepWhere(p)
}

func (list *TestLogEventList) doKeepWhere(p func(*TestLogEvent) bool) *TestLogEventList {
	result := make([]*TestLogEvent, 0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result = append(result, v)
		}
	}

	list.m = result
	return list
}

//-------------------------------------------------------------------------------------------------

// Take returns a slice of TestLogEventList containing the leading n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
func (list *TestLogEventList) Take(n int) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return list
	}

	result := MakeTestLogEventList(0, 0)
	result.m = list.m[0:n]
	return result
}

// Drop returns a slice of TestLogEventList without the leading n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *TestLogEventList) Drop(n int) *TestLogEventList {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	if n >= len(list.m) {
		return nil
	}

	result := MakeTestLogEventList(0, 0)
	result.m = list.m[n:]
	return result
}

// TakeLast returns a slice of TestLogEventList containing the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, the whole original list is returned.
//
// The original list is not modified.
func (list *TestLogEventList) TakeLast(n int) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return list
	}

	result := MakeTestLogEventList(0, 0)
	result.m = list.m[l-n:]
	return result
}

// DropLast returns a slice of TestLogEventList without the trailing n elements of the source list.
// If n is greater than or equal to the size of the list, an empty list is returned.
//
// The original list is not modified.
func (list *TestLogEventList) DropLast(n int) *TestLogEventList {
	if list == nil || n == 0 {
		return list
	}

	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if n >= l {
		return nil
	}

	result := MakeTestLogEventList(0, 0)
	result.m = list.m[:l-n]
	return result
}

// TakeWhile returns a new TestLogEventList containing the leading elements of the source list. Whilst the
// predicate p returns true, elements are added to the result. Once predicate p returns false, all remaining
// elements are excluded.
//
// The original list is not modified.
func (list *TestLogEventList) TakeWhile(p func(*TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeTestLogEventList(0, 0)
	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		} else {
			return result
		}
	}
	return result
}

// DropWhile returns a new TestLogEventList containing the trailing elements of the source list. Whilst the
// predicate p returns true, elements are excluded from the result. Once predicate p returns false, all remaining
// elements are added.
//
// The original list is not modified.
func (list *TestLogEventList) DropWhile(p func(*TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeTestLogEventList(0, 0)
	adding := false

	for _, v := range list.m {
		if adding || !p(v) {
			adding = true
			result.m = append(result.m, v)
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

// Find returns the first TestLogEvent that returns true for predicate p.
// False is returned if none match.
func (list *TestLogEventList) Find(p func(*TestLogEvent) bool) (*TestLogEvent, bool) {
	if list == nil {
		return nil, false
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			return v, true
		}
	}

	return nil, false
}

// Filter returns a new TestLogEventList whose elements return true for predicate p.
//
// The original list is not modified. See also DoKeepWhere (which does modify the original list).
func (list *TestLogEventList) Filter(p func(*TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeTestLogEventList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			result.m = append(result.m, v)
		}
	}

	return result
}

// Partition returns two new TestLogEventLists whose elements return true or false for the predicate, p.
// The first result consists of all elements that satisfy the predicate and the second result consists of
// all elements that don't. The relative order of the elements in the results is the same as in the
// original list.
//
// The original list is not modified
func (list *TestLogEventList) Partition(p func(*TestLogEvent) bool) (*TestLogEventList, *TestLogEventList) {
	if list == nil {
		return nil, nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	matching := MakeTestLogEventList(0, len(list.m))
	others := MakeTestLogEventList(0, len(list.m))

	for _, v := range list.m {
		if p(v) {
			matching.m = append(matching.m, v)
		} else {
			others.m = append(others.m, v)
		}
	}

	return matching, others
}

// Map returns a new TestLogEventList by transforming every element with function f.
// The resulting list is the same size as the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *TestLogEventList) Map(f func(*TestLogEvent) *TestLogEvent) *TestLogEventList {
	if list == nil {
		return nil
	}

	result := MakeTestLogEventList(len(list.m), len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		result.m[i] = f(v)
	}

	return result
}

// FlatMap returns a new TestLogEventList by transforming every element with function f that
// returns zero or more items in a slice. The resulting list may have a different size to the original list.
// The original list is not modified.
//
// This is a domain-to-range mapping function. For bespoke transformations to other types, copy and modify
// this method appropriately.
func (list *TestLogEventList) FlatMap(f func(*TestLogEvent) []*TestLogEvent) *TestLogEventList {
	if list == nil {
		return nil
	}

	result := MakeTestLogEventList(0, len(list.m))
	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		result.m = append(result.m, f(v)...)
	}

	return result
}

// CountBy gives the number elements of TestLogEventList that return true for the predicate p.
func (list *TestLogEventList) CountBy(p func(*TestLogEvent) bool) (result int) {
	if list == nil {
		return 0
	}

	list.s.RLock()
	defer list.s.RUnlock()

	for _, v := range list.m {
		if p(v) {
			result++
		}
	}
	return
}

// Fold aggregates all the values in the list using a supplied function, starting from some initial value.
func (list *TestLogEventList) Fold(initial *TestLogEvent, fn func(*TestLogEvent, *TestLogEvent) *TestLogEvent) *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	m := initial
	for _, v := range list.m {
		m = fn(m, v)
	}

	return m
}

// MinBy returns an element of TestLogEventList containing the minimum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such
// element is returned. Panics if there are no elements.
func (list *TestLogEventList) MinBy(less func(*TestLogEvent, *TestLogEvent) bool) *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the minimum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[i], list.m[m]) {
			m = i
		}
	}

	return list.m[m]
}

// MaxBy returns an element of TestLogEventList containing the maximum value, when compared to other elements
// using a passed func defining ‘less’. In the case of multiple items being equally maximal, the first such
// element is returned. Panics if there are no elements.
func (list *TestLogEventList) MaxBy(less func(*TestLogEvent, *TestLogEvent) bool) *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	l := len(list.m)
	if l == 0 {
		panic("Cannot determine the maximum of an empty list.")
	}

	m := 0
	for i := 1; i < l; i++ {
		if less(list.m[m], list.m[i]) {
			m = i
		}
	}

	return list.m[m]
}

// DistinctBy returns a new TestLogEventList whose elements are unique, where equality is defined by the equal function.
func (list *TestLogEventList) DistinctBy(equal func(*TestLogEvent, *TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.RLock()
	defer list.s.RUnlock()

	result := MakeTestLogEventList(0, len(list.m))
Outer:
	for _, v := range list.m {
		for _, r := range result.m {
			if equal(v, r) {
				continue Outer
			}
		}
		result.m = append(result.m, v)
	}
	return result
}

// IndexWhere finds the index of the first element satisfying predicate p. If none exists, -1 is returned.
func (list *TestLogEventList) IndexWhere(p func(*TestLogEvent) bool) int {
	return list.IndexWhere2(p, 0)
}

// IndexWhere2 finds the index of the first element satisfying predicate p at or after some start index.
// If none exists, -1 is returned.
func (list *TestLogEventList) IndexWhere2(p func(*TestLogEvent) bool, from int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	for i, v := range list.m {
		if i >= from && p(v) {
			return i
		}
	}
	return -1
}

// LastIndexWhere finds the index of the last element satisfying predicate p.
// If none exists, -1 is returned.
func (list *TestLogEventList) LastIndexWhere(p func(*TestLogEvent) bool) int {
	return list.LastIndexWhere2(p, -1)
}

// LastIndexWhere2 finds the index of the last element satisfying predicate p at or before some start index.
// If none exists, -1 is returned.
func (list *TestLogEventList) LastIndexWhere2(p func(*TestLogEvent) bool, before int) int {
	list.s.RLock()
	defer list.s.RUnlock()

	if before < 0 {
		before = len(list.m)
	}
	for i := len(list.m) - 1; i >= 0; i-- {
		v := list.m[i]
		if i <= before && p(v) {
			return i
		}
	}
	return -1
}

//-------------------------------------------------------------------------------------------------

type sortableTestLogEventList struct {
	less func(i, j *TestLogEvent) bool
	m    []*TestLogEvent
}

func (sl sortableTestLogEventList) Less(i, j int) bool {
	return sl.less(sl.m[i], sl.m[j])
}

func (sl sortableTestLogEventList) Len() int {
	return len(sl.m)
}

func (sl sortableTestLogEventList) Swap(i, j int) {
	sl.m[i], sl.m[j] = sl.m[j], sl.m[i]
}

// SortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
func (list *TestLogEventList) SortBy(less func(i, j *TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Sort(sortableTestLogEventList{less, list.m})
	return list
}

// StableSortBy alters the list so that the elements are sorted by a specified ordering.
// Sorting happens in-place; the modified list is returned.
// The algorithm keeps the original order of equal elements.
func (list *TestLogEventList) StableSortBy(less func(i, j *TestLogEvent) bool) *TestLogEventList {
	if list == nil {
		return nil
	}

	list.s.Lock()
	defer list.s.Unlock()

	sort.Stable(sortableTestLogEventList{less, list.m})
	return list
}

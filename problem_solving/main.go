package main

import (
	"container/list"
	"fmt"
	"sync"
)

// Track represents a music track
type Track struct {
	ID     int
	Title  string
	Artist string
}

// LRUCache represents the LRU cache for music streaming
type LRUCache struct {
	capacity int
	tracks   map[int]*list.Element
	list     *list.List
	mu       sync.Mutex
}

// NewLRUCache creates a new LRUCache with the specified capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		tracks:   make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves a track from the cache, updating its position
func (c *LRUCache) Get(trackID int) *Track {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.tracks[trackID]; exists {
		// Move the track to the front (most recently used)
		c.list.MoveToFront(elem)
		return elem.Value.(*Track)
	}

	return nil
}

// Add adds a track to the cache, evicting the LRU track if necessary
func (c *LRUCache) Add(track *Track) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.tracks[track.ID]; exists {
		// Track already in cache, move it to the front
		c.list.MoveToFront(elem)
	} else {
		// Check if the cache is full
		if len(c.tracks) >= c.capacity {
			// Evict the least recently used track
			lruElem := c.list.Back()
			if lruElem != nil {
				lruTrack := c.list.Remove(lruElem).(*Track)
				delete(c.tracks, lruTrack.ID)
			}
		}

		// Add the new track to the front
		elem := c.list.PushFront(track)
		c.tracks[track.ID] = elem
	}
}

// Print prints the contents of the cache
func (c *LRUCache) Print() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for elem := c.list.Front(); elem != nil; elem = elem.Next() {
		track := elem.Value.(*Track)
		fmt.Printf("Track ID: %d, Title: %s, Artist: %s\n", track.ID, track.Title, track.Artist)
	}
}

func main() {
	// Example usage
	cache := NewLRUCache(3)

	// Adding tracks to the cache
	cache.Add(&Track{ID: 1, Title: "Song 1", Artist: "Artist 1"})
	cache.Add(&Track{ID: 2, Title: "Song 2", Artist: "Artist 2"})
	cache.Add(&Track{ID: 3, Title: "Song 3", Artist: "Artist 3"})

	// Retrieving a track
	track := cache.Get(2)
	if track != nil {
		fmt.Printf("Retrieved Track ID: %d, Title: %s, Artist: %s\n", track.ID, track.Title, track.Artist)
	}

	// Adding more tracks to trigger eviction
	cache.Add(&Track{ID: 4, Title: "Song 4", Artist: "Artist 4"})
	cache.Add(&Track{ID: 5, Title: "Song 5", Artist: "Artist 5"})

	track = cache.Get(2)
	if track != nil {
		fmt.Printf("Retrieved Track ID: %d, Title: %s, Artist: %s\n", track.ID, track.Title, track.Artist)
	}
	// Printing the contents of the cache
	cache.Print()
}

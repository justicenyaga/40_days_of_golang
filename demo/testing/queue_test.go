package queue

import "testing"

func TestAddQueue(t *testing.T) {
	queue := New(3)

	for i := 0; i < 3; i++ {
		if len(queue.items) != i {
			t.Errorf("Incorrect queue element count: %v, want: %v", len(queue.items), i)
		}

		if !queue.Append(i) {
			t.Errorf("failed to append item %v to the queue", i)
		}
	}

	// Queue already full at this point
	if queue.Append(4) {
		t.Errorf("Should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	queue := New(3)
	for i := 0; i < 3; i++ {
		queue.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := queue.Next()
		if !ok {
			t.Errorf("Should be able to get item from the queue")
		}
		if item != i {
			t.Errorf("Got item in wrong order: %v, want: %v", item, i)
		}
	}

	// Queue already empty at this point
	item, ok := queue.Next()
	if ok {
		t.Errorf("Should not be able to retrieve an item from an empty queue, got: %v", item)
	}
}

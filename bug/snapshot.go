package bug

import (
	"fmt"
	"time"

	"github.com/MichaelMure/git-bug/entity"
	"github.com/MichaelMure/git-bug/identity"
)

// Snapshot is a compiled form of the Bug data structure used for storage and merge
type Snapshot struct {
	id entity.Id

	status       Status
	title        string
	comments     []Comment
	labels       []Label
	author       identity.Interface
	actors       []identity.Interface
	participants []identity.Interface
	CreateTime   time.Time

	timeline []TimelineItem

	Operations []Operation
}

// Return the Bug identifier
func (snap *Snapshot) Id() entity.Id {
	if snap.id == "" {
		// simply panic as it would be a coding error (no id provided at construction)
		panic("no id")
	}
	return snap.id
}

// Return the bugs status
func (snap *Snapshot) Status() Status {
	return snap.status
}

// Return the bugs title
func (snap *Snapshot) Title() string {
	return snap.title
}

// Return the bugs attached comments
func (snap *Snapshot) Comments() []Comment {
	return snap.comments
}

// Return the bugs assigned labels
func (snap *Snapshot) Labels() []Label {
	return snap.labels
}

// Return the bugs author
func (snap *Snapshot) Author() identity.Interface {
	return snap.author
}

// Return the bugs actors
func (snap *Snapshot) Actors() []identity.Interface {
	return snap.actors
}

// Return the bugs participants
func (snap *Snapshot) Participants() []identity.Interface {
	return snap.participants
}

// Return the bugs timeline
func (snap *Snapshot) Timeline() []TimelineItem {
	return snap.timeline
}

// Return the last time a bug was modified
func (snap *Snapshot) EditTime() time.Time {
	if len(snap.Operations) == 0 {
		return time.Unix(0, 0)
	}

	return snap.Operations[len(snap.Operations)-1].Time()
}

// GetCreateMetadata return the creation metadata
func (snap *Snapshot) GetCreateMetadata(key string) (string, bool) {
	return snap.Operations[0].GetMetadata(key)
}

// SearchTimelineItem will search in the timeline for an item matching the given hash
func (snap *Snapshot) SearchTimelineItem(id entity.Id) (TimelineItem, error) {
	for i := range snap.timeline {
		if snap.timeline[i].Id() == id {
			return snap.timeline[i], nil
		}
	}

	return nil, fmt.Errorf("timeline item not found")
}

// SearchComment will search for a comment matching the given hash
func (snap *Snapshot) SearchComment(id entity.Id) (*Comment, error) {
	for _, c := range snap.comments {
		if c.id == id {
			return &c, nil
		}
	}

	return nil, fmt.Errorf("comment item not found")
}

// UpdateComment will replace the comment with the given id by another comment
func (snap *Snapshot) UpdateComment(id entity.Id, comment Comment) error {
	for i := range snap.comments {
		if snap.comments[i].id == id {
			snap.comments[i] = comment
			return nil
		}
	}

	return fmt.Errorf("comment item not found")
}

// Change current status to the new status
func (snap *Snapshot) setStatusTo(newStatus Status) {
	snap.status = newStatus
}

// Change current title to the new title
func (snap *Snapshot) changeTitleTo(newTitle string) {
	snap.title = newTitle
}

// Append the supplied comment to the snapshots comments
func (snap *Snapshot) appendComment(comment Comment) {
	snap.comments = append(snap.comments, comment)
}

// Add the supplied label to the snapshots labels
func (snap *Snapshot) addLabel(label Label) {
	if !snap.HasLabel(label) {
		snap.labels = append(snap.labels, label)
	}
}

// Remove the supplied label from the snapshots labels
func (snap *Snapshot) removeLabel(removed Label) {
	for i, label := range snap.labels {
		if label == removed {
			snap.labels[i] = snap.labels[len(snap.labels)-1]
			snap.labels = snap.labels[:len(snap.labels)-1]
		}
	}
}

// append the operation author to the actors list
func (snap *Snapshot) addActor(actor identity.Interface) {
	for _, a := range snap.actors {
		if actor.Id() == a.Id() {
			return
		}
	}

	snap.actors = append(snap.actors, actor)
}

// append the operation author to the participants list
func (snap *Snapshot) addParticipant(participant identity.Interface) {
	for _, p := range snap.participants {
		if participant.Id() == p.Id() {
			return
		}
	}

	snap.participants = append(snap.participants, participant)
}

// Append the supplied timelineitem to the snapshots timeline
func (snap *Snapshot) appendTimelineItem(item TimelineItem) {
	snap.timeline = append(snap.timeline, item)
}

// HasParticipant return true if the id is a participant
func (snap *Snapshot) HasParticipant(id entity.Id) bool {
	for _, p := range snap.participants {
		if p.Id() == id {
			return true
		}
	}
	return false
}

// HasAnyParticipant return true if one of the ids is a participant
func (snap *Snapshot) HasAnyParticipant(ids ...entity.Id) bool {
	for _, id := range ids {
		if snap.HasParticipant(id) {
			return true
		}
	}
	return false
}

// HasActor return true if the id is a actor
func (snap *Snapshot) HasActor(id entity.Id) bool {
	for _, p := range snap.actors {
		if p.Id() == id {
			return true
		}
	}
	return false
}

// HasAnyActor return true if one of the ids is a actor
func (snap *Snapshot) HasAnyActor(ids ...entity.Id) bool {
	for _, id := range ids {
		if snap.HasActor(id) {
			return true
		}
	}
	return false
}

func (snap *Snapshot) HasLabel(label Label) bool {
	for _, l := range snap.labels {
		if l == label {
			return true
		}
	}
	return false
}

// Sign post method for gqlgen
func (snap *Snapshot) IsAuthored() {}

package github

import "github.com/shurcooL/githubv4"

type pageInfo struct {
	EndCursor       githubv4.String
	HasNextPage     bool
	StartCursor     githubv4.String
	HasPreviousPage bool
}

type actor struct {
	Typename  githubv4.String `graphql:"__typename"`
	Login     githubv4.String
	AvatarUrl githubv4.String
	User      struct {
		Name  *githubv4.String
		Email githubv4.String
	} `graphql:"... on User"`
	Organization struct {
		Name  *githubv4.String
		Email *githubv4.String
	} `graphql:"... on Organization"`
}

type actorEvent struct {
	Id        githubv4.ID
	CreatedAt githubv4.DateTime
	Actor     *actor
}

type authorEvent struct {
	Id        githubv4.ID
	CreatedAt githubv4.DateTime
	Author    *actor
}

type userContentEdit struct {
	Id        githubv4.ID
	CreatedAt githubv4.DateTime
	UpdatedAt githubv4.DateTime
	EditedAt  githubv4.DateTime
	Editor    *actor
	DeletedAt *githubv4.DateTime
	DeletedBy *actor
	Diff      *githubv4.String
}

type issueComment struct {
	authorEvent // NOTE: contains Id
	Body        githubv4.String
	Url         githubv4.URI
}

type timelineItem struct {
	Typename githubv4.String `graphql:"__typename"`

	// issue
	IssueComment issueComment `graphql:"... on IssueComment"`

	// Label
	LabeledEvent struct {
		actorEvent
		Label struct {
			// Color githubv4.String
			Name githubv4.String
		}
	} `graphql:"... on LabeledEvent"`
	UnlabeledEvent struct {
		actorEvent
		Label struct {
			// Color githubv4.String
			Name githubv4.String
		}
	} `graphql:"... on UnlabeledEvent"`

	// Status
	ClosedEvent struct {
		actorEvent
		// Url githubv4.URI
	} `graphql:"... on  ClosedEvent"`
	ReopenedEvent struct {
		actorEvent
	} `graphql:"... on  ReopenedEvent"`

	// Title
	RenamedTitleEvent struct {
		actorEvent
		CurrentTitle  githubv4.String
		PreviousTitle githubv4.String
	} `graphql:"... on RenamedTitleEvent"`
}

type ghostQuery struct {
	User struct {
		Login     githubv4.String
		AvatarUrl githubv4.String
		Name      *githubv4.String
	} `graphql:"user(login: $login)"`
}

type labelsQuery struct {
	Repository struct {
		Labels struct {
			Nodes []struct {
				ID          string `graphql:"id"`
				Name        string `graphql:"name"`
				Color       string `graphql:"color"`
				Description string `graphql:"description"`
			}
			PageInfo pageInfo
		} `graphql:"labels(first: $first, after: $after)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type loginQuery struct {
	Viewer struct {
		Login string `graphql:"login"`
	} `graphql:"viewer"`
}

type issueQuery struct {
	Repository struct {
		Issues struct {
			Nodes    []issue
			PageInfo pageInfo
		} `graphql:"issues(first: $issueFirst, after: $issueAfter, orderBy: {field: CREATED_AT, direction: ASC}, filterBy: {since: $issueSince})"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type issue struct {
	authorEvent
	Title  string
	Number githubv4.Int
	Body   githubv4.String
	Url    githubv4.URI
}

type issueEditQuery struct {
	Node struct {
		Typename githubv4.String `graphql:"__typename"`
		Issue    struct {
			UserContentEdits struct {
				Nodes      []userContentEdit
				TotalCount githubv4.Int
				PageInfo   pageInfo
			} `graphql:"userContentEdits(last: $issueEditLast, before: $issueEditBefore)"`
		} `graphql:"... on Issue"`
	} `graphql:"node(id: $gqlNodeId)"`
}

type timelineQuery struct {
	Node struct {
		Typename githubv4.String `graphql:"__typename"`
		Issue    struct {
			TimelineItems struct {
				Nodes    []timelineItem
				PageInfo pageInfo
			} `graphql:"timelineItems(first: $timelineFirst, after: $timelineAfter)"`
		} `graphql:"... on Issue"`
	} `graphql:"node(id: $gqlNodeId)"`
}

type commentEditQuery struct {
	Node struct {
		Typename     githubv4.String `graphql:"__typename"`
		IssueComment struct {
			UserContentEdits struct {
				Nodes    []userContentEdit
				PageInfo pageInfo
			} `graphql:"userContentEdits(last: $commentEditLast, before: $commentEditBefore)"`
		} `graphql:"... on IssueComment"`
	} `graphql:"node(id: $gqlNodeId)"`
}

package models

type githubEventType int

const (
	CommitCommentEvent            githubEventType = iota
	CreateEvent
	DeleteEvent
	DeploymentEvent
	DeploymentStatusEvent
	DownloadEvent
	FollowEvent
	ForkEvent
	ForkApplyEvent
	GistEvent
	GollumEvent
	IssueCommentEvent
	IssuesEvent
	LabelEvent
	MemberEvent
	MembershipEvent
	MilestoneEvent
	OrganizationEvent
	OrgBlockEvent
	PageBuildEvent
	ProjectCardEvent
	ProjectColumnEvent
	ProjectEvent
	PublicEvent
	PullRequestEvent
	PullRequestReviewEvent
	PullRequestReviewCommentEvent
	PushEvent
	ReleaseEvent
	RepositoryEvent
	StatusEvent
	TeamEvent
	TeamAddEvent
	WatchEvent
)

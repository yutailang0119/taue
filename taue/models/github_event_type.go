package models

type githubEventType int

const (
	//commitComment githubEventType = iota
	//create
	//delete
	//deployment
	//deploymentStatus
	//download
	//follow
	//fork
	//forkApply
	//gist
	//gollum
	//issueComment
	//issues
	//label
	//member
	//membership
	//milestone
	//organization
	//orgBlock
	//pageBuild
	//projectCard
	//projectColumn
	//project
	//public
	//pullRequest
	//pullRequestReview
	//pullRequestReviewComment
	//push
	//release
	//repository
	//status
	//team
	//teamAdd
	//match

	CommitCommentEvent githubEventType = iota
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

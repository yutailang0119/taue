package models

import (
	"time"
	"errors"
)

// GitHubEvent is https://api.github.com/users/:username/event
type GitHubEvent struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Actor     githubActor `json:"actor"`
	Repo      githubRepo  `json:"repo"`
	CreatedAt string      `json:"created_at"`
}

type githubActor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type githubRepo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (githubEvent GitHubEvent) createdAt() time.Time {
	const githubTimeLayout = "2006-01-02T15:04:05Z"
	t, _ := time.Parse(githubTimeLayout, githubEvent.CreatedAt)
	return t
}

func (githubEvent GitHubEvent) isTodayContribute() bool {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	thisCreatedAt := githubEvent.createdAt()
	return thisCreatedAt.Year() == yesterday.Year() && thisCreatedAt.Month() == yesterday.Month() && thisCreatedAt.Day() == yesterday.Day()
}

func (githubEvent GitHubEvent) isCountable() bool {
	eventType, err := githubEvent.eventType()
	if err != nil {
		return false
	}

	return eventType == CommitCommentEvent ||
		eventType == GistEvent ||
		eventType == IssuesEvent ||
		eventType == PullRequestEvent ||
		eventType == PullRequestReviewCommentEvent ||
		eventType == PullRequestReviewEvent ||
		eventType == PushEvent
}

func (githubEvent GitHubEvent) eventType() (eventType githubEventType, err error) {
	switch githubEvent.Type {
	case CommitCommentEvent.String():
		eventType = CommitCommentEvent
	case CreateEvent.String():
		eventType = CreateEvent
	case DeleteEvent.String():
		eventType = DeleteEvent
	case DeploymentEvent.String():
		eventType = DeploymentEvent
	case DeploymentStatusEvent.String():
		eventType = DeploymentStatusEvent
	case DownloadEvent.String():
		eventType = DownloadEvent
	case FollowEvent.String():
		eventType = FollowEvent
	case ForkEvent.String():
		eventType = ForkEvent
	case ForkApplyEvent.String():
		eventType = ForkApplyEvent
	case GistEvent.String():
		eventType = GistEvent
	case GollumEvent.String():
		eventType = GollumEvent
	case IssueCommentEvent.String():
		eventType = IssueCommentEvent
	case IssuesEvent.String():
		eventType = IssuesEvent
	case LabelEvent.String():
		eventType = LabelEvent
	case MemberEvent.String():
		eventType = MemberEvent
	case MembershipEvent.String():
		eventType = MembershipEvent
	case MilestoneEvent.String():
		eventType = MilestoneEvent
	case OrganizationEvent.String():
		eventType = OrganizationEvent
	case OrgBlockEvent.String():
		eventType = OrgBlockEvent
	case PageBuildEvent.String():
		eventType = PageBuildEvent
	case ProjectCardEvent.String():
		eventType = ProjectCardEvent
	case ProjectColumnEvent.String():
		eventType = ProjectColumnEvent
	case ProjectEvent.String():
		eventType = ProjectEvent
	case PublicEvent.String():
		eventType = PublicEvent
	case PullRequestEvent.String():
		eventType = PullRequestEvent
	case PullRequestReviewEvent.String():
		eventType = PullRequestReviewEvent
	case PullRequestReviewCommentEvent.String():
		eventType = PullRequestReviewCommentEvent
	case PushEvent.String():
		eventType = PushEvent
	case ReleaseEvent.String():
		eventType = ReleaseEvent
	case RepositoryEvent.String():
		eventType = RepositoryEvent
	case StatusEvent.String():
		eventType = StatusEvent
	case TeamEvent.String():
		eventType = TeamEvent
	case TeamAddEvent.String():
		eventType = TeamAddEvent
	case WatchEvent.String():
		eventType = WatchEvent
	default:
		err = errors.New("Unkown GitHub Event")
	}

	return eventType, err
}

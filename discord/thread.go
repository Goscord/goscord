package discord

import "time"

type ThreadMetadata struct {
	Archived            bool      `json:"archived"`
	AutoArchiveDuration int       `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time `json:"archive_timestamp"`
	Locked              bool      `json:"locked"`
	Invitable           bool      `json:"invitable,omitempty"`
	CreateTimestamp     time.Time `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	Id            string    `json:"id,omitempty"`
	UserId        string    `json:"user_id,omitempty"`
	JoinTimestamp time.Time `json:"join_timestamp"`
	Flags         int       `json:"flags"`
}

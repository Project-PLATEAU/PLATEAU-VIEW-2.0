package asset

import (
	"strings"
)

type ArchiveExtractionStatus string

const (
	ArchiveExtractionStatusSkipped    ArchiveExtractionStatus = "skipped"
	ArchiveExtractionStatusPending    ArchiveExtractionStatus = "pending"
	ArchiveExtractionStatusInProgress ArchiveExtractionStatus = "in_progress"
	ArchiveExtractionStatusDone       ArchiveExtractionStatus = "done"
	ArchiveExtractionStatusFailed     ArchiveExtractionStatus = "failed"
)

func ArchiveExtractionStatusFrom(s string) (ArchiveExtractionStatus, bool) {
	ss := strings.ToLower(s)
	switch ArchiveExtractionStatus(ss) {
	case ArchiveExtractionStatusSkipped:
		return ArchiveExtractionStatusSkipped, true
	case ArchiveExtractionStatusPending:
		return ArchiveExtractionStatusPending, true
	case ArchiveExtractionStatusInProgress:
		return ArchiveExtractionStatusInProgress, true
	case ArchiveExtractionStatusDone:
		return ArchiveExtractionStatusDone, true
	case ArchiveExtractionStatusFailed:
		return ArchiveExtractionStatusFailed, true
	default:
		return ArchiveExtractionStatus(""), false
	}
}

func ArchiveExtractionStatusFromRef(s *string) *ArchiveExtractionStatus {
	if s == nil {
		return nil
	}

	ss, ok := ArchiveExtractionStatusFrom(*s)
	if !ok {
		return nil
	}
	return &ss
}

func (s ArchiveExtractionStatus) String() string {
	return string(s)
}

func (s *ArchiveExtractionStatus) StringRef() *string {
	if s == nil {
		return nil
	}
	s2 := string(*s)
	return &s2
}

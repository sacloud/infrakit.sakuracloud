package sacloud

import "time"

// PropCreatedAt 作成日時内包型
type propCreatedAt struct {
	CreatedAt *time.Time `json:",omitempty"` // 作成日時
}

// GetCreatedAt 作成日時 取得
func (p *propCreatedAt) GetCreatedAt() *time.Time {
	return p.CreatedAt
}

// PropModifiedAt 変更日時内包型
type PropModifiedAt struct {
	// CreatedAt 変更日時
	ModifiedAt *time.Time `json:",omitempty"`
}

// GetModifiedAt 変更日時 取得
func (p *PropModifiedAt) GetModifiedAt() *time.Time {
	return p.ModifiedAt
}

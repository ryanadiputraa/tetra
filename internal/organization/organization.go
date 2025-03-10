package organization

import (
	"context"
	"fmt"
	"time"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/user"
)

const (
	invitationMailBody = `
<!DOCTYPE html>
<html lang="id">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Undangan Inventra</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f7;
      color: #333;
      margin: 0;
      padding: 20px;
    }

    a {
      text-decoration: none;
      color: #ffffff !important;
    }

    .container {
      max-width: 600px;
      margin: 0 auto;
      background: #ffffff;
      padding: 20px;
      border-radius: 8px;
      box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    }

    .content {
      text-align: center;
    }

    .button {
      display: inline-block;
      margin-top: 20px;
      padding: 10px 20px;
      font-size: 16px;
      color: #ffffff;
      background-color: #4682AB;
      text-decoration: none;
      border-radius: 5px;
    }

    .footer {
      margin-top: 20px;
      text-align: center;
      font-size: 12px;
      color: #666666;
    }
  </style>
</head>

<body>
  <div class="container">
    <div class="content">
      <p>Halo!</p>
      <p>Kamu telah diundang untuk bergabung dengan %s di Inventra. Silakan klik tombol Terima Undangan di bawah untuk bergabung.</p>
      <a href="%s" class="button">Terima Undangan</a>
    </div>
    <div class="footer">
      <p>Copyright © %d Inventra. All Right Reserved.</p>
    </div>
  </div>
</body>

</html>
`
)

type Organization struct {
	ID                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OwnerID           int       `json:"-" gorm:"notNull"`
	Owner             user.User `json:"owner" gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name              string    `json:"name" gorm:"type:varchar(100);notNull"`
	CreatedAt         time.Time `json:"created_at" gorm:"notNull"`
	SubscriptionEndAt time.Time `json:"subscription_end_at" gorm:"notNull"`
	Members           []Member  `json:"-" `
}

type Member struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrganizationID int       `json:"organization_id" gorm:"notNull;constraint:OnDelete:CASCADE"`
	UserID         int       `json:"user_id" gorm:"notNull"`
	User           user.User `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Role           string    `json:"role" gorm:"type:varchar(10);notNull"`
	CreatedAt      time.Time `json:"created_at" gorm:"notNull"`
}

type MemberData struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type OrganizationPayload struct {
	Name string `json:"name" validate:"required"`
}

type InvitePayload struct {
	Email string `json:"email" validate:"required,email"`
}

type AcceptInvitationPayload struct {
	Code string `json:"code" validate:"required"`
}

type ChangeMemberPayload struct {
	Role string `json:"role" validate:"required"`
}

func New(Name string, userID int) Organization {
	return Organization{
		OwnerID:           userID,
		Name:              Name,
		CreatedAt:         time.Now().UTC(),
		SubscriptionEndAt: time.Now().AddDate(0, 1, 0).UTC(),
	}
}

func NewMember(organizationID, userID int, role auth.Role) Member {
	return Member{
		OrganizationID: organizationID,
		UserID:         userID,
		Role:           string(role),
		CreatedAt:      time.Now().UTC(),
	}
}

func GenrateInvitationMailBody(organizationName, domain, inviteCode string) string {
	link := domain + "/join/" + inviteCode
	year := time.Now().Year()
	return fmt.Sprintf(invitationMailBody, organizationName, link, year)
}

type OrganizationService interface {
	Create(ctx context.Context, Name string, userID int) (Organization, error)
	GetByID(ctx context.Context, organizationID int) (Organization, error)
	IsSubscriptionValid(ctx context.Context, organizationID int) (bool, error)
	Delete(ctx context.Context, organizationID, userID int) error
	ListMember(ctx context.Context, organizationID int) ([]MemberData, error)
	InviteUser(ctx context.Context, organizationID int, email string) error
	Join(ctx context.Context, organizationID, userID int) (Member, error)
	RemoveMember(ctx context.Context, organizationID, memberID int) error
	ChangeMemberRole(ctx context.Context, organizationID, memberID int, role string) error
	Leave(ctx context.Context, organizationID, memberID int) error
}

type OrganizationRepository interface {
	Save(ctx context.Context, organization Organization) (Organization, error)
	FindByID(ctx context.Context, organizationID int) (Organization, error)
	Delete(ctx context.Context, organizationID, userID int) error
	AddMember(ctx context.Context, member Member) (Member, error)
	FetchMembers(ctx context.Context, organizationID int) ([]MemberData, error)
	DeleteMember(ctx context.Context, organizationID, memberID int) error
	UpdateMemberRole(ctx context.Context, organizationID, memberID int, role string) error
}

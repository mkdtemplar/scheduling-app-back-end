package api

import (
	"fmt"
	"html"
	"net/http"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/models/dto"
	"scheduling-app-back-end/internal/repository/interfaces"
	"scheduling-app-back-end/internal/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func NewDailyAssignmentsHandler(positionsRepo interfaces.IPositionsRepository,
	dailyRepo interfaces.IDailyAssignmentsRepository) *DailyAssignmentsHandler {
	return &DailyAssignmentsHandler{
		positionsRepo: positionsRepo,
		dailyRepo:     dailyRepo,
		//from:          from,
	}
}

func (da *DailyAssignmentsHandler) SendToTeam(c *gin.Context) {
	var req dto.SendDailyAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload", "bind_error": err.Error()})
		return
	}

	req.PositionName = strings.TrimSpace(req.PositionName)
	req.Description = strings.TrimSpace(req.Description)

	if req.PositionName == "" && req.PositionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position_name or position_id is required"})
		return
	}
	if req.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description is required"})
		return
	}

	//from := da.from
	senderEmail := ""
	if v, ok := c.Get("admin_mail"); ok {
		if s, ok2 := v.(string); ok2 {
			senderEmail = strings.TrimSpace(s)
		}
	}
	positionName := req.PositionName
	if positionName == "" && req.PositionID > 0 {
		pos, err := da.positionsRepo.GetPositionByID(c.Request.Context(), req.PositionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "position not found"})
			return
		}
		positionName = strings.TrimSpace(pos.PositionName)
		if positionName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "position name is empty"})
			return
		}
	}

	users, err := da.dailyRepo.GetUsersByPositionName(c.Request.Context(), positionName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to load users"})
		return
	}

	subject := fmt.Sprintf("Daily Assignment — %s", positionName)
	content := buildDailyAssignmentHTML(positionName, req.Description)

	sent := 0
	for _, u := range users {
		to := strings.TrimSpace(u.Email)
		if to == "" {
			continue
		}

		err = utils.SendMsgErr(models.MailData{
			To:      to,
			From:    senderEmail,
			Subject: subject,
			Content: content,
		})
		if err == nil {
			sent++
		}
	}

	saveErr := da.dailyRepo.CreateAssignment(c.Request.Context(), &models.DailyAssignment{
		SenderEmail:  senderEmail,
		PositionName: positionName,
		Description:  req.Description,
		SentCount:    sent,
		CreatedAt:    time.Time{},
	})

	if saveErr != nil {
		c.JSON(http.StatusOK, gin.H{"sent_count": sent, "warning": "emails sent but failed to save assignment in database"})
		return
	}

	c.JSON(http.StatusOK, dto.SendDailyAssignmentResponse{SentCount: sent})
}

func buildDailyAssignmentHTML(positionName, description string) string {
	pos := html.EscapeString(positionName)
	desc := html.EscapeString(description)

	return fmt.Sprintf(`
<div style="font-family: Arial, Helvetica, sans-serif; background:#0b1220; padding:18px;">
  <div style="max-width:680px; margin:0 auto; background:#0f1b2e; border:1px solid rgba(255,255,255,0.12); border-radius:14px; overflow:hidden;">
    <div style="padding:14px 16px; background:linear-gradient(135deg, rgba(59,130,246,0.95), rgba(14,165,233,0.85)); color:#fff;">
      <div style="font-size:18px; font-weight:800;">Daily Assignment</div>
      <div style="opacity:0.9; margin-top:4px;">Team: <b>%s</b></div>
    </div>
    <div style="padding:16px; color:rgba(255,255,255,0.92);">
      <div style="font-size:14px; opacity:0.9; margin-bottom:10px;">Message:</div>
      <div style="white-space:pre-wrap; background:rgba(255,255,255,0.06); border:1px solid rgba(255,255,255,0.10); border-radius:12px; padding:12px; line-height:1.5;">%s</div>
      <div style="margin-top:14px; font-size:12px; opacity:0.75;">— Scheduling App</div>
    </div>
  </div>
</div>`, pos, desc)
}

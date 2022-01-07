package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"gitlab.ushareit.me/sgt/hawkeye/network-monitor/src/models"
)

func Page() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//if !ctx.Request().Header.IsGet() {
		//	return nil
		//}

		var (
			pageSize int = 10
			curPage  int = 1
		)

		if ps, err := strconv.Atoi(ctx.Query("limit")); err == nil && ps > 0 {
			pageSize = ps
		}
		if p, err := strconv.Atoi(ctx.Query("page")); err == nil && p > 0 {
			curPage = p
		}

		offset := (curPage - 1) * pageSize
		if offset < 0 {
			offset = 0
		}

		page := models.Page{
			PageSize: pageSize,
			Offset:   offset,
			Page:     curPage,
			Query:    ctx.Query("query"),
		}

		switch ctx.Query("sort") {
		case "asc":
			page.Sort = "asc"
		case "desc":
			page.Sort = "desc"
		}

		switch ctx.Query("order_by") {
		case "name":
			page.OrderBy = "name"
		case "create_time":
			page.OrderBy = "create_time"
		default:
			page.OrderBy = "update_time"
		}

		ctx.Context().SetUserValue("page", page)

		return ctx.Next()
	}
}

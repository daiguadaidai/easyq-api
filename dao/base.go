package dao

import (
	"fmt"
	"strings"
)

/* 获取like子句, 只需要字符串类型数据
 *
 * obj: 对象
 * names: 需要生成like的对象
 */
func GetLikeClausesByKeyWords(keywords string, names ...string) []string {
	// 将对象转化为 map
	likeClauses := make([]string, 0, len(names))
	for _, name := range names {
		if len(keywords) <= 0 { // 空字符串
			continue
		}

		likeClause := fmt.Sprintf("%s LIKE '%%%v%%'", name, keywords)
		likeClauses = append(likeClauses, likeClause)
	}

	return likeClauses
}

func JoinOrClauses(clauses ...string) string {
	noEmptyClauses := make([]string, 0, len(clauses))
	for _, clause := range clauses {
		if len(clause) == 0 {
			continue
		}
		noEmptyClauses = append(noEmptyClauses, clause)
	}

	if len(noEmptyClauses) == 0 {
		return ""
	} else if len(noEmptyClauses) == 1 {
		return noEmptyClauses[0]
	} else {
		return fmt.Sprintf("(%s)", strings.Join(noEmptyClauses, " OR "))
	}
}

func JoinAndClauses(clauses ...string) string {
	noEmptyClauses := make([]string, 0, len(clauses))
	for _, clause := range clauses {
		if len(clause) == 0 {
			continue
		}
		noEmptyClauses = append(noEmptyClauses, clause)
	}

	if len(noEmptyClauses) == 0 {
		return ""
	} else if len(noEmptyClauses) == 1 {
		return noEmptyClauses[0]
	} else {
		return fmt.Sprintf("(%s)", strings.Join(noEmptyClauses, " AND "))
	}
}

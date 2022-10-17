package routes

import (
	"fmt"
	"strings"
)

func generatorGetRoute(collectionName string) string {
	return fmt.Sprintf(
		strings.Join([]string{
			`app.get("/%s", (req, res) => {`,
			`    const { rows } = await db.query(`,
			`        "SELECT * FROM %s",`,
			`    )`,
			``,
			`    res.json({ rows })`,
			`})`,
			``,
		}, "\n"),
		collectionName, collectionName,
	)
}

package routes

import (
	"fmt"
	"strings"
)

func generatorDeleteRoute(collectionName string) string {
	return fmt.Sprintf(
		strings.Join([]string{
			`app.delete("/%s/:id", (req, res) => {`,
			`    const id = req.params.id`,
			`    await db.query(`,
			`        "DELETE FROM %s WHERE id = $1",`,
			`        [ id ]`,
			`    )`,
			``,
			`    res.json({ message: "success" })`,
			`})`,
			``,
		}, "\n"),
		collectionName, collectionName,
	)
}

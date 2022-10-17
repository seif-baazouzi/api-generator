package routes

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"fmt"
	"strings"
)

func generatorPostRoute(collectionName string, collection config.Collection) string {
	columns := strings.Join(utils.GetCollectionFields(collection), ", ")
	columnsPlaceHolders := strings.Join(utils.GenerateSqlPlaceholders(len(collection)), ", ")

	return fmt.Sprintf(
		strings.Join([]string{
			`app.post("/%s", (req, res) => {`,
			`    const { %s } = req.body`,
			`    const { rows } = await db.query(`,
			`        "INSERT INTO %s (%s) VALUES (%s) RETURNING id",`,
			`        [ %s ]`,
			`    )`,
			``,
			`    res.json({ message: "success", id: rows[0].id })`,
			`})`,
			``,
		}, "\n"),
		collectionName, columns, collectionName, columns, columnsPlaceHolders, columns,
	)
}

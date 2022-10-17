package routes

import (
	"api-generator/src/config"
	"api-generator/src/utils"
	"fmt"
	"strings"
)

func generatorPutRoute(collectionName string, collection config.Collection) string {
	columns := strings.Join(utils.GetCollectionFields(collection), ", ")
	columnsPlaceHolders := strings.Join(utils.GenerateSqlPlaceholders(len(collection)), ", ")
	idPlaceholder := len(collection) + 1

	return fmt.Sprintf(
		strings.Join([]string{
			`app.put("/%s/:id", (req, res) => {`,
			`    const id = req.params.id`,
			`    const { %s } = req.body`,
			`    await db.query(`,
			`        "UPDATE %s SET (%s) = %s WHERE id = $%d",`,
			`        [ %s, id ]`,
			`    )`,
			``,
			`    res.json({ message: "success" })`,
			`})`,
			``,
		}, "\n"),
		collectionName, columns, collectionName, columns, columnsPlaceHolders, idPlaceholder, columns,
	)
}

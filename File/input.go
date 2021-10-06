package File

import "pasarwarga/models"



type DetailFile struct {
	PdfFile string `uri:"id" binding:"required"`
	User    models.Users
}


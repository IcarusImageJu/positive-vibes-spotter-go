package render

import (
	"fmt"
	"os/exec"
	logger "positive-vibes-spotter/log"
	"positive-vibes-spotter/utils"
)

// CreateImageWithCaption creates an image with the specified caption and saves it to the output path.
func CreateImageWithCaption(caption string, outputPath string) {
	logger.Info("Création de l’image avec la légende")
	
	stdout, err := logger.Writer()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Erreur lors de l'ouverture du fichier de log pour stdout: %v", err))
	}
	stderr, err := logger.Writer()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Erreur lors de l'ouverture du fichier de log pour stderr: %v", err))
	}
	
	cmd := exec.Command("convert", "-background", "black", "-fill", "white", "-font", "Arial", "-pointsize", "72",
	"-gravity", "southwest", "-extent", "1280x720", "-size", "1200x600", "caption:"+caption,
	"-bordercolor", "black", "-border", "100x100", "-gravity", "southwest", "-extent", "1280x720+50+50", outputPath)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	if err != nil {
		logger.Fatal(err)
	}
}


// DisplayImageWithNoise displays the image at the specified output path and adds noise at intervals to prevent pixel burn-in.
func DisplayImageWithNoise(outputPath string) {
	logger.Info("Affichage de l’image")
	
	cmd := exec.Command("fim", "-a", "--quiet", outputPath)
	err := cmd.Start()
	if err != nil {
		logger.Fatal(err)
	}
	
}

func Render(caption string, base64image string, outputPath string) {
	utils.CheckInstall("fim", "fim")
	utils.CheckInstall("convert", "imagemagick")
	utils.CheckAndInstallFonts()
	CreateImageWithCaption(caption, outputPath)
	DisplayImageWithNoise(outputPath)
}
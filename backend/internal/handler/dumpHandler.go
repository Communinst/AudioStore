package handler

import (
	"AudioShare/backend/internal/config"
	"AudioShare/backend/internal/entity"
	"AudioShare/backend/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type DumpHandler struct {
	dump service.DumpServiceInterface
}

func NewDumpHandler(dump service.DumpServiceInterface) *DumpHandler {
	return &DumpHandler{
		dump: dump,
	}
}

// @Summary Create database dump
// @Description Creates a new PostgreSQL database dump using pg_dump
// @Tags dump
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Dump created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /dump/create [post]
func (this *DumpHandler) CreateDump(c *gin.Context) {
	dumpCfg := config.MustLoadDumpConfig()
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02_15-04-05")
	filePath := filepath.Join(dumpCfg.Dir, fmt.Sprintf("%s_%s", dumpCfg.Prefix, timeString))

	fmt.Printf("%s\n", filePath)

	cmd := exec.Command("docker", "exec", dumpCfg.ContainerName, "pg_dump", "-U", dumpCfg.Username, "-F", "c", dumpCfg.DbName)
	outputFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileInfo, err := outputFile.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fileSize := fileInfo.Size()

	err = this.dump.InsertDump(c.Request.Context(), filePath, fileSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dump created successfully"})
}

// @Summary Restore database dump
// @Description Restores a PostgreSQL database from a dump file using pg_restore
// @Tags dump
// @Accept json
// @Produce json
// @Param request body entity.Dump true "Dump file information"
// @Success 200 {object} map[string]string "Dump restored successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /dump/restore [post]
func (this *DumpHandler) RestoreDump(c *gin.Context) {
	var fileName entity.Dump

	if err := c.ShouldBindJSON(&fileName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dumpCfg := config.MustLoadDumpConfig()

	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02_15-04-05")
	filePath := fmt.Sprintf("%s_%s", dumpCfg.RestorePrefix, timeString)

	copyCmd := exec.Command("docker", "cp", fileName.Filename, fmt.Sprintf("%s:%s", dumpCfg.ContainerName, filePath))
	if err := copyCmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	restoreCmd := exec.Command("docker", "exec", dumpCfg.ContainerName, "pg_restore", "-U", dumpCfg.Username, "--clean", "-d", dumpCfg.DbName, filePath)
	if err := restoreCmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dump restored successfully"})
}

// @Summary Get all dumps
// @Description Retrieves a list of all database dumps
// @Tags dump
// @Accept json
// @Produce json
// @Success 200 {array} entity.Dump "List of dumps"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /dump/list [get]
func (this *DumpHandler) GetAllDumps(c *gin.Context) {
	dumps, err := this.dump.GetAllDumps(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%v", dumps)
	c.JSON(http.StatusOK, dumps)
}

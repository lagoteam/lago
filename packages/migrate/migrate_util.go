package migrate

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// AddMigration 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func AddMigration(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}

// FilePut 将数据存入文件
func FilePut(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// FileExists 判断文件是否存在
func FileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetMigrationFileName(fullFilename string) string {
	//_, fullFilename, _, _ := runtime.Caller(0)
	//fmt.Println(fullFilename)
	var filename string
	filename = path.Base(fullFilename)
	//fmt.Println("filename=", filename)
	//fmt.Println("filenameExtension=", FileNameWithoutExtension(filename))
	return FileNameWithoutExtension(filename)
}

// SuccessConsole 打印一条成功消息，绿色输出
func SuccessConsole(msg string) {
	colorOutConsole(msg, "green")
}

// ErrorConsole 打印一条报错消息，红色输出
func ErrorConsole(msg string) {
	colorOutConsole(msg, "red")
}

// WarningConsole 打印一条提示消息，黄色输出
func WarningConsole(msg string) {
	colorOutConsole(msg, "yellow")
}

// ExitConsole 打印一条报错消息，并退出 os.ExitConsole(1)
func ExitConsole(msg string) {
	ErrorConsole(msg)
	os.Exit(1)
}

// ExitIfConsole 语法糖，自带 err != nil 判断
func ExitIfConsole(err error) {
	if err != nil {
		ExitConsole(err.Error())
	}
}

// colorOutConsole 内部使用，设置高亮颜色
func colorOutConsole(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}

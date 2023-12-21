package config

// 写入内核配置文件，暂未使用

import (
	"bufio"
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"os"
	"reflect"
	"strings"
)

// CoreWrite 写入配置文件
func CoreWrite(configFilePath string, newConfig *coreConfig) error {
	newConfigMap := make(map[string]string)
	newConfigWriteFlag := make(map[string]bool)

	// 通过反射获取结构体中的yaml字段值和值内容
	valueType := reflect.TypeOf(*newConfig)
	value := reflect.ValueOf(*newConfig)
	var k, v string
	for i := 0; i < valueType.NumField(); i++ {
		k = valueType.Field(i).Tag.Get("yaml")
		v = value.Field(i).String()
		newConfigMap[k] = v
		newConfigWriteFlag[k] = false
	}

	log.Debugln("%+v", newConfigMap)
	// map[external-controller:0.0.0.0:7899 external-ui:./ui/ secret:hi]

	// 打开文件以供读写
	file, err := os.OpenFile(configFilePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// 创建一个用于写入文件的缓冲写入器
	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		// 检查行是否以特定字符串开头
		for k, v := range newConfigMap {
			if strings.HasPrefix(line, k) {
				// 如果匹配，则替换该行内容
				line = k + ": " + v
				newConfigWriteFlag[k] = true
			}
		}

		lines = append(lines, line)
	}

	// 清空文件内容
	err = file.Truncate(0)
	if err != nil {
		//fmt.Println("Error truncating file:", err)
		return err
	}

	// 将文件指针移动到文件开头
	_, err = file.Seek(0, 0)
	if err != nil {
		//fmt.Println("Error seek:", err)
		return err
	}

	// 将处理后的行写入文件
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			//fmt.Println("Error writing to file:", err)
			return err
		}
	}

	// 将未写入的配置写入文件末尾
	for k, v := range newConfigWriteFlag {
		if !v {
			line := k + ": " + newConfigMap[k]
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				//fmt.Println("Error writing to file:", err)
				return err
			}
		}
	}

	// 刷新缓冲并确保所有写入操作完成
	if err := writer.Flush(); err != nil {
		//fmt.Println("Error flushing writer:", err)
		return err
	}

	if err := scanner.Err(); err != nil {
		//fmt.Println("Error scanning file:", err)
		return err
	}

	//fmt.Println("File updated successfully.")

	return nil
}

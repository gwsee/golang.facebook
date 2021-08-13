package http

import (
	"encoding/json"
	"os"
	"time"
)

func saveFile(urls, method string, postData map[string]string, by []byte, errs error) (fileName string, err error) {
	var saveObj []byte
	saveObj = append(saveObj, []byte(urls)...)
	saveObj = append(saveObj, '\n')
	saveObj = append(saveObj, []byte(method)...)
	saveObj = append(saveObj, '\n')
	buf, _ := json.Marshal(postData)
	saveObj = append(saveObj, buf...)
	saveObj = append(saveObj, '\n')
	if SaveActionLogAns || errs != nil {
		saveObj = append(saveObj, by...)
		saveObj = append(saveObj, '\n')
	}
	saveObj = append(saveObj, []byte("facebook_request_finished")...)
	saveObj = append(saveObj, '\n')
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		if errs == nil {
			logFilePath = dir + SaveActionDir + "/facebook_log/" + now.Format("200601") + "/" + now.Format("02") + "/"
		} else {
			logFilePath = dir + SaveActionDir + "/facebook_log/" + now.Format("200601") + "/" + now.Format("02") + "_errs/"
		}
	}
	if _, err = os.Stat(logFilePath); os.IsNotExist(err) {
		os.MkdirAll(logFilePath, 0777)
		os.Chmod(logFilePath, 0777)
	}
	logFileName := now.Format("1504") + ".log"
	fileName = logFilePath + logFileName
	fileObj, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer fileObj.Close()
	fileObj.Write(saveObj)
	//保存文件结束！！
	return
}

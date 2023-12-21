package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/config"
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"github.com/xxnuo/mihomo-launcher/res"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type ResponseBody struct {
	CoreType   string `json:"core_type"`
	BinPath    string `json:"bin_path"`
	ConfigDir  string `json:"config_dir"`
	ConfigFile string `json:"config_file"`
	LogFile    string `json:"log_file"`
	ExtArgs    string `json:"ext_args"`
}

type JsonResponse struct {
	Code uint64        `json:"code"`
	Msg  string        `json:"msg"`
	Data *ResponseBody `json:"data"`
}

var (
	CoreStatus = Unavailable // 0: unavailable, 1: running, 2: stopped
)

// EnsureServiceRunning 检查服务是否运行及端口占用情况，如果未运行则启动服务
// 初始化 res.EnsureServiceExes 后执行
func EnsureServiceRunning() error {
	err := error(nil)
	CoreStatus = Unavailable
	if utils.IsPortAvailable(ServicePort) {
		// 端口未被占用，服务未运行，重新安装服务
		//_ = res.ExecUninstallService()
		log.Infoln("未检测到服务运行，初次使用请允许管理员权限弹窗安装服务")
		err = res.ExecInstallService()
		if err == nil {
			CoreStatus = Running
		}

	} else {
		// 端口被占用
		// 检查是否是本程序的服务
		status, err := CheckService()
		if status.Msg == "ok" {
			CoreStatus = Running
		} else {
			err = fmt.Errorf("port %d is occupied by other program: %w", ServicePort, err)
		}

		if status.Code == 400 {
			CoreStatus = Stopped
		}
	}
	return err
}

func CheckService() (*JsonResponse, error) {
	url := fmt.Sprintf("%s/get_clash", ServiceURL)
	resp, err := http.Get(url)
	if err != nil {
		CoreStatus = Unavailable
		return nil, fmt.Errorf("failed to connect to the Clash Verge Service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response JsonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the Clash Verge Service response: %w", err)
	}

	if response.Code == 400 {
		CoreStatus = Stopped
	} else if response.Code == 0 {
		CoreStatus = Running
	} else {
		CoreStatus = Unavailable
	}

	return &response, nil
}

func RestartCore() error {
	//status, err := CheckService()
	//if err != nil {
	//	return err
	//}
	//if status.Code == 0 {
	//	err = StopCore()
	//	if err != nil {
	//		return err
	//	}
	//	time.Sleep(1 * time.Second)
	//	return RunCore()
	//}
	//return nil

	// 服务内部会自动重启
	return RunCore()
}

func RunCore() error {
	status, err := CheckService()
	if err != nil {
		// 服务未运行
		CoreStatus = Unavailable
		return err
	}
	if status.Code == 0 {
		// 内核已运行
		CoreStatus = Running
		return nil
	}

	url := fmt.Sprintf("%s/start_clash", ServiceURL)
	data := &ResponseBody{
		CoreType:   "clash-meta-alpha",
		BinPath:    res.CoreExePath,
		ConfigDir:  config.CoreConfigDir,
		ConfigFile: filepath.Join(config.CoreConfigDir, config.CoreConfigFilename),
		LogFile:    filepath.Join(config.CoreConfigDir, "launcher.log"),
		ExtArgs:    strings.Join(config.LauncherConfig.CoreArgs.GetArgs(), " "),
	}

	jsonData, err := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("failed to post %s :%w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var response JsonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to parse the Clash Verge Service response: %w", err)
	}

	if response.Code != 0 {
		CoreStatus = Stopped
		return fmt.Errorf("failed to start the Clash Verge Service: %s", response.Msg)
	} else {
		CoreStatus = Running
	}

	return nil
}

func StopCore() error {
	status, err := CheckService()
	if err != nil {
		return err
	}
	if status.Code == 400 {
		//	clash not executed
		CoreStatus = Stopped
		return nil
	}

	//if status.Code == 0 {
	// clash is running
	//}

	url := fmt.Sprintf("%s/stop_clash", ServiceURL)
	resp, err := http.Post(url, "application/json", nil)

	if err != nil {
		return fmt.Errorf("failed to post %s :%w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var response JsonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to parse the Clash Verge Service response: %w", err)
	}

	if response.Code != 0 {
		return fmt.Errorf("failed to stop the Clash Verge Service: %s", response.Msg)
	} else {
		CoreStatus = Stopped
	}

	return nil
}

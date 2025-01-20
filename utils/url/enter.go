package url

import "MikuMikuCloudDrive/config"

func ConcatWebUrl(url string) string {
	appConfig := config.ReadAppConfig()
	return appConfig.WebUrl + "/" + url
}

package frontend

import "github.com/freonservice/freon/pkg/api"

const (
	userStatusNotActive = "not active"
	userStatusBanned    = "banned"

	userRoleModerator = "moderator"
	userRoleAdmin     = "admin"
)

func getUserStatusByInteger(status api.UserStatus) string {
	switch status { //nolint:exhaustive
	case api.UserStatus_USER_NOT_ACTIVE:
		return userStatusNotActive
	case api.UserStatus_USER_IS_BANNED:
		return userStatusBanned
	default:
		return "active"
	}
}

func getUserStatusByString(status string) api.UserStatus {
	switch status {
	case userStatusNotActive:
		return api.UserStatus_USER_NOT_ACTIVE
	case userStatusBanned:
		return api.UserStatus_USER_IS_BANNED
	default:
		return api.UserStatus_USER_ACTIVE
	}
}

func getUserRoleByInteger(role api.UserRole) string {
	switch role { //nolint:exhaustive
	case api.UserRole_USER_ROLE_ADMIN:
		return userRoleAdmin
	case api.UserRole_USER_ROLE_MODERATOR:
		return userRoleModerator
	default:
		return "translator"
	}
}

func getUserRoleByString(role string) api.UserRole {
	switch role {
	case userRoleAdmin:
		return api.UserRole_USER_ROLE_ADMIN
	case userRoleModerator:
		return api.UserRole_USER_ROLE_MODERATOR
	default:
		return api.UserRole_USER_ROLE_TRANSLATOR
	}
}

func getPlatformByString(role string) api.PlatformType {
	switch role {
	case "web": //nolint:goconst
		return api.PlatformType_PLATFORM_TYPE_WEB
	case "ios": //nolint:goconst
		return api.PlatformType_PLATFORM_TYPE_IOS
	default:
		return api.PlatformType_PLATFORM_TYPE_ANDROID
	}
}

func getPlatformByInteger(platformType int64) string {
	switch platformType {
	case int64(api.PlatformType_PLATFORM_TYPE_WEB):
		return "web"
	case int64(api.PlatformType_PLATFORM_TYPE_IOS):
		return "ios"
	default:
		return "android"
	}
}

func getStorageTypeByInteger(storageType int64) string {
	switch storageType {
	case int64(api.StorageType_STORAGE_TYPE_LOCAL):
		return "local"
	default:
		return "external"
	}
}

func getStatusByInteger(status api.Status) string {
	switch status { //nolint:exhaustive
	case api.Status_ACTIVE:
		return "Active" //nolint:goconst
	default:
		return "Not active"
	}
}

func getTranslationStatus(status api.TranslationStatus) string {
	switch status { //nolint:exhaustive
	case api.TranslationStatus_TRANSLATION_ACTIVE:
		return "Active"
	default:
		return "Hidden"
	}
}

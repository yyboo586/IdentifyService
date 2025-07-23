package consts

const (
	// CachePrefix 应用缓存数据前缀
	CachePrefix = "APP:"

	CacheModelMem   = "memory"
	CacheModelRedis = "redis"
	CacheModelDist  = "dist"

	// CacheSysDict 字典缓存菜单KEY
	CacheSysDict = CachePrefix + "sysDict"

	// CacheSysDictTag 字典缓存标签
	CacheSysDictTag = CachePrefix + "sysDictTag"
	// CacheSysConfigTag 系统参数配置
	CacheSysConfigTag = CachePrefix + "sysConfigTag"
)

const (
	// CacheSysAuthMenu 缓存菜单key
	CacheSysAuthMenu = CachePrefix + "sysAuthMenu"
	// Org 缓存部门key
	CacheOrg = CachePrefix + "Org"

	// CacheSysRole 角色缓存key
	CacheSysRole = CachePrefix + "sysRole"
	// CacheSysWebSet 站点配置缓存key
	CacheSysWebSet = CachePrefix + "sysWebSet"
	// CacheSysCmsMenu cms缓存key
	CacheSysCmsMenu = CachePrefix + "sysCmsMenu"

	// CacheSysAuthTag 权限缓存TAG标签
	CacheSysAuthTag = CachePrefix + "sysAuthTag"
	// CacheSysModelTag 模型缓存标签
	CacheSysModelTag = CachePrefix + "sysModelTag"
	// CacheSysCmsTag cms缓存标签
	CacheSysCmsTag = CachePrefix + "sysCmsTag"
)

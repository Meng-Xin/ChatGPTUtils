package idefine

type IScenes interface {
	// GetScenesID 获取场景ID
	GetScenesID() int
	//  SetScenes 设置场景
	SetScenes(interface{}) error
}

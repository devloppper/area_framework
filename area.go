package area_framework

import "context"

// area 区域
type area interface {

	// OnReceiveHandler 消息接受处理器
	OnReceiveHandler(ctx context.Context, msg *Message)

	// OnStartHandler 区域初始化
	OnStartHandler(ctx context.Context)

	// OnStopHandler 区域结束工具
	OnStopHandler(ctx context.Context)

	// GetRound 获取x,y,w,h
	GetRound() (int, int, int, int)

	// Move 区域移动
	Move(x, y, deltaW, deltaY int)
}

// AreaPanel 区域操作面板
type AreaPanel interface {

	// WithBoardCaster 添加一个广播器
	WithBoardCaster(bc *BoardCaster)

	// NewArea 新建一个区域
	NewArea(x, y, w, h int) (*area, error)

	// Move 移动该Panel下的所有Area
	Move(x, y, delaW, deltaY int)

	// OnStartHandler 当区域初始化 操作面板进行操作 在所有区域开始之后
	OnStartHandler(ctx context.Context)

	// OnStopHandler 当区域结束 操作面板需要执行的操作 在所有区域全部停止后
	OnStopHandler(ctx context.Context)
}

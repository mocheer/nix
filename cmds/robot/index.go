package robot

import "time"

// ComputerRobot
// 电脑机器人，根据提供的指令，自动模拟鼠标、键盘的行为事件，控制应用程序，可截图录屏分析
type ComputerRobot struct {
	actions []robotAction
}

type robotAction struct {
	RunType string
	At      time.Duration
}

func (m *ComputerRobot) AddAction(action robotAction) {
	m.actions = append(m.actions, action)
}

func (m *ComputerRobot) RemoveAction() {

}

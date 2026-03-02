package loader

func (model Model) View() string {
	return "Connecting...\t" + model.bar.View()
}
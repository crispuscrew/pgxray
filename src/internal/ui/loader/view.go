package loader

func (model Model) View() string {
	return "Connecting...\n\n" + model.bar.ViewAs(model.percent)
}
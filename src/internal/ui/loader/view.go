package loader

func (model Model) View() string {
	if model.Desc != "" { return model.Desc + "\t" + model.bar.View() }
	return model.bar.View()
}
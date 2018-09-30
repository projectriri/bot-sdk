package botsdk

type Bots []Bot

func (bots Bots) GetUpdatesChan(bufferSize int) (UpdateChannel, UpdateErrorChannel, error) {
	updc := make(chan Update)
	errc := make(chan error)
	upda := make([]UpdateChannel, 0, len(bots))
	erra := make([]UpdateErrorChannel, 0, len(bots))
	for _, bot := range bots {
		ch, ech, err := bot.GetUpdatesChan(bufferSize)
		if err != nil {
			return nil, nil, err
		}
		upda = append(upda, ch)
		erra = append(erra, ech)
	}
	for i := 0; i < len(upda); i++ {
		go func(ch UpdateChannel, ech UpdateErrorChannel) {
			for {
				select {
				case update := <-ch:
					updc <- update
				case e := <-ech:
					errc <- e
				}
			}
		}(upda[i], erra[i])
	}
	return updc, errc, nil
}

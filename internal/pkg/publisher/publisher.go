package publisher

const (
	// Telegram const.
	Telegram = "telegram"
	// Stdout const.
	Stdout = "stdout"
)

// Item struct.
type Item struct {
	Title              string
	URL                string
	DiscountPercentage *string
	Price              *float64
}

// Publisher is a publisher interface.
type Publisher interface {
	Publish(item Item)
}

// Publishers contains the available publishers.
type Publishers struct {
	TelegramPublisher *TelegramPublisher
	StdoutPublisher   *StdoutPublisher
}

// NewPublishers creates a new Publishers struct.
func NewPublishers() Publishers {
	return Publishers{
		NewTelegramPublisher(),
		NewStdoutPublisher(),
	}
}

// GetPublisher returns a new Publisher based on a provided name.
func (v *Publishers) GetPublisher(scraperName string) Publisher {
	switch scraperName {
	case Stdout:
		return v.StdoutPublisher
	case Telegram:
		return v.TelegramPublisher
	default:
		return nil
	}
}
